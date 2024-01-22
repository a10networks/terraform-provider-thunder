

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNumberTopkSourcesOper struct {
    
    Oper DdosDstZoneIpProtoProtoNumberTopkSourcesOperOper `json:"oper"`
    ZoneName string 
    ProtocolNum string 

}
type DataDdosDstZoneIpProtoProtoNumberTopkSourcesOper struct {
    DtDdosDstZoneIpProtoProtoNumberTopkSourcesOper DdosDstZoneIpProtoProtoNumberTopkSourcesOper `json:"topk-sources"`
}


type DdosDstZoneIpProtoProtoNumberTopkSourcesOperOper struct {
    Indicators []DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicators `json:"indicators"`
    EntryList []DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicatorsSources `json:"sources"`
}


type DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZoneIpProtoProtoNumberTopkSourcesOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNumberTopkSourcesOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-number/" +p.ProtocolNum + "/topk-sources/oper"
}

func (p *DdosDstZoneIpProtoProtoNumberTopkSourcesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNumberTopkSourcesOper,error) {
logger.Println("DdosDstZoneIpProtoProtoNumberTopkSourcesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNumberTopkSourcesOper
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
