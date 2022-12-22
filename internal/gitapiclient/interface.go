package gitapiclient

type iGitClient interface {
	setApiHost(name string)
	setToken(token string)
	setPath(path string)
	GetApiHost() string
	GetToken() string
	GetPath() string
}
