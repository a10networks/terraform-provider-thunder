

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNumberLevelIndicator struct {
	Inst struct {

    DataPacketSize int `json:"data-packet-size"`
    Score int `json:"score"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
    SrcViolationActions string `json:"src-violation-actions"`
    Type string `json:"type"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
    ZoneViolationActions string `json:"zone-violation-actions"`
    ZoneName string 
    LevelNum string 
    ProtocolNum string 

	} `json:"indicator"`
}

func (p *DdosDstZoneIpProtoProtoNumberLevelIndicator) GetId() string{
    return p.Inst.Type
}

func (p *DdosDstZoneIpProtoProtoNumberLevelIndicator) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/ip-proto/proto-number/" +p.Inst.ProtocolNum + "/level/" +p.Inst.LevelNum + "/indicator"
}

func (p *DdosDstZoneIpProtoProtoNumberLevelIndicator) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumberLevelIndicator::Post")
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

func (p *DdosDstZoneIpProtoProtoNumberLevelIndicator) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumberLevelIndicator::Get")
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
func (p *DdosDstZoneIpProtoProtoNumberLevelIndicator) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumberLevelIndicator::Put")
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

func (p *DdosDstZoneIpProtoProtoNumberLevelIndicator) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumberLevelIndicator::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
