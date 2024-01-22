

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryTopkDestinationsOper struct {
    
    Oper DdosDstEntryTopkDestinationsOperOper `json:"oper"`
    DstEntryName string 

}
type DataDdosDstEntryTopkDestinationsOper struct {
    DtDdosDstEntryTopkDestinationsOper DdosDstEntryTopkDestinationsOper `json:"topk-destinations"`
}


type DdosDstEntryTopkDestinationsOperOper struct {
    Indicators []DdosDstEntryTopkDestinationsOperOperIndicators `json:"indicators"`
    EntryList []DdosDstEntryTopkDestinationsOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstEntryTopkDestinationsOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Destinations []DdosDstEntryTopkDestinationsOperOperIndicatorsDestinations `json:"destinations"`
}


type DdosDstEntryTopkDestinationsOperOperIndicatorsDestinations struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstEntryTopkDestinationsOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstEntryTopkDestinationsOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstEntryTopkDestinationsOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstEntryTopkDestinationsOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryTopkDestinationsOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/topk-destinations/oper"
}

func (p *DdosDstEntryTopkDestinationsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryTopkDestinationsOper,error) {
logger.Println("DdosDstEntryTopkDestinationsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryTopkDestinationsOper
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
