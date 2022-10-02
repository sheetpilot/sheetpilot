package main

import (
	"fmt"
	"github.com/sheetpilot/sheetpilot/github"
	"github.com/sheetpilot/sheetpilot/updateManifest"
	"k8s.io/cli-runtime/pkg/printers"
	"log"
	"os"
)

func main() {
	pat := ""
	gitURL := "https://github.com/sheetpilot/sample-deployment.git"
	owner := "dtherhtun"
	app := "admin"
	var numOfReplica int32 = 10

	tempDir, err := github.Clone(gitURL, pat)
	if err != nil {
		fmt.Println(err)
	}

	m, _ := findDeployment(tempDir)

	yamlData, err := os.ReadFile(m[app])
	if err != nil {
		log.Println(err)
	}
	obj, err := updateManifest.CheckDeployment(app, yamlData)
	if err != nil {
		log.Println(err)
	}
	Obj := updateManifest.UpdateReplicas(obj, numOfReplica)
	Obj = updateManifest.UpdateResources(Obj, 700, 1)
	newFile, err := os.Create(m[app])
	if err != nil {
		log.Println(err)
	}
	y := printers.YAMLPrinter{}
	defer newFile.Close()
	y.PrintObj(Obj, newFile)

	if err = github.Commit(tempDir); err != nil {
		fmt.Println(err)
	}

	err, cleanup := github.Push(tempDir, owner, pat)
	if err != nil {
		fmt.Println(err)
	}
	defer cleanup()
}
