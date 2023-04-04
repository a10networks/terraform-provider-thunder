package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwAlgRtsp struct {
	DefaultPortDisable FwAlgRtspInstance `json:"rtsp,omitempty"`
}

type FwAlgRtspInstance struct {
	DefaultPortDisable string                    `json:"default-port-disable,omitempty"`
	Counters1          []FwAlgRtspSamplingEnable `json:"sampling-enable,omitempty"`
	UUID               string                    `json:"uuid,omitempty"`
}

type FwAlgRtspSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostFwAlgRtsp(id string, inst FwAlgRtsp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwAlgRtsp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/alg/rtsp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwAlgRtsp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("PostFwAlgRtsp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwAlgRtsp(id string, host string) (*FwAlgRtsp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwAlgRtsp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/alg/rtsp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwAlgRtsp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetFwAlgRtsp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
