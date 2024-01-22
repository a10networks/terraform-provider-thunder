

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortPatternRecognitionPuDetailsOper struct {
    
    Oper DdosDstEntryPortPatternRecognitionPuDetailsOperOper `json:"oper"`
    Protocol string 
    PortNum string 
    DstEntryName string 

}
type DataDdosDstEntryPortPatternRecognitionPuDetailsOper struct {
    DtDdosDstEntryPortPatternRecognitionPuDetailsOper DdosDstEntryPortPatternRecognitionPuDetailsOper `json:"pattern-recognition-pu-details"`
}


type DdosDstEntryPortPatternRecognitionPuDetailsOperOper struct {
    AllFilters []DdosDstEntryPortPatternRecognitionPuDetailsOperOperAllFilters `json:"all-filters"`
}


type DdosDstEntryPortPatternRecognitionPuDetailsOperOperAllFilters struct {
    ProcessingUnit string `json:"processing-unit"`
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstEntryPortPatternRecognitionPuDetailsOperOperAllFiltersFilterList `json:"filter-list"`
}


type DdosDstEntryPortPatternRecognitionPuDetailsOperOperAllFiltersFilterList struct {
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}

func (p *DdosDstEntryPortPatternRecognitionPuDetailsOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortPatternRecognitionPuDetailsOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port/" +p.PortNum + "+" +p.Protocol + "/pattern-recognition-pu-details/oper"
}

func (p *DdosDstEntryPortPatternRecognitionPuDetailsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortPatternRecognitionPuDetailsOper,error) {
logger.Println("DdosDstEntryPortPatternRecognitionPuDetailsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortPatternRecognitionPuDetailsOper
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
