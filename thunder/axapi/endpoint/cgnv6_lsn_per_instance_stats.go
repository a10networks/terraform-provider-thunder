

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnPerInstanceStats struct {
    
    Stats Cgnv6LsnPerInstanceStatsStats `json:"stats"`

}
type DataCgnv6LsnPerInstanceStats struct {
    DtCgnv6LsnPerInstanceStats Cgnv6LsnPerInstanceStats `json:"per-instance"`
}


type Cgnv6LsnPerInstanceStatsStats struct {
    KeyName string `json:"key-name"`
    Data_session_created int `json:"data_session_created"`
    Data_session_freed int `json:"data_session_freed"`
    Tcp_fullcone_created int `json:"tcp_fullcone_created"`
    Tcp_fullcone_freed int `json:"tcp_fullcone_freed"`
    Udp_fullcone_created int `json:"udp_fullcone_created"`
    Udp_fullcone_freed int `json:"udp_fullcone_freed"`
    User_quota_created int `json:"user_quota_created"`
    User_quota_put_in_del_q int `json:"user_quota_put_in_del_q"`
    Tcp_allocated int `json:"tcp_allocated"`
    Tcp_freed int `json:"tcp_freed"`
    Udp_allocated int `json:"udp_allocated"`
    Udp_freed int `json:"udp_freed"`
    Icmp_allocated int `json:"icmp_allocated"`
    Icmp_freed int `json:"icmp_freed"`
}

func (p *Cgnv6LsnPerInstanceStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnPerInstanceStats) getPath() string{
    return "cgnv6/lsn/per-instance/stats"
}

func (p *Cgnv6LsnPerInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnPerInstanceStats,error) {
logger.Println("Cgnv6LsnPerInstanceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnPerInstanceStats
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
