package main

import (
	"net"
	"os"

	"github.com/sheetpilot/sheetpilot/internal/controller"
	"github.com/sheetpilot/sheetpilot/internal/helper"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

var log *logrus.Entry

func init() {
	l := logrus.New()

	log = l.WithFields(logrus.Fields{
		"app": map[string]string{
			"host": os.Getenv("HOST"),
		},
	})
}

func main() {
	srv := grpc.NewServer()

	scaleController := controller.NewScaleController(srv)
	scaleController.RegisterService()

	listener, conErr := net.Listen("tcp", helper.GetEnv("PORT", ":10001"))
	log.Infof("Sheet Pilot Git Push Service Listening at http://%s", helper.GetEnv("APP_HOST_ADDRESS", "127.0.0.1:10001"))

	if conErr != nil {
		panic(conErr)
	}

	if conn := srv.Serve(listener); conn != nil {
		panic(conn)
	}

}

// package main

// import (
// 	"fmt"
// 	"github.com/sheetpilot/sheetpilot/github"
// 	"github.com/sheetpilot/sheetpilot/updateManifest"
// 	"k8s.io/cli-runtime/pkg/printers"
// 	"log"
// 	"os"
// )

// func main() {
// 	pat := ""
// 	gitURL := "https://github.com/sheetpilot/sample-deployment.git"
// 	owner := "dtherhtun"
// 	app := "admin"
// 	var numOfReplica int32 = 10

// 	tempDir, err := github.Clone(gitURL, pat)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	appFiles, err := findDeployment(tempDir)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	yamlData, err := os.ReadFile(appFiles[app])
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	obj, err := updateManifest.CheckDeployment(app, yamlData)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	updatedObj := updateManifest.UpdateResourceValues(obj, 700, 1, 1000, 2, numOfReplica)

// 	newFile, err := os.Create(appFiles[app])
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	y := printers.YAMLPrinter{}
// 	defer newFile.Close()
// 	y.PrintObj(updatedObj, newFile)

// 	if err = github.Commit(tempDir); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	err, cleanup := github.Push(tempDir, owner, pat)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer cleanup()
// }

// TestCode
//func main() {
//	var numOfReplica int32 = 10
//	yamlData, err := os.ReadFile("nginx-deployment.yaml")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	obj, err := updateManifest.CheckDeployment("my-nginx", yamlData)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	updatedObj := updateManifest.UpdateResourceValues(obj, "6", "1100Mi", "5", "", numOfReplica)
//
//	newFile, err := os.Create("new-nginx-deployment.yaml")
//	if err != nil {
//		log.Println(err)
//	}
//	updateManifest.PrintDeployment(newFile, updatedObj)
//}
