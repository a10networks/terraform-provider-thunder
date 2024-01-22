

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy struct {
	Inst struct {

    Action string `json:"action"`
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneTemplate DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyZoneTemplate `json:"zone-template"`
    ZoneName string 
    ClassListName string 
    Protocol string 
    SrcBasedPolicyName string 

	} `json:"class-list-overflow-policy"`
}


type DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyZoneTemplate struct {
    IcmpV4 string `json:"icmp-v4"`
    IcmpV6 string `json:"icmp-v6"`
    IpProto string `json:"ip-proto"`
    Encap string `json:"encap"`
}

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy) GetId() string{
    return p.Inst.DummyName
}

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/ip-proto/proto-name/" +p.Inst.Protocol + "/src-based-policy/" +p.Inst.SrcBasedPolicyName + "/policy-class-list/" +p.Inst.ClassListName + "/class-list-overflow-policy"
}

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy::Post")
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

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy::Get")
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
func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy::Put")
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

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
