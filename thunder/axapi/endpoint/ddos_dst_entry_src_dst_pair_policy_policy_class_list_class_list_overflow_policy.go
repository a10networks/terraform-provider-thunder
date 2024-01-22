

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy struct {
	Inst struct {

    AppTypeSrcDstOverflowList []DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowList `json:"app-type-src-dst-overflow-list"`
    Bypass int `json:"bypass"`
    DummyName string `json:"dummy-name"`
    ExceedLogCfg DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyExceedLogCfg `json:"exceed-log-cfg"`
    Glid string `json:"glid"`
    L4TypeSrcDstOverflowList []DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowList `json:"l4-type-src-dst-overflow-list"`
    LogPeriodic int `json:"log-periodic"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    SrcBasedPolicyName string 
    ClassListName string 
    DstEntryName string 

	} `json:"class-list-overflow-policy"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyTemplate struct {
    Logging string `json:"logging"`
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy) GetId() string{
    return p.Inst.DummyName
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/src-dst-pair-policy/" +p.Inst.SrcBasedPolicyName + "/policy-class-list/" +p.Inst.ClassListName + "/class-list-overflow-policy"
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy::Post")
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

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy::Get")
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
func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy::Put")
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

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
