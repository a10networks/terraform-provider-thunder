

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntrySrcDstPairClassList struct {
	Inst struct {

    AppTypeSrcDstList []DdosDstEntrySrcDstPairClassListAppTypeSrcDstList `json:"app-type-src-dst-list"`
    CidList []DdosDstEntrySrcDstPairClassListCidList `json:"cid-list"`
    ClassListName string `json:"class-list-name"`
    ExceedLogCfg DdosDstEntrySrcDstPairClassListExceedLogCfg `json:"exceed-log-cfg"`
    L4TypeSrcDstList []DdosDstEntrySrcDstPairClassListL4TypeSrcDstList `json:"l4-type-src-dst-list"`
    LogPeriodic int `json:"log-periodic"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DstEntryName string 

	} `json:"src-dst-pair-class-list"`
}


type DdosDstEntrySrcDstPairClassListAppTypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairClassListAppTypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairClassListAppTypeSrcDstListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntrySrcDstPairClassListCidList struct {
    CidNum int `json:"cid-num"`
    ExceedLogCfg DdosDstEntrySrcDstPairClassListCidListExceedLogCfg `json:"exceed-log-cfg"`
    LogPeriodic int `json:"log-periodic"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    L4TypeSrcDstCidList []DdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidList `json:"l4-type-src-dst-cid-list"`
    AppTypeSrcDstCidList []DdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidList `json:"app-type-src-dst-cid-list"`
}


type DdosDstEntrySrcDstPairClassListCidListExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntrySrcDstPairClassListExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairClassListL4TypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairClassListL4TypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairClassListL4TypeSrcDstListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}

func (p *DdosDstEntrySrcDstPairClassList) GetId() string{
    return p.Inst.ClassListName
}

func (p *DdosDstEntrySrcDstPairClassList) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/src-dst-pair-class-list"
}

func (p *DdosDstEntrySrcDstPairClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairClassList::Post")
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

func (p *DdosDstEntrySrcDstPairClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairClassList::Get")
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
func (p *DdosDstEntrySrcDstPairClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairClassList::Put")
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

func (p *DdosDstEntrySrcDstPairClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPairClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
