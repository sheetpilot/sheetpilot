package gitapiclient

import (
	"fmt"

	"github.com/sheetpilot/sheetpilot/internal/httpclient"
)

type gitApiClientFactory struct {
	clientName string
	token      string
}

func GitApiClientFactory(factoryArgs *gitApiClientFactory) (func(repoPath string) iGitClient, error) {
	client := new(gitApiClientFactory)
	return client.GetGitApiClient(factoryArgs)
}

func (f *gitApiClientFactory) GetGitApiClient(args *gitApiClientFactory) (func(repoPath string) iGitClient, error) {
	if args.clientName == "github" {
		client := NewGitHubClient("https://api.github.com/", args.token)
		client.githubapiClient(httpclient.NewHttpClient())

		return client.githubapiClient(httpclient.NewHttpClient()), nil
	}

	// add new client name what ever you need
	return nil, fmt.Errorf("Wrong Client passed")
}
