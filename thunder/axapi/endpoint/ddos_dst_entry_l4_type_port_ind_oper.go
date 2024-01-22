

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryL4TypePortIndOper struct {
    
    Oper DdosDstEntryL4TypePortIndOperOper `json:"oper"`
    Protocol string 
    DstEntryName string 

}
type DataDdosDstEntryL4TypePortIndOper struct {
    DtDdosDstEntryL4TypePortIndOper DdosDstEntryL4TypePortIndOper `json:"port-ind"`
}


type DdosDstEntryL4TypePortIndOperOper struct {
    Indicators []DdosDstEntryL4TypePortIndOperOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
}


type DdosDstEntryL4TypePortIndOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    EntryMaximum string `json:"entry-maximum"`
    EntryMinimum string `json:"entry-minimum"`
    EntryNonZeroMinimum string `json:"entry-non-zero-minimum"`
    EntryAverage string `json:"entry-average"`
    SrcMaximum string `json:"src-maximum"`
}

func (p *DdosDstEntryL4TypePortIndOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryL4TypePortIndOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/l4-type/" +p.Protocol + "/port-ind/oper"
}

func (p *DdosDstEntryL4TypePortIndOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryL4TypePortIndOper,error) {
logger.Println("DdosDstEntryL4TypePortIndOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryL4TypePortIndOper
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
