

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstDynamicEntryOverflowPolicy struct {
	Inst struct {

    DefaultAddressType string `json:"default-address-type"`
    DropDisable int `json:"drop-disable"`
    DropDisableFwdImmediate int `json:"drop-disable-fwd-immediate"`
    ExceedLogCfg DdosDstDynamicEntryOverflowPolicyExceedLogCfg `json:"exceed-log-cfg"`
    ExceedLogDepCfg DdosDstDynamicEntryOverflowPolicyExceedLogDepCfg `json:"exceed-log-dep-cfg"`
    Glid string `json:"glid"`
    InboundForwardDscp int `json:"inbound-forward-dscp"`
    IpProtoList []DdosDstDynamicEntryOverflowPolicyIpProtoList `json:"ip-proto-list"`
    L4TypeList []DdosDstDynamicEntryOverflowPolicyL4TypeList `json:"l4-type-list"`
    LogPeriodic int `json:"log-periodic"`
    OutboundForwardDscp int `json:"outbound-forward-dscp"`
    PortList []DdosDstDynamicEntryOverflowPolicyPortList `json:"port-list"`
    SrcPortList []DdosDstDynamicEntryOverflowPolicySrcPortList `json:"src-port-list"`
    Template DdosDstDynamicEntryOverflowPolicyTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"dynamic-entry-overflow-policy"`
}


type DdosDstDynamicEntryOverflowPolicyExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
    WithSflowSample int `json:"with-sflow-sample"`
}


type DdosDstDynamicEntryOverflowPolicyExceedLogDepCfg struct {
    ExceedLogEnable int `json:"exceed-log-enable"`
    LogWithSflowDep int `json:"log-with-sflow-dep"`
}


type DdosDstDynamicEntryOverflowPolicyIpProtoList struct {
    PortNum int `json:"port-num"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstDynamicEntryOverflowPolicyIpProtoListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstDynamicEntryOverflowPolicyIpProtoListTemplate struct {
    Other string `json:"other"`
}


type DdosDstDynamicEntryOverflowPolicyL4TypeList struct {
    Protocol string `json:"protocol"`
    Glid string `json:"glid"`
    Deny int `json:"deny"`
    MaxRexmitSynPerFlow int `json:"max-rexmit-syn-per-flow"`
    SynAuth string `json:"syn-auth" dval:"send-rst"`
    SynCookie int `json:"syn-cookie"`
    TcpResetClient int `json:"tcp-reset-client"`
    TcpResetServer int `json:"tcp-reset-server"`
    DropOnNoPortMatch string `json:"drop-on-no-port-match" dval:"enable"`
    Stateful int `json:"stateful"`
    TunnelDecap DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecap `json:"tunnel-decap"`
    TunnelRateLimit DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelRateLimit `json:"tunnel-rate-limit"`
    DropFragPkt int `json:"drop-frag-pkt"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecap struct {
    IpDecap int `json:"ip-decap"`
    GreDecap int `json:"gre-decap"`
    KeyCfg []DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecapKeyCfg `json:"key-cfg"`
}


type DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecapKeyCfg struct {
    Key string `json:"key"`
}


type DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelRateLimit struct {
    IpRateLimit int `json:"ip-rate-limit"`
    GreRateLimit int `json:"gre-rate-limit"`
}


type DdosDstDynamicEntryOverflowPolicyPortList struct {
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstDynamicEntryOverflowPolicyPortListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstDynamicEntryOverflowPolicyPortListTemplate struct {
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
}


type DdosDstDynamicEntryOverflowPolicySrcPortList struct {
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstDynamicEntryOverflowPolicySrcPortListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstDynamicEntryOverflowPolicySrcPortListTemplate struct {
    SrcUdp string `json:"src-udp"`
    SrcTcp string `json:"src-tcp"`
}


type DdosDstDynamicEntryOverflowPolicyTemplate struct {
    Logging string `json:"logging"`
}

func (p *DdosDstDynamicEntryOverflowPolicy) GetId() string{
    return p.Inst.DefaultAddressType
}

func (p *DdosDstDynamicEntryOverflowPolicy) getPath() string{
    return "ddos/dst/dynamic-entry-overflow-policy"
}

func (p *DdosDstDynamicEntryOverflowPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicy::Post")
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

func (p *DdosDstDynamicEntryOverflowPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicy::Get")
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
func (p *DdosDstDynamicEntryOverflowPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicy::Put")
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

func (p *DdosDstDynamicEntryOverflowPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
