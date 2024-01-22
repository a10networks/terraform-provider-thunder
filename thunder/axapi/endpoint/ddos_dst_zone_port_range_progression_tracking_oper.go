

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortRangeProgressionTrackingOper struct {
    
    Oper DdosDstZonePortRangeProgressionTrackingOperOper `json:"oper"`
    Protocol string 
    ZoneName string 
    PortRangeEnd string 
    PortRangeStart string 

}
type DataDdosDstZonePortRangeProgressionTrackingOper struct {
    DtDdosDstZonePortRangeProgressionTrackingOper DdosDstZonePortRangeProgressionTrackingOper `json:"progression-tracking"`
}


type DdosDstZonePortRangeProgressionTrackingOperOper struct {
    Indicators []DdosDstZonePortRangeProgressionTrackingOperOperIndicators `json:"indicators"`
    LearningDetails int `json:"learning-details"`
    LearningBrief int `json:"learning-brief"`
    RecommendedTemplate int `json:"recommended-template"`
    TemplateDebugTable int `json:"template-debug-table"`
}


type DdosDstZonePortRangeProgressionTrackingOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}

func (p *DdosDstZonePortRangeProgressionTrackingOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortRangeProgressionTrackingOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port-range/" +p.PortRangeStart + "+" +p.PortRangeEnd + "+" +p.Protocol + "/progression-tracking/oper"
}

func (p *DdosDstZonePortRangeProgressionTrackingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortRangeProgressionTrackingOper,error) {
logger.Println("DdosDstZonePortRangeProgressionTrackingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortRangeProgressionTrackingOper
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
