

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneTopkDestinationsOper struct {
    
    Oper DdosDstZoneTopkDestinationsOperOper `json:"oper"`
    ZoneName string 

}
type DataDdosDstZoneTopkDestinationsOper struct {
    DtDdosDstZoneTopkDestinationsOper DdosDstZoneTopkDestinationsOper `json:"topk-destinations"`
}


type DdosDstZoneTopkDestinationsOperOper struct {
    Indicators []DdosDstZoneTopkDestinationsOperOperIndicators `json:"indicators"`
    EntryList []DdosDstZoneTopkDestinationsOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZoneTopkDestinationsOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Destinations []DdosDstZoneTopkDestinationsOperOperIndicatorsDestinations `json:"destinations"`
}


type DdosDstZoneTopkDestinationsOperOperIndicatorsDestinations struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZoneTopkDestinationsOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZoneTopkDestinationsOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstZoneTopkDestinationsOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZoneTopkDestinationsOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneTopkDestinationsOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/topk-destinations/oper"
}

func (p *DdosDstZoneTopkDestinationsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneTopkDestinationsOper,error) {
logger.Println("DdosDstZoneTopkDestinationsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneTopkDestinationsOper
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
