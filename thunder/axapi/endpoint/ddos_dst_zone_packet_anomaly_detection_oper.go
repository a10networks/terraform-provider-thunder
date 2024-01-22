

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePacketAnomalyDetectionOper struct {
    
    Oper DdosDstZonePacketAnomalyDetectionOperOper `json:"oper"`
    ZoneName string 

}
type DataDdosDstZonePacketAnomalyDetectionOper struct {
    DtDdosDstZonePacketAnomalyDetectionOper DdosDstZonePacketAnomalyDetectionOper `json:"packet-anomaly-detection"`
}


type DdosDstZonePacketAnomalyDetectionOperOper struct {
    Indicators []DdosDstZonePacketAnomalyDetectionOperOperIndicators `json:"indicators"`
    DataSource string `json:"data-source"`
}


type DdosDstZonePacketAnomalyDetectionOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    Threshold string `json:"threshold"`
    IsAnomaly int `json:"is-anomaly"`
}

func (p *DdosDstZonePacketAnomalyDetectionOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePacketAnomalyDetectionOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/packet-anomaly-detection/oper"
}

func (p *DdosDstZonePacketAnomalyDetectionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePacketAnomalyDetectionOper,error) {
logger.Println("DdosDstZonePacketAnomalyDetectionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePacketAnomalyDetectionOper
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
