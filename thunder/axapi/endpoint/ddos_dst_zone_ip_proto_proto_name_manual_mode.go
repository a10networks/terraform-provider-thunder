

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNameManualMode struct {
	Inst struct {

    Config string `json:"config"`
    GlidAction string `json:"glid-action"`
    SrcDefaultGlid string `json:"src-default-glid"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneTemplate DdosDstZoneIpProtoProtoNameManualModeZoneTemplate `json:"zone-template"`
    ZoneName string 
    Protocol string 

	} `json:"manual-mode"`
}


type DdosDstZoneIpProtoProtoNameManualModeZoneTemplate struct {
    IcmpV4 string `json:"icmp-v4"`
    IcmpV6 string `json:"icmp-v6"`
    IpProto string `json:"ip-proto"`
    Encap string `json:"encap"`
}

func (p *DdosDstZoneIpProtoProtoNameManualMode) GetId() string{
    return p.Inst.Config
}

func (p *DdosDstZoneIpProtoProtoNameManualMode) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/ip-proto/proto-name/" +p.Inst.Protocol + "/manual-mode"
}

func (p *DdosDstZoneIpProtoProtoNameManualMode) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameManualMode::Post")
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

func (p *DdosDstZoneIpProtoProtoNameManualMode) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameManualMode::Get")
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
func (p *DdosDstZoneIpProtoProtoNameManualMode) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameManualMode::Put")
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

func (p *DdosDstZoneIpProtoProtoNameManualMode) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameManualMode::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
