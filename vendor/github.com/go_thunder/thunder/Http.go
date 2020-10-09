package go_thunder

import (
	"crypto/tls"
	"io"
	"net/http"
	"util"
)

func DoHttp(method string, host string, body io.Reader, headers map[string]string) (*http.Response, error) {

	logger := util.GetLoggerInstance()

	logger.Println("[INFO] inside do http")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, host, body)
	if err != nil {
		logger.Println("[INFO] error creating new request")
	}

	for property, value := range headers {
		req.Header.Set(property, value)
	}

	resp, err := client.Do(req)

	logger.Println("[INFO] status -" + resp.Status)

	if err != nil {
		logger.Println("[INFO] error ")
	}

	return resp, err
}
