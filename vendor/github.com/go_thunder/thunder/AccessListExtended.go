package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type AccessListExtended struct {
	Extd AccessListExtendedInstance `json:"extended,omitempty"`
}

type AccessListExtendedInstance struct {
	Extd       int               `json:"extd,omitempty"`
	ExtdSeqNum []AccessListRules `json:"rules,omitempty"`
	UUID       string            `json:"uuid,omitempty"`
}

type AccessListRules struct {
	AclLog                 int    `json:"acl-log,omitempty"`
	AnyCode                int    `json:"any-code,omitempty"`
	AnyType                int    `json:"any-type,omitempty"`
	Dscp                   int    `json:"dscp,omitempty"`
	DstAny                 int    `json:"dst-any,omitempty"`
	DstEq                  int    `json:"dst-eq,omitempty"`
	DstGt                  int    `json:"dst-gt,omitempty"`
	DstHost                string `json:"dst-host,omitempty"`
	DstLt                  int    `json:"dst-lt,omitempty"`
	DstMask                string `json:"dst-mask,omitempty"`
	DstObjectGroup         string `json:"dst-object-group,omitempty"`
	DstPortEnd             int    `json:"dst-port-end,omitempty"`
	DstRange               int    `json:"dst-range,omitempty"`
	DstSubnet              string `json:"dst-subnet,omitempty"`
	Established            int    `json:"established,omitempty"`
	Ethernet               int    `json:"ethernet,omitempty"`
	ExtdAction             string `json:"extd-action,omitempty"`
	ExtdRemark             string `json:"extd-remark,omitempty"`
	ExtdSeqNum             int    `json:"extd-seq-num,omitempty"`
	Fragments              int    `json:"fragments,omitempty"`
	IP                     int    `json:"ip,omitempty"`
	Icmp                   int    `json:"icmp,omitempty"`
	IcmpCode               int    `json:"icmp-code,omitempty"`
	IcmpType               int    `json:"icmp-type,omitempty"`
	ServiceObjGroup        string `json:"service-obj-group,omitempty"`
	SpecialCode            string `json:"special-code,omitempty"`
	SpecialType            string `json:"special-type,omitempty"`
	SrcAny                 int    `json:"src-any,omitempty"`
	SrcEq                  int    `json:"src-eq,omitempty"`
	SrcGt                  int    `json:"src-gt,omitempty"`
	SrcHost                string `json:"src-host,omitempty"`
	SrcLt                  int    `json:"src-lt,omitempty"`
	SrcMask                string `json:"src-mask,omitempty"`
	SrcObjectGroup         string `json:"src-object-group,omitempty"`
	SrcPortEnd             int    `json:"src-port-end,omitempty"`
	SrcRange               int    `json:"src-range,omitempty"`
	SrcSubnet              string `json:"src-subnet,omitempty"`
	TCP                    int    `json:"tcp,omitempty"`
	TransparentSessionOnly int    `json:"transparent-session-only,omitempty"`
	Trunk                  int    `json:"trunk,omitempty"`
	UDP                    int    `json:"udp,omitempty"`
	Vlan                   int    `json:"vlan,omitempty"`
}

func PostAccessListExtended(id string, inst AccessListExtended, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostAccessListExtended")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/access-list/extended", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m AccessListExtended
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostAccessListExtended", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetAccessListExtended(id string, name1 string, host string) (*AccessListExtended, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetAccessListExtended")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/access-list/extended/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m AccessListExtended
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetAccessListExtended", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutAccessListExtended(id string, name1 string, inst AccessListExtended, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutAccessListExtended")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/access-list/extended/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m AccessListExtended
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutAccessListExtended", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteAccessListExtended(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteAccessListExtended")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/access-list/extended/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m AccessListExtended
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteAccessListExtended", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
