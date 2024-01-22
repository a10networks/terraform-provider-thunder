

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneDetectionPacketAnomalyDetection struct {
	Inst struct {

    Configuration string `json:"configuration"`
    IndicatorList []DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList `json:"indicator-list"`
    Toggle string `json:"toggle" dval:"enable"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"packet-anomaly-detection"`
}


type DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList struct {
    Type string `json:"type"`
    ThresholdNum int `json:"threshold-num" dval:"100"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DdosDstZoneDetectionPacketAnomalyDetection) GetId() string{
    return "1"
}

func (p *DdosDstZoneDetectionPacketAnomalyDetection) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/detection/packet-anomaly-detection"
}

func (p *DdosDstZoneDetectionPacketAnomalyDetection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionPacketAnomalyDetection::Post")
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

func (p *DdosDstZoneDetectionPacketAnomalyDetection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionPacketAnomalyDetection::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *DdosDstZoneDetectionPacketAnomalyDetection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionPacketAnomalyDetection::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *DdosDstZoneDetectionPacketAnomalyDetection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionPacketAnomalyDetection::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
