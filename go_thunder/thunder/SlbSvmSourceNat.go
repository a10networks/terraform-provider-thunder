package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SvmSourceNat struct {
	Pool SvmSourceNatInstance `json:"svm-source-nat,omitempty"`
}
type SvmSourceNatInstance struct {
	Pool string `json:"pool,omitempty"`
}

func PostSlbSvmSourceNat(id string, inst SvmSourceNat, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbSvmSourceNat")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/svm-source-nat", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SvmSourceNat
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbSvmSourceNat REQ RES..........................", m)
			err := check_api_status("PostSlbSvmSourceNat", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbSvmSourceNat(id string, host string) (*SvmSourceNat, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbSvmSourceNat")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/svm-source-nat/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SvmSourceNat
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbSvmSourceNat REQ RES..........................", m)
			err := check_api_status("GetSlbSvmSourceNat", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
