package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type PlayerIdGlobal struct {
	Counters1 PlayerIdGlobalInstance `json:"player-id-global,omitempty"`
}
type SamplingEnable27 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type PlayerIdGlobalInstance struct {
	MinExpiration         int                `json:"min-expiration,omitempty"`
	PktActivityExpiration int                `json:"pkt-activity-expiration,omitempty"`
	ForcePassive          int                `json:"force-passive,omitempty"`
	AbsMaxExpiration      int                `json:"abs-max-expiration,omitempty"`
	Counters1             []SamplingEnable27 `json:"sampling-enable,omitempty"`
	EnforcementTimer      int                `json:"enforcement-timer,omitempty"`
	Enable64BitPlayerID   int                `json:"enable-64bit-player-id,omitempty"`
}

func PostSlbPlayerIdGlobal(id string, inst PlayerIdGlobal, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbPlayerIdGlobal")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/player-id-global", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m PlayerIdGlobal
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbPlayerIdGlobal REQ RES..........................", m)
			err := check_api_status("PostSlbPlayerIdGlobal", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbPlayerIdGlobal(id string, host string) (*PlayerIdGlobal, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbPlayerIdGlobal")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/player-id-global/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m PlayerIdGlobal
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbPlayerIdGlobal REQ RES..........................", m)
			err := check_api_status("GetSlbPlayerIdGlobal", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
