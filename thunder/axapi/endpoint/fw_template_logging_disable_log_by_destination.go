package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type FwTemplateLoggingDisableLogByDestination struct {
	Inst struct {
		Icmp int `json:"icmp"`

		IpList []FwTemplateLoggingDisableLogByDestinationIpList `json:"ip-list"`

		Ip6List []FwTemplateLoggingDisableLogByDestinationIp6List `json:"ip6-list"`

		Others int `json:"others"`

		TcpList []FwTemplateLoggingDisableLogByDestinationTcpList `json:"tcp-list"`

		UdpList []FwTemplateLoggingDisableLogByDestinationUdpList `json:"udp-list"`

		Uuid string `json:"uuid"`

		Logging_name string
	} `json:"disable-log-by-destination"`
}

type FwTemplateLoggingDisableLogByDestinationIpList struct {
	Ipv4Addr string                                                  `json:"ipv4-addr"`
	TcpList  []FwTemplateLoggingDisableLogByDestinationIpListTcpList `json:"tcp-list"`
	UdpList  []FwTemplateLoggingDisableLogByDestinationIpListUdpList `json:"udp-list"`
	Icmp     int                                                     `json:"icmp"`
	Others   int                                                     `json:"others"`
	Uuid     string                                                  `json:"uuid"`
	UserTag  string                                                  `json:"user-tag"`
}

type FwTemplateLoggingDisableLogByDestinationIpListTcpList struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingDisableLogByDestinationIpListUdpList struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

type FwTemplateLoggingDisableLogByDestinationIp6List struct {
	Ipv6Addr string                                                   `json:"ipv6-addr"`
	TcpList  []FwTemplateLoggingDisableLogByDestinationIp6ListTcpList `json:"tcp-list"`
	UdpList  []FwTemplateLoggingDisableLogByDestinationIp6ListUdpList `json:"udp-list"`
	Icmp     int                                                      `json:"icmp"`
	Others   int                                                      `json:"others"`
	Uuid     string                                                   `json:"uuid"`
	UserTag  string                                                   `json:"user-tag"`
}

type FwTemplateLoggingDisableLogByDestinationIp6ListTcpList struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingDisableLogByDestinationIp6ListUdpList struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

type FwTemplateLoggingDisableLogByDestinationTcpList struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingDisableLogByDestinationUdpList struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

func (p *FwTemplateLoggingDisableLogByDestination) GetId() string {
	return "1"
}

func (p *FwTemplateLoggingDisableLogByDestination) getPath() string {
	return "fw/template/logging/" + p.Inst.Logging_name + "/disable-log-by-destination"
}

func (p *FwTemplateLoggingDisableLogByDestination) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLoggingDisableLogByDestination::Post")
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

func (p *FwTemplateLoggingDisableLogByDestination) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLoggingDisableLogByDestination::Get")
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
func (p *FwTemplateLoggingDisableLogByDestination) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLoggingDisableLogByDestination::Put")
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

func (p *FwTemplateLoggingDisableLogByDestination) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLoggingDisableLogByDestination::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
