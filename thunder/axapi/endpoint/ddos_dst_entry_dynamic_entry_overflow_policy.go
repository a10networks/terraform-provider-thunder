

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryDynamicEntryOverflowPolicy struct {
	Inst struct {

    AppTypeSrcDstList []DdosDstEntryDynamicEntryOverflowPolicyAppTypeSrcDstList `json:"app-type-src-dst-list"`
    Bypass int `json:"bypass"`
    DummyName string `json:"dummy-name"`
    ExceedLogCfg DdosDstEntryDynamicEntryOverflowPolicyExceedLogCfg `json:"exceed-log-cfg"`
    Glid string `json:"glid"`
    L4TypeSrcDstList []DdosDstEntryDynamicEntryOverflowPolicyL4TypeSrcDstList `json:"l4-type-src-dst-list"`
    LogPeriodic int `json:"log-periodic"`
    Template DdosDstEntryDynamicEntryOverflowPolicyTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DstEntryName string 

	} `json:"dynamic-entry-overflow-policy"`
}


type DdosDstEntryDynamicEntryOverflowPolicyAppTypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntryDynamicEntryOverflowPolicyAppTypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntryDynamicEntryOverflowPolicyAppTypeSrcDstListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntryDynamicEntryOverflowPolicyExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntryDynamicEntryOverflowPolicyL4TypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntryDynamicEntryOverflowPolicyL4TypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntryDynamicEntryOverflowPolicyL4TypeSrcDstListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntryDynamicEntryOverflowPolicyTemplate struct {
    Logging string `json:"logging"`
}

func (p *DdosDstEntryDynamicEntryOverflowPolicy) GetId() string{
    return p.Inst.DummyName
}

func (p *DdosDstEntryDynamicEntryOverflowPolicy) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/dynamic-entry-overflow-policy"
}

func (p *DdosDstEntryDynamicEntryOverflowPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryDynamicEntryOverflowPolicy::Post")
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

func (p *DdosDstEntryDynamicEntryOverflowPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryDynamicEntryOverflowPolicy::Get")
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
func (p *DdosDstEntryDynamicEntryOverflowPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryDynamicEntryOverflowPolicy::Put")
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

func (p *DdosDstEntryDynamicEntryOverflowPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryDynamicEntryOverflowPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
