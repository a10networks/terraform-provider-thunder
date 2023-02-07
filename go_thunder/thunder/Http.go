package go_thunder

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"time"
	"util"
)

func DoHttp(method string, host string, body io.Reader, headers map[string]string) (*http.Response, error) {

	logger := util.GetLoggerInstance()
	//defer SessionLogOff(host, headers)
	logger.Println("[INFO] inside do http")
	//logger.Println("*****"  + method + "    "+ host + "******" )
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//client := &http.Client{Transport: tr}
	client := &http.Client{Transport: tr, Timeout: 30 * time.Second}
	req, err := http.NewRequest(method, host, body)
	//logger.Println("req new request - ", req.StatusCode)
	if err != nil {
		logger.Println("[INFO] error creating new request")
	}

	for property, value := range headers {
		req.Header.Set(property, value)
	}
	resp, err := client.Do(req)
	//logger.Println("DoHttp response body ---> ", resp)
	if err != nil {
		logger.Println("[INFO] error from Http.go ")
	}

	return resp, err
}

func SessionLogOff(host string, headers map[string]string) {
	logger := util.GetLoggerInstance()

	u, err := url.Parse(host)
	hostadd := u.Host
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 30 * time.Second}
	// logger.Println("host---->  " + host)
	// logger.Println("GET", "https://"+hostadd+"/axapi/v3/logoff")
	// logger.Println("headers---------> ")
	// logger.Println(headers)

	req, err := http.NewRequest("GET", "https://"+hostadd+"/axapi/v3/logoff", nil)
	if err != nil {
		logger.Println(err)
		logger.Println("[WARNING] clearing session log fail")
	}

	for property, value := range headers {
		req.Header.Set(property, value)
	}
	resp, err := client.Do(req)
	logger.Println("resp--------->")
	logger.Println(resp)
	if err != nil {
		logger.Println(err)
		logger.Println("[WARNING] RESP - clearing session log fail ")
	}

	logger.Println("[INFO] clear session log ")
}
