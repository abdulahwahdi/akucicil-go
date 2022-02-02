package akucicil

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// httpRequest data model
type akucicilHttp struct {
	httpClient *http.Client
}

// newRequest function for intialize httpRequest object
// Paramter, timeout in time.Duration
func newRequest(timeout time.Duration) *akucicilHttp {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &akucicilHttp{
		httpClient: &http.Client{Timeout: time.Second * timeout, Transport: tr},
	}
}

// newReq function for initalize http request,
// paramters, http method, uri path, body, and headers
func (bb *akucicilHttp) newReq(method string, fullPath string, body io.Reader, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

// exec private function for call http request
func (bb *akucicilHttp) exec(method, path string, body io.Reader, v interface{}, headers map[string]string) error {
	req, err := bb.newReq(method, path, body, headers)

	if err != nil {
		return err
	}

	res, errHttpClient := bb.httpClient.Do(req)
	if errHttpClient != nil {
		return err
	}

	defer res.Body.Close()

	if v != nil {
		return json.NewDecoder(res.Body).Decode(v)
	}

	return nil
}
