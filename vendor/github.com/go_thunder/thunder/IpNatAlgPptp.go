package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type NatAlgPptp struct {
	UUID NatAlgPptpInstance `json:"pptp,omitempty"`
}
type SamplingEnable48 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type NatAlgPptpInstance struct {
	Pptp      string             `json:"pptp,omitempty"`
	Counters1 []SamplingEnable48 `json:"sampling-enable,omitempty"`
	UUID      string             `json:"uuid,omitempty"`
}

func PostIpNatAlgPptp(id string, inst NatAlgPptp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpNatAlgPptp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/nat/alg/pptp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NatAlgPptp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpNatAlgPptp REQ RES..........................", m)
			err := check_api_status("PostIpNatAlgPptp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpNatAlgPptp(id string, host string) (*NatAlgPptp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpNatAlgPptp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/nat/alg/pptp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NatAlgPptp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetIpNatAlgPptp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
