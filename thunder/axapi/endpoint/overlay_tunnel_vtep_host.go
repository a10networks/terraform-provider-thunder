

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type OverlayTunnelVtepHost struct {
	Inst struct {

    IpAddr string `json:"ip-addr"`
    OverlayMacAddr string `json:"overlay-mac-addr"`
    RemoteVtep string `json:"remote-vtep"`
    Uuid string `json:"uuid"`
    Vni int `json:"vni"`
    Id1 string 

	} `json:"host"`
}

func (p *OverlayTunnelVtepHost) GetId() string{
    return p.Inst.IpAddr+"+"+p.Inst.OverlayMacAddr+"+"+strconv.Itoa(p.Inst.Vni)+"+"+p.Inst.RemoteVtep
}

func (p *OverlayTunnelVtepHost) getPath() string{
    return "overlay-tunnel/vtep/" +p.Inst.Id1 + "/host"
}

func (p *OverlayTunnelVtepHost) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepHost::Post")
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

func (p *OverlayTunnelVtepHost) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepHost::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *OverlayTunnelVtepHost) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepHost::Put")
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

func (p *OverlayTunnelVtepHost) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepHost::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
