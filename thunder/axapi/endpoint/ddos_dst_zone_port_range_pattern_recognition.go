

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortRangePatternRecognition struct {
	Inst struct {

    Algorithm string `json:"algorithm"`
    AppPayloadOffset int `json:"app-payload-offset"`
    CaptureTraffic string `json:"capture-traffic"`
    FilterInactiveThreshold int `json:"filter-inactive-threshold"`
    FilterThreshold int `json:"filter-threshold"`
    Mode string `json:"mode"`
    Sensitivity string `json:"sensitivity"`
    TriggeredBy string `json:"triggered-by"`
    Uuid string `json:"uuid"`
    Protocol string 
    ZoneName string 
    PortRangeEnd string 
    PortRangeStart string 

	} `json:"pattern-recognition"`
}

func (p *DdosDstZonePortRangePatternRecognition) GetId() string{
    return "1"
}

func (p *DdosDstZonePortRangePatternRecognition) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/port-range/" +p.Inst.PortRangeStart + "+" +p.Inst.PortRangeEnd + "+" +p.Inst.Protocol + "/pattern-recognition"
}

func (p *DdosDstZonePortRangePatternRecognition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangePatternRecognition::Post")
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

func (p *DdosDstZonePortRangePatternRecognition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangePatternRecognition::Get")
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
func (p *DdosDstZonePortRangePatternRecognition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangePatternRecognition::Put")
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

func (p *DdosDstZonePortRangePatternRecognition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangePatternRecognition::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
