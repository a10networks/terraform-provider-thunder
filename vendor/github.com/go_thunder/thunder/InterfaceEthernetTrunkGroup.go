package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"strconv"
	"util"
)

type EthernetTrunkGroup struct {
	UUID EthernetTrunkGroupInstance `json:"trunk-group,omitempty"`
}

type UdldTimeoutCfg struct {
	Slow int `json:"slow,omitempty"`
	Fast int `json:"fast,omitempty"`
}
type EthernetTrunkGroupInstance struct {
	UUID         string         `json:"uuid,omitempty"`
	TrunkNumber  int            `json:"trunk-number,omitempty"`
	UserTag      string         `json:"user-tag,omitempty"`
	Slow         UdldTimeoutCfg `json:"udld-timeout-cfg,omitempty"`
	Mode         string         `json:"mode,omitempty"`
	Timeout      string         `json:"timeout,omitempty"`
	Type         string         `json:"type,omitempty"`
	AdminKey     int            `json:"admin-key,omitempty"`
	PortPriority int            `json:"port-priority,omitempty"`
}

func PostInterfaceEthernetTrunkGroup(id string, idNum int, inst EthernetTrunkGroup, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceEthernetTrunkGroup")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ethernet/"+strconv.Itoa(idNum)+"/trunk-group", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m EthernetTrunkGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostInterfaceEthernetTrunkGroup REQ RES..........................", m)
			err := check_api_status("PostInterfaceEthernetTrunkGroup", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceEthernetTrunkGroup(id string, idNum string, name string, host string) (*EthernetTrunkGroup, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceEthernetTrunkGroup")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ethernet/"+idNum+"/trunk-group/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m EthernetTrunkGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetInterfaceEthernetTrunkGroup REQ RES..........................", m)
			err := check_api_status("GetInterfaceEthernetTrunkGroup", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutInterfaceEthernetTrunkGroup(id string, idNum int, name int, inst EthernetTrunkGroup, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutInterfaceEthernetTrunkGroup")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/interface/ethernet/"+strconv.Itoa(idNum)+"/trunk-group/"+strconv.Itoa(name), bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m EthernetTrunkGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("PutInterfaceEthernetTrunkGroup", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteInterfaceEthernetTrunkGroup(id string, idNum string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteInterfaceEthernetTrunkGroup")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/interface/ethernet/"+idNum+"/trunk-group/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m EthernetTrunkGroup
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
