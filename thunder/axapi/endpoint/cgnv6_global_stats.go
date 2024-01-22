

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6GlobalStats struct {
    
    Stats Cgnv6GlobalStatsStats `json:"stats"`

}
type DataCgnv6GlobalStats struct {
    DtCgnv6GlobalStats Cgnv6GlobalStats `json:"global"`
}


type Cgnv6GlobalStatsStats struct {
    TcpTotalPortsAllocated int `json:"tcp-total-ports-allocated"`
    UdpTotalPortsAllocated int `json:"udp-total-ports-allocated"`
    IcmpTotalPortsAllocated int `json:"icmp-total-ports-allocated"`
}

func (p *Cgnv6GlobalStats) GetId() string{
    return "1"
}

func (p *Cgnv6GlobalStats) getPath() string{
    return "cgnv6/global/stats"
}

func (p *Cgnv6GlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6GlobalStats,error) {
logger.Println("Cgnv6GlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6GlobalStats
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
