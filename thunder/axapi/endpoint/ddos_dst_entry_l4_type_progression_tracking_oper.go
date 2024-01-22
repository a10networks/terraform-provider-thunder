

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryL4TypeProgressionTrackingOper struct {
    
    Oper DdosDstEntryL4TypeProgressionTrackingOperOper `json:"oper"`
    Protocol string 
    DstEntryName string 

}
type DataDdosDstEntryL4TypeProgressionTrackingOper struct {
    DtDdosDstEntryL4TypeProgressionTrackingOper DdosDstEntryL4TypeProgressionTrackingOper `json:"progression-tracking"`
}


type DdosDstEntryL4TypeProgressionTrackingOperOper struct {
    Indicators []DdosDstEntryL4TypeProgressionTrackingOperOperIndicators `json:"indicators"`
}


type DdosDstEntryL4TypeProgressionTrackingOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}

func (p *DdosDstEntryL4TypeProgressionTrackingOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryL4TypeProgressionTrackingOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/l4-type/" +p.Protocol + "/progression-tracking/oper"
}

func (p *DdosDstEntryL4TypeProgressionTrackingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryL4TypeProgressionTrackingOper,error) {
logger.Println("DdosDstEntryL4TypeProgressionTrackingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryL4TypeProgressionTrackingOper
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
