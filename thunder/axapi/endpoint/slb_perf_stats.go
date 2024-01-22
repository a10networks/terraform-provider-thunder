

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbPerfStats struct {
    
    Stats SlbPerfStatsStats `json:"stats"`

}
type DataSlbPerfStats struct {
    DtSlbPerfStats SlbPerfStats `json:"perf"`
}


type SlbPerfStatsStats struct {
    TotalThroughputBitsPerSec int `json:"total-throughput-bits-per-sec"`
    L4ConnsPerSec int `json:"l4-conns-per-sec"`
    L7ConnsPerSec int `json:"l7-conns-per-sec"`
    L7TransPerSec int `json:"l7-trans-per-sec"`
    SslConnsPerSec int `json:"ssl-conns-per-sec"`
    IpNatConnsPerSec int `json:"ip-nat-conns-per-sec"`
    TotalNewConnsPerSec int `json:"total-new-conns-per-sec"`
    TotalCurrConns int `json:"total-curr-conns"`
    L4Bandwidth int `json:"l4-bandwidth"`
    L7Bandwidth int `json:"l7-bandwidth"`
    ServSslConnsPerSec int `json:"serv-ssl-conns-per-sec"`
    FwConnsPerSec int `json:"fw-conns-per-sec"`
    GifwConnsPerSec int `json:"gifw-conns-per-sec"`
    L7ProxyConnsPerSec int `json:"l7-proxy-conns-per-sec"`
    L7ProxyTransPerSec int `json:"l7-proxy-trans-per-sec"`
}

func (p *SlbPerfStats) GetId() string{
    return "1"
}

func (p *SlbPerfStats) getPath() string{
    return "slb/perf/stats"
}

func (p *SlbPerfStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbPerfStats,error) {
logger.Println("SlbPerfStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbPerfStats
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
