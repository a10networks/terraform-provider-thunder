package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type IpAccessList struct {
	IPAccessListInstanceName IPAccessListInstance `json:"access-list,omitempty"`
}

type IPAccessListInstance struct {
	IPAccessListInstanceName        string                      `json:"name,omitempty"`
	IPAccessListInstanceRulesSeqNum []IPAccessListInstanceRules `json:"rules,omitempty"`
	IPAccessListInstanceUUID        string                      `json:"uuid,omitempty"`
	IPAccessListInstanceUserTag     string                      `json:"user-tag,omitempty"`
}

type IPAccessListInstanceRules struct {
	IPAccessListInstanceRulesAclLog                 int    `json:"acl-log,omitempty"`
	IPAccessListInstanceRulesAction                 string `json:"action,omitempty"`
	IPAccessListInstanceRulesAnyCode                int    `json:"any-code,omitempty"`
	IPAccessListInstanceRulesAnyType                int    `json:"any-type,omitempty"`
	IPAccessListInstanceRulesDscp                   int    `json:"dscp,omitempty"`
	IPAccessListInstanceRulesDstAny                 int    `json:"dst-any,omitempty"`
	IPAccessListInstanceRulesDstEq                  int    `json:"dst-eq,omitempty"`
	IPAccessListInstanceRulesDstGt                  int    `json:"dst-gt,omitempty"`
	IPAccessListInstanceRulesDstHost                string `json:"dst-host,omitempty"`
	IPAccessListInstanceRulesDstLt                  int    `json:"dst-lt,omitempty"`
	IPAccessListInstanceRulesDstMask                string `json:"dst-mask,omitempty"`
	IPAccessListInstanceRulesDstObjectGroup         string `json:"dst-object-group,omitempty"`
	IPAccessListInstanceRulesDstPortEnd             int    `json:"dst-port-end,omitempty"`
	IPAccessListInstanceRulesDstRange               int    `json:"dst-range,omitempty"`
	IPAccessListInstanceRulesDstSubnet              string `json:"dst-subnet,omitempty"`
	IPAccessListInstanceRulesEstablished            int    `json:"established,omitempty"`
	IPAccessListInstanceRulesEthernet               int    `json:"ethernet,omitempty"`
	IPAccessListInstanceRulesFragments              int    `json:"fragments,omitempty"`
	IPAccessListInstanceRulesGeoLocation            string `json:"geo-location,omitempty"`
	IPAccessListInstanceRulesIP                     int    `json:"ip,omitempty"`
	IPAccessListInstanceRulesIcmp                   int    `json:"icmp,omitempty"`
	IPAccessListInstanceRulesIcmpCode               int    `json:"icmp-code,omitempty"`
	IPAccessListInstanceRulesIcmpType               int    `json:"icmp-type,omitempty"`
	IPAccessListInstanceRulesRemark                 string `json:"remark,omitempty"`
	IPAccessListInstanceRulesSeqNum                 int    `json:"seq-num,omitempty"`
	IPAccessListInstanceRulesServiceObjGroup        string `json:"service-obj-group,omitempty"`
	IPAccessListInstanceRulesSpecialCode            string `json:"special-code,omitempty"`
	IPAccessListInstanceRulesSpecialType            string `json:"special-type,omitempty"`
	IPAccessListInstanceRulesSrcAny                 int    `json:"src-any,omitempty"`
	IPAccessListInstanceRulesSrcEq                  int    `json:"src-eq,omitempty"`
	IPAccessListInstanceRulesSrcGt                  int    `json:"src-gt,omitempty"`
	IPAccessListInstanceRulesSrcHost                string `json:"src-host,omitempty"`
	IPAccessListInstanceRulesSrcLt                  int    `json:"src-lt,omitempty"`
	IPAccessListInstanceRulesSrcMask                string `json:"src-mask,omitempty"`
	IPAccessListInstanceRulesSrcObjectGroup         string `json:"src-object-group,omitempty"`
	IPAccessListInstanceRulesSrcPortEnd             int    `json:"src-port-end,omitempty"`
	IPAccessListInstanceRulesSrcRange               int    `json:"src-range,omitempty"`
	IPAccessListInstanceRulesSrcSubnet              string `json:"src-subnet,omitempty"`
	IPAccessListInstanceRulesTCP                    int    `json:"tcp,omitempty"`
	IPAccessListInstanceRulesTransparentSessionOnly int    `json:"transparent-session-only,omitempty"`
	IPAccessListInstanceRulesTrunk                  int    `json:"trunk,omitempty"`
	IPAccessListInstanceRulesUDP                    int    `json:"udp,omitempty"`
	IPAccessListInstanceRulesVlan                   int    `json:"vlan,omitempty"`
}

func PostIpAccessList(id string, inst IpAccessList, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpAccessList")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/access-list", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpAccessList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostIpAccessList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpAccessList(id string, name1 string, host string) (*IpAccessList, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpAccessList")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/access-list/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpAccessList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetIpAccessList", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutIpAccessList(id string, name1 string, inst IpAccessList, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutIpAccessList")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/ip/access-list/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpAccessList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutIpAccessList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteIpAccessList(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteIpAccessList")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/ip/access-list/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpAccessList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteIpAccessList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
