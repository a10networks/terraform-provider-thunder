

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneDetectionVictimIpDetectionOper struct {
    
    Oper DdosDstZoneDetectionVictimIpDetectionOperOper `json:"oper"`
    ZoneName string 

}
type DataDdosDstZoneDetectionVictimIpDetectionOper struct {
    DtDdosDstZoneDetectionVictimIpDetectionOper DdosDstZoneDetectionVictimIpDetectionOper `json:"victim-ip-detection"`
}


type DdosDstZoneDetectionVictimIpDetectionOperOper struct {
    IpEntryList []DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryList `json:"ip-entry-list"`
    IpEntryCount int `json:"ip-entry-count"`
    TotalIpEntryCount int `json:"total-ip-entry-count"`
    ActiveList int `json:"active-list"`
    VictimList int `json:"victim-list"`
    Ipv4Ip string `json:"ipv4-ip"`
    Ipv6Ip string `json:"ipv6-ip"`
}


type DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryList struct {
    IpAddressStr string `json:"ip-address-str"`
    Indicators []DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicators `json:"indicators"`
    IsLearningDone int `json:"is-learning-done"`
    IsHistogramLearningDone int `json:"is-histogram-learning-done"`
    IsIpAnomaly int `json:"is-ip-anomaly"`
    Is_static_threshold int `json:"is_static_threshold"`
    EscalationTimestamp string `json:"escalation-timestamp"`
    DeEscalationTimestamp string `json:"de-escalation-timestamp"`
}


type DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Value []DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicatorsValue `json:"value"`
    IsAnomaly int `json:"is-anomaly"`
}


type DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicatorsValue struct {
    Current string `json:"current"`
    Threshold string `json:"threshold"`
}

func (p *DdosDstZoneDetectionVictimIpDetectionOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneDetectionVictimIpDetectionOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/detection/victim-ip-detection/oper"
}

func (p *DdosDstZoneDetectionVictimIpDetectionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneDetectionVictimIpDetectionOper,error) {
logger.Println("DdosDstZoneDetectionVictimIpDetectionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneDetectionVictimIpDetectionOper
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
