package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type AccessListStandard struct {
	Std AccessListStandardInstance `json:"standard,omitempty"`
}

type AccessListStandardInstance struct {
	Std    int                  `json:"std,omitempty"`
	SeqNum []AccessListStdrules `json:"stdrules,omitempty"`
	UUID   string               `json:"uuid,omitempty"`
}

type AccessListStdrules struct {
	Action                 string `json:"action,omitempty"`
	Any                    int    `json:"any,omitempty"`
	Host                   string `json:"host,omitempty"`
	Log                    int    `json:"log,omitempty"`
	RevSubnetMask          string `json:"rev-subnet-mask,omitempty"`
	SeqNum                 int    `json:"seq-num,omitempty"`
	StdRemark              string `json:"std-remark,omitempty"`
	Subnet                 string `json:"subnet,omitempty"`
	TransparentSessionOnly int    `json:"transparent-session-only,omitempty"`
}

func PostAccessListStandard(id string, inst AccessListStandard, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostAccessListStandard")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/access-list/standard", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m AccessListStandard
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostAccessListStandard", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetAccessListStandard(id string, name1 string, host string) (*AccessListStandard, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetAccessListStandard")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/access-list/standard/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m AccessListStandard
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetAccessListStandard", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutAccessListStandard(id string, name1 string, inst AccessListStandard, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutAccessListStandard")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/access-list/standard/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m AccessListStandard
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutAccessListStandard", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteAccessListStandard(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteAccessListStandard")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/access-list/standard/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m AccessListStandard
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteAccessListStandard", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
