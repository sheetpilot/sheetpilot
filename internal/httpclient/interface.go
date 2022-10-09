package httpclient

import "net/http"

type iHttpClient interface {
	setHost(host string) iHttpClient
	setHeader(req *http.Request, queryParams map[string]interface{}) iHttpClient
	setBearerToken(token string) iHttpClient
	setTimeout(seconds int64) iHttpClient
	Get(path string, queryParams map[string]interface{}) ([]byte, error)
	Post(path string, queryParams map[string]interface{}, body map[string]interface{}) ([]byte, error)
}
