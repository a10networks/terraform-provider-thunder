

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortRangeSrcBasedPolicy struct {
	Inst struct {

    PolicyClassListList []DdosDstZonePortRangeSrcBasedPolicyPolicyClassListList `json:"policy-class-list-list"`
    SrcBasedPolicyName string `json:"src-based-policy-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneName string 
    Protocol string 
    PortRangeStart string 
    PortRangeEnd string 

	} `json:"src-based-policy"`
}


type DdosDstZonePortRangeSrcBasedPolicyPolicyClassListList struct {
    ClassListName string `json:"class-list-name"`
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    Action string `json:"action"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    ZoneTemplate DdosDstZonePortRangeSrcBasedPolicyPolicyClassListListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []DdosDstZonePortRangeSrcBasedPolicyPolicyClassListListSamplingEnable `json:"sampling-enable"`
    ClassListOverflowPolicyList []DdosDstZonePortRangeSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}


type DdosDstZonePortRangeSrcBasedPolicyPolicyClassListListZoneTemplate struct {
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


type DdosDstZonePortRangeSrcBasedPolicyPolicyClassListListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZonePortRangeSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    ZoneTemplate DdosDstZonePortRangeSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZonePortRangeSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
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

func (p *DdosDstZonePortRangeSrcBasedPolicy) GetId() string{
    return url.QueryEscape(p.Inst.SrcBasedPolicyName)
}

func (p *DdosDstZonePortRangeSrcBasedPolicy) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/port-range/" +p.Inst.PortRangeStart + "+" +p.Inst.PortRangeEnd + "+" +p.Inst.Protocol + "/src-based-policy"
}

func (p *DdosDstZonePortRangeSrcBasedPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangeSrcBasedPolicy::Post")
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

func (p *DdosDstZonePortRangeSrcBasedPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangeSrcBasedPolicy::Get")
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
func (p *DdosDstZonePortRangeSrcBasedPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangeSrcBasedPolicy::Put")
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

func (p *DdosDstZonePortRangeSrcBasedPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangeSrcBasedPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
