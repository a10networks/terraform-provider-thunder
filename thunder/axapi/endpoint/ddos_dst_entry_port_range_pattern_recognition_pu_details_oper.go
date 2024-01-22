

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortRangePatternRecognitionPuDetailsOper struct {
    
    Oper DdosDstEntryPortRangePatternRecognitionPuDetailsOperOper `json:"oper"`
    Protocol string 
    PortRangeEnd string 
    PortRangeStart string 
    DstEntryName string 

}
type DataDdosDstEntryPortRangePatternRecognitionPuDetailsOper struct {
    DtDdosDstEntryPortRangePatternRecognitionPuDetailsOper DdosDstEntryPortRangePatternRecognitionPuDetailsOper `json:"pattern-recognition-pu-details"`
}


type DdosDstEntryPortRangePatternRecognitionPuDetailsOperOper struct {
    AllFilters []DdosDstEntryPortRangePatternRecognitionPuDetailsOperOperAllFilters `json:"all-filters"`
}


type DdosDstEntryPortRangePatternRecognitionPuDetailsOperOperAllFilters struct {
    ProcessingUnit string `json:"processing-unit"`
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstEntryPortRangePatternRecognitionPuDetailsOperOperAllFiltersFilterList `json:"filter-list"`
}


type DdosDstEntryPortRangePatternRecognitionPuDetailsOperOperAllFiltersFilterList struct {
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}

func (p *DdosDstEntryPortRangePatternRecognitionPuDetailsOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortRangePatternRecognitionPuDetailsOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port-range/" +p.PortRangeStart + "+" +p.PortRangeEnd + "+" +p.Protocol + "/pattern-recognition-pu-details/oper"
}

func (p *DdosDstEntryPortRangePatternRecognitionPuDetailsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortRangePatternRecognitionPuDetailsOper,error) {
logger.Println("DdosDstEntryPortRangePatternRecognitionPuDetailsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortRangePatternRecognitionPuDetailsOper
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
