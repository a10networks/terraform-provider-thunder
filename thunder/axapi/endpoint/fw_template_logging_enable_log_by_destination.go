package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type FwTemplateLoggingEnableLogByDestination struct {
	Inst struct {
		Icmp int `json:"icmp"`

		IpList []FwTemplateLoggingEnableLogByDestinationIpList `json:"ip-list"`

		Ip6List []FwTemplateLoggingEnableLogByDestinationIp6List `json:"ip6-list"`

		Others int `json:"others"`

		TcpList []FwTemplateLoggingEnableLogByDestinationTcpList `json:"tcp-list"`

		UdpList []FwTemplateLoggingEnableLogByDestinationUdpList `json:"udp-list"`

		Uuid string `json:"uuid"`

		Logging_name string
	} `json:"enable-log-by-destination"`
}

type FwTemplateLoggingEnableLogByDestinationIpList struct {
	Ipv4Addr string                                                 `json:"ipv4-addr"`
	TcpList  []FwTemplateLoggingEnableLogByDestinationIpListTcpList `json:"tcp-list"`
	UdpList  []FwTemplateLoggingEnableLogByDestinationIpListUdpList `json:"udp-list"`
	Icmp     int                                                    `json:"icmp"`
	Others   int                                                    `json:"others"`
	Uuid     string                                                 `json:"uuid"`
	UserTag  string                                                 `json:"user-tag"`
}

type FwTemplateLoggingEnableLogByDestinationIpListTcpList struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingEnableLogByDestinationIpListUdpList struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

type FwTemplateLoggingEnableLogByDestinationIp6List struct {
	Ipv6Addr string                                                  `json:"ipv6-addr"`
	TcpList  []FwTemplateLoggingEnableLogByDestinationIp6ListTcpList `json:"tcp-list"`
	UdpList  []FwTemplateLoggingEnableLogByDestinationIp6ListUdpList `json:"udp-list"`
	Icmp     int                                                     `json:"icmp"`
	Others   int                                                     `json:"others"`
	Uuid     string                                                  `json:"uuid"`
	UserTag  string                                                  `json:"user-tag"`
}

type FwTemplateLoggingEnableLogByDestinationIp6ListTcpList struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingEnableLogByDestinationIp6ListUdpList struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

type FwTemplateLoggingEnableLogByDestinationTcpList struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingEnableLogByDestinationUdpList struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

func (p *FwTemplateLoggingEnableLogByDestination) GetId() string {
	return "1"
}

func (p *FwTemplateLoggingEnableLogByDestination) getPath() string {
	return "fw/template/logging/" + p.Inst.Logging_name + "/enable-log-by-destination"
}

func (p *FwTemplateLoggingEnableLogByDestination) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLoggingEnableLogByDestination::Post")
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

func (p *FwTemplateLoggingEnableLogByDestination) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLoggingEnableLogByDestination::Get")
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
func (p *FwTemplateLoggingEnableLogByDestination) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLoggingEnableLogByDestination::Put")
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

func (p *FwTemplateLoggingEnableLogByDestination) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLoggingEnableLogByDestination::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
