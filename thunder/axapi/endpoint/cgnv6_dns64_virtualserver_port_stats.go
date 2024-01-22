

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Dns64VirtualserverPortStats struct {
    
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Stats Cgnv6Dns64VirtualserverPortStatsStats `json:"stats"`
    Name string 

}
type DataCgnv6Dns64VirtualserverPortStats struct {
    DtCgnv6Dns64VirtualserverPortStats Cgnv6Dns64VirtualserverPortStats `json:"port"`
}


type Cgnv6Dns64VirtualserverPortStatsStats struct {
    Curr_conn int `json:"curr_conn"`
    Total_l4_conn int `json:"total_l4_conn"`
    Total_l7_conn int `json:"total_l7_conn"`
    Toatal_tcp_conn int `json:"toatal_tcp_conn"`
    Total_conn int `json:"total_conn"`
    Total_fwd_bytes int `json:"total_fwd_bytes"`
    Total_fwd_pkts int `json:"total_fwd_pkts"`
    Total_rev_bytes int `json:"total_rev_bytes"`
    Total_rev_pkts int `json:"total_rev_pkts"`
    Total_dns_pkts int `json:"total_dns_pkts"`
    Total_mf_dns_pkts int `json:"total_mf_dns_pkts"`
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Compression_bytes_before int `json:"compression_bytes_before"`
    Compression_bytes_after int `json:"compression_bytes_after"`
    Compression_hit int `json:"compression_hit"`
    Compression_miss int `json:"compression_miss"`
    Compression_miss_no_client int `json:"compression_miss_no_client"`
    Compression_miss_template_exclusion int `json:"compression_miss_template_exclusion"`
    Curr_req int `json:"curr_req"`
    Total_req int `json:"total_req"`
    Total_req_succ int `json:"total_req_succ"`
    Peak_conn int `json:"peak_conn"`
    Curr_conn_rate int `json:"curr_conn_rate"`
    Last_rsp_time int `json:"last_rsp_time"`
    Fastest_rsp_time int `json:"fastest_rsp_time"`
    Slowest_rsp_time int `json:"slowest_rsp_time"`
}

func (p *Cgnv6Dns64VirtualserverPortStats) GetId() string{
    return "1"
}

func (p *Cgnv6Dns64VirtualserverPortStats) getPath() string{
    
    return "cgnv6/dns64-virtualserver/" +p.Name + "/port/" +strconv.Itoa(p.PortNumber)+"+"+p.Protocol+"/stats"
}

func (p *Cgnv6Dns64VirtualserverPortStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Dns64VirtualserverPortStats,error) {
logger.Println("Cgnv6Dns64VirtualserverPortStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Dns64VirtualserverPortStats
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
