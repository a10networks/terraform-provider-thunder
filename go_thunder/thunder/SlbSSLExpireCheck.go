package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SSLExpireCheck struct {
	UUID SSLExpireCheckInstance `json:"ssl-expire-check,omitempty"`
}

type Exception struct {
	Action          string `json:"action,omitempty"`
	CertificateName string `json:"certificate-name,omitempty"`
}

type SSLExpireCheckInstance struct {
	Action                Exception `json:"exception,omitempty"`
	UUID                  string    `json:"uuid,omitempty"`
	SslExpireEmailAddress string    `json:"ssl-expire-email-address,omitempty"`
	IntervalDays          int       `json:"interval-days,omitempty"`
	ExpireAddress1        string    `json:"expire-address1,omitempty"`
	Before                int       `json:"before,omitempty"`
}

func PostSlbSSLExpireCheck(id string, inst SSLExpireCheck, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbSSLExpireCheck")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/ssl-expire-check", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SSLExpireCheck
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbSSLExpireCheck REQ RES..........................", m)
			err := check_api_status("PostSlbSSLExpireCheck", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbSSLExpireCheck(id string, host string) (*SSLExpireCheck, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbSSLExpireCheck")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/ssl-expire-check/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SSLExpireCheck
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbSSLExpireCheck REQ RES..........................", m)
			err := check_api_status("GetSlbSSLExpireCheck", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
