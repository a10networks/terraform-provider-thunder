

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsCaaRecordStats struct {
    
    CriticalFlag int `json:"critical-flag"`
    PropertyTag string `json:"property-tag"`
    Rdata string `json:"rdata"`
    Stats GslbZoneServiceDnsCaaRecordStatsStats `json:"stats"`
    ServicePort string 
    ServiceName string 
    Name string 

}
type DataGslbZoneServiceDnsCaaRecordStats struct {
    DtGslbZoneServiceDnsCaaRecordStats GslbZoneServiceDnsCaaRecordStats `json:"dns-caa-record"`
}


type GslbZoneServiceDnsCaaRecordStatsStats struct {
    Hits int `json:"hits"`
}

func (p *GslbZoneServiceDnsCaaRecordStats) GetId() string{
    return "1"
}

func (p *GslbZoneServiceDnsCaaRecordStats) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +p.ServicePort + "+" +p.ServiceName + "/dns-caa-record/" +strconv.Itoa(p.CriticalFlag)+"+"+p.PropertyTag+"+"+p.Rdata+"/stats"
}

func (p *GslbZoneServiceDnsCaaRecordStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceDnsCaaRecordStats,error) {
logger.Println("GslbZoneServiceDnsCaaRecordStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceDnsCaaRecordStats
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
