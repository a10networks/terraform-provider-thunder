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
	//logger.Println("*****"  + method + "    "+ host + "******" )
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, host, body)
	//logger.Println("req new request - ", req.StatusCode)

	if err != nil {
		logger.Println("[INFO] error creating new request")
	}

	for property, value := range headers {
		req.Header.Set(property, value)
	}
	resp, err := client.Do(req)
	logger.Println("req new request - ", resp.StatusCode)

	if err != nil {
		logger.Println("[INFO] error ")
	}

	return resp, err
}
