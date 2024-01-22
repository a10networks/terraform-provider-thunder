

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortRangeTopkDestinationsOper struct {
    
    Oper DdosDstZonePortRangeTopkDestinationsOperOper `json:"oper"`
    Protocol string 
    ZoneName string 
    PortRangeEnd string 
    PortRangeStart string 

}
type DataDdosDstZonePortRangeTopkDestinationsOper struct {
    DtDdosDstZonePortRangeTopkDestinationsOper DdosDstZonePortRangeTopkDestinationsOper `json:"topk-destinations"`
}


type DdosDstZonePortRangeTopkDestinationsOperOper struct {
    Indicators []DdosDstZonePortRangeTopkDestinationsOperOperIndicators `json:"indicators"`
    EntryList []DdosDstZonePortRangeTopkDestinationsOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZonePortRangeTopkDestinationsOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Destinations []DdosDstZonePortRangeTopkDestinationsOperOperIndicatorsDestinations `json:"destinations"`
}


type DdosDstZonePortRangeTopkDestinationsOperOperIndicatorsDestinations struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZonePortRangeTopkDestinationsOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZonePortRangeTopkDestinationsOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstZonePortRangeTopkDestinationsOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZonePortRangeTopkDestinationsOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortRangeTopkDestinationsOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port-range/" +p.PortRangeStart + "+" +p.PortRangeEnd + "+" +p.Protocol + "/topk-destinations/oper"
}

func (p *DdosDstZonePortRangeTopkDestinationsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortRangeTopkDestinationsOper,error) {
logger.Println("DdosDstZonePortRangeTopkDestinationsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortRangeTopkDestinationsOper
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
