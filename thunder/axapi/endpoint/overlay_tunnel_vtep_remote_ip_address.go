

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type OverlayTunnelVtepRemoteIpAddress struct {
	Inst struct {

    Encap string `json:"encap"`
    GreKeepalive OverlayTunnelVtepRemoteIpAddressGreKeepalive1081 `json:"gre-keepalive"`
    IpAddress string `json:"ip-address"`
    UseGreKey OverlayTunnelVtepRemoteIpAddressUseGreKey1082 `json:"use-gre-key"`
    UseLif OverlayTunnelVtepRemoteIpAddressUseLif1083 `json:"use-lif"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VniList []OverlayTunnelVtepRemoteIpAddressVniList `json:"vni-list"`
    Id1 string 

	} `json:"remote-ip-address"`
}


type OverlayTunnelVtepRemoteIpAddressGreKeepalive1081 struct {
    RetryTime int `json:"retry-time" dval:"10"`
    RetryCount int `json:"retry-count"`
    Uuid string `json:"uuid"`
}


type OverlayTunnelVtepRemoteIpAddressUseGreKey1082 struct {
    GreKey int `json:"gre-key"`
    Uuid string `json:"uuid"`
}


type OverlayTunnelVtepRemoteIpAddressUseLif1083 struct {
    Partition string `json:"partition"`
    Lif string `json:"lif"`
    Uuid string `json:"uuid"`
}


type OverlayTunnelVtepRemoteIpAddressVniList struct {
    Segment int `json:"segment"`
    Uuid string `json:"uuid"`
}

func (p *OverlayTunnelVtepRemoteIpAddress) GetId() string{
    return p.Inst.IpAddress
}

func (p *OverlayTunnelVtepRemoteIpAddress) getPath() string{
    return "overlay-tunnel/vtep/" +p.Inst.Id1 + "/remote-ip-address"
}

func (p *OverlayTunnelVtepRemoteIpAddress) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepRemoteIpAddress::Post")
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

func (p *OverlayTunnelVtepRemoteIpAddress) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepRemoteIpAddress::Get")
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
func (p *OverlayTunnelVtepRemoteIpAddress) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepRemoteIpAddress::Put")
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

func (p *OverlayTunnelVtepRemoteIpAddress) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtepRemoteIpAddress::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
