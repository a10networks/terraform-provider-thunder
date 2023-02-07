package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerSNMPv1V2cUserOid struct {
	OidVal SnmpServerSNMPv1V2cUserOidInstance `json:"oid,omitempty"`
}

type SnmpServerSNMPv1V2cUserOidInstance struct {
	OidVal  string                        `json:"oid-val,omitempty"`
	DNSHost SnmpServerSNMPv1V2cUserRemote `json:"remote,omitempty"`
	UUID    string                        `json:"uuid,omitempty"`
	UserTag string                        `json:"user-tag,omitempty"`
}

type SnmpServerSNMPv1V2cUserRemote struct {
	DNSHost  []SnmpServerSNMPv1V2cUserHostList `json:"host-list,omitempty"`
	Ipv4Host []SnmpServerSNMPv1V2cUserIpv4List `json:"ipv4-list,omitempty"`
	Ipv6Host []SnmpServerSNMPv1V2cUserIpv6List `json:"ipv6-list,omitempty"`
}

type SnmpServerSNMPv1V2cUserHostList struct {
	DNSHost  string `json:"dns-host,omitempty"`
	Ipv4Mask string `json:"ipv4-mask,omitempty"`
}

type SnmpServerSNMPv1V2cUserIpv4List struct {
	Ipv4Host string `json:"ipv4-host,omitempty"`
	Ipv4Mask string `json:"ipv4-mask,omitempty"`
}

type SnmpServerSNMPv1V2cUserIpv6List struct {
	Ipv6Host string `json:"ipv6-host,omitempty"`
	Ipv6Mask int    `json:"ipv6-mask,omitempty"`
}

func PostSnmpServerSNMPv1V2cUserOid(id string, name1 string, inst SnmpServerSNMPv1V2cUserOid, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerSNMPv1V2cUserOid")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/SNMPv1-v2c/user/"+name1+"/oid", bytes.NewReader(payloadBytes), headers)
	logger.Println("POST-------->   https://" + host + "/axapi/v3/snmp-server/SNMPv1-v2c/user/" + name1 + "/oid/")
	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerSNMPv1V2cUserOid
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerSNMPv1V2cUserOid", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func GetSnmpServerSNMPv1V2cUserOid(id string, name1 string, name2 string, host string) (*SnmpServerSNMPv1V2cUserOid, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerSNMPv1V2cUserOid")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/SNMPv1-v2c/user/"+name1+"/oid/"+name2, nil, headers)
	logger.Println("GET -------->   https://" + host + "/axapi/v3/snmp-server/SNMPv1-v2c/user/" + name1 + "/oid/" + name2)
	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerSNMPv1V2cUserOid
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerSNMPv1V2cUserOid", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSnmpServerSNMPv1V2cUserOid(id string, name1 string, name2 string, inst SnmpServerSNMPv1V2cUserOid, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSnmpServerSNMPv1V2cUserOid")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/snmp-server/SNMPv1-v2c/user/"+name1+"/oid/"+name2, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerSNMPv1V2cUserOid
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSnmpServerSNMPv1V2cUserOid", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func DeleteSnmpServerSNMPv1V2cUserOid(id string, name1 string, name2 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSnmpServerSNMPv1V2cUserOid")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/snmp-server/SNMPv1-v2c/user/"+name1+"/oid/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerSNMPv1V2cUserOid
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
