

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstInterfaceIpv6 struct {
	Inst struct {

    Addr string `json:"addr"`
    IpProtoList []DdosDstInterfaceIpv6IpProtoList `json:"ip-proto-list"`
    L4TypeList []DdosDstInterfaceIpv6L4TypeList `json:"l4-type-list"`
    LogEnable int `json:"log-enable"`
    PortList []DdosDstInterfaceIpv6PortList `json:"port-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"interface-ipv6"`
}


type DdosDstInterfaceIpv6IpProtoList struct {
    PortNum int `json:"port-num"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstInterfaceIpv6L4TypeList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    TunnelDecap DdosDstInterfaceIpv6L4TypeListTunnelDecap `json:"tunnel-decap"`
    TunnelRateLimit DdosDstInterfaceIpv6L4TypeListTunnelRateLimit `json:"tunnel-rate-limit"`
    DropFragPkt int `json:"drop-frag-pkt"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstInterfaceIpv6L4TypeListTunnelDecap struct {
    IpDecap int `json:"ip-decap"`
    GreDecap int `json:"gre-decap"`
    KeyCfg []DdosDstInterfaceIpv6L4TypeListTunnelDecapKeyCfg `json:"key-cfg"`
}


type DdosDstInterfaceIpv6L4TypeListTunnelDecapKeyCfg struct {
    Key string `json:"key"`
}


type DdosDstInterfaceIpv6L4TypeListTunnelRateLimit struct {
    IpRateLimit int `json:"ip-rate-limit"`
    GreRateLimit int `json:"gre-rate-limit"`
}


type DdosDstInterfaceIpv6PortList struct {
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DdosDstInterfaceIpv6) GetId() string{
    return p.Inst.Addr
}

func (p *DdosDstInterfaceIpv6) getPath() string{
    return "ddos/dst/interface-ipv6"
}

func (p *DdosDstInterfaceIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpv6::Post")
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

func (p *DdosDstInterfaceIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpv6::Get")
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
func (p *DdosDstInterfaceIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpv6::Put")
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

func (p *DdosDstInterfaceIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
