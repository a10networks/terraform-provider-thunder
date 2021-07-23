package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerCommunityRead struct {
	UUID SnmpServerCommunityReadInstance `json:"read,omitempty"`
}

type SnmpServerCommunityReadInstance struct {
	OidVal   []SnmpServerCommunityOidList `json:"oid-list,omitempty"`
	HostList SnmpServerCommunityRemote    `json:"remote,omitempty"`
	UUID     string                       `json:"uuid,omitempty"`
	User     string                       `json:"user,omitempty"`
	UserTag  string                       `json:"user-tag,omitempty"`
}

type SnmpServerCommunityOidList struct {
	OidVal   string                    `json:"oid-val,omitempty"`
	HostList SnmpServerCommunityRemote `json:"remote,omitempty"`
	UUID     string                    `json:"uuid,omitempty"`
	UserTag  string                    `json:"user-tag,omitempty"`
}

type SnmpServerCommunityRemote struct {
	DNSHost  []SnmpServerCommunityHostList `json:"host-list,omitempty"`
	Ipv4Host []SnmpServerCommunityIpv4List `json:"ipv4-list,omitempty"`
	Ipv6Host []SnmpServerCommunityIpv6List `json:"ipv6-list,omitempty"`
}

type SnmpServerCommunityHostList struct {
	DNSHost  string `json:"dns-host,omitempty"`
	Ipv4Mask string `json:"ipv4-mask,omitempty"`
}

type SnmpServerCommunityIpv4List struct {
	Ipv4Host string `json:"ipv4-host,omitempty"`
	Ipv4Mask string `json:"ipv4-mask,omitempty"`
}

type SnmpServerCommunityIpv6List struct {
	Ipv6Host string `json:"ipv6-host,omitempty"`
	Ipv6Mask int    `json:"ipv6-mask,omitempty"`
}

func PostSnmpServerCommunityRead(id string, inst SnmpServerCommunityRead, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerCommunityRead")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/community/read", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerCommunityRead
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerCommunityRead", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func GetSnmpServerCommunityRead(id string, name string, host string) (*SnmpServerCommunityRead, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerCommunityRead")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/community/read/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerCommunityRead
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerCommunityRead", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSnmpServerCommunityRead(id string, name string, inst SnmpServerCommunityRead, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSnmpServerCommunityRead")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/snmp-server/community/read/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerCommunityRead
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSnmpServerCommunityRead", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func DeleteSnmpServerCommunityRead(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSnmpServerCommunityRead")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/snmp-server/community/read/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerCommunityRead
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
