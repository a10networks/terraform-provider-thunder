package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Imapproxy struct {
	Counters1 ImapproxyInstance `json:"imapproxy,omitempty"`
}
type SamplingEnable41 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type ImapproxyInstance struct {
	Counters1 []SamplingEnable41 `json:"sampling-enable,omitempty"`
}

func PostSlbImapproxy(id string, inst Imapproxy, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbImap-proxy")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/imap-proxy", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Imapproxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbImapproxy REQ RES..........................", m)
			err := check_api_status("PostSlbImapproxy", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func GetSlbImapproxy(id string, host string) (*Imapproxy, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbImap-proxy")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/imap-proxy/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Imapproxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbImapproxy REQ RES..........................", m)
			err := check_api_status("GetSlbImapproxy", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
