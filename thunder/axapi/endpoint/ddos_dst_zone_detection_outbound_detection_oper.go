

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneDetectionOutboundDetectionOper struct {
    
    Oper DdosDstZoneDetectionOutboundDetectionOperOper `json:"oper"`
    TopkSourceSubnet DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnet `json:"topk-source-subnet"`
    ZoneName string 

}
type DataDdosDstZoneDetectionOutboundDetectionOper struct {
    DtDdosDstZoneDetectionOutboundDetectionOper DdosDstZoneDetectionOutboundDetectionOper `json:"outbound-detection"`
}


type DdosDstZoneDetectionOutboundDetectionOperOper struct {
    DiscoveryTimestamp string `json:"discovery-timestamp"`
    EntryList []DdosDstZoneDetectionOutboundDetectionOperOperEntryList `json:"entry-list"`
}


type DdosDstZoneDetectionOutboundDetectionOperOperEntryList struct {
    LocationType string `json:"location-type"`
    LocationName string `json:"location-name"`
    Indicators []DdosDstZoneDetectionOutboundDetectionOperOperEntryListIndicators `json:"indicators"`
    DataSource string `json:"data-source"`
    Anomaly string `json:"anomaly"`
    AnomalyTimestamp string `json:"anomaly-timestamp"`
    InitialLearning string `json:"initial-learning"`
    ActiveTime int `json:"active-time"`
}


type DdosDstZoneDetectionOutboundDetectionOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    NonZeroMinimum string `json:"non-zero-minimum"`
    Average string `json:"average"`
    AdaptiveThreshold string `json:"adaptive-threshold"`
}


type DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnet struct {
    Oper DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOper `json:"oper"`
}


type DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOper struct {
    EntryList []DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryList `json:"entry-list"`
}


type DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryList struct {
    LocationType string `json:"location-type"`
    LocationName string `json:"location-name"`
    Indicators []DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicators `json:"indicators"`
}


type DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    SourceSubnets []DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicatorsSourceSubnets `json:"source-subnets"`
}


type DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicatorsSourceSubnets struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}

func (p *DdosDstZoneDetectionOutboundDetectionOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneDetectionOutboundDetectionOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/detection/outbound-detection/oper"
}

func (p *DdosDstZoneDetectionOutboundDetectionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneDetectionOutboundDetectionOper,error) {
logger.Println("DdosDstZoneDetectionOutboundDetectionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneDetectionOutboundDetectionOper
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
