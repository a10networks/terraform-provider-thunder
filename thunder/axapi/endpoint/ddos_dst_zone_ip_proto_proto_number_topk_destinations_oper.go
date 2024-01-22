

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNumberTopkDestinationsOper struct {
    
    Oper DdosDstZoneIpProtoProtoNumberTopkDestinationsOperOper `json:"oper"`
    ZoneName string 
    ProtocolNum string 

}
type DataDdosDstZoneIpProtoProtoNumberTopkDestinationsOper struct {
    DtDdosDstZoneIpProtoProtoNumberTopkDestinationsOper DdosDstZoneIpProtoProtoNumberTopkDestinationsOper `json:"topk-destinations"`
}


type DdosDstZoneIpProtoProtoNumberTopkDestinationsOperOper struct {
    Indicators []DdosDstZoneIpProtoProtoNumberTopkDestinationsOperOperIndicators `json:"indicators"`
    EntryList []DdosDstZoneIpProtoProtoNumberTopkDestinationsOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZoneIpProtoProtoNumberTopkDestinationsOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Destinations []DdosDstZoneIpProtoProtoNumberTopkDestinationsOperOperIndicatorsDestinations `json:"destinations"`
}


type DdosDstZoneIpProtoProtoNumberTopkDestinationsOperOperIndicatorsDestinations struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZoneIpProtoProtoNumberTopkDestinationsOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZoneIpProtoProtoNumberTopkDestinationsOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstZoneIpProtoProtoNumberTopkDestinationsOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZoneIpProtoProtoNumberTopkDestinationsOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNumberTopkDestinationsOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-number/" +p.ProtocolNum + "/topk-destinations/oper"
}

func (p *DdosDstZoneIpProtoProtoNumberTopkDestinationsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNumberTopkDestinationsOper,error) {
logger.Println("DdosDstZoneIpProtoProtoNumberTopkDestinationsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNumberTopkDestinationsOper
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
