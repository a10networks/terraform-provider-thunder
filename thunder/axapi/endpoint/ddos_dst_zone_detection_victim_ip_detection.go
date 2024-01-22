

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneDetectionVictimIpDetection struct {
	Inst struct {

    Configuration string `json:"configuration"`
    HistogramToggle string `json:"histogram-toggle" dval:"histogram-disable"`
    IndicatorList []DdosDstZoneDetectionVictimIpDetectionIndicatorList `json:"indicator-list"`
    Toggle string `json:"toggle" dval:"disable"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"victim-ip-detection"`
}


type DdosDstZoneDetectionVictimIpDetectionIndicatorList struct {
    Type string `json:"type"`
    IpThresholdNum int `json:"ip-threshold-num"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DdosDstZoneDetectionVictimIpDetection) GetId() string{
    return "1"
}

func (p *DdosDstZoneDetectionVictimIpDetection) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/detection/victim-ip-detection"
}

func (p *DdosDstZoneDetectionVictimIpDetection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionVictimIpDetection::Post")
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

func (p *DdosDstZoneDetectionVictimIpDetection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionVictimIpDetection::Get")
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
func (p *DdosDstZoneDetectionVictimIpDetection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionVictimIpDetection::Put")
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

func (p *DdosDstZoneDetectionVictimIpDetection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionVictimIpDetection::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
