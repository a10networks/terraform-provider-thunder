package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type Cgnv6TemplateLoggingEnableLogByDestination struct {
	Inst struct {
		Icmp int `json:"icmp"`

		IpList []Cgnv6TemplateLoggingEnableLogByDestinationIpList `json:"ip-list"`

		Ip6List []Cgnv6TemplateLoggingEnableLogByDestinationIp6List `json:"ip6-list"`

		Others int `json:"others"`

		TcpList []Cgnv6TemplateLoggingEnableLogByDestinationTcpList `json:"tcp-list"`

		UdpList []Cgnv6TemplateLoggingEnableLogByDestinationUdpList `json:"udp-list"`

		Uuid string `json:"uuid"`

		Logging_name string
	} `json:"enable-log-by-destination"`
}

type Cgnv6TemplateLoggingEnableLogByDestinationIpList struct {
	Ipv4Addr string                                                    `json:"ipv4-addr"`
	TcpList  []Cgnv6TemplateLoggingEnableLogByDestinationIpListTcpList `json:"tcp-list"`
	UdpList  []Cgnv6TemplateLoggingEnableLogByDestinationIpListUdpList `json:"udp-list"`
	Icmp     int                                                       `json:"icmp"`
	Others   int                                                       `json:"others"`
	Uuid     string                                                    `json:"uuid"`
	UserTag  string                                                    `json:"user-tag"`
}

type Cgnv6TemplateLoggingEnableLogByDestinationIpListTcpList struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type Cgnv6TemplateLoggingEnableLogByDestinationIpListUdpList struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

type Cgnv6TemplateLoggingEnableLogByDestinationIp6List struct {
	Ipv6Addr string                                                     `json:"ipv6-addr"`
	TcpList  []Cgnv6TemplateLoggingEnableLogByDestinationIp6ListTcpList `json:"tcp-list"`
	UdpList  []Cgnv6TemplateLoggingEnableLogByDestinationIp6ListUdpList `json:"udp-list"`
	Icmp     int                                                        `json:"icmp"`
	Others   int                                                        `json:"others"`
	Uuid     string                                                     `json:"uuid"`
	UserTag  string                                                     `json:"user-tag"`
}

type Cgnv6TemplateLoggingEnableLogByDestinationIp6ListTcpList struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type Cgnv6TemplateLoggingEnableLogByDestinationIp6ListUdpList struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

type Cgnv6TemplateLoggingEnableLogByDestinationTcpList struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type Cgnv6TemplateLoggingEnableLogByDestinationUdpList struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

func (p *Cgnv6TemplateLoggingEnableLogByDestination) GetId() string {
	return "1"
}

func (p *Cgnv6TemplateLoggingEnableLogByDestination) getPath() string {
	return "cgnv6/template/logging/" + p.Inst.Logging_name + "/enable-log-by-destination"
}

func (p *Cgnv6TemplateLoggingEnableLogByDestination) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("Cgnv6TemplateLoggingEnableLogByDestination::Post")
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

func (p *Cgnv6TemplateLoggingEnableLogByDestination) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("Cgnv6TemplateLoggingEnableLogByDestination::Get")
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
func (p *Cgnv6TemplateLoggingEnableLogByDestination) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("Cgnv6TemplateLoggingEnableLogByDestination::Put")
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

func (p *Cgnv6TemplateLoggingEnableLogByDestination) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("Cgnv6TemplateLoggingEnableLogByDestination::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
