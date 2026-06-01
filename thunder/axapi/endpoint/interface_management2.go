package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type InterfaceManagement2 struct {
	Inst struct {
		AccessList InterfaceManagement2AccessList `json:"access-list"`

		Action string `json:"action" dval:"enable"`

		BroadcastRateLimit InterfaceManagement2BroadcastRateLimit `json:"broadcast-rate-limit"`

		Duplexity string `json:"duplexity" dval:"auto"`

		FlowControl int `json:"flow-control"`

		Ip InterfaceManagement2Ip `json:"ip"`

		Ipv6 []InterfaceManagement2Ipv6 `json:"ipv6"`

		Mtu int `json:"mtu"`

		SamplingEnable []InterfaceManagement2SamplingEnable `json:"sampling-enable"`

		Speed string `json:"speed" dval:"auto"`

		Uuid string `json:"uuid"`
	} `json:"management2"`
}

type InterfaceManagement2AccessList struct {
	AclId   int    `json:"acl-id"`
	AclName string `json:"acl-name"`
}

type InterfaceManagement2BroadcastRateLimit struct {
	BcastRateLimitEnable int `json:"bcast-rate-limit-enable"`
	Rate                 int `json:"rate" dval:"500"`
}

type InterfaceManagement2Ip struct {
	Ipv4Address            string `json:"ipv4-address"`
	Ipv4Netmask            string `json:"ipv4-netmask"`
	ControlAppsUseMgmtPort int    `json:"control-apps-use-mgmt-port"`
	DefaultGateway         string `json:"default-gateway"`
}

type InterfaceManagement2Ipv6 struct {
	Ipv6Addr           string `json:"ipv6-addr"`
	AddressType        string `json:"address-type"`
	V6AclName          string `json:"v6-acl-name"`
	Inbound            int    `json:"inbound"`
	DefaultIpv6Gateway string `json:"default-ipv6-gateway"`
}

type InterfaceManagement2SamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *InterfaceManagement2) GetId() string {
	return "1"
}

func (p *InterfaceManagement2) getPath() string {
	return "interface/management2"
}

func (p *InterfaceManagement2) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceManagement2::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
	return err
}

func (p *InterfaceManagement2) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceManagement2::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *InterfaceManagement2) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceManagement2::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
	return err
}

func (p *InterfaceManagement2) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceManagement2::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
