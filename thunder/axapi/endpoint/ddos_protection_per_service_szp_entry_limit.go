package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosProtectionPerServiceSzpEntryLimit struct {
	Inst struct {
		DnsTcpLimit int `json:"dns-tcp-limit"`

		DnsUdpLimit int `json:"dns-udp-limit"`

		HttpLimit int `json:"http-limit"`

		IpProtoCustomLimit int `json:"ip-proto-custom-limit"`

		IpProtoGreLimit int `json:"ip-proto-gre-limit"`

		IpProtoIcmpV4Limit int `json:"ip-proto-icmp-v4-limit"`

		IpProtoIcmpV6Limit int `json:"ip-proto-icmp-v6-limit"`

		IpProtoIpv4EncapLimit int `json:"ip-proto-ipv4-encap-limit"`

		IpProtoIpv6EncapLimit int `json:"ip-proto-ipv6-encap-limit"`

		IpProtoOtherLimit int `json:"ip-proto-other-limit"`

		QuicLimit int `json:"quic-limit"`

		SipTcpLimit int `json:"sip-tcp-limit"`

		SipUdpLimit int `json:"sip-udp-limit"`

		SslL4Limit int `json:"ssl-l4-limit"`

		TcpLimit int `json:"tcp-limit"`

		UdpLimit int `json:"udp-limit"`

		Uuid string `json:"uuid"`
	} `json:"per-service-szp-entry-limit"`
}

func (p *DdosProtectionPerServiceSzpEntryLimit) GetId() string {
	return "1"
}

func (p *DdosProtectionPerServiceSzpEntryLimit) getPath() string {
	return "ddos/protection/per-service-szp-entry-limit"
}

func (p *DdosProtectionPerServiceSzpEntryLimit) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosProtectionPerServiceSzpEntryLimit::Post")
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

func (p *DdosProtectionPerServiceSzpEntryLimit) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosProtectionPerServiceSzpEntryLimit::Get")
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
func (p *DdosProtectionPerServiceSzpEntryLimit) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosProtectionPerServiceSzpEntryLimit::Put")
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

func (p *DdosProtectionPerServiceSzpEntryLimit) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosProtectionPerServiceSzpEntryLimit::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
