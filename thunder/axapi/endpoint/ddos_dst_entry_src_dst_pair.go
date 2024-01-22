

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntrySrcDstPair struct {
	Inst struct {

    AppTypeSrcDstList []DdosDstEntrySrcDstPairAppTypeSrcDstList `json:"app-type-src-dst-list"`
    Bypass int `json:"bypass"`
    Default int `json:"default"`
    ExceedLogCfg DdosDstEntrySrcDstPairExceedLogCfg `json:"exceed-log-cfg"`
    Glid string `json:"glid"`
    L4TypeSrcDstList []DdosDstEntrySrcDstPairL4TypeSrcDstList `json:"l4-type-src-dst-list"`
    LogPeriodic int `json:"log-periodic"`
    Template DdosDstEntrySrcDstPairTemplate `json:"template"`
    Uuid string `json:"uuid"`
    DstEntryName string 

	} `json:"src-dst-pair"`
}


type DdosDstEntrySrcDstPairAppTypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairAppTypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairAppTypeSrcDstListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntrySrcDstPairExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairL4TypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairL4TypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairL4TypeSrcDstListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntrySrcDstPairTemplate struct {
    Logging string `json:"logging"`
}

func (p *DdosDstEntrySrcDstPair) GetId() string{
    return "1"
}

func (p *DdosDstEntrySrcDstPair) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/src-dst-pair"
}

func (p *DdosDstEntrySrcDstPair) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPair::Post")
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

func (p *DdosDstEntrySrcDstPair) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPair::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *DdosDstEntrySrcDstPair) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPair::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *DdosDstEntrySrcDstPair) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntrySrcDstPair::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
