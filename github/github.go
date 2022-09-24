package github

import (
	"fmt"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
	"time"
)

func Clone(repo, pat string) (string, error) {

	tempDir, err := os.MkdirTemp("", "sheet-pilot")
	if err != nil {
		return "", fmt.Errorf("can not create temporary directory: %w", err)
	}
	fmt.Println(tempDir)

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

	return tempDir, nil
}

func Commit(repoPath string) error {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("can not open temporary git repo directory: %w", err)
	}

	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("something wrong with git repo: %w", err)
	}

	if _, err := w.Add("."); err != nil {
		return fmt.Errorf("can not add: %w", err)
	}

	fmt.Println(w.Status())
	if _, err := w.Commit("updated", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "sheetPilot",
			Email: "info@sheetpilot.io",
			When:  time.Now(),
		},
	}); err != nil {
		return fmt.Errorf("can not commit: %w", err)
	}

	return nil
}

func Push(repoPath, owner, pat string) error {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("can not open temporary git repo directory: %w", err)
	}

	if err = r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: owner,
			Password: pat,
		},
		Progress: os.Stdout,
	}); err != nil {
		return fmt.Errorf("can not push: %w", err)
	}

	return nil
}
