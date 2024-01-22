

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryL4Type struct {
	Inst struct {

    Deny int `json:"deny"`
    DetectionEnable int `json:"detection-enable"`
    DisableSynAuth int `json:"disable-syn-auth"`
    DropFragPkt int `json:"drop-frag-pkt"`
    DropOnNoPortMatch string `json:"drop-on-no-port-match" dval:"enable"`
    EnableTopK int `json:"enable-top-k"`
    Glid string `json:"glid"`
    GlidExceedAction DdosDstEntryL4TypeGlidExceedAction `json:"glid-exceed-action"`
    IpFilteringPolicy string `json:"ip-filtering-policy"`
    IpFilteringPolicyOper DdosDstEntryL4TypeIpFilteringPolicyOper158 `json:"ip-filtering-policy-oper"`
    MaxRexmitSynPerFlow int `json:"max-rexmit-syn-per-flow"`
    MaxRexmitSynPerFlowExceedAction string `json:"max-rexmit-syn-per-flow-exceed-action"`
    PortInd DdosDstEntryL4TypePortInd159 `json:"port-ind"`
    ProgressionTracking DdosDstEntryL4TypeProgressionTracking161 `json:"progression-tracking"`
    Protocol string `json:"protocol"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    Stateful int `json:"stateful"`
    SynAuth string `json:"syn-auth" dval:"send-rst"`
    SynCookie int `json:"syn-cookie"`
    TcpResetClient int `json:"tcp-reset-client"`
    TcpResetServer int `json:"tcp-reset-server"`
    Template DdosDstEntryL4TypeTemplate `json:"template"`
    TopkNumRecords int `json:"topk-num-records" dval:"20"`
    TopkSources DdosDstEntryL4TypeTopkSources162 `json:"topk-sources"`
    TunnelDecap DdosDstEntryL4TypeTunnelDecap `json:"tunnel-decap"`
    TunnelRateLimit DdosDstEntryL4TypeTunnelRateLimit `json:"tunnel-rate-limit"`
    UndefinedPortHitStatistics DdosDstEntryL4TypeUndefinedPortHitStatistics `json:"undefined-port-hit-statistics"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DstEntryName string 

	} `json:"l4-type"`
}


type DdosDstEntryL4TypeGlidExceedAction struct {
    StatelessEncapActionCfg DdosDstEntryL4TypeGlidExceedActionStatelessEncapActionCfg `json:"stateless-encap-action-cfg"`
}


type DdosDstEntryL4TypeGlidExceedActionStatelessEncapActionCfg struct {
    StatelessEncapAction string `json:"stateless-encap-action"`
    EncapTemplate string `json:"encap-template"`
}


type DdosDstEntryL4TypeIpFilteringPolicyOper158 struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryL4TypePortInd159 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []DdosDstEntryL4TypePortIndSamplingEnable160 `json:"sampling-enable"`
}


type DdosDstEntryL4TypePortIndSamplingEnable160 struct {
    Counters1 string `json:"counters1"`
}


type DdosDstEntryL4TypeProgressionTracking161 struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryL4TypeTemplate struct {
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntryL4TypeTopkSources162 struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryL4TypeTunnelDecap struct {
    IpDecap int `json:"ip-decap"`
    GreDecap int `json:"gre-decap"`
    KeyCfg []DdosDstEntryL4TypeTunnelDecapKeyCfg `json:"key-cfg"`
}


type DdosDstEntryL4TypeTunnelDecapKeyCfg struct {
    Key string `json:"key"`
}


type DdosDstEntryL4TypeTunnelRateLimit struct {
    IpRateLimit int `json:"ip-rate-limit"`
    GreRateLimit int `json:"gre-rate-limit"`
}


type DdosDstEntryL4TypeUndefinedPortHitStatistics struct {
    UndefinedPortHitStatistics int `json:"undefined-port-hit-statistics"`
    ResetInterval int `json:"reset-interval" dval:"60"`
}

func (p *DdosDstEntryL4Type) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosDstEntryL4Type) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/l4-type"
}

func (p *DdosDstEntryL4Type) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryL4Type::Post")
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

func (p *DdosDstEntryL4Type) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryL4Type::Get")
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
func (p *DdosDstEntryL4Type) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryL4Type::Put")
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

func (p *DdosDstEntryL4Type) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryL4Type::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
