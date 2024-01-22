

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServicePatternRecognition struct {
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
    ZoneName string 
    PortNum string 
    Protocol string 

	} `json:"pattern-recognition"`
}

func (p *DdosDstZonePortZoneServicePatternRecognition) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServicePatternRecognition) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/port/zone-service/" +p.Inst.PortNum + "+" +p.Inst.Protocol + "/pattern-recognition"
}

func (p *DdosDstZonePortZoneServicePatternRecognition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServicePatternRecognition::Post")
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

func (p *DdosDstZonePortZoneServicePatternRecognition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServicePatternRecognition::Get")
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
func (p *DdosDstZonePortZoneServicePatternRecognition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServicePatternRecognition::Put")
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

func (p *DdosDstZonePortZoneServicePatternRecognition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServicePatternRecognition::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
