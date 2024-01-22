

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceOtherPatternRecognitionOper struct {
    
    Oper DdosDstZonePortZoneServiceOtherPatternRecognitionOperOper `json:"oper"`
    PortOther string 
    ZoneName string 
    Protocol string 

}
type DataDdosDstZonePortZoneServiceOtherPatternRecognitionOper struct {
    DtDdosDstZonePortZoneServiceOtherPatternRecognitionOper DdosDstZonePortZoneServiceOtherPatternRecognitionOper `json:"pattern-recognition"`
}


type DdosDstZonePortZoneServiceOtherPatternRecognitionOperOper struct {
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstZonePortZoneServiceOtherPatternRecognitionOperOperFilterList `json:"filter-list"`
}


type DdosDstZonePortZoneServiceOtherPatternRecognitionOperOperFilterList struct {
    ProcessingUnit string `json:"processing-unit"`
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}

func (p *DdosDstZonePortZoneServiceOtherPatternRecognitionOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceOtherPatternRecognitionOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service-other/" +p.PortOther + "+" +p.Protocol + "/pattern-recognition/oper"
}

func (p *DdosDstZonePortZoneServiceOtherPatternRecognitionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceOtherPatternRecognitionOper,error) {
logger.Println("DdosDstZonePortZoneServiceOtherPatternRecognitionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceOtherPatternRecognitionOper
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
