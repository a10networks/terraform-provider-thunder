package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type FwServiceGroup struct {
	Protocol FwServiceGroupInstance `json:"service-group-instance,omitempty"`
}

type FwServiceGroupInstance struct {
	HealthCheck string                    `json:"health-check,omitempty"`
	Port        []FwServiceMemberList     `json:"member-list,omitempty"`
	Name        string                    `json:"name,omitempty"`
	Protocol    string                    `json:"protocol,omitempty"`
	Counters1   []FwServiceSamplingEnable `json:"sampling-enable,omitempty"`
	UUID        string                    `json:"uuid,omitempty"`
	UserTag     string                    `json:"user-tag,omitempty"`
}

type FwServiceMemberList struct {
	Name      string                    `json:"name,omitempty"`
	Port      int                       `json:"port,omitempty"`
	Counters1 []FwServiceSamplingEnable `json:"sampling-enable,omitempty"`
	UUID      string                    `json:"uuid,omitempty"`
	UserTag   string                    `json:"user-tag,omitempty"`
}

type FwServiceSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostFwServiceGroup(id string, inst FwServiceGroup, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwServiceGroup")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/service-group", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServiceGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetFwServiceGroup(id string, name string, host string) (*FwServiceGroup, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwServiceGroup")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/service-group/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServiceGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}

}

func PutFwServiceGroup(id string, name string, inst FwServiceGroup, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutFwServiceGroup")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/fw/service-group/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServiceGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteFwServiceGroup(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteFwServiceGroup")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/fw/service-group/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServiceGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}
	return nil
}
