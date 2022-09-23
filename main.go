package main

import (
	"github.com/sheetpilot/sheetpilot/updateManifest"
	"k8s.io/cli-runtime/pkg/printers"
	"log"
	"os"
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
	updatedObj := updateManifest.UpdateReplicas(obj, numOfReplica)
	newFile, err := os.Create(fname)
	if err != nil {
		log.Println(err)
	}
	y := printers.YAMLPrinter{}
	defer newFile.Close()
	y.PrintObj(updatedObj, newFile)
}