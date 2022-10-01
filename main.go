package main

import (
	"fmt"
	"github.com/sheetpilot/sheetpilot/github"
	"github.com/sheetpilot/sheetpilot/updateManifest"
	"k8s.io/cli-runtime/pkg/printers"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fname := "nginx-deployment.yaml"
	deployName := "my-nginx"
	var numOfReplica int32 = 5
	yamlData, err := os.ReadFile(fname)
	if err != nil {
		log.Println(err)
	}
	obj, err := updateManifest.CheckDeployment(deployName, yamlData)
	if err != nil {
		log.Println(err)
	}
	Obj := updateManifest.UpdateReplicas(obj, numOfReplica)
	Obj = updateManifest.UpdateResources(Obj, 500, 2)
	newFile, err := os.Create(fname)
	if err != nil {
		log.Println(err)
	}
	y := printers.YAMLPrinter{}
	defer newFile.Close()
	y.PrintObj(Obj, newFile)

	pat := ""
	gitURL := "https://github.com/sheetpilot/sample-deployment.git"
	owner := "dtherhtun"

	tempDir, err := github.Clone(gitURL, pat)
	if err != nil {
		fmt.Println(err)
	}
	testfile := filepath.Join(tempDir, "testfile.txt")
	err = os.WriteFile(testfile, []byte("hello world!"), 0644)
	if err != nil {
		fmt.Println("can not create test file")
	}
	if err = github.Commit(tempDir); err != nil {
		fmt.Println(err)
	}

	if err = github.Push(tempDir, owner, pat); err != nil {
		fmt.Println(err)
	}
	os.Remove(tempDir)
}
