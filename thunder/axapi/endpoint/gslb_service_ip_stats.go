

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbServiceIpStats struct {
    
    NodeName string `json:"node-name"`
    PortList []GslbServiceIpStatsPortList `json:"port-list"`
    Stats GslbServiceIpStatsStats `json:"stats"`

}
type DataGslbServiceIpStats struct {
    DtGslbServiceIpStats GslbServiceIpStats `json:"service-ip"`
}


type GslbServiceIpStatsPortList struct {
    PortNum int `json:"port-num"`
    PortProto string `json:"port-proto"`
    Stats GslbServiceIpStatsPortListStats `json:"stats"`
}


type GslbServiceIpStatsPortListStats struct {
    Active int `json:"active"`
    Current int `json:"current"`
}


type GslbServiceIpStatsStats struct {
    Hits int `json:"hits"`
    Recent int `json:"recent"`
}

func (p *GslbServiceIpStats) GetId() string{
    return "1"
}

func (p *GslbServiceIpStats) getPath() string{
    
    return "gslb/service-ip/"+p.NodeName+"/stats"
}

func (p *GslbServiceIpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbServiceIpStats,error) {
logger.Println("GslbServiceIpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbServiceIpStats
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
