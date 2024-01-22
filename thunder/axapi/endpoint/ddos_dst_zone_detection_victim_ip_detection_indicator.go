

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneDetectionVictimIpDetectionIndicator struct {
	Inst struct {

    IpThresholdNum int `json:"ip-threshold-num"`
    Type string `json:"type"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"indicator"`
}

func (p *DdosDstZoneDetectionVictimIpDetectionIndicator) GetId() string{
    return p.Inst.Type
}

func (p *DdosDstZoneDetectionVictimIpDetectionIndicator) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/detection/victim-ip-detection/indicator"
}

func (p *DdosDstZoneDetectionVictimIpDetectionIndicator) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionVictimIpDetectionIndicator::Post")
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

func (p *DdosDstZoneDetectionVictimIpDetectionIndicator) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionVictimIpDetectionIndicator::Get")
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
func (p *DdosDstZoneDetectionVictimIpDetectionIndicator) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionVictimIpDetectionIndicator::Put")
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

func (p *DdosDstZoneDetectionVictimIpDetectionIndicator) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionVictimIpDetectionIndicator::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
