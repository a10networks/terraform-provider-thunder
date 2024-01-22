

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntrySrcDstPairPolicyPolicyClassList struct {
	Inst struct {

    AppTypeSrcDstList []DdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstList `json:"app-type-src-dst-list"`
    Bypass int `json:"bypass"`
    ClassListName string `json:"class-list-name"`
    ClassListOverflowPolicyList []DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
    ExceedLogCfg DdosDstEntrySrcDstPairPolicyPolicyClassListExceedLogCfg `json:"exceed-log-cfg"`
    Glid string `json:"glid"`
    L4TypeSrcDstList []DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstList `json:"l4-type-src-dst-list"`
    LogPeriodic int `json:"log-periodic"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    SamplingEnable []DdosDstEntrySrcDstPairPolicyPolicyClassListSamplingEnable `json:"sampling-enable"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    SrcBasedPolicyName string 
    DstEntryName string 

	} `json:"policy-class-list"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Bypass int `json:"bypass"`
    ExceedLogCfg DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListExceedLogCfg `json:"exceed-log-cfg"`
    LogPeriodic int `json:"log-periodic"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListTemplate `json:"template"`
    Glid string `json:"glid"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    L4TypeSrcDstOverflowList []DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowList `json:"l4-type-src-dst-overflow-list"`
    AppTypeSrcDstOverflowList []DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowList `json:"app-type-src-dst-overflow-list"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListTemplate struct {
    Logging string `json:"logging"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListTemplate struct {
    Logging string `json:"logging"`
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassList) GetId() string{
    return url.QueryEscape(p.Inst.ClassListName)
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassList) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/src-dst-pair-policy/" +p.Inst.SrcBasedPolicyName + "/policy-class-list"
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassList::Post")
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

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassList::Get")
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
func (p *DdosDstEntrySrcDstPairPolicyPolicyClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassList::Put")
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

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
