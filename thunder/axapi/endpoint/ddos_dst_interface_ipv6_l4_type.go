

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstInterfaceIpv6L4Type struct {
	Inst struct {

    Deny int `json:"deny"`
    DropFragPkt int `json:"drop-frag-pkt"`
    Glid string `json:"glid"`
    Protocol string `json:"protocol"`
    TunnelDecap DdosDstInterfaceIpv6L4TypeTunnelDecap `json:"tunnel-decap"`
    TunnelRateLimit DdosDstInterfaceIpv6L4TypeTunnelRateLimit `json:"tunnel-rate-limit"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Addr string 

	} `json:"l4-type"`
}


type DdosDstInterfaceIpv6L4TypeTunnelDecap struct {
    IpDecap int `json:"ip-decap"`
    GreDecap int `json:"gre-decap"`
    KeyCfg []DdosDstInterfaceIpv6L4TypeTunnelDecapKeyCfg `json:"key-cfg"`
}


type DdosDstInterfaceIpv6L4TypeTunnelDecapKeyCfg struct {
    Key string `json:"key"`
}


type DdosDstInterfaceIpv6L4TypeTunnelRateLimit struct {
    IpRateLimit int `json:"ip-rate-limit"`
    GreRateLimit int `json:"gre-rate-limit"`
}

func (p *DdosDstInterfaceIpv6L4Type) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosDstInterfaceIpv6L4Type) getPath() string{
    return "ddos/dst/interface-ipv6/" +p.Inst.Addr + "/l4-type"
}

func (p *DdosDstInterfaceIpv6L4Type) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpv6L4Type::Post")
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

func (p *DdosDstInterfaceIpv6L4Type) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpv6L4Type::Get")
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
func (p *DdosDstInterfaceIpv6L4Type) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpv6L4Type::Put")
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

func (p *DdosDstInterfaceIpv6L4Type) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpv6L4Type::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
