

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosNetworkObjectStats struct {
    
    ObjectName string `json:"object-name"`
    Stats DdosNetworkObjectStatsStats `json:"stats"`

}
type DataDdosNetworkObjectStats struct {
    DtDdosNetworkObjectStats DdosNetworkObjectStats `json:"network-object"`
}


type DdosNetworkObjectStatsStats struct {
    Subnet_learned int `json:"subnet_learned"`
    Subnet_aged int `json:"subnet_aged"`
    Subnet_create_fail int `json:"subnet_create_fail"`
    Ip_learned int `json:"ip_learned"`
    Ip_aged int `json:"ip_aged"`
    Ip_create_fail int `json:"ip_create_fail"`
    Service_learned int `json:"service_learned"`
    Service_aged int `json:"service_aged"`
    Service_create_fail int `json:"service_create_fail"`
}

func (p *DdosNetworkObjectStats) GetId() string{
    return "1"
}

func (p *DdosNetworkObjectStats) getPath() string{
    
    return "ddos/network-object/"+p.ObjectName+"/stats"
}

func (p *DdosNetworkObjectStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosNetworkObjectStats,error) {
logger.Println("DdosNetworkObjectStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosNetworkObjectStats
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
