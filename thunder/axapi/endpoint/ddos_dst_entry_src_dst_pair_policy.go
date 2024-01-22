

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntrySrcDstPairPolicy struct {
	Inst struct {

    PolicyClassListList []DdosDstEntrySrcDstPairPolicyPolicyClassListList `json:"policy-class-list-list"`
    SrcBasedPolicyName string `json:"src-based-policy-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DstEntryName string 

	} `json:"src-dst-pair-policy"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListList struct {
    ClassListName string `json:"class-list-name"`
    Bypass int `json:"bypass"`
    ExceedLogCfg DdosDstEntrySrcDstPairPolicyPolicyClassListListExceedLogCfg `json:"exceed-log-cfg"`
    LogPeriodic int `json:"log-periodic"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListListTemplate `json:"template"`
    Glid string `json:"glid"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []DdosDstEntrySrcDstPairPolicyPolicyClassListListSamplingEnable `json:"sampling-enable"`
    L4TypeSrcDstList []DdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstList `json:"l4-type-src-dst-list"`
    AppTypeSrcDstList []DdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstList `json:"app-type-src-dst-list"`
    ClassListOverflowPolicyList []DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListTemplate struct {
    Logging string `json:"logging"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Bypass int `json:"bypass"`
    ExceedLogCfg DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListExceedLogCfg `json:"exceed-log-cfg"`
    LogPeriodic int `json:"log-periodic"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListTemplate `json:"template"`
    Glid string `json:"glid"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    L4TypeSrcDstOverflowList []DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList `json:"l4-type-src-dst-overflow-list"`
    AppTypeSrcDstOverflowList []DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList `json:"app-type-src-dst-overflow-list"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListTemplate struct {
    Logging string `json:"logging"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}

func (p *DdosDstEntrySrcDstPairPolicy) GetId() string{
    return url.QueryEscape(p.Inst.SrcBasedPolicyName)
}

func (p *DdosDstEntrySrcDstPairPolicy) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/src-dst-pair-policy"
}

func (p *DdosDstEntrySrcDstPairPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicy::Post")
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

func (p *DdosDstEntrySrcDstPairPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicy::Get")
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
func (p *DdosDstEntrySrcDstPairPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicy::Put")
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

func (p *DdosDstEntrySrcDstPairPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
