package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Partition struct {
	PartitionName PartitionInstance `json:"partition,omitempty"`
}

type PartitionInstance struct {
	ApplicationType    string     `json:"application-type,omitempty"`
	ID                 int        `json:"id,omitempty"`
	PartitionName      string     `json:"partition-name,omitempty"`
	Vlan               SharedVlan `json:"shared-vlan,omitempty"`
	ResourceAccounting Template   `json:"template,omitempty"`
	UUID               string     `json:"uuid,omitempty"`
	UserTag            string     `json:"user-tag,omitempty"`
}

type SharedVlan struct {
	MgmtFloatingIPAddress string `json:"mgmt-floating-ip-address,omitempty"`
	UUID                  string `json:"uuid,omitempty"`
	Vlan                  int    `json:"vlan,omitempty"`
	Vrid                  int    `json:"vrid,omitempty"`
}

type Template struct {
	ResourceAccounting string `json:"resource-accounting,omitempty"`
	UUID               string `json:"uuid,omitempty"`
}

func PostPartition(id string, inst Partition, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostPartition")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/partition", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Partition
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostPartition", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetPartition(id string, name1 string, host string) (*Partition, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetPartition")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/partition/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Partition
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetPartition", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func DeletePartition(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeletePartition")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/partition/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Partition
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeletePartition", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
