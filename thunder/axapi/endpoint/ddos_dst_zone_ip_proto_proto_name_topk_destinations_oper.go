

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNameTopkDestinationsOper struct {
    
    Oper DdosDstZoneIpProtoProtoNameTopkDestinationsOperOper `json:"oper"`
    ZoneName string 
    Protocol string 

}
type DataDdosDstZoneIpProtoProtoNameTopkDestinationsOper struct {
    DtDdosDstZoneIpProtoProtoNameTopkDestinationsOper DdosDstZoneIpProtoProtoNameTopkDestinationsOper `json:"topk-destinations"`
}


type DdosDstZoneIpProtoProtoNameTopkDestinationsOperOper struct {
    Indicators []DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicators `json:"indicators"`
    EntryList []DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Destinations []DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicatorsDestinations `json:"destinations"`
}


type DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicatorsDestinations struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZoneIpProtoProtoNameTopkDestinationsOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNameTopkDestinationsOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-name/" +p.Protocol + "/topk-destinations/oper"
}

func (p *DdosDstZoneIpProtoProtoNameTopkDestinationsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNameTopkDestinationsOper,error) {
logger.Println("DdosDstZoneIpProtoProtoNameTopkDestinationsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNameTopkDestinationsOper
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
