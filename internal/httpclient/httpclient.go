package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type httpClient struct {
	client      *http.Client
	host        string
	bearerToken string
}

func NewHttpClient() IHttpClient {
	apiClient := new(httpClient)

	apiClient.client = new(http.Client)
	apiClient.client.Timeout = 30 * time.Second
	return apiClient
}

func (h *httpClient) setHost(host string) IHttpClient {
	h.host = host
	return h
}

func (h *httpClient) setBearerToken(token string) IHttpClient {
	h.bearerToken = token
	return h
}

func (h *httpClient) setHeader(req *http.Request, queryParams map[string]interface{}) IHttpClient {

	keys := make([]string, 0)
	for k, _ := range queryParams {
		keys = append(keys, k)
	}

	for _, key := range keys {
		req.Header.Set(key, fmt.Sprint(queryParams[key]))
	}

	return h
}

func (h *httpClient) setTimeout(timeout int64) IHttpClient {
	h.client.Timeout = time.Duration(timeout) * time.Second
	return h
}

func (h *httpClient) Get(path string, queryParams map[string]interface{}) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, h.host+path, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.URL.RawQuery = getUrlQueryStringFromQueryParams(req, queryParams)

	resp, err := h.client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request")
		return make([]byte, 0), err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseBody, err

}

func (h *httpClient) Post(path string, queryParams map[string]interface{}, body map[string]interface{}) ([]byte, error) {
	jsonBuffer, jsonError := prepareJsonPayload(body)
	if jsonError != nil {
		log.Fatal(jsonError)
	}

	req, err := http.NewRequest(http.MethodPost, h.host+path, jsonBuffer)
	if err != nil {
		log.Fatal(err)
	}

	req.URL.RawQuery = getUrlQueryStringFromQueryParams(req, queryParams)

	resp, err := h.client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request")
		return make([]byte, 0), err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseBody, err
}

func prepareJsonPayload(values map[string]interface{}) (*bytes.Buffer, error) {
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return bytes.NewBuffer(jsonData), nil

}

func getUrlQueryStringFromQueryParams(req *http.Request, queryParams map[string]interface{}) string {
	q := req.URL.Query()
	keys := make([]string, 0)
	for k, _ := range queryParams {
		keys = append(keys, k)
	}

	for _, key := range keys {
		q.Add(key, fmt.Sprint(queryParams[key]))
	}

	return q.Encode()
}
