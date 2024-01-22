

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstDefault struct {
	Inst struct {

    Age int `json:"age"`
    ApplyPolicyOnOverflow int `json:"apply-policy-on-overflow"`
    DefaultAddressType string `json:"default-address-type"`
    Deny int `json:"deny"`
    Disable int `json:"disable"`
    DropDisable int `json:"drop-disable"`
    DropDisableFwdImmediate int `json:"drop-disable-fwd-immediate"`
    DropFragPkt int `json:"drop-frag-pkt"`
    ExceedLogCfg DdosDstDefaultExceedLogCfg `json:"exceed-log-cfg"`
    ExceedLogDepCfg DdosDstDefaultExceedLogDepCfg `json:"exceed-log-dep-cfg"`
    Glid string `json:"glid"`
    InboundForwardDscp int `json:"inbound-forward-dscp"`
    IpProtoList []DdosDstDefaultIpProtoList `json:"ip-proto-list"`
    L4TypeList []DdosDstDefaultL4TypeList `json:"l4-type-list"`
    LogPeriodic int `json:"log-periodic"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    OutboundForwardDscp int `json:"outbound-forward-dscp"`
    PortList []DdosDstDefaultPortList `json:"port-list"`
    SrcPortList []DdosDstDefaultSrcPortList `json:"src-port-list"`
    Template DdosDstDefaultTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"default"`
}


type DdosDstDefaultExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
    WithSflowSample int `json:"with-sflow-sample"`
}


type DdosDstDefaultExceedLogDepCfg struct {
    ExceedLogEnable int `json:"exceed-log-enable"`
    LogWithSflowDep int `json:"log-with-sflow-dep"`
}


type DdosDstDefaultIpProtoList struct {
    PortNum int `json:"port-num"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstDefaultIpProtoListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstDefaultIpProtoListTemplate struct {
    Other string `json:"other"`
}


type DdosDstDefaultL4TypeList struct {
    Protocol string `json:"protocol"`
    Glid string `json:"glid"`
    Deny int `json:"deny"`
    MaxRexmitSynPerFlow int `json:"max-rexmit-syn-per-flow"`
    DisableSynAuth int `json:"disable-syn-auth"`
    SynAuth string `json:"syn-auth" dval:"send-rst"`
    SynCookie int `json:"syn-cookie"`
    TcpResetClient int `json:"tcp-reset-client"`
    TcpResetServer int `json:"tcp-reset-server"`
    DropOnNoPortMatch string `json:"drop-on-no-port-match" dval:"enable"`
    Stateful int `json:"stateful"`
    TunnelDecap DdosDstDefaultL4TypeListTunnelDecap `json:"tunnel-decap"`
    TunnelRateLimit DdosDstDefaultL4TypeListTunnelRateLimit `json:"tunnel-rate-limit"`
    DropFragPkt int `json:"drop-frag-pkt"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstDefaultL4TypeListTunnelDecap struct {
    IpDecap int `json:"ip-decap"`
    GreDecap int `json:"gre-decap"`
    KeyCfg []DdosDstDefaultL4TypeListTunnelDecapKeyCfg `json:"key-cfg"`
}


type DdosDstDefaultL4TypeListTunnelDecapKeyCfg struct {
    Key string `json:"key"`
}


type DdosDstDefaultL4TypeListTunnelRateLimit struct {
    IpRateLimit int `json:"ip-rate-limit"`
    GreRateLimit int `json:"gre-rate-limit"`
}


type DdosDstDefaultPortList struct {
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstDefaultPortListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstDefaultPortListTemplate struct {
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
}


type DdosDstDefaultSrcPortList struct {
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstDefaultSrcPortListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstDefaultSrcPortListTemplate struct {
    SrcUdp string `json:"src-udp"`
    SrcTcp string `json:"src-tcp"`
}


type DdosDstDefaultTemplate struct {
    Logging string `json:"logging"`
}

func (p *DdosDstDefault) GetId() string{
    return p.Inst.DefaultAddressType
}

func (p *DdosDstDefault) getPath() string{
    return "ddos/dst/default"
}

func (p *DdosDstDefault) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDefault::Post")
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

func (p *DdosDstDefault) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDefault::Get")
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
func (p *DdosDstDefault) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDefault::Put")
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

func (p *DdosDstDefault) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDefault::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
