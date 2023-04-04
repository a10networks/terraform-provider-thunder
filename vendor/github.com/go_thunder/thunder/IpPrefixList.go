package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type IpPrefixList struct {
	IPPrefixListInstanceName IPPrefixListInstance `json:"prefix-list,omitempty"`
}

type IPPrefixListInstance struct {
	IPPrefixListInstanceName     string                      `json:"name,omitempty"`
	IPPrefixListInstanceRulesSeq []IPPrefixListInstanceRules `json:"rules,omitempty"`
	IPPrefixListInstanceUUID     string                      `json:"uuid,omitempty"`
}

type IPPrefixListInstanceRules struct {
	IPPrefixListInstanceRulesAction      string `json:"action,omitempty"`
	IPPrefixListInstanceRulesAny         int    `json:"any,omitempty"`
	IPPrefixListInstanceRulesDescription string `json:"description,omitempty"`
	IPPrefixListInstanceRulesGe          int    `json:"ge,omitempty"`
	IPPrefixListInstanceRulesIpaddr      string `json:"ipaddr,omitempty"`
	IPPrefixListInstanceRulesLe          int    `json:"le,omitempty"`
	IPPrefixListInstanceRulesSeq         int    `json:"seq,omitempty"`
}

func PostIpPrefixList(id string, inst IpPrefixList, host string) error {

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
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/prefix-list", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpPrefixList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostIpPrefixList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpPrefixList(id string, name1 string, host string) (*IpPrefixList, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpPrefixList")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/prefix-list/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpPrefixList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetIpPrefixList", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutIpPrefixList(id string, name1 string, inst IpPrefixList, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutIpPrefixList")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/ip/prefix-list/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpPrefixList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutIpPrefixList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteIpPrefixList(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteIpPrefixList")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/ip/prefix-list/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpPrefixList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteIpPrefixList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
