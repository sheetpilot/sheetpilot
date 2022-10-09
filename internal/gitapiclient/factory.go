package gitapiclient

import "fmt"

func GetGitApiClient(clientName string) (iGitClient, error) {
	if clientName == "github" {
		client := newGithubClient("actual repo path should be here /repo/path", "Dummy Token")
		client.setApiHost("https://api.github.com/")
		client.setToken("token")

		return client, nil
	}

	// add new client name what ever you need
	return nil, fmt.Errorf("Wrong Client passed")
}
