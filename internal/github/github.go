package github

import (
	"fmt"
	"os"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/sirupsen/logrus"
)

func Clone(repo, pat string, log *logrus.Entry) (string, error) {

	tempDir, err := os.MkdirTemp("", "sheet-pilot")
	if err != nil {
		return "", fmt.Errorf("can not create temporary directory: %w", err)
	}

	log.Infof("cloning %s repo into the %s path", repo, tempDir)

	if _, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "",
			Password: pat,
		},
		URL:      repo,
		Progress: os.Stdout,
	}); err != nil {
		return "", fmt.Errorf("can not clone repo: %w", err)
	}
	log.Infof("%s repo successfully cloned into %s", repo, tempDir)

	return tempDir, nil
}

func Commit(repoPath string, log *logrus.Entry) error {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("can not open temporary git repo directory: %w", err)
	}

	commitWorkTree, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("something wrong with git repo: %w", err)
	}

	if _, err := commitWorkTree.Add("."); err != nil {
		return fmt.Errorf("can not add: %w", err)
	}

	if _, err := commitWorkTree.Commit("updated", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "sheetPilot",
			Email: "info@sheetpilot.io",
			When:  time.Now(),
		},
	}); err != nil {
		return fmt.Errorf("can not commit: %w", err)
	}

	log.Infof("git commit completed")

	return nil
}

func Push(repoPath, owner, pat string, log *logrus.Entry) (error, func(log *logrus.Entry)) {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("can not open temporary git repo directory: %w", err), nil
	}

	if err = r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: owner,
			Password: pat,
		},
		Progress: os.Stdout,
	}); err != nil {
		return fmt.Errorf("can not push: %w", err), nil
	}
	log.Infof("git push completed")

	return nil, func(log *logrus.Entry) {
		os.RemoveAll(repoPath)
		log.Infof("cleanup temporary directory")
	}
}
