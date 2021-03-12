package web

import (
	"fmt"
	"io"
	"net/http"
)

type HttpService interface {
	Send(method, url string, body io.Reader, headers map[string]string) (*http.Response, error)
}

type httpService struct {
	httpClient http.Client
}

func NewHttpService() HttpService {

	client := http.Client{

	}

	return &httpService{
		httpClient: client,
	}
}

func (httpService *httpService) Send(method, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if request.Body != nil {
		defer request.Body.Close()
	}

	httpService.addHeaders(request, headers)

	response, err := httpService.httpClient.Do(request)

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return response, nil
}

func (httpService *httpService) addHeaders(request *http.Request, headers map[string]string) {
	for headerKey, headerValue := range headers {
		request.Header.Set(headerKey, headerValue)
	}
}
