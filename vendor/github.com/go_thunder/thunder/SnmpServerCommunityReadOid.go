package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerCommunityReadOid struct {
	Remote SnmpServerCommunityReadOidInstance `json:"oid,omitempty"`
}

type SnmpServerCommunityReadOidInstance struct {
	OidVal   string                        `json:"oid-val,omitempty"`
	HostList SnmpServerCommunityReadRemote `json:"remote,omitempty"`
	UUID     string                        `json:"uuid,omitempty"`
	UserTag  string                        `json:"user-tag,omitempty"`
}

type SnmpServerCommunityReadRemote struct {
	DNSHost  []SnmpServerCommunityReadHostList `json:"host-list,omitempty"`
	Ipv4Host []SnmpServerCommunityReadIpv4List `json:"ipv4-list,omitempty"`
	Ipv6Host []SnmpServerCommunityReadIpv6List `json:"ipv6-list,omitempty"`
}

type SnmpServerCommunityReadHostList struct {
	DNSHost  string `json:"dns-host,omitempty"`
	Ipv4Mask string `json:"ipv4-mask,omitempty"`
}

type SnmpServerCommunityReadIpv4List struct {
	Ipv4Host string `json:"ipv4-host,omitempty"`
	Ipv4Mask string `json:"ipv4-mask,omitempty"`
}

type SnmpServerCommunityReadIpv6List struct {
	Ipv6Host string `json:"ipv6-host,omitempty"`
	Ipv6Mask int    `json:"ipv6-mask,omitempty"`
}

func PostSnmpServerCommunityReadOid(id string, name1 string, inst SnmpServerCommunityReadOid, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerCommunityReadOid")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/community/read/"+name1+"/oid", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerCommunityReadOid
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerCommunityReadOid", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func GetSnmpServerCommunityReadOid(id string, name1 string, name2 string, host string) (*SnmpServerCommunityReadOid, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerCommunityReadOid")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/community/read/"+name1+"/oid/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerCommunityReadOid
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerCommunityReadOid", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSnmpServerCommunityReadOid(id string, name1 string, name2 string, inst SnmpServerCommunityReadOid, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSnmpServerCommunityReadOid")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/snmp-server/community/read/"+name1+"/oid/"+name2, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerCommunityReadOid
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSnmpServerCommunityReadOid", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func DeleteSnmpServerCommunityReadOid(id string, name1 string, name2 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSnmpServerCommunityReadOid")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/snmp-server/community/read/"+name1+"/oid/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerCommunityReadOid
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
