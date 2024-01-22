

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstInterfaceIp struct {
	Inst struct {

    Addr string `json:"addr"`
    IpProtoList []DdosDstInterfaceIpIpProtoList `json:"ip-proto-list"`
    L4TypeList []DdosDstInterfaceIpL4TypeList `json:"l4-type-list"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    PortList []DdosDstInterfaceIpPortList `json:"port-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"interface-ip"`
}


type DdosDstInterfaceIpIpProtoList struct {
    PortNum int `json:"port-num"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstInterfaceIpL4TypeList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    TunnelDecap DdosDstInterfaceIpL4TypeListTunnelDecap `json:"tunnel-decap"`
    TunnelRateLimit DdosDstInterfaceIpL4TypeListTunnelRateLimit `json:"tunnel-rate-limit"`
    DropFragPkt int `json:"drop-frag-pkt"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstInterfaceIpL4TypeListTunnelDecap struct {
    IpDecap int `json:"ip-decap"`
    GreDecap int `json:"gre-decap"`
    KeyCfg []DdosDstInterfaceIpL4TypeListTunnelDecapKeyCfg `json:"key-cfg"`
}


type DdosDstInterfaceIpL4TypeListTunnelDecapKeyCfg struct {
    Key string `json:"key"`
}


type DdosDstInterfaceIpL4TypeListTunnelRateLimit struct {
    IpRateLimit int `json:"ip-rate-limit"`
    GreRateLimit int `json:"gre-rate-limit"`
}


type DdosDstInterfaceIpPortList struct {
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DdosDstInterfaceIp) GetId() string{
    return p.Inst.Addr
}

func (p *DdosDstInterfaceIp) getPath() string{
    return "ddos/dst/interface-ip"
}

func (p *DdosDstInterfaceIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIp::Post")
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

func (p *DdosDstInterfaceIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIp::Get")
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
func (p *DdosDstInterfaceIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIp::Put")
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

func (p *DdosDstInterfaceIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstInterfaceIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
