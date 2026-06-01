package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type OverlayTunnelVtepRemoteIpv6AddressGreKeepalive struct {
	Inst struct {
		RetryCount int `json:"retry-count"`

		RetryTime int `json:"retry-time" dval:"10"`

		Uuid string `json:"uuid"`

		Ipv6Address string

		Id1 string
	} `json:"gre-keepalive"`
}

func (p *OverlayTunnelVtepRemoteIpv6AddressGreKeepalive) GetId() string {
	return "1"
}

func (p *OverlayTunnelVtepRemoteIpv6AddressGreKeepalive) getPath() string {
	return "overlay-tunnel/vtep/" + p.Inst.Id1 + "/remote-ipv6-address/" + p.Inst.Ipv6Address + "/gre-keepalive"
}

func (p *OverlayTunnelVtepRemoteIpv6AddressGreKeepalive) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("OverlayTunnelVtepRemoteIpv6AddressGreKeepalive::Post")
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

func (p *OverlayTunnelVtepRemoteIpv6AddressGreKeepalive) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("OverlayTunnelVtepRemoteIpv6AddressGreKeepalive::Get")
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
func (p *OverlayTunnelVtepRemoteIpv6AddressGreKeepalive) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("OverlayTunnelVtepRemoteIpv6AddressGreKeepalive::Put")
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

func (p *OverlayTunnelVtepRemoteIpv6AddressGreKeepalive) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("OverlayTunnelVtepRemoteIpv6AddressGreKeepalive::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
