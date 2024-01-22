

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServicePatternRecognitionOper struct {
    
    Oper DdosDstZonePortZoneServicePatternRecognitionOperOper `json:"oper"`
    ZoneName string 
    PortNum string 
    Protocol string 

}
type DataDdosDstZonePortZoneServicePatternRecognitionOper struct {
    DtDdosDstZonePortZoneServicePatternRecognitionOper DdosDstZonePortZoneServicePatternRecognitionOper `json:"pattern-recognition"`
}


type DdosDstZonePortZoneServicePatternRecognitionOperOper struct {
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstZonePortZoneServicePatternRecognitionOperOperFilterList `json:"filter-list"`
}


type DdosDstZonePortZoneServicePatternRecognitionOperOperFilterList struct {
    ProcessingUnit string `json:"processing-unit"`
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}

func (p *DdosDstZonePortZoneServicePatternRecognitionOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServicePatternRecognitionOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service/" +p.PortNum + "+" +p.Protocol + "/pattern-recognition/oper"
}

func (p *DdosDstZonePortZoneServicePatternRecognitionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServicePatternRecognitionOper,error) {
logger.Println("DdosDstZonePortZoneServicePatternRecognitionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServicePatternRecognitionOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
