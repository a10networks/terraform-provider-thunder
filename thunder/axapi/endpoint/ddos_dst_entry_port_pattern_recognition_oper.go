

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortPatternRecognitionOper struct {
    
    Oper DdosDstEntryPortPatternRecognitionOperOper `json:"oper"`
    Protocol string 
    PortNum string 
    DstEntryName string 

}
type DataDdosDstEntryPortPatternRecognitionOper struct {
    DtDdosDstEntryPortPatternRecognitionOper DdosDstEntryPortPatternRecognitionOper `json:"pattern-recognition"`
}


type DdosDstEntryPortPatternRecognitionOperOper struct {
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstEntryPortPatternRecognitionOperOperFilterList `json:"filter-list"`
}


type DdosDstEntryPortPatternRecognitionOperOperFilterList struct {
    ProcessingUnit string `json:"processing-unit"`
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}

func (p *DdosDstEntryPortPatternRecognitionOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortPatternRecognitionOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port/" +p.PortNum + "+" +p.Protocol + "/pattern-recognition/oper"
}

func (p *DdosDstEntryPortPatternRecognitionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortPatternRecognitionOper,error) {
logger.Println("DdosDstEntryPortPatternRecognitionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortPatternRecognitionOper
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
