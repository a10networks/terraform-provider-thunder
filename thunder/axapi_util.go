package thunder

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"
	"util"
)

func doHttp(method string, host string, body io.Reader, headers map[string]string) (*http.Response, error) {
	logger := util.GetLoggerInstance()
	logger.Println(method, ">", host)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 30 * time.Second}
	req, err := http.NewRequest(method, host, body)
	if err != nil {
		logger.Println("failed to create new request")
	}

	for property, value := range headers {
		req.Header.Set(property, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Println("failed at http.Client.Do()")
	}
	return resp, err
}

func GetReqHeader(auth_token string) map[string]string {
	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = auth_token
	return headers
}
