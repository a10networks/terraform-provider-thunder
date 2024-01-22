

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6ServerStats struct {
    
    Name string `json:"name"`
    PortList []Cgnv6ServerStatsPortList `json:"port-list"`
    Stats Cgnv6ServerStatsStats `json:"stats"`

}
type DataCgnv6ServerStats struct {
    DtCgnv6ServerStats Cgnv6ServerStats `json:"server"`
}


type Cgnv6ServerStatsPortList struct {
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Stats Cgnv6ServerStatsPortListStats `json:"stats"`
}


type Cgnv6ServerStatsPortListStats struct {
    Curr_conn int `json:"curr_conn"`
    Curr_req int `json:"curr_req"`
    Total_req int `json:"total_req"`
    Total_req_succ int `json:"total_req_succ"`
    Total_fwd_bytes int `json:"total_fwd_bytes"`
    Total_fwd_pkts int `json:"total_fwd_pkts"`
    Total_rev_bytes int `json:"total_rev_bytes"`
    Total_rev_pkts int `json:"total_rev_pkts"`
    Total_conn int `json:"total_conn"`
    Last_total_conn int `json:"last_total_conn"`
    Peak_conn int `json:"peak_conn"`
    Es_resp_200 int `json:"es_resp_200"`
    Es_resp_300 int `json:"es_resp_300"`
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Es_resp_other int `json:"es_resp_other"`
    Es_req_count int `json:"es_req_count"`
    Es_resp_count int `json:"es_resp_count"`
    Es_resp_invalid_http int `json:"es_resp_invalid_http"`
    Total_rev_pkts_inspected int `json:"total_rev_pkts_inspected"`
    Total_rev_pkts_inspected_good_status_code int `json:"total_rev_pkts_inspected_good_status_code"`
    Response_time int `json:"response_time"`
    Fastest_rsp_time int `json:"fastest_rsp_time"`
    Slowest_rsp_time int `json:"slowest_rsp_time"`
}


type Cgnv6ServerStatsStats struct {
    CurrConn int `json:"curr-conn"`
    TotalConn int `json:"total-conn"`
    FwdPkt int `json:"fwd-pkt"`
    RevPkt int `json:"rev-pkt"`
    PeakConn int `json:"peak-conn"`
}

func (p *Cgnv6ServerStats) GetId() string{
    return "1"
}

func (p *Cgnv6ServerStats) getPath() string{
    
    return "cgnv6/server/"+p.Name+"/stats"
}

func (p *Cgnv6ServerStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6ServerStats,error) {
logger.Println("Cgnv6ServerStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6ServerStats
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
