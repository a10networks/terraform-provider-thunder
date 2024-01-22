

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassList struct {
	Inst struct {

    Action string `json:"action"`
    ClassListName string `json:"class-list-name"`
    ClassListOverflowPolicyList []DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    SamplingEnable []DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneTemplate DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListZoneTemplate `json:"zone-template"`
    SrcBasedPolicyName string 
    ZoneName string 
    ProtocolNum string 

	} `json:"policy-class-list"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListClassListOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    ZoneTemplate DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListClassListOverflowPolicyListZoneTemplate struct {
    IpProto string `json:"ip-proto"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListZoneTemplate struct {
    Logging string `json:"logging"`
    IpProto string `json:"ip-proto"`
}

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassList) GetId() string{
    return url.QueryEscape(p.Inst.ClassListName)
}

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassList) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/ip-proto/proto-number/" +p.Inst.ProtocolNum + "/src-based-policy/" +p.Inst.SrcBasedPolicyName + "/policy-class-list"
}

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassList::Post")
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

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassList::Get")
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
func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassList::Put")
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

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
