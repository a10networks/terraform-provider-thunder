package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type PrefixList struct {
	UUID PrefixListInstance `json:"prefix-list,omitempty"`
}

type Rules struct {
	Le          int    `json:"le,omitempty"`
	Description string `json:"description,omitempty"`
	Seq         int    `json:"seq,omitempty"`
	Ipaddr      string `json:"ipaddr,omitempty"`
	Ge          int    `json:"ge,omitempty"`
	Action      string `json:"action,omitempty"`
	Any         int    `json:"any,omitempty"`
}
type PrefixListInstance struct {
	Le   []Rules `json:"rules,omitempty"`
	Name string  `json:"name,omitempty"`
	UUID string  `json:"uuid,omitempty"`
}

func PostIpPrefixList(id string, inst PrefixList, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpPrefixList")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/prefix-list", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m PrefixList
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpPrefixList REQ RES..........................", m)
			err := check_api_status("PostIpPrefixList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpPrefixList(id string, name string, host string) (*PrefixList, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpPrefixList")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/prefix-list/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m PrefixList
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIpPrefixList REQ RES..........................", m)
			err := check_api_status("GetIpPrefixList", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
