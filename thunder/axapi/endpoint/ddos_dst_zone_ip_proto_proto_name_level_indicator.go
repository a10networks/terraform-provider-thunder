

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNameLevelIndicator struct {
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
    Protocol string 

	} `json:"indicator"`
}

func (p *DdosDstZoneIpProtoProtoNameLevelIndicator) GetId() string{
    return p.Inst.Type
}

func (p *DdosDstZoneIpProtoProtoNameLevelIndicator) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/ip-proto/proto-name/" +p.Inst.Protocol + "/level/" +p.Inst.LevelNum + "/indicator"
}

func (p *DdosDstZoneIpProtoProtoNameLevelIndicator) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameLevelIndicator::Post")
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

func (p *DdosDstZoneIpProtoProtoNameLevelIndicator) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameLevelIndicator::Get")
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
func (p *DdosDstZoneIpProtoProtoNameLevelIndicator) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameLevelIndicator::Put")
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

func (p *DdosDstZoneIpProtoProtoNameLevelIndicator) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNameLevelIndicator::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
