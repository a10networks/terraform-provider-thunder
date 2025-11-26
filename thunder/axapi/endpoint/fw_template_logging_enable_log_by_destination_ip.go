package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type FwTemplateLoggingEnableLogByDestinationIp struct {
	Inst struct {
		Icmp int `json:"icmp"`

		Ipv4Addr string `json:"ipv4-addr"`

		Others int `json:"others"`

		TcpList []FwTemplateLoggingEnableLogByDestinationIpTcpList `json:"tcp-list"`

		UdpList []FwTemplateLoggingEnableLogByDestinationIpUdpList `json:"udp-list"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		Logging_name string
	} `json:"ip"`
}

type FwTemplateLoggingEnableLogByDestinationIpTcpList struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingEnableLogByDestinationIpUdpList struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

func (p *FwTemplateLoggingEnableLogByDestinationIp) GetId() string {
	return url.QueryEscape(p.Inst.Ipv4Addr)
}

func (p *FwTemplateLoggingEnableLogByDestinationIp) getPath() string {
	return "fw/template/logging/" + p.Inst.Logging_name + "/enable-log-by-destination/ip"
}

func (p *FwTemplateLoggingEnableLogByDestinationIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLoggingEnableLogByDestinationIp::Post")
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

func (p *FwTemplateLoggingEnableLogByDestinationIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLoggingEnableLogByDestinationIp::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *FwTemplateLoggingEnableLogByDestinationIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLoggingEnableLogByDestinationIp::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
	return err
}

func (p *FwTemplateLoggingEnableLogByDestinationIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLoggingEnableLogByDestinationIp::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
