

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceTopkDestinationsOper struct {
    
    Oper DdosDstZonePortZoneServiceTopkDestinationsOperOper `json:"oper"`
    ZoneName string 
    PortNum string 
    Protocol string 

}
type DataDdosDstZonePortZoneServiceTopkDestinationsOper struct {
    DtDdosDstZonePortZoneServiceTopkDestinationsOper DdosDstZonePortZoneServiceTopkDestinationsOper `json:"topk-destinations"`
}


type DdosDstZonePortZoneServiceTopkDestinationsOperOper struct {
    Indicators []DdosDstZonePortZoneServiceTopkDestinationsOperOperIndicators `json:"indicators"`
    EntryList []DdosDstZonePortZoneServiceTopkDestinationsOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZonePortZoneServiceTopkDestinationsOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Destinations []DdosDstZonePortZoneServiceTopkDestinationsOperOperIndicatorsDestinations `json:"destinations"`
}


type DdosDstZonePortZoneServiceTopkDestinationsOperOperIndicatorsDestinations struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZonePortZoneServiceTopkDestinationsOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZonePortZoneServiceTopkDestinationsOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstZonePortZoneServiceTopkDestinationsOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZonePortZoneServiceTopkDestinationsOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceTopkDestinationsOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service/" +p.PortNum + "+" +p.Protocol + "/topk-destinations/oper"
}

func (p *DdosDstZonePortZoneServiceTopkDestinationsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceTopkDestinationsOper,error) {
logger.Println("DdosDstZonePortZoneServiceTopkDestinationsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceTopkDestinationsOper
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
