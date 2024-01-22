

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNameProgressionTrackingOper struct {
    
    Oper DdosDstZoneIpProtoProtoNameProgressionTrackingOperOper `json:"oper"`
    ZoneName string 
    Protocol string 

}
type DataDdosDstZoneIpProtoProtoNameProgressionTrackingOper struct {
    DtDdosDstZoneIpProtoProtoNameProgressionTrackingOper DdosDstZoneIpProtoProtoNameProgressionTrackingOper `json:"progression-tracking"`
}


type DdosDstZoneIpProtoProtoNameProgressionTrackingOperOper struct {
    Indicators []DdosDstZoneIpProtoProtoNameProgressionTrackingOperOperIndicators `json:"indicators"`
    LearningDetails int `json:"learning-details"`
    LearningBrief int `json:"learning-brief"`
    RecommendedTemplate int `json:"recommended-template"`
    TemplateDebugTable int `json:"template-debug-table"`
}


type DdosDstZoneIpProtoProtoNameProgressionTrackingOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}

func (p *DdosDstZoneIpProtoProtoNameProgressionTrackingOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNameProgressionTrackingOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-name/" +p.Protocol + "/progression-tracking/oper"
}

func (p *DdosDstZoneIpProtoProtoNameProgressionTrackingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNameProgressionTrackingOper,error) {
logger.Println("DdosDstZoneIpProtoProtoNameProgressionTrackingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNameProgressionTrackingOper
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
