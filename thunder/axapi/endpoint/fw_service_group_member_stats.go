

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type FwServiceGroupMemberStats struct {
    
    Name string `json:"name"`
    Port int `json:"port"`
    Stats FwServiceGroupMemberStatsStats `json:"stats"`

}
type DataFwServiceGroupMemberStats struct {
    DtFwServiceGroupMemberStats FwServiceGroupMemberStats `json:"member"`
}


type FwServiceGroupMemberStatsStats struct {
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

func (p *FwServiceGroupMemberStats) GetId() string{
    return "1"
}

func (p *FwServiceGroupMemberStats) getPath() string{
    
    return "fw/service-group/"+p.Name+"/member/"+p.Name+"+" +strconv.Itoa(p.Port)+"/stats"
}

func (p *FwServiceGroupMemberStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwServiceGroupMemberStats,error) {
logger.Println("FwServiceGroupMemberStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwServiceGroupMemberStats
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
