

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper struct {
    
    Oper DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOper `json:"oper"`
    ZoneName string 

}
type DataDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper struct {
    DtDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper `json:"topk-source-subnet"`
}


type DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOper struct {
    EntryList []DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryList `json:"entry-list"`
}


type DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryList struct {
    LocationType string `json:"location-type"`
    LocationName string `json:"location-name"`
    Indicators []DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicators `json:"indicators"`
}


type DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    SourceSubnets []DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicatorsSourceSubnets `json:"source-subnets"`
}


type DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicatorsSourceSubnets struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}

func (p *DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/detection/outbound-detection/topk-source-subnet/oper"
}

func (p *DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper,error) {
logger.Println("DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper
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
