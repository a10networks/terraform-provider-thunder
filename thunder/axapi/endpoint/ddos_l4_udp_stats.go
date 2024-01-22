

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosL4UdpStats struct {
    
    Stats DdosL4UdpStatsStats `json:"stats"`

}
type DataDdosL4UdpStats struct {
    DtDdosL4UdpStats DdosL4UdpStats `json:"l4-udp"`
}


type DdosL4UdpStatsStats struct {
    Udp_sess_create int `json:"udp_sess_create"`
    Inudp int `json:"inudp"`
    Instateless int `json:"instateless"`
    Udp_total_drop int `json:"udp_total_drop"`
    Udp_drop_dst int `json:"udp_drop_dst"`
    Udp_drop_src int `json:"udp_drop_src"`
    Udp_drop_black_user_cfg_src int `json:"udp_drop_black_user_cfg_src"`
    Udp_src_dst_drop int `json:"udp_src_dst_drop"`
    Udp_drop_black_user_cfg_src_dst int `json:"udp_drop_black_user_cfg_src_dst"`
    Udp_port_zero_drop int `json:"udp_port_zero_drop"`
    Udp_wellknown_src_port_drop int `json:"udp_wellknown_src_port_drop"`
    Udp_retry_start int `json:"udp_retry_start"`
    Udp_retry_pass int `json:"udp_retry_pass"`
    Udp_retry_fail int `json:"udp_retry_fail"`
    Udp_retry_timeout int `json:"udp_retry_timeout"`
    Udp_payload_too_big_drop int `json:"udp_payload_too_big_drop"`
    Udp_payload_too_small_drop int `json:"udp_payload_too_small_drop"`
    Ntp_monlist_req_drop int `json:"ntp_monlist_req_drop"`
    Ntp_monlist_resp_drop int `json:"ntp_monlist_resp_drop"`
    Udp_conn_prate_drop int `json:"udp_conn_prate_drop"`
    Dst_udp_filter_match int `json:"dst_udp_filter_match"`
    Dst_udp_filter_not_match int `json:"dst_udp_filter_not_match"`
    Dst_udp_filter_action_blacklist int `json:"dst_udp_filter_action_blacklist"`
    Dst_udp_filter_action_drop int `json:"dst_udp_filter_action_drop"`
    Dst_udp_filter_action_default_pass int `json:"dst_udp_filter_action_default_pass"`
    Dst_udp_filter_action_whitelist int `json:"dst_udp_filter_action_whitelist"`
    Udp_auth_pass int `json:"udp_auth_pass"`
    Src_udp_filter_match int `json:"src_udp_filter_match"`
    Src_udp_filter_not_match int `json:"src_udp_filter_not_match"`
    Src_udp_filter_action_blacklist int `json:"src_udp_filter_action_blacklist"`
    Src_udp_filter_action_drop int `json:"src_udp_filter_action_drop"`
    Src_udp_filter_action_default_pass int `json:"src_udp_filter_action_default_pass"`
    Src_udp_filter_action_whitelist int `json:"src_udp_filter_action_whitelist"`
    Src_dst_udp_filter_match int `json:"src_dst_udp_filter_match"`
    Src_dst_udp_filter_not_match int `json:"src_dst_udp_filter_not_match"`
    Src_dst_udp_filter_action_blacklist int `json:"src_dst_udp_filter_action_blacklist"`
    Src_dst_udp_filter_action_drop int `json:"src_dst_udp_filter_action_drop"`
    Src_dst_udp_filter_action_default_pass int `json:"src_dst_udp_filter_action_default_pass"`
    Src_dst_udp_filter_action_whitelist int `json:"src_dst_udp_filter_action_whitelist"`
    Udp_wellknown_src_port int `json:"udp_wellknown_src_port"`
    Udp_wellknown_src_port_bl int `json:"udp_wellknown_src_port_bl"`
    Udp_retry_pass_wl int `json:"udp_retry_pass_wl"`
    Udp_retry_fail_bl int `json:"udp_retry_fail_bl"`
    Udp_payload_too_big int `json:"udp_payload_too_big"`
    Udp_payload_too_big_bl int `json:"udp_payload_too_big_bl"`
    Udp_payload_too_small int `json:"udp_payload_too_small"`
    Udp_payload_too_small_bl int `json:"udp_payload_too_small_bl"`
    Ntp_monlist_req int `json:"ntp_monlist_req"`
    Ntp_monlist_req_bl int `json:"ntp_monlist_req_bl"`
    Ntp_monlist_resp int `json:"ntp_monlist_resp"`
    Ntp_monlist_resp_bl int `json:"ntp_monlist_resp_bl"`
    Udp_conn_prate_exceed int `json:"udp_conn_prate_exceed"`
    Udp_conn_prate_bl int `json:"udp_conn_prate_bl"`
    Udp_any_exceed int `json:"udp_any_exceed"`
    Udp_drop_bl int `json:"udp_drop_bl"`
    Udp_frag_rcvd int `json:"udp_frag_rcvd"`
    Udp_frag_drop int `json:"udp_frag_drop"`
    Udp_auth_drop int `json:"udp_auth_drop"`
    Udp_total_bytes_rcv int `json:"udp_total_bytes_rcv"`
    Udp_total_bytes_drop int `json:"udp_total_bytes_drop"`
    Udp_retry_gap_drop int `json:"udp_retry_gap_drop"`
}

func (p *DdosL4UdpStats) GetId() string{
    return "1"
}

func (p *DdosL4UdpStats) getPath() string{
    return "ddos/l4-udp/stats"
}

func (p *DdosL4UdpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosL4UdpStats,error) {
logger.Println("DdosL4UdpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosL4UdpStats
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
