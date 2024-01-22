

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortRangePatternRecognitionPuDetailsOper struct {
    
    Oper DdosDstZonePortRangePatternRecognitionPuDetailsOperOper `json:"oper"`
    Protocol string 
    ZoneName string 
    PortRangeEnd string 
    PortRangeStart string 

}
type DataDdosDstZonePortRangePatternRecognitionPuDetailsOper struct {
    DtDdosDstZonePortRangePatternRecognitionPuDetailsOper DdosDstZonePortRangePatternRecognitionPuDetailsOper `json:"pattern-recognition-pu-details"`
}


type DdosDstZonePortRangePatternRecognitionPuDetailsOperOper struct {
    AllFilters []DdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFilters `json:"all-filters"`
}


type DdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFilters struct {
    ProcessingUnit string `json:"processing-unit"`
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFiltersFilterList `json:"filter-list"`
}


type DdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFiltersFilterList struct {
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}

func (p *DdosDstZonePortRangePatternRecognitionPuDetailsOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortRangePatternRecognitionPuDetailsOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port-range/" +p.PortRangeStart + "+" +p.PortRangeEnd + "+" +p.Protocol + "/pattern-recognition-pu-details/oper"
}

func (p *DdosDstZonePortRangePatternRecognitionPuDetailsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortRangePatternRecognitionPuDetailsOper,error) {
logger.Println("DdosDstZonePortRangePatternRecognitionPuDetailsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortRangePatternRecognitionPuDetailsOper
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
