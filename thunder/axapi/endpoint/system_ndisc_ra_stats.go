

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemNdiscRaStats struct {
    
    Stats SystemNdiscRaStatsStats `json:"stats"`

}
type DataSystemNdiscRaStats struct {
    DtSystemNdiscRaStats SystemNdiscRaStats `json:"ndisc-ra"`
}


type SystemNdiscRaStatsStats struct {
    Good_recv int `json:"good_recv"`
    Periodic_sent int `json:"periodic_sent"`
    Rate_limit int `json:"rate_limit"`
    Bad_hop_limit int `json:"bad_hop_limit"`
    Truncated int `json:"truncated"`
    Bad_icmpv6_csum int `json:"bad_icmpv6_csum"`
    Bad_icmpv6_code int `json:"bad_icmpv6_code"`
    Bad_icmpv6_option int `json:"bad_icmpv6_option"`
    L2_addr_and_unspec int `json:"l2_addr_and_unspec"`
    No_free_buffers int `json:"no_free_buffers"`
}

func (p *SystemNdiscRaStats) GetId() string{
    return "1"
}

func (p *SystemNdiscRaStats) getPath() string{
    return "system/ndisc-ra/stats"
}

func (p *SystemNdiscRaStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemNdiscRaStats,error) {
logger.Println("SystemNdiscRaStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemNdiscRaStats
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
