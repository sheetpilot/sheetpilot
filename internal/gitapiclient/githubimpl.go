package gitapiclient

type githubClientImpl struct {
	apiHost string
	token   string
	path    string
}

func (g *githubClientImpl) setApiHost(apiHost string) {
	g.apiHost = apiHost
}

func (g *githubClientImpl) GetApiHost() string {
	return g.apiHost
}

func (g *githubClientImpl) setToken(token string) {
	g.token = token
}

func (g *githubClientImpl) GetToken() string {
	return g.token
}

func (g *githubClientImpl) setPath(path string) {
	g.path = path
}

func (g *githubClientImpl) GetPath() string {
	return g.path
}

func (g *githubClientImpl) GetFileContentFromRepo() {

}
