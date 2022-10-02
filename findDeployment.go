package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func findDeployment(repo string) (map[string]string, error) {
	var files []string
	appFiles := make(map[string]string)

	err := filepath.Walk(repo, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Base(path) == "deployment.yaml" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		baseDir := strings.Split(file, "/")
		app := baseDir[len(baseDir)-2 : len(baseDir)-1][0]
		appFiles[app] = file
	}

	return appFiles, nil
}
