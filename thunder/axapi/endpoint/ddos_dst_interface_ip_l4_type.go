

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstInterfaceIpL4Type struct {
	Inst struct {

    Deny int `json:"deny"`
    DropFragPkt int `json:"drop-frag-pkt"`
    Glid string `json:"glid"`
    Protocol string `json:"protocol"`
    TunnelDecap DdosDstInterfaceIpL4TypeTunnelDecap `json:"tunnel-decap"`
    TunnelRateLimit DdosDstInterfaceIpL4TypeTunnelRateLimit `json:"tunnel-rate-limit"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Addr string 

	} `json:"l4-type"`
}


type DdosDstInterfaceIpL4TypeTunnelDecap struct {
    IpDecap int `json:"ip-decap"`
    GreDecap int `json:"gre-decap"`
    KeyCfg []DdosDstInterfaceIpL4TypeTunnelDecapKeyCfg `json:"key-cfg"`
}


type DdosDstInterfaceIpL4TypeTunnelDecapKeyCfg struct {
    Key string `json:"key"`
}


type DdosDstInterfaceIpL4TypeTunnelRateLimit struct {
    IpRateLimit int `json:"ip-rate-limit"`
    GreRateLimit int `json:"gre-rate-limit"`
}

func (p *DdosDstInterfaceIpL4Type) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosDstInterfaceIpL4Type) getPath() string{
    return "ddos/dst/interface-ip/" +p.Inst.Addr + "/l4-type"
}

func (p *DdosDstInterfaceIpL4Type) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpL4Type::Post")
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

func (p *DdosDstInterfaceIpL4Type) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpL4Type::Get")
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
func (p *DdosDstInterfaceIpL4Type) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpL4Type::Put")
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

func (p *DdosDstInterfaceIpL4Type) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpL4Type::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
