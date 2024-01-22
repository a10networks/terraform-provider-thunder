

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryL4TypeTopkSourcesOper struct {
    
    Oper DdosDstEntryL4TypeTopkSourcesOperOper `json:"oper"`
    Protocol string 
    DstEntryName string 

}
type DataDdosDstEntryL4TypeTopkSourcesOper struct {
    DtDdosDstEntryL4TypeTopkSourcesOper DdosDstEntryL4TypeTopkSourcesOper `json:"topk-sources"`
}


type DdosDstEntryL4TypeTopkSourcesOperOper struct {
    Indicators []DdosDstEntryL4TypeTopkSourcesOperOperIndicators `json:"indicators"`
    EntryList []DdosDstEntryL4TypeTopkSourcesOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstEntryL4TypeTopkSourcesOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstEntryL4TypeTopkSourcesOperOperIndicatorsSources `json:"sources"`
}


type DdosDstEntryL4TypeTopkSourcesOperOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstEntryL4TypeTopkSourcesOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstEntryL4TypeTopkSourcesOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstEntryL4TypeTopkSourcesOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstEntryL4TypeTopkSourcesOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryL4TypeTopkSourcesOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/l4-type/" +p.Protocol + "/topk-sources/oper"
}

func (p *DdosDstEntryL4TypeTopkSourcesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryL4TypeTopkSourcesOper,error) {
logger.Println("DdosDstEntryL4TypeTopkSourcesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryL4TypeTopkSourcesOper
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
