package axapi

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/clarketm/json"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GenRequestHeaderNoToken() map[string]string {
	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	return headers
}

func GenRequestHeader(authToken string) map[string]string {
	var headers = GenRequestHeaderNoToken()
	headers["Authorization"] = authToken
	return headers
}

type AxApiResponse struct {
	Inst struct {
		Status string `json:"status"`
		Err    struct {
			Code int    `json:"code"`
			From string `json:"from"`
			Msg  string `json:"msg"`
		} `json:"err"`
	} `json:"response"`
}

func IsIpv4NetmaskBrief(val interface{}, key string) (warns []string, errs []error) {
	value := val.(string)
	if false == strings.HasPrefix(value, "/") {
		errs = append(errs, fmt.Errorf("must be /0 to /32, got: %s", value))
	} else {
		tokens := strings.Split(value, "/")
		prefixLen, err := strconv.Atoi(tokens[1])
		if err != nil {
			errs = append(errs, fmt.Errorf("must be /0 to /32, got: %s", value))
		} else if 0 > prefixLen || prefixLen > 32 {
			errs = append(errs, fmt.Errorf("must be /0 to /32, got: %s", value))
		}
	}
	return
}

func IsIpv6AddressPlen(val interface{}, key string) (warns []string, errs []error) {
	value := val.(string)
	if false == strings.Contains(value, ".") && strings.Contains(value, "/") {
		tokens := strings.Split(value, "/")
		prefixLen, err := strconv.Atoi(tokens[1])
		if err == nil && 0 <= prefixLen && prefixLen <= 128 && nil != net.ParseIP(tokens[0]) {
			return
		}
	}
	//this check doesn't work for IPv4-mapped format. But ACOS doesn't suppor it
	errs = append(errs, fmt.Errorf("must be A:B:C:D:E:F:G:H/nn, got: %s", value))
	return
}

func IsAbsolutePath(val interface{}, key string) (warns []string, errs []error) {
	value := val.(string)
	valueLen := len(value)
	if strings.HasPrefix(value, "/") && 1 <= valueLen && valueLen <= 127 {
		return
	}
	errs = append(errs, fmt.Errorf("must be absolute path and length 1 ~ 127, got: %s (len=%d)", value, valueLen))
	return
}

func getAxApiResponseMsg(resp []uint8) string {
	var axResp AxApiResponse
	stringData := string(resp)
	json.Unmarshal([]byte(stringData), &axResp)
	if axResp.Inst.Status == "fail" {
		failMsg := "axapi failure:" + axResp.Inst.Err.From + ":err=" + axResp.Inst.Err.Msg
		return failMsg
	}
	return ""
}

func Request(method string, host string, endpoint string, instId string, reqBytes []byte, headers map[string]string, logger *ThunderLog) (*http.Response, []byte, error) {
	URL := "https://" + host + "/axapi/v3/" + endpoint
	if instId != "" {
		URL = URL + "/" + instId
	}
	logger.Println(method, ">", URL)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 30 * time.Second}
	req, err := http.NewRequest(method, URL, bytes.NewReader(reqBytes))
	if err != nil {
		logger.Println("Failed to create HTTP request")
		return nil, nil, err
	}
	for property, value := range headers {
		req.Header.Set(property, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Println("HTTP request failed:", err)
		return nil, nil, err
	}
	var respBytes []byte
	if resp != nil {
		logger.Println("HTTP status", resp.Status)
		respBytes, _ = ioutil.ReadAll(resp.Body)
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			err = errors.New(getAxApiResponseMsg(respBytes))
		}
		logger.Println("axApi response:\n", string(respBytes))
	} else {
		logger.Println("Got no HTTP response")
		return nil, nil, errors.New("Got no HTTP response")
	}
	return resp, respBytes, err
}

func SendPost(host string, endpoint string, reqBytes []byte, headers map[string]string, logger *ThunderLog) (*http.Response, []byte, error) {
	return Request("POST", host, endpoint, "", reqBytes, headers, logger)
}

func SendGet(host string, endpoint string, instId string, reqBytes []byte, headers map[string]string, logger *ThunderLog) (*http.Response, []byte, error) {
	return Request("GET", host, endpoint, instId, reqBytes, headers, logger)
}

func SendPut(host string, endpoint string, instId string, reqBytes []byte, headers map[string]string, logger *ThunderLog) (*http.Response, []byte, error) {
	return Request("PUT", host, endpoint, instId, reqBytes, headers, logger)
}

func SendDelete(host string, endpoint string, instId string, reqBytes []byte, headers map[string]string, logger *ThunderLog) (*http.Response, []byte, error) {
	return Request("DELETE", host, endpoint, instId, reqBytes, headers, logger)
}

func GenUrlForDownload(protocol string, host string, path string, username string) string {
	url := protocol + "://"
	if len(username) > 0 {
		url += username + "@"
	}
	url += host
	if protocol == "scp" {
		url += ":"
	}
	url += path
	return url
}
