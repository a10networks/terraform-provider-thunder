

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNameTopkSourcesOper struct {
    
    Oper DdosDstZoneIpProtoProtoNameTopkSourcesOperOper `json:"oper"`
    ZoneName string 
    Protocol string 

}
type DataDdosDstZoneIpProtoProtoNameTopkSourcesOper struct {
    DtDdosDstZoneIpProtoProtoNameTopkSourcesOper DdosDstZoneIpProtoProtoNameTopkSourcesOper `json:"topk-sources"`
}


type DdosDstZoneIpProtoProtoNameTopkSourcesOperOper struct {
    Indicators []DdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicators `json:"indicators"`
    EntryList []DdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicatorsSources `json:"sources"`
}


type DdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZoneIpProtoProtoNameTopkSourcesOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNameTopkSourcesOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-name/" +p.Protocol + "/topk-sources/oper"
}

func (p *DdosDstZoneIpProtoProtoNameTopkSourcesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNameTopkSourcesOper,error) {
logger.Println("DdosDstZoneIpProtoProtoNameTopkSourcesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNameTopkSourcesOper
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
