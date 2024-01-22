

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type OverlayTunnelVtepLocalIpAddress struct {
	Inst struct {

    IpAddress string `json:"ip-address"`
    Uuid string `json:"uuid"`
    VniList []OverlayTunnelVtepLocalIpAddressVniList `json:"vni-list"`
    Id1 string 

	} `json:"local-ip-address"`
}


type OverlayTunnelVtepLocalIpAddressVniList struct {
    Segment int `json:"segment"`
    Partition string `json:"partition"`
    Gateway int `json:"gateway"`
    Lif string `json:"lif"`
    Uuid string `json:"uuid"`
}

func (p *OverlayTunnelVtepLocalIpAddress) GetId() string{
    return "1"
}

func (p *OverlayTunnelVtepLocalIpAddress) getPath() string{
    return "overlay-tunnel/vtep/" +p.Inst.Id1 + "/local-ip-address"
}

func (p *OverlayTunnelVtepLocalIpAddress) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepLocalIpAddress::Post")
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

func (p *OverlayTunnelVtepLocalIpAddress) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepLocalIpAddress::Get")
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
func (p *OverlayTunnelVtepLocalIpAddress) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepLocalIpAddress::Put")
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

func (p *OverlayTunnelVtepLocalIpAddress) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepLocalIpAddress::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
