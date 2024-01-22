

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type OverlayTunnelVtepRemoteIpv6AddressUseLif struct {
	Inst struct {

    Lif string `json:"lif"`
    Partition string `json:"partition"`
    Uuid string `json:"uuid"`
    Id1 string 
    Ipv6Address string 

	} `json:"use-lif"`
}

func (p *OverlayTunnelVtepRemoteIpv6AddressUseLif) GetId() string{
    return "1"
}

func (p *OverlayTunnelVtepRemoteIpv6AddressUseLif) getPath() string{
    return "overlay-tunnel/vtep/" +p.Inst.Id1 + "/remote-ipv6-address/" +p.Inst.Ipv6Address + "/use-lif"
}

func (p *OverlayTunnelVtepRemoteIpv6AddressUseLif) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepRemoteIpv6AddressUseLif::Post")
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

func (p *OverlayTunnelVtepRemoteIpv6AddressUseLif) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepRemoteIpv6AddressUseLif::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *OverlayTunnelVtepRemoteIpv6AddressUseLif) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepRemoteIpv6AddressUseLif::Put")
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

func (p *OverlayTunnelVtepRemoteIpv6AddressUseLif) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepRemoteIpv6AddressUseLif::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
