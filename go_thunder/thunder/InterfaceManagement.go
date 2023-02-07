package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Management struct {
	UUID ManagementInstance `json:"management,omitempty"`
}

type TxDot1Cfg struct {
	LinkAggregation int `json:"link-aggregation,omitempty"`
	Vlan            int `json:"vlan,omitempty"`
	TxDot1Tlvs      int `json:"tx-dot1-tlvs,omitempty"`
}
type NotificationCfg struct {
	Notification int `json:"notification,omitempty"`
	NotifEnable  int `json:"notif-enable,omitempty"`
}
type EnableCfg struct {
	Rx       int `json:"rx,omitempty"`
	Tx       int `json:"tx,omitempty"`
	RtEnable int `json:"rt-enable,omitempty"`
}
type TxTlvsCfg struct {
	SystemCapabilities int `json:"system-capabilities,omitempty"`
	SystemDescription  int `json:"system-description,omitempty"`
	ManagementAddress  int `json:"management-address,omitempty"`
	TxTlvs             int `json:"tx-tlvs,omitempty"`
	Exclude            int `json:"exclude,omitempty"`
	PortDescription    int `json:"port-description,omitempty"`
	SystemName         int `json:"system-name,omitempty"`
}
type Lldp struct {
	LinkAggregation    TxDot1Cfg       `json:"tx-dot1-cfg,omitempty"`
	Notification       NotificationCfg `json:"notification-cfg,omitempty"`
	Rx                 EnableCfg       `json:"enable-cfg,omitempty"`
	SystemCapabilities TxTlvsCfg       `json:"tx-tlvs-cfg,omitempty"`
	UUID               string          `json:"uuid,omitempty"`
}
type BroadcastRateLimit struct {
	Rate                 int `json:"rate,omitempty"`
	BcastRateLimitEnable int `json:"bcast-rate-limit-enable,omitempty"`
}
type ManagementIP struct {
	Dhcp                   int    `json:"dhcp,omitempty"`
	Ipv4Address            string `json:"ipv4-address,omitempty"`
	ControlAppsUseMgmtPort int    `json:"control-apps-use-mgmt-port,omitempty"`
	DefaultGateway         string `json:"default-gateway,omitempty"`
	Ipv4Netmask            string `json:"ipv4-netmask,omitempty"`
}
type SecondaryIP struct {
	Ipv4Netmask            string `json:"ipv4-netmask,omitempty"`
	ControlAppsUseMgmtPort int    `json:"control-apps-use-mgmt-port,omitempty"`
	SecondaryIP            int    `json:"secondary-ip,omitempty"`
	DefaultGateway         string `json:"default-gateway,omitempty"`
	Dhcp                   int    `json:"dhcp,omitempty"`
	Ipv4Address            string `json:"ipv4-address,omitempty"`
}
type AccessList struct {
	ACLName string `json:"acl-name,omitempty"`
	ACLID   int    `json:"acl-id,omitempty"`
}
type ManagementSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}
type Ipv6 struct {
	DefaultIpv6Gateway string `json:"default-ipv6-gateway,omitempty"`
	Inbound            int    `json:"inbound,omitempty"`
	AddressType        string `json:"address-type,omitempty"`
	Ipv6Addr           string `json:"ipv6-addr,omitempty"`
	V6ACLName          string `json:"v6-acl-name,omitempty"`
}
type ManagementInstance struct {
	LinkAggregation    Lldp                       `json:"lldp,omitempty"`
	FlowControl        int                        `json:"flow-control,omitempty"`
	Rate               BroadcastRateLimit         `json:"broadcast-rate-limit,omitempty"`
	UUID               string                     `json:"uuid,omitempty"`
	Duplexity          string                     `json:"duplexity,omitempty"`
	Dhcp               ManagementIP               `json:"ip,omitempty"`
	Ipv4Netmask        SecondaryIP                `json:"secondary-ip,omitempty"`
	ACLName            AccessList                 `json:"access-list,omitempty"`
	Counters1          []ManagementSamplingEnable `json:"sampling-enable,omitempty"`
	DefaultIpv6Gateway []Ipv6                     `json:"ipv6,omitempty"`
	Action             string                     `json:"action,omitempty"`
	Speed              string                     `json:"speed,omitempty"`
}

func PostInterfaceManagement(id string, inst Management, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceManagement")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/management", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Management
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostInterfaceManagement REQ RES..........................", m)
			err := check_api_status("PostInterfaceManagement", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceManagement(id string, host string) (*Management, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceManagement")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/management/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Management
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetInterfaceManagement REQ RES..........................", m)
			err := check_api_status("GetInterfaceManagement", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
