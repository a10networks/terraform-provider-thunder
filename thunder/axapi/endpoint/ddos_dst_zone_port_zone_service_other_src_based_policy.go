

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceOtherSrcBasedPolicy struct {
	Inst struct {

    PolicyClassListList []DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListList `json:"policy-class-list-list"`
    SrcBasedPolicyName string `json:"src-based-policy-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    PortOther string 
    ZoneName string 
    Protocol string 

	} `json:"src-based-policy"`
}


type DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListList struct {
    ClassListName string `json:"class-list-name"`
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    Action string `json:"action"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    ZoneTemplate DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListListSamplingEnable `json:"sampling-enable"`
    ClassListOverflowPolicyList []DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}


type DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListListZoneTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
    Logging string `json:"logging"`
}


type DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    ZoneTemplate DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
    Logging string `json:"logging"`
}

func (p *DdosDstZonePortZoneServiceOtherSrcBasedPolicy) GetId() string{
    return url.QueryEscape(p.Inst.SrcBasedPolicyName)
}

func (p *DdosDstZonePortZoneServiceOtherSrcBasedPolicy) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/port/zone-service-other/" +p.Inst.PortOther + "+" +p.Inst.Protocol + "/src-based-policy"
}

func (p *DdosDstZonePortZoneServiceOtherSrcBasedPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherSrcBasedPolicy::Post")
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

func (p *DdosDstZonePortZoneServiceOtherSrcBasedPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherSrcBasedPolicy::Get")
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
func (p *DdosDstZonePortZoneServiceOtherSrcBasedPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherSrcBasedPolicy::Put")
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

func (p *DdosDstZonePortZoneServiceOtherSrcBasedPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherSrcBasedPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
