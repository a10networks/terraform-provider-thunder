

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsSrvRecordStats struct {
    
    Port int `json:"port"`
    SrvName string `json:"srv-name"`
    Stats GslbZoneServiceDnsSrvRecordStatsStats `json:"stats"`
    ServicePort string 
    ServiceName string 
    Name string 

}
type DataGslbZoneServiceDnsSrvRecordStats struct {
    DtGslbZoneServiceDnsSrvRecordStats GslbZoneServiceDnsSrvRecordStats `json:"dns-srv-record"`
}


type GslbZoneServiceDnsSrvRecordStatsStats struct {
    Hits int `json:"hits"`
}

func (p *GslbZoneServiceDnsSrvRecordStats) GetId() string{
    return "1"
}

func (p *GslbZoneServiceDnsSrvRecordStats) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +p.ServicePort + "+" +p.ServiceName + "/dns-srv-record/"+p.SrvName+"+" +strconv.Itoa(p.Port)+"/stats"
}

func (p *GslbZoneServiceDnsSrvRecordStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceDnsSrvRecordStats,error) {
logger.Println("GslbZoneServiceDnsSrvRecordStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceDnsSrvRecordStats
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
