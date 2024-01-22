

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortProgressionTrackingOper struct {
    
    Oper DdosDstEntryPortProgressionTrackingOperOper `json:"oper"`
    Protocol string 
    PortNum string 
    DstEntryName string 

}
type DataDdosDstEntryPortProgressionTrackingOper struct {
    DtDdosDstEntryPortProgressionTrackingOper DdosDstEntryPortProgressionTrackingOper `json:"progression-tracking"`
}


type DdosDstEntryPortProgressionTrackingOperOper struct {
    Indicators []DdosDstEntryPortProgressionTrackingOperOperIndicators `json:"indicators"`
}


type DdosDstEntryPortProgressionTrackingOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}

func (p *DdosDstEntryPortProgressionTrackingOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortProgressionTrackingOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port/" +p.PortNum + "+" +p.Protocol + "/progression-tracking/oper"
}

func (p *DdosDstEntryPortProgressionTrackingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortProgressionTrackingOper,error) {
logger.Println("DdosDstEntryPortProgressionTrackingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortProgressionTrackingOper
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
