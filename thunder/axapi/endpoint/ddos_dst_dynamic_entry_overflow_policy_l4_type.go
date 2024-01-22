

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstDynamicEntryOverflowPolicyL4Type struct {
	Inst struct {

    Deny int `json:"deny"`
    DropFragPkt int `json:"drop-frag-pkt"`
    DropOnNoPortMatch string `json:"drop-on-no-port-match" dval:"enable"`
    Glid string `json:"glid"`
    MaxRexmitSynPerFlow int `json:"max-rexmit-syn-per-flow"`
    Protocol string `json:"protocol"`
    Stateful int `json:"stateful"`
    SynAuth string `json:"syn-auth" dval:"send-rst"`
    SynCookie int `json:"syn-cookie"`
    TcpResetClient int `json:"tcp-reset-client"`
    TcpResetServer int `json:"tcp-reset-server"`
    TunnelDecap DdosDstDynamicEntryOverflowPolicyL4TypeTunnelDecap `json:"tunnel-decap"`
    TunnelRateLimit DdosDstDynamicEntryOverflowPolicyL4TypeTunnelRateLimit `json:"tunnel-rate-limit"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DefaultAddressType string 

	} `json:"l4-type"`
}


type DdosDstDynamicEntryOverflowPolicyL4TypeTunnelDecap struct {
    IpDecap int `json:"ip-decap"`
    GreDecap int `json:"gre-decap"`
    KeyCfg []DdosDstDynamicEntryOverflowPolicyL4TypeTunnelDecapKeyCfg `json:"key-cfg"`
}


type DdosDstDynamicEntryOverflowPolicyL4TypeTunnelDecapKeyCfg struct {
    Key string `json:"key"`
}


type DdosDstDynamicEntryOverflowPolicyL4TypeTunnelRateLimit struct {
    IpRateLimit int `json:"ip-rate-limit"`
    GreRateLimit int `json:"gre-rate-limit"`
}

func (p *DdosDstDynamicEntryOverflowPolicyL4Type) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosDstDynamicEntryOverflowPolicyL4Type) getPath() string{
    return "ddos/dst/dynamic-entry-overflow-policy/" +p.Inst.DefaultAddressType + "/l4-type"
}

func (p *DdosDstDynamicEntryOverflowPolicyL4Type) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicyL4Type::Post")
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

func (p *DdosDstDynamicEntryOverflowPolicyL4Type) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicyL4Type::Get")
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
func (p *DdosDstDynamicEntryOverflowPolicyL4Type) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicyL4Type::Put")
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

func (p *DdosDstDynamicEntryOverflowPolicyL4Type) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicyL4Type::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
