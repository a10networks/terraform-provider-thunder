package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerSNMPv1V2cUser struct {
	User SnmpServerSNMPv1V2cUserInstance `json:"user,omitempty"`
}

type SnmpServerSNMPv1V2cUserInstance struct {
	Encrypted string                       `json:"encrypted,omitempty"`
	OidVal    []SnmpServerSNMPv1V2cOidList `json:"oid-list,omitempty"`
	Passwd    string                       `json:"passwd,omitempty"`
	HostList  SnmpServerSNMPv1V2cRemote    `json:"remote,omitempty"`
	UUID      string                       `json:"uuid,omitempty"`
	User      string                       `json:"user,omitempty"`
	UserTag   string                       `json:"user-tag,omitempty"`
}

type SnmpServerSNMPv1V2cOidList struct {
	OidVal   string                    `json:"oid-val,omitempty"`
	HostList SnmpServerSNMPv1V2cRemote `json:"remote,omitempty"`
	UUID     string                    `json:"uuid,omitempty"`
	UserTag  string                    `json:"user-tag,omitempty"`
}

type SnmpServerSNMPv1V2cRemote struct {
	DNSHost  []SnmpServerSNMPv1V2cHostList `json:"host-list,omitempty"`
	Ipv4Host []SnmpServerSNMPv1V2cIpv4List `json:"ipv4-list,omitempty"`
	Ipv6Host []SnmpServerSNMPv1V2cIpv6List `json:"ipv6-list,omitempty"`
}

type SnmpServerSNMPv1V2cHostList struct {
	DNSHost  string `json:"dns-host,omitempty"`
	Ipv4Mask string `json:"ipv4-mask,omitempty"`
}

type SnmpServerSNMPv1V2cIpv4List struct {
	Ipv4Host string `json:"ipv4-host,omitempty"`
	Ipv4Mask string `json:"ipv4-mask,omitempty"`
}

type SnmpServerSNMPv1V2cIpv6List struct {
	Ipv6Host string `json:"ipv6-host,omitempty"`
	Ipv6Mask int    `json:"ipv6-mask,omitempty"`
}

func PostSnmpServerSNMPv1V2cUser(id string, inst SnmpServerSNMPv1V2cUser, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerSNMPv1V2cUser")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/SNMPv1-v2c/user", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerSNMPv1V2cUser
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("POSTSnmpServerSNMPv1V2cUser", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSnmpServerSNMPv1V2cUser(id string, name string, host string) (*SnmpServerSNMPv1V2cUser, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerSNMPv1V2cUser")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/SNMPv1-v2c/user/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerSNMPv1V2cUser
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetSnmpServerSNMPv1V2cUser", data)
			if err != nil {
				return nil, err
			}

			return &m, nil
		}
	}

}

func PutSnmpServerSNMPv1V2cUser(id string, name string, inst SnmpServerSNMPv1V2cUser, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSnmpServerSNMPv1V2cUser")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/snmp-server/SNMPv1-v2c/user/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerSNMPv1V2cUser
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PUT REQ RES..........................", m)
			err := check_api_status("PutSnmpServerSNMPv1V2cUser", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSnmpServerSNMPv1V2cUser(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSnmpServerSNMPv1V2cUser")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/snmp-server/SNMPv1-v2c/user/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerSNMPv1V2cUser
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] DELETE REQ RES..........................", m)

		}
	}
	return nil
}
