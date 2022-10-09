package gitapiclient

type gitHubClient struct {
	githubClientImpl
}

func newGithubClient(repoPath string, token string) iGitClient {

	ghClient := new(gitHubClient)

	apiClient := new(githubClientImpl)
	apiClient.setApiHost("https://api.github.com")
	apiClient.setPath(repoPath)
	apiClient.setToken(token)
	ghClient.githubClientImpl = *apiClient
	return ghClient
}
