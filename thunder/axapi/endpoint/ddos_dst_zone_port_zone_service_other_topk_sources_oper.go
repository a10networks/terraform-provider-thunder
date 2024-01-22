

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceOtherTopkSourcesOper struct {
    
    Oper DdosDstZonePortZoneServiceOtherTopkSourcesOperOper `json:"oper"`
    PortOther string 
    ZoneName string 
    Protocol string 

}
type DataDdosDstZonePortZoneServiceOtherTopkSourcesOper struct {
    DtDdosDstZonePortZoneServiceOtherTopkSourcesOper DdosDstZonePortZoneServiceOtherTopkSourcesOper `json:"topk-sources"`
}


type DdosDstZonePortZoneServiceOtherTopkSourcesOperOper struct {
    Indicators []DdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicators `json:"indicators"`
    EntryList []DdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicatorsSources `json:"sources"`
}


type DdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZonePortZoneServiceOtherTopkSourcesOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceOtherTopkSourcesOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service-other/" +p.PortOther + "+" +p.Protocol + "/topk-sources/oper"
}

func (p *DdosDstZonePortZoneServiceOtherTopkSourcesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceOtherTopkSourcesOper,error) {
logger.Println("DdosDstZonePortZoneServiceOtherTopkSourcesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceOtherTopkSourcesOper
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
