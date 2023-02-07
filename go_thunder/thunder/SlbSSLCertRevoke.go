package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SSLCertRevoke struct {
	Counters1 SSLCertRevokeInstance `json:"ssl-cert-revoke,omitempty"`
}

type SamplingEnable37 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type SSLCertRevokeInstance struct {
	Counters1 []SamplingEnable37 `json:"sampling-enable,omitempty"`
}

func PostSlbSSLCertRevoke(id string, inst SSLCertRevoke, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbSSLCertRevoke")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/ssl-cert-revoke", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SSLCertRevoke
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbSSLCertRevoke REQ RES..........................", m)
			err := check_api_status("PostSlbSSLCertRevoke", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbSSLCertRevoke(id string, host string) (*SSLCertRevoke, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbSSLCertRevoke")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/ssl-cert-revoke/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SSLCertRevoke
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbSSLCertRevoke REQ RES..........................", m)
			err := check_api_status("GetSlbSSLCertRevoke", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
