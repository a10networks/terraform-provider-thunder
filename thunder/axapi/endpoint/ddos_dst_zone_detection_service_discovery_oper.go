

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneDetectionServiceDiscoveryOper struct {
    
    Oper DdosDstZoneDetectionServiceDiscoveryOperOper `json:"oper"`
    ZoneName string 

}
type DataDdosDstZoneDetectionServiceDiscoveryOper struct {
    DtDdosDstZoneDetectionServiceDiscoveryOper DdosDstZoneDetectionServiceDiscoveryOper `json:"service-discovery"`
}


type DdosDstZoneDetectionServiceDiscoveryOperOper struct {
    DiscoveredServiceList []DdosDstZoneDetectionServiceDiscoveryOperOperDiscoveredServiceList `json:"discovered-service-list"`
}


type DdosDstZoneDetectionServiceDiscoveryOperOperDiscoveredServiceList struct {
    Port int `json:"port"`
    Protocol string `json:"protocol"`
    Rate int `json:"rate"`
}

func (p *DdosDstZoneDetectionServiceDiscoveryOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneDetectionServiceDiscoveryOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/detection/service-discovery/oper"
}

func (p *DdosDstZoneDetectionServiceDiscoveryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneDetectionServiceDiscoveryOper,error) {
logger.Println("DdosDstZoneDetectionServiceDiscoveryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneDetectionServiceDiscoveryOper
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
