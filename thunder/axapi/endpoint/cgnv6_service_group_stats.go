

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6ServiceGroupStats struct {
    
    MemberList []Cgnv6ServiceGroupStatsMemberList `json:"member-list"`
    Name string `json:"name"`
    Stats Cgnv6ServiceGroupStatsStats `json:"stats"`

}
type DataCgnv6ServiceGroupStats struct {
    DtCgnv6ServiceGroupStats Cgnv6ServiceGroupStats `json:"service-group"`
}


type Cgnv6ServiceGroupStatsMemberList struct {
    Name string `json:"name"`
    Port int `json:"port"`
    Stats Cgnv6ServiceGroupStatsMemberListStats `json:"stats"`
}


type Cgnv6ServiceGroupStatsMemberListStats struct {
    Curr_conn int `json:"curr_conn"`
    Total_fwd_bytes int `json:"total_fwd_bytes"`
    Total_fwd_pkts int `json:"total_fwd_pkts"`
    Total_rev_bytes int `json:"total_rev_bytes"`
    Total_rev_pkts int `json:"total_rev_pkts"`
    Total_conn int `json:"total_conn"`
    Total_rev_pkts_inspected int `json:"total_rev_pkts_inspected"`
    Total_rev_pkts_inspected_status_code_2xx int `json:"total_rev_pkts_inspected_status_code_2xx"`
    Total_rev_pkts_inspected_status_code_non_5xx int `json:"total_rev_pkts_inspected_status_code_non_5xx"`
    Curr_req int `json:"curr_req"`
    Total_req int `json:"total_req"`
    Total_req_succ int `json:"total_req_succ"`
    Peak_conn int `json:"peak_conn"`
    Response_time int `json:"response_time"`
    Fastest_rsp_time int `json:"fastest_rsp_time"`
    Slowest_rsp_time int `json:"slowest_rsp_time"`
    Curr_ssl_conn int `json:"curr_ssl_conn"`
    Total_ssl_conn int `json:"total_ssl_conn"`
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    State_flaps int `json:"state_flaps"`
}


type Cgnv6ServiceGroupStatsStats struct {
    Server_selection_fail_drop int `json:"server_selection_fail_drop"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Service_peak_conn int `json:"service_peak_conn"`
}

func (p *Cgnv6ServiceGroupStats) GetId() string{
    return "1"
}

func (p *Cgnv6ServiceGroupStats) getPath() string{
    
    return "cgnv6/service-group/"+p.Name+"/stats"
}

func (p *Cgnv6ServiceGroupStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6ServiceGroupStats,error) {
logger.Println("Cgnv6ServiceGroupStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6ServiceGroupStats
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
