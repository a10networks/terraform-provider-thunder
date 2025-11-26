package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type OverlayTunnelVtepLocalIpv6AddressVni struct {
	Inst struct {
		Gateway int `json:"gateway"`

		Lif string `json:"lif"`

		Partition string `json:"partition"`

		Segment int `json:"segment"`

		Uuid string `json:"uuid"`

		Id1 string
	} `json:"vni"`
}

func (p *OverlayTunnelVtepLocalIpv6AddressVni) GetId() string {
	return strconv.Itoa(p.Inst.Segment)
}

func (p *OverlayTunnelVtepLocalIpv6AddressVni) getPath() string {
	return "overlay-tunnel/vtep/" + p.Inst.Id1 + "/local-ipv6-address/vni"
}

func (p *OverlayTunnelVtepLocalIpv6AddressVni) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("OverlayTunnelVtepLocalIpv6AddressVni::Post")
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

func (p *OverlayTunnelVtepLocalIpv6AddressVni) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("OverlayTunnelVtepLocalIpv6AddressVni::Get")
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
func (p *OverlayTunnelVtepLocalIpv6AddressVni) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("OverlayTunnelVtepLocalIpv6AddressVni::Put")
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

func (p *OverlayTunnelVtepLocalIpv6AddressVni) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("OverlayTunnelVtepLocalIpv6AddressVni::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
