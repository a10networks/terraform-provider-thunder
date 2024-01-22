

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNumberSrcBasedPolicy struct {
	Inst struct {

    PolicyClassListList []DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListList `json:"policy-class-list-list"`
    SrcBasedPolicyName string `json:"src-based-policy-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneName string 
    ProtocolNum string 

	} `json:"src-based-policy"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListList struct {
    ClassListName string `json:"class-list-name"`
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    ZoneTemplate DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListSamplingEnable `json:"sampling-enable"`
    ClassListOverflowPolicyList []DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListZoneTemplate struct {
    Logging string `json:"logging"`
    IpProto string `json:"ip-proto"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    ZoneTemplate DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
    IpProto string `json:"ip-proto"`
}

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicy) GetId() string{
    return url.QueryEscape(p.Inst.SrcBasedPolicyName)
}

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicy) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/ip-proto/proto-number/" +p.Inst.ProtocolNum + "/src-based-policy"
}

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumberSrcBasedPolicy::Post")
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

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumberSrcBasedPolicy::Get")
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
func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumberSrcBasedPolicy::Put")
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

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumberSrcBasedPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
