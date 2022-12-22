package gitapiclient

import "github.com/sheetpilot/sheetpilot/internal/httpclient"

type gitHubClient struct {
	githubClientImpl
}

func NewGitHubClient(host string, token string) gitHubClient {
	apiclient := new(gitHubClient)

	apiclient.setApiHost(host)
	apiclient.setToken(token)
	return *apiclient
}

func (gh *gitHubClient) githubapiClient(httpClient httpclient.IHttpClient) func(repoPath string) iGitClient {

	return func(repoPath string) iGitClient {
		apiClient := new(githubClientImpl)
		apiClient.setApiHost(gh.apiHost)
		apiClient.setToken(gh.token)
		apiClient.setHttpClient(httpClient)
		apiClient.setPath(repoPath)

		return apiClient
	}
}
