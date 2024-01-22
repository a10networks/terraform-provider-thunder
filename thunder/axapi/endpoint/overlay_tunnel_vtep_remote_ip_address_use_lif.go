

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type OverlayTunnelVtepRemoteIpAddressUseLif struct {
	Inst struct {

    Lif string `json:"lif"`
    Partition string `json:"partition"`
    Uuid string `json:"uuid"`
    Id1 string 
    IpAddress string 

	} `json:"use-lif"`
}

func (p *OverlayTunnelVtepRemoteIpAddressUseLif) GetId() string{
    return "1"
}

func (p *OverlayTunnelVtepRemoteIpAddressUseLif) getPath() string{
    return "overlay-tunnel/vtep/" +p.Inst.Id1 + "/remote-ip-address/" +p.Inst.IpAddress + "/use-lif"
}

func (p *OverlayTunnelVtepRemoteIpAddressUseLif) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepRemoteIpAddressUseLif::Post")
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

func (p *OverlayTunnelVtepRemoteIpAddressUseLif) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepRemoteIpAddressUseLif::Get")
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
func (p *OverlayTunnelVtepRemoteIpAddressUseLif) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepRemoteIpAddressUseLif::Put")
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

func (p *OverlayTunnelVtepRemoteIpAddressUseLif) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepRemoteIpAddressUseLif::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
