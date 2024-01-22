

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortTopkSourcesOper struct {
    
    Oper DdosDstEntryPortTopkSourcesOperOper `json:"oper"`
    Protocol string 
    PortNum string 
    DstEntryName string 

}
type DataDdosDstEntryPortTopkSourcesOper struct {
    DtDdosDstEntryPortTopkSourcesOper DdosDstEntryPortTopkSourcesOper `json:"topk-sources"`
}


type DdosDstEntryPortTopkSourcesOperOper struct {
    Indicators []DdosDstEntryPortTopkSourcesOperOperIndicators `json:"indicators"`
    EntryList []DdosDstEntryPortTopkSourcesOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstEntryPortTopkSourcesOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstEntryPortTopkSourcesOperOperIndicatorsSources `json:"sources"`
}


type DdosDstEntryPortTopkSourcesOperOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstEntryPortTopkSourcesOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstEntryPortTopkSourcesOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstEntryPortTopkSourcesOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstEntryPortTopkSourcesOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortTopkSourcesOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port/" +p.PortNum + "+" +p.Protocol + "/topk-sources/oper"
}

func (p *DdosDstEntryPortTopkSourcesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortTopkSourcesOper,error) {
logger.Println("DdosDstEntryPortTopkSourcesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortTopkSourcesOper
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
