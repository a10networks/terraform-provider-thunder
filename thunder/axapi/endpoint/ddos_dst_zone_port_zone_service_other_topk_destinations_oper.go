

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceOtherTopkDestinationsOper struct {
    
    Oper DdosDstZonePortZoneServiceOtherTopkDestinationsOperOper `json:"oper"`
    PortOther string 
    ZoneName string 
    Protocol string 

}
type DataDdosDstZonePortZoneServiceOtherTopkDestinationsOper struct {
    DtDdosDstZonePortZoneServiceOtherTopkDestinationsOper DdosDstZonePortZoneServiceOtherTopkDestinationsOper `json:"topk-destinations"`
}


type DdosDstZonePortZoneServiceOtherTopkDestinationsOperOper struct {
    Indicators []DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicators `json:"indicators"`
    EntryList []DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Destinations []DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicatorsDestinations `json:"destinations"`
}


type DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicatorsDestinations struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZonePortZoneServiceOtherTopkDestinationsOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceOtherTopkDestinationsOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service-other/" +p.PortOther + "+" +p.Protocol + "/topk-destinations/oper"
}

func (p *DdosDstZonePortZoneServiceOtherTopkDestinationsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceOtherTopkDestinationsOper,error) {
logger.Println("DdosDstZonePortZoneServiceOtherTopkDestinationsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceOtherTopkDestinationsOper
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
