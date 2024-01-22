

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortRangeTopkSourcesOper struct {
    
    Oper DdosDstZonePortRangeTopkSourcesOperOper `json:"oper"`
    Protocol string 
    ZoneName string 
    PortRangeEnd string 
    PortRangeStart string 

}
type DataDdosDstZonePortRangeTopkSourcesOper struct {
    DtDdosDstZonePortRangeTopkSourcesOper DdosDstZonePortRangeTopkSourcesOper `json:"topk-sources"`
}


type DdosDstZonePortRangeTopkSourcesOperOper struct {
    Indicators []DdosDstZonePortRangeTopkSourcesOperOperIndicators `json:"indicators"`
    EntryList []DdosDstZonePortRangeTopkSourcesOperOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZonePortRangeTopkSourcesOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstZonePortRangeTopkSourcesOperOperIndicatorsSources `json:"sources"`
}


type DdosDstZonePortRangeTopkSourcesOperOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZonePortRangeTopkSourcesOperOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZonePortRangeTopkSourcesOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstZonePortRangeTopkSourcesOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZonePortRangeTopkSourcesOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortRangeTopkSourcesOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port-range/" +p.PortRangeStart + "+" +p.PortRangeEnd + "+" +p.Protocol + "/topk-sources/oper"
}

func (p *DdosDstZonePortRangeTopkSourcesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortRangeTopkSourcesOper,error) {
logger.Println("DdosDstZonePortRangeTopkSourcesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortRangeTopkSourcesOper
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
