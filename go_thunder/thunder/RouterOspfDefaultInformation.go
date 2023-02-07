package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterOspfDefaultInformation struct {
	Always RouterOspfDefaultInformationInstance `json:"default-information,omitempty"`
}

type RouterOspfDefaultInformationInstance struct {
	Always     int    `json:"always,omitempty"`
	Metric     int    `json:"metric,omitempty"`
	MetricType int    `json:"metric-type,omitempty"`
	Originate  int    `json:"originate,omitempty"`
	RouteMap   string `json:"route-map,omitempty"`
	UUID       string `json:"uuid,omitempty"`
}

func PostRouterOspfDefaultInformation(id string, name1 string, inst RouterOspfDefaultInformation, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterOspfDefaultInformation")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/ospf/"+name1+"/default-information", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterOspfDefaultInformation
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterOspfDefaultInformation", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterOspfDefaultInformation(id string, name1 string, host string) (*RouterOspfDefaultInformation, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterOspfDefaultInformation")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/ospf/"+name1+"/default-information", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterOspfDefaultInformation
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterOspfDefaultInformation", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
