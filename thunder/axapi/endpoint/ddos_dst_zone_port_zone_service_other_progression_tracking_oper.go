

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceOtherProgressionTrackingOper struct {
    
    Oper DdosDstZonePortZoneServiceOtherProgressionTrackingOperOper `json:"oper"`
    PortOther string 
    ZoneName string 
    Protocol string 

}
type DataDdosDstZonePortZoneServiceOtherProgressionTrackingOper struct {
    DtDdosDstZonePortZoneServiceOtherProgressionTrackingOper DdosDstZonePortZoneServiceOtherProgressionTrackingOper `json:"progression-tracking"`
}


type DdosDstZonePortZoneServiceOtherProgressionTrackingOperOper struct {
    Indicators []DdosDstZonePortZoneServiceOtherProgressionTrackingOperOperIndicators `json:"indicators"`
    LearningDetails int `json:"learning-details"`
    LearningBrief int `json:"learning-brief"`
    RecommendedTemplate int `json:"recommended-template"`
    TemplateDebugTable int `json:"template-debug-table"`
}


type DdosDstZonePortZoneServiceOtherProgressionTrackingOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}

func (p *DdosDstZonePortZoneServiceOtherProgressionTrackingOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceOtherProgressionTrackingOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service-other/" +p.PortOther + "+" +p.Protocol + "/progression-tracking/oper"
}

func (p *DdosDstZonePortZoneServiceOtherProgressionTrackingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceOtherProgressionTrackingOper,error) {
logger.Println("DdosDstZonePortZoneServiceOtherProgressionTrackingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceOtherProgressionTrackingOper
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
