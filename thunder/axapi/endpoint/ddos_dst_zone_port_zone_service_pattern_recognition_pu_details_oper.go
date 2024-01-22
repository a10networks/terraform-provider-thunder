

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServicePatternRecognitionPuDetailsOper struct {
    
    Oper DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOper `json:"oper"`
    ZoneName string 
    PortNum string 
    Protocol string 

}
type DataDdosDstZonePortZoneServicePatternRecognitionPuDetailsOper struct {
    DtDdosDstZonePortZoneServicePatternRecognitionPuDetailsOper DdosDstZonePortZoneServicePatternRecognitionPuDetailsOper `json:"pattern-recognition-pu-details"`
}


type DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOper struct {
    AllFilters []DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFilters `json:"all-filters"`
}


type DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFilters struct {
    ProcessingUnit string `json:"processing-unit"`
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFiltersFilterList `json:"filter-list"`
}


type DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFiltersFilterList struct {
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}

func (p *DdosDstZonePortZoneServicePatternRecognitionPuDetailsOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServicePatternRecognitionPuDetailsOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service/" +p.PortNum + "+" +p.Protocol + "/pattern-recognition-pu-details/oper"
}

func (p *DdosDstZonePortZoneServicePatternRecognitionPuDetailsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServicePatternRecognitionPuDetailsOper,error) {
logger.Println("DdosDstZonePortZoneServicePatternRecognitionPuDetailsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServicePatternRecognitionPuDetailsOper
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
