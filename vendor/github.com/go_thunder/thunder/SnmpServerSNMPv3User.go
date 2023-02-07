package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerSNMPv3User struct {
	Username SnmpServerSNMPv3UserInstance `json:"user,omitempty"`
}

type SnmpServerSNMPv3UserInstance struct {
	AuthVal         string `json:"auth-val,omitempty"`
	Encpasswd       string `json:"encpasswd,omitempty"`
	Group           string `json:"group,omitempty"`
	Passwd          string `json:"passwd,omitempty"`
	Priv            string `json:"priv,omitempty"`
	PrivPwEncrypted string `json:"priv-pw-encrypted,omitempty"`
	PwEncrypted     string `json:"pw-encrypted,omitempty"`
	UUID            string `json:"uuid,omitempty"`
	Username        string `json:"username,omitempty"`
	V3              string `json:"v3,omitempty"`
}

func PostSnmpServerSNMPv3User(id string, inst SnmpServerSNMPv3User, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerSNMPv3User")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/SNMPv3/user", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerSNMPv3User
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerSNMPv3User", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSnmpServerSNMPv3User(id string, name string, host string) (*SnmpServerSNMPv3User, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerSNMPv3User")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/SNMPv3/user/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerSNMPv3User
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerSNMPv3User", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSnmpServerSNMPv3User(id string, name string, inst SnmpServerSNMPv3User, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSnmpServerSNMPv3User")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/snmp-server/SNMPv3/user/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerSNMPv3User
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSnmpServerSNMPv3User", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSnmpServerSNMPv3User(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSnmpServerSNMPv3User")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/snmp-server/SNMPv3/user/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerSNMPv3User
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)

		}
	}
	return nil
}
