

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortRangeProgressionTrackingOper struct {
    
    Oper DdosDstEntryPortRangeProgressionTrackingOperOper `json:"oper"`
    Protocol string 
    PortRangeEnd string 
    PortRangeStart string 
    DstEntryName string 

}
type DataDdosDstEntryPortRangeProgressionTrackingOper struct {
    DtDdosDstEntryPortRangeProgressionTrackingOper DdosDstEntryPortRangeProgressionTrackingOper `json:"progression-tracking"`
}


type DdosDstEntryPortRangeProgressionTrackingOperOper struct {
    Indicators []DdosDstEntryPortRangeProgressionTrackingOperOperIndicators `json:"indicators"`
}


type DdosDstEntryPortRangeProgressionTrackingOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}

func (p *DdosDstEntryPortRangeProgressionTrackingOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortRangeProgressionTrackingOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port-range/" +p.PortRangeStart + "+" +p.PortRangeEnd + "+" +p.Protocol + "/progression-tracking/oper"
}

func (p *DdosDstEntryPortRangeProgressionTrackingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortRangeProgressionTrackingOper,error) {
logger.Println("DdosDstEntryPortRangeProgressionTrackingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortRangeProgressionTrackingOper
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
