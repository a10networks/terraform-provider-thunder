

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatPerInstanceStats struct {
    
    Stats Cgnv6FixedNatPerInstanceStatsStats `json:"stats"`

}
type DataCgnv6FixedNatPerInstanceStats struct {
    DtCgnv6FixedNatPerInstanceStats Cgnv6FixedNatPerInstanceStats `json:"per-instance"`
}


type Cgnv6FixedNatPerInstanceStatsStats struct {
    KeyName string `json:"key-name"`
    Data_session_created int `json:"data_session_created"`
    Data_session_freed int `json:"data_session_freed"`
    Tcp_fullcone_created int `json:"tcp_fullcone_created"`
    Tcp_fullcone_freed int `json:"tcp_fullcone_freed"`
    Udp_fullcone_created int `json:"udp_fullcone_created"`
    Udp_fullcone_freed int `json:"udp_fullcone_freed"`
    Tcp_allocated int `json:"tcp_allocated"`
    Tcp_freed int `json:"tcp_freed"`
    Udp_allocated int `json:"udp_allocated"`
    Udp_freed int `json:"udp_freed"`
    Icmp_allocated int `json:"icmp_allocated"`
    Icmp_freed int `json:"icmp_freed"`
    Total_nat_in_use int `json:"total_nat_in_use"`
    ActiveSubscriberAdded int `json:"active-subscriber-added"`
    ActiveSubscriberRemoved int `json:"active-subscriber-removed"`
}

func (p *Cgnv6FixedNatPerInstanceStats) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatPerInstanceStats) getPath() string{
    return "cgnv6/fixed-nat/per-instance/stats"
}

func (p *Cgnv6FixedNatPerInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatPerInstanceStats,error) {
logger.Println("Cgnv6FixedNatPerInstanceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatPerInstanceStats
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
