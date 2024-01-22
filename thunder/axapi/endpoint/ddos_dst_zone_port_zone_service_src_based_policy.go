

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceSrcBasedPolicy struct {
	Inst struct {

    PolicyClassListList []DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListList `json:"policy-class-list-list"`
    SrcBasedPolicyName string `json:"src-based-policy-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneName string 
    PortNum string 
    Protocol string 

	} `json:"src-based-policy"`
}


type DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListList struct {
    ClassListName string `json:"class-list-name"`
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    ZoneTemplate DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListSamplingEnable `json:"sampling-enable"`
    ClassListOverflowPolicyList []DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}


type DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListZoneTemplate struct {
    Quic string `json:"quic"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
    Logging string `json:"logging"`
}


type DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    ZoneTemplate DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
    Quic string `json:"quic"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
    Logging string `json:"logging"`
}

func (p *DdosDstZonePortZoneServiceSrcBasedPolicy) GetId() string{
    return url.QueryEscape(p.Inst.SrcBasedPolicyName)
}

func (p *DdosDstZonePortZoneServiceSrcBasedPolicy) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/port/zone-service/" +p.Inst.PortNum + "+" +p.Inst.Protocol + "/src-based-policy"
}

func (p *DdosDstZonePortZoneServiceSrcBasedPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceSrcBasedPolicy::Post")
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

func (p *DdosDstZonePortZoneServiceSrcBasedPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceSrcBasedPolicy::Get")
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
func (p *DdosDstZonePortZoneServiceSrcBasedPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceSrcBasedPolicy::Put")
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

func (p *DdosDstZonePortZoneServiceSrcBasedPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceSrcBasedPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
