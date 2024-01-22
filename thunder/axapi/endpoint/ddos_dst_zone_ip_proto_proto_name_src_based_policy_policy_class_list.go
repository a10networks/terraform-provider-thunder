

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassList struct {
	Inst struct {

    Action string `json:"action"`
    ClassListName string `json:"class-list-name"`
    ClassListOverflowPolicyList []DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    SamplingEnable []DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneTemplate DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListZoneTemplate `json:"zone-template"`
    SrcBasedPolicyName string 
    ZoneName string 
    Protocol string 

	} `json:"policy-class-list"`
}


type DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    ZoneTemplate DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyListZoneTemplate struct {
    IcmpV4 string `json:"icmp-v4"`
    IcmpV6 string `json:"icmp-v6"`
    IpProto string `json:"ip-proto"`
    Encap string `json:"encap"`
}


type DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListZoneTemplate struct {
    Logging string `json:"logging"`
    IcmpV4 string `json:"icmp-v4"`
    IcmpV6 string `json:"icmp-v6"`
    IpProto string `json:"ip-proto"`
    Encap string `json:"encap"`
}

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassList) GetId() string{
    return url.QueryEscape(p.Inst.ClassListName)
}

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassList) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/ip-proto/proto-name/" +p.Inst.Protocol + "/src-based-policy/" +p.Inst.SrcBasedPolicyName + "/policy-class-list"
}

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassList::Post")
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

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassList::Get")
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
func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassList::Put")
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

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
