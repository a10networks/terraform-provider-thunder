

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortRangePortIndOper struct {
    
    Oper DdosDstEntryPortRangePortIndOperOper `json:"oper"`
    Protocol string 
    PortRangeEnd string 
    PortRangeStart string 
    DstEntryName string 

}
type DataDdosDstEntryPortRangePortIndOper struct {
    DtDdosDstEntryPortRangePortIndOper DdosDstEntryPortRangePortIndOper `json:"port-ind"`
}


type DdosDstEntryPortRangePortIndOperOper struct {
    Indicators []DdosDstEntryPortRangePortIndOperOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
}


type DdosDstEntryPortRangePortIndOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    EntryMaximum string `json:"entry-maximum"`
    EntryMinimum string `json:"entry-minimum"`
    EntryNonZeroMinimum string `json:"entry-non-zero-minimum"`
    EntryAverage string `json:"entry-average"`
    SrcMaximum string `json:"src-maximum"`
}

func (p *DdosDstEntryPortRangePortIndOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortRangePortIndOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port-range/" +p.PortRangeStart + "+" +p.PortRangeEnd + "+" +p.Protocol + "/port-ind/oper"
}

func (p *DdosDstEntryPortRangePortIndOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortRangePortIndOper,error) {
logger.Println("DdosDstEntryPortRangePortIndOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortRangePortIndOper
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
