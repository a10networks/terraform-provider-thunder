

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type GslbServiceIpPortStats struct {
    
    PortNum int `json:"port-num"`
    PortProto string `json:"port-proto"`
    Stats GslbServiceIpPortStatsStats `json:"stats"`
    NodeName string 

}
type DataGslbServiceIpPortStats struct {
    DtGslbServiceIpPortStats GslbServiceIpPortStats `json:"port"`
}


type GslbServiceIpPortStatsStats struct {
    Active int `json:"active"`
    Current int `json:"current"`
}

func (p *GslbServiceIpPortStats) GetId() string{
    return "1"
}

func (p *GslbServiceIpPortStats) getPath() string{
    
    return "gslb/service-ip/" +p.NodeName + "/port/" +strconv.Itoa(p.PortNum)+"+"+p.PortProto+"/stats"
}

func (p *GslbServiceIpPortStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbServiceIpPortStats,error) {
logger.Println("GslbServiceIpPortStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbServiceIpPortStats
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
