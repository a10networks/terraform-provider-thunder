

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type OverlayTunnelOptions struct {
	Inst struct {

    FragmentationModeInner int `json:"fragmentation-mode-inner"`
    GatewayMac string `json:"gateway-mac"`
    IpDscpPreserve int `json:"ip-dscp-preserve"`
    NvgreDisableFlowId int `json:"nvgre-disable-flow-id"`
    NvgreKeyModeLower24 int `json:"nvgre-key-mode-lower24"`
    SrcPortRange OverlayTunnelOptionsSrcPortRange1080 `json:"src-port-range"`
    TcpMssAdjustDisable int `json:"tcp-mss-adjust-disable"`
    Uuid string `json:"uuid"`
    VxlanDestPort int `json:"vxlan-dest-port"`

	} `json:"options"`
}


type OverlayTunnelOptionsSrcPortRange1080 struct {
    MinPort int `json:"min-port" dval:"1"`
    MaxPort int `json:"max-port" dval:"65535"`
    Uuid string `json:"uuid"`
}

func (p *OverlayTunnelOptions) GetId() string{
    return "1"
}

func (p *OverlayTunnelOptions) getPath() string{
    return "overlay-tunnel/options"
}

func (p *OverlayTunnelOptions) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelOptions::Post")
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

func (p *OverlayTunnelOptions) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelOptions::Get")
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
func (p *OverlayTunnelOptions) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelOptions::Put")
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

func (p *OverlayTunnelOptions) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelOptions::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
