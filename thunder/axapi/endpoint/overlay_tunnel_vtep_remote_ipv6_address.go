package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type OverlayTunnelVtepRemoteIpv6Address struct {
	Inst struct {
		ClassList string `json:"class-list"`

		Encap string `json:"encap"`

		GreKeepalive OverlayTunnelVtepRemoteIpv6AddressGreKeepalive1164 `json:"gre-keepalive"`

		Ipv6Address string `json:"ipv6-address"`

		UseGreKey OverlayTunnelVtepRemoteIpv6AddressUseGreKey1165 `json:"use-gre-key"`

		UseLif OverlayTunnelVtepRemoteIpv6AddressUseLif1166 `json:"use-lif"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		VniList []OverlayTunnelVtepRemoteIpv6AddressVniList `json:"vni-list"`

		Id1 string
	} `json:"remote-ipv6-address"`
}

type OverlayTunnelVtepRemoteIpv6AddressGreKeepalive1164 struct {
	RetryTime  int    `json:"retry-time" dval:"10"`
	RetryCount int    `json:"retry-count"`
	Uuid       string `json:"uuid"`
}

type OverlayTunnelVtepRemoteIpv6AddressUseGreKey1165 struct {
	GreKey int    `json:"gre-key"`
	Uuid   string `json:"uuid"`
}

type OverlayTunnelVtepRemoteIpv6AddressUseLif1166 struct {
	Partition string `json:"partition"`
	Lif       string `json:"lif"`
	Uuid      string `json:"uuid"`
}

type OverlayTunnelVtepRemoteIpv6AddressVniList struct {
	Segment int    `json:"segment"`
	Uuid    string `json:"uuid"`
}

func (p *OverlayTunnelVtepRemoteIpv6Address) GetId() string {
	return p.Inst.Ipv6Address
}

func (p *OverlayTunnelVtepRemoteIpv6Address) getPath() string {
	return "overlay-tunnel/vtep/" + p.Inst.Id1 + "/remote-ipv6-address"
}

func (p *OverlayTunnelVtepRemoteIpv6Address) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("OverlayTunnelVtepRemoteIpv6Address::Post")
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

func (p *OverlayTunnelVtepRemoteIpv6Address) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("OverlayTunnelVtepRemoteIpv6Address::Get")
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
func (p *OverlayTunnelVtepRemoteIpv6Address) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("OverlayTunnelVtepRemoteIpv6Address::Put")
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

func (p *OverlayTunnelVtepRemoteIpv6Address) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("OverlayTunnelVtepRemoteIpv6Address::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
