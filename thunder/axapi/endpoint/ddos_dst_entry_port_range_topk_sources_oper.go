

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortRangeTopkSourcesOper struct {
    
    Oper DdosDstEntryPortRangeTopkSourcesOperOper `json:"oper"`
    Protocol string 
    PortRangeEnd string 
    PortRangeStart string 
    DstEntryName string 

}
type DataDdosDstEntryPortRangeTopkSourcesOper struct {
    DtDdosDstEntryPortRangeTopkSourcesOper DdosDstEntryPortRangeTopkSourcesOper `json:"topk-sources"`
}


type DdosDstEntryPortRangeTopkSourcesOperOper struct {
    Indicators []DdosDstEntryPortRangeTopkSourcesOperOperIndicators `json:"indicators"`
    EntryList []DdosDstEntryPortRangeTopkSourcesOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstEntryPortRangeTopkSourcesOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstEntryPortRangeTopkSourcesOperOperIndicatorsSources `json:"sources"`
}


type DdosDstEntryPortRangeTopkSourcesOperOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstEntryPortRangeTopkSourcesOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstEntryPortRangeTopkSourcesOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstEntryPortRangeTopkSourcesOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstEntryPortRangeTopkSourcesOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortRangeTopkSourcesOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port-range/" +p.PortRangeStart + "+" +p.PortRangeEnd + "+" +p.Protocol + "/topk-sources/oper"
}

func (p *DdosDstEntryPortRangeTopkSourcesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortRangeTopkSourcesOper,error) {
logger.Println("DdosDstEntryPortRangeTopkSourcesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortRangeTopkSourcesOper
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
