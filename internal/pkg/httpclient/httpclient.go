package httpclient

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type HttpClient struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
	apiSecret  string
}

type RequestData interface {
	Path() (string, error)
	Encode() (string, error)
}

type ResponseData interface {
	Decode(body io.ReadCloser) error
	SetBody(body io.ReadCloser)
	SetStatusCode(code int)
}

func New(baseURL, apiKey, apiSecret string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL:   baseURL,
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (h *HttpClient) Get(requestData RequestData, responseData ResponseData) error {

	requestDataPath, err := requestData.Path()
	if err != nil {
		return err
	}

	requestURL := h.baseURL + requestDataPath
	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(h.apiKey, h.apiSecret)
	request.Header.Set("Accept", "application/json")

	response, err := h.httpClient.Do(request)
	if err != nil {
		return err
	}

	responseData.SetStatusCode(response.StatusCode)

	if isJSONData(response.Header.Get("Content-Type")) {
		return responseData.Decode(response.Body)
	}

	responseData.SetBody(response.Body)

	return nil
}

func (h *HttpClient) Post(requestData RequestData, responseData ResponseData) error {

	requestDataPath, err := requestData.Path()
	if err != nil {
		return err
	}

	requestDataJSON, err := requestData.Encode()
	if err != nil {
		return err
	}

	requestURL := h.baseURL + requestDataPath
	request, err := http.NewRequest("POST", requestURL, strings.NewReader(requestDataJSON))
	if err != nil {
		return err
	}

	request.SetBasicAuth(h.apiKey, h.apiSecret)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err := h.httpClient.Do(request)
	if err != nil {
		return err
	}

	responseData.SetStatusCode(response.StatusCode)

	if isJSONData(response.Header.Get("Content-Type")) {
		return responseData.Decode(response.Body)
	}

	responseData.SetBody(response.Body)

	return nil
}

func (h *HttpClient) Upload(requestData RequestData, responseData ResponseData, filename string) error {

	requestDataPath, err := requestData.Path()
	if err != nil {
		return err
	}

	multipartData := &bytes.Buffer{}
	multipartContentType, err := buildMultipart(multipartData, filename)
	if err != nil {
		return err
	}

	requestURL := h.baseURL + requestDataPath
	request, err := http.NewRequest("POST", requestURL, bytes.NewReader(multipartData.Bytes()))
	if err != nil {
		return err
	}

	request.SetBasicAuth(h.apiKey, h.apiSecret)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", multipartContentType)

	response, err := h.httpClient.Do(request)
	if err != nil {
		return err
	}

	responseData.SetStatusCode(response.StatusCode)

	if isJSONData(response.Header.Get("Content-Type")) {
		return responseData.Decode(response.Body)
	}

	responseData.SetBody(response.Body)

	return nil

}

func (h *HttpClient) Delete(requestData RequestData, responseData ResponseData) error {

	requestDataPath, err := requestData.Path()
	if err != nil {
		return err
	}

	requestURL := h.baseURL + requestDataPath
	request, err := http.NewRequest("DELETE", requestURL, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(h.apiKey, h.apiSecret)

	response, err := h.httpClient.Do(request)
	if err != nil {
		return err
	}

	responseData.SetStatusCode(response.StatusCode)

	if isJSONData(response.Header.Get("Content-Type")) {
		return responseData.Decode(response.Body)
	}

	responseData.SetBody(response.Body)

	return nil
}

// support methods

func isJSONData(contentType string) bool {

	if contentType == "" {
		return false
	}

	for _, v := range strings.Split(contentType, ";") {
		if strings.TrimSpace(v) == "application/json" {
			return true
		}
	}

	return false
}

func buildMultipart(body *bytes.Buffer, filename string) (string, error) {

	multipartWriter := multipart.NewWriter(body)

	part, err := multipartWriter.CreateFormFile("image", filepath.Base(filename))
	if err != nil {
		return "", err
	}
	defer multipartWriter.Close()

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	return multipartWriter.FormDataContentType(), nil
}
