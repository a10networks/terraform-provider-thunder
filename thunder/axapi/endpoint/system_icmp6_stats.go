

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIcmp6Stats struct {
    
    Stats SystemIcmp6StatsStats `json:"stats"`

}
type DataSystemIcmp6Stats struct {
    DtSystemIcmp6Stats SystemIcmp6Stats `json:"icmp6"`
}


type SystemIcmp6StatsStats struct {
    In_msgs int `json:"in_msgs"`
    In_errors int `json:"in_errors"`
    In_dest_un_reach int `json:"in_dest_un_reach"`
    In_pkt_too_big int `json:"in_pkt_too_big"`
    In_time_exceeds int `json:"in_time_exceeds"`
    In_param_prob int `json:"in_param_prob"`
    In_echoes int `json:"in_echoes"`
    In_exho_reply int `json:"in_exho_reply"`
    In_grp_mem_query int `json:"in_grp_mem_query"`
    In_grp_mem_resp int `json:"in_grp_mem_resp"`
    In_grp_mem_reduction int `json:"in_grp_mem_reduction"`
    In_router_sol int `json:"in_router_sol"`
    In_ra int `json:"in_ra"`
    In_ns int `json:"in_ns"`
    In_na int `json:"in_na"`
    In_redirect int `json:"in_redirect"`
    Out_msg int `json:"out_msg"`
    Out_dst_un_reach int `json:"out_dst_un_reach"`
    Out_pkt_too_big int `json:"out_pkt_too_big"`
    Out_time_exceeds int `json:"out_time_exceeds"`
    Out_param_prob int `json:"out_param_prob"`
    Out_echo_req int `json:"out_echo_req"`
    Out_echo_replies int `json:"out_echo_replies"`
    Out_rs int `json:"out_rs"`
    Out_ra int `json:"out_ra"`
    Out_ns int `json:"out_ns"`
    Out_na int `json:"out_na"`
    Out_redirects int `json:"out_redirects"`
    Out_mem_resp int `json:"out_mem_resp"`
    Out_mem_reductions int `json:"out_mem_reductions"`
    Err_rs int `json:"err_rs"`
    Err_ra int `json:"err_ra"`
    Err_ns int `json:"err_ns"`
    Err_na int `json:"err_na"`
    Err_redirects int `json:"err_redirects"`
    Err_echoes int `json:"err_echoes"`
    Err_echo_replies int `json:"err_echo_replies"`
}

func (p *SystemIcmp6Stats) GetId() string{
    return "1"
}

func (p *SystemIcmp6Stats) getPath() string{
    return "system/icmp6/stats"
}

func (p *SystemIcmp6Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemIcmp6Stats,error) {
logger.Println("SystemIcmp6Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemIcmp6Stats
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
