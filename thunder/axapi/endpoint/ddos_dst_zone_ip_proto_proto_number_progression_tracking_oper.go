

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNumberProgressionTrackingOper struct {
    
    Oper DdosDstZoneIpProtoProtoNumberProgressionTrackingOperOper `json:"oper"`
    ZoneName string 
    ProtocolNum string 

}
type DataDdosDstZoneIpProtoProtoNumberProgressionTrackingOper struct {
    DtDdosDstZoneIpProtoProtoNumberProgressionTrackingOper DdosDstZoneIpProtoProtoNumberProgressionTrackingOper `json:"progression-tracking"`
}


type DdosDstZoneIpProtoProtoNumberProgressionTrackingOperOper struct {
    Indicators []DdosDstZoneIpProtoProtoNumberProgressionTrackingOperOperIndicators `json:"indicators"`
    LearningDetails int `json:"learning-details"`
    LearningBrief int `json:"learning-brief"`
    RecommendedTemplate int `json:"recommended-template"`
    TemplateDebugTable int `json:"template-debug-table"`
}


type DdosDstZoneIpProtoProtoNumberProgressionTrackingOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}

func (p *DdosDstZoneIpProtoProtoNumberProgressionTrackingOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNumberProgressionTrackingOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-number/" +p.ProtocolNum + "/progression-tracking/oper"
}

func (p *DdosDstZoneIpProtoProtoNumberProgressionTrackingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNumberProgressionTrackingOper,error) {
logger.Println("DdosDstZoneIpProtoProtoNumberProgressionTrackingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNumberProgressionTrackingOper
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
