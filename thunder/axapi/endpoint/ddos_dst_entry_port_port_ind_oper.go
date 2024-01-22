

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortPortIndOper struct {
    
    Oper DdosDstEntryPortPortIndOperOper `json:"oper"`
    Protocol string 
    PortNum string 
    DstEntryName string 

}
type DataDdosDstEntryPortPortIndOper struct {
    DtDdosDstEntryPortPortIndOper DdosDstEntryPortPortIndOper `json:"port-ind"`
}


type DdosDstEntryPortPortIndOperOper struct {
    Indicators []DdosDstEntryPortPortIndOperOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
}


type DdosDstEntryPortPortIndOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    EntryMaximum string `json:"entry-maximum"`
    EntryMinimum string `json:"entry-minimum"`
    EntryNonZeroMinimum string `json:"entry-non-zero-minimum"`
    EntryAverage string `json:"entry-average"`
    SrcMaximum string `json:"src-maximum"`
}

func (p *DdosDstEntryPortPortIndOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortPortIndOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port/" +p.PortNum + "+" +p.Protocol + "/port-ind/oper"
}

func (p *DdosDstEntryPortPortIndOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortPortIndOper,error) {
logger.Println("DdosDstEntryPortPortIndOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortPortIndOper
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
