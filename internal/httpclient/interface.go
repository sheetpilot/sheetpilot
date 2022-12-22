package httpclient

import "net/http"

type IHttpClient interface {
	setHost(host string) IHttpClient
	setHeader(req *http.Request, queryParams map[string]interface{}) IHttpClient
	setBearerToken(token string) IHttpClient
	setTimeout(seconds int64) IHttpClient
	Get(path string, queryParams map[string]interface{}) ([]byte, error)
	Post(path string, queryParams map[string]interface{}, body map[string]interface{}) ([]byte, error)
}
