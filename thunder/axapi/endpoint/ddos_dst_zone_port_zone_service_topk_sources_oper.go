

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceTopkSourcesOper struct {
    
    Oper DdosDstZonePortZoneServiceTopkSourcesOperOper `json:"oper"`
    ZoneName string 
    PortNum string 
    Protocol string 

}
type DataDdosDstZonePortZoneServiceTopkSourcesOper struct {
    DtDdosDstZonePortZoneServiceTopkSourcesOper DdosDstZonePortZoneServiceTopkSourcesOper `json:"topk-sources"`
}


type DdosDstZonePortZoneServiceTopkSourcesOperOper struct {
    Indicators []DdosDstZonePortZoneServiceTopkSourcesOperOperIndicators `json:"indicators"`
    EntryList []DdosDstZonePortZoneServiceTopkSourcesOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZonePortZoneServiceTopkSourcesOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstZonePortZoneServiceTopkSourcesOperOperIndicatorsSources `json:"sources"`
}


type DdosDstZonePortZoneServiceTopkSourcesOperOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZonePortZoneServiceTopkSourcesOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZonePortZoneServiceTopkSourcesOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstZonePortZoneServiceTopkSourcesOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZonePortZoneServiceTopkSourcesOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceTopkSourcesOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service/" +p.PortNum + "+" +p.Protocol + "/topk-sources/oper"
}

func (p *DdosDstZonePortZoneServiceTopkSourcesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceTopkSourcesOper,error) {
logger.Println("DdosDstZonePortZoneServiceTopkSourcesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceTopkSourcesOper
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
