

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosL4TcpStats struct {
    
    Stats DdosL4TcpStatsStats `json:"stats"`

}
type DataDdosL4TcpStats struct {
    DtDdosL4TcpStats DdosL4TcpStats `json:"l4-tcp"`
}


type DdosL4TcpStatsStats struct {
    Tcp_sess_create int `json:"tcp_sess_create"`
    Intcp int `json:"intcp"`
    Tcp_syn_rcvd int `json:"tcp_syn_rcvd"`
    Tcp_invalid_syn_rcvd int `json:"tcp_invalid_syn_rcvd"`
    Tcp_syn_ack_rcvd int `json:"tcp_syn_ack_rcvd"`
    Tcp_ack_rcvd int `json:"tcp_ack_rcvd"`
    Tcp_fin_rcvd int `json:"tcp_fin_rcvd"`
    Tcp_rst_rcvd int `json:"tcp_rst_rcvd"`
    Tcp_outrst int `json:"tcp_outrst"`
    Tcp_reset_client int `json:"tcp_reset_client"`
    Tcp_reset_server int `json:"tcp_reset_server"`
    Tcp_syn_rate int `json:"tcp_syn_rate"`
    Tcp_total_drop int `json:"tcp_total_drop"`
    Tcp_dst_drop int `json:"tcp_dst_drop"`
    Tcp_src_drop int `json:"tcp_src_drop"`
    Tcp_drop_black_user_cfg_src int `json:"tcp_drop_black_user_cfg_src"`
    Tcp_src_dst_drop int `json:"tcp_src_dst_drop"`
    Tcp_drop_black_user_cfg_src_dst int `json:"tcp_drop_black_user_cfg_src_dst"`
    Tcp_port_zero_drop int `json:"tcp_port_zero_drop"`
    Tcp_syncookie_sent int `json:"tcp_syncookie_sent"`
    Tcp_syncookie_sent_fail int `json:"tcp_syncookie_sent_fail"`
    Tcp_syncookie_check_fail int `json:"tcp_syncookie_check_fail"`
    Tcp_syncookie_hw_missing int `json:"tcp_syncookie_hw_missing"`
    Tcp_syncookie_fail_bl int `json:"tcp_syncookie_fail_bl"`
    Tcp_syncookie_pass int `json:"tcp_syncookie_pass"`
    Syn_auth_pass int `json:"syn_auth_pass"`
    Syn_auth_skip int `json:"syn_auth_skip"`
    Tcp_action_on_ack_start int `json:"tcp_action_on_ack_start"`
    Tcp_action_on_ack_matched int `json:"tcp_action_on_ack_matched"`
    Tcp_action_on_ack_passed int `json:"tcp_action_on_ack_passed"`
    Tcp_action_on_ack_failed int `json:"tcp_action_on_ack_failed"`
    Tcp_action_on_ack_timeout int `json:"tcp_action_on_ack_timeout"`
    Tcp_action_on_ack_reset int `json:"tcp_action_on_ack_reset"`
    Tcp_ack_no_syn int `json:"tcp_ack_no_syn"`
    Tcp_out_of_seq int `json:"tcp_out_of_seq"`
    Tcp_zero_window int `json:"tcp_zero_window"`
    Tcp_retransmit int `json:"tcp_retransmit"`
    Tcp_rexmit_syn_limit_drop int `json:"tcp_rexmit_syn_limit_drop"`
    Tcp_zero_window_bl int `json:"tcp_zero_window_bl"`
    Tcp_out_of_seq_bl int `json:"tcp_out_of_seq_bl"`
    Tcp_retransmit_bl int `json:"tcp_retransmit_bl"`
    Tcp_rexmit_syn_limit_bl int `json:"tcp_rexmit_syn_limit_bl"`
    Tcp_per_conn_prate_exceed int `json:"tcp_per_conn_prate_exceed"`
    Tcp_action_on_ack_gap_drop int `json:"tcp_action_on_ack_gap_drop"`
    Tcp_action_on_ack_gap_pass int `json:"tcp_action_on_ack_gap_pass"`
    Tcp_action_on_syn_start int `json:"tcp_action_on_syn_start"`
    Tcp_action_on_syn_passed int `json:"tcp_action_on_syn_passed"`
    Tcp_action_on_syn_failed int `json:"tcp_action_on_syn_failed"`
    Tcp_action_on_syn_timeout int `json:"tcp_action_on_syn_timeout"`
    Tcp_action_on_syn_reset int `json:"tcp_action_on_syn_reset"`
    Tcp_action_on_syn_gap_drop int `json:"tcp_action_on_syn_gap_drop"`
    Tcp_action_on_syn_gap_pass int `json:"tcp_action_on_syn_gap_pass"`
    Tcp_unauth_rst_drop int `json:"tcp_unauth_rst_drop"`
    Dst_tcp_filter_match int `json:"dst_tcp_filter_match"`
    Dst_tcp_filter_not_match int `json:"dst_tcp_filter_not_match"`
    Dst_tcp_filter_action_blacklist int `json:"dst_tcp_filter_action_blacklist"`
    Dst_tcp_filter_action_drop int `json:"dst_tcp_filter_action_drop"`
    Dst_tcp_filter_action_default_pass int `json:"dst_tcp_filter_action_default_pass"`
    Tcp_concurrent int `json:"tcp_concurrent"`
    Dst_tcp_filter_action_whitelist int `json:"dst_tcp_filter_action_whitelist"`
    Src_tcp_filter_match int `json:"src_tcp_filter_match"`
    Src_tcp_filter_not_match int `json:"src_tcp_filter_not_match"`
    Src_tcp_filter_action_blacklist int `json:"src_tcp_filter_action_blacklist"`
    Src_tcp_filter_action_drop int `json:"src_tcp_filter_action_drop"`
    Src_tcp_filter_action_default_pass int `json:"src_tcp_filter_action_default_pass"`
    Src_tcp_filter_action_whitelist int `json:"src_tcp_filter_action_whitelist"`
    Src_dst_tcp_filter_match int `json:"src_dst_tcp_filter_match"`
    Src_dst_tcp_filter_not_match int `json:"src_dst_tcp_filter_not_match"`
    Src_dst_tcp_filter_action_blacklist int `json:"src_dst_tcp_filter_action_blacklist"`
    Src_dst_tcp_filter_action_drop int `json:"src_dst_tcp_filter_action_drop"`
    Src_dst_tcp_filter_action_default_pass int `json:"src_dst_tcp_filter_action_default_pass"`
    Src_dst_tcp_filter_action_whitelist int `json:"src_dst_tcp_filter_action_whitelist"`
    Syn_auth_pass_wl int `json:"syn_auth_pass_wl"`
    Tcp_out_of_seq_drop int `json:"tcp_out_of_seq_drop"`
    Tcp_zero_window_drop int `json:"tcp_zero_window_drop"`
    Tcp_retransmit_drop int `json:"tcp_retransmit_drop"`
    Tcp_per_conn_prate_exceed_bl int `json:"tcp_per_conn_prate_exceed_bl"`
    Tcp_any_exceed int `json:"tcp_any_exceed"`
    Tcp_drop_bl int `json:"tcp_drop_bl"`
    Tcp_frag_rcvd int `json:"tcp_frag_rcvd"`
    Tcp_frag_drop int `json:"tcp_frag_drop"`
    Tcp_auth_drop int `json:"tcp_auth_drop"`
    Tcp_auth_resp int `json:"tcp_auth_resp"`
    Tcp_total_bytes_rcv int `json:"tcp_total_bytes_rcv"`
    Tcp_total_bytes_drop int `json:"tcp_total_bytes_drop"`
    Tcp_action_on_ack_bl int `json:"tcp_action_on_ack_bl"`
    Tcp_action_on_syn_bl int `json:"tcp_action_on_syn_bl"`
    Tcp_per_conn_ofo_rate_exceed_drop int `json:"tcp_per_conn_ofo_rate_exceed_drop"`
    Tcp_per_conn_ofo_rate_exceed_bl int `json:"tcp_per_conn_ofo_rate_exceed_bl"`
    Tcp_per_conn_rexmit_rate_exceed_drop int `json:"tcp_per_conn_rexmit_rate_exceed_drop"`
    Tcp_per_conn_rexmit_rate_exceed_bl int `json:"tcp_per_conn_rexmit_rate_exceed_bl"`
    Tcp_per_conn_zwindow_rate_exceed_drop int `json:"tcp_per_conn_zwindow_rate_exceed_drop"`
    Tcp_per_conn_zwindow_rate_exceed_bl int `json:"tcp_per_conn_zwindow_rate_exceed_bl"`
    Tcp_syn_tfo_rcvd int `json:"tcp_syn_tfo_rcvd"`
    Tcp_progression_violation_exceed int `json:"tcp_progression_violation_exceed"`
    Tcp_progression_violation_exceed_bl int `json:"tcp_progression_violation_exceed_bl"`
    Tcp_progression_violation_exceed_drop int `json:"tcp_progression_violation_exceed_drop"`
    Tcp_progression_violation_exceed_reset int `json:"tcp_progression_violation_exceed_reset"`
    Tcp_auth_rst int `json:"tcp_auth_rst"`
}

func (p *DdosL4TcpStats) GetId() string{
    return "1"
}

func (p *DdosL4TcpStats) getPath() string{
    return "ddos/l4-tcp/stats"
}

func (p *DdosL4TcpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosL4TcpStats,error) {
logger.Println("DdosL4TcpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosL4TcpStats
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
