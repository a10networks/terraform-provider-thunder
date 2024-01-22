

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceProgressionTrackingOper struct {
    
    Oper DdosDstZonePortZoneServiceProgressionTrackingOperOper `json:"oper"`
    ZoneName string 
    PortNum string 
    Protocol string 

}
type DataDdosDstZonePortZoneServiceProgressionTrackingOper struct {
    DtDdosDstZonePortZoneServiceProgressionTrackingOper DdosDstZonePortZoneServiceProgressionTrackingOper `json:"progression-tracking"`
}


type DdosDstZonePortZoneServiceProgressionTrackingOperOper struct {
    Indicators []DdosDstZonePortZoneServiceProgressionTrackingOperOperIndicators `json:"indicators"`
    LearningDetails int `json:"learning-details"`
    LearningBrief int `json:"learning-brief"`
    RecommendedTemplate int `json:"recommended-template"`
    TemplateDebugTable int `json:"template-debug-table"`
}


type DdosDstZonePortZoneServiceProgressionTrackingOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}

func (p *DdosDstZonePortZoneServiceProgressionTrackingOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceProgressionTrackingOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service/" +p.PortNum + "+" +p.Protocol + "/progression-tracking/oper"
}

func (p *DdosDstZonePortZoneServiceProgressionTrackingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceProgressionTrackingOper,error) {
logger.Println("DdosDstZonePortZoneServiceProgressionTrackingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceProgressionTrackingOper
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
