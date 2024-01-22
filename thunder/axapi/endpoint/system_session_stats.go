

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemSessionStats struct {
    
    Stats SystemSessionStatsStats `json:"stats"`

}
type DataSystemSessionStats struct {
    DtSystemSessionStats SystemSessionStats `json:"session"`
}


type SystemSessionStatsStats struct {
    Total_l4_conn int `json:"total_l4_conn"`
    Conn_counter int `json:"conn_counter"`
    Conn_freed_counter int `json:"conn_freed_counter"`
    Total_l4_packet_count int `json:"total_l4_packet_count"`
    Total_l7_packet_count int `json:"total_l7_packet_count"`
    Total_l4_conn_proxy int `json:"total_l4_conn_proxy"`
    Total_l7_conn int `json:"total_l7_conn"`
    Total_tcp_conn int `json:"total_tcp_conn"`
    Curr_free_conn int `json:"curr_free_conn"`
    Tcp_est_counter int `json:"tcp_est_counter"`
    Tcp_half_open_counter int `json:"tcp_half_open_counter"`
    Tcp_half_close_counter int `json:"tcp_half_close_counter"`
    Udp_counter int `json:"udp_counter"`
    Ip_counter int `json:"ip_counter"`
    Other_counter int `json:"other_counter"`
    Reverse_nat_tcp_counter int `json:"reverse_nat_tcp_counter"`
    Reverse_nat_udp_counter int `json:"reverse_nat_udp_counter"`
    Tcp_syn_half_open_counter int `json:"tcp_syn_half_open_counter"`
    Conn_smp_alloc_counter int `json:"conn_smp_alloc_counter"`
    Conn_smp_free_counter int `json:"conn_smp_free_counter"`
    Conn_smp_aged_counter int `json:"conn_smp_aged_counter"`
    Ssl_count_curr int `json:"ssl_count_curr"`
    Ssl_count_total int `json:"ssl_count_total"`
    Server_ssl_count_curr int `json:"server_ssl_count_curr"`
    Server_ssl_count_total int `json:"server_ssl_count_total"`
    Client_ssl_reuse_total int `json:"client_ssl_reuse_total"`
    Server_ssl_reuse_total int `json:"server_ssl_reuse_total"`
    Total_ip_nat_conn int `json:"total_ip_nat_conn"`
    Total_l2l3_conn int `json:"total_l2l3_conn"`
    Conn_type_0_available int `json:"conn_type_0_available"`
    Conn_type_1_available int `json:"conn_type_1_available"`
    Conn_type_2_available int `json:"conn_type_2_available"`
    Conn_type_3_available int `json:"conn_type_3_available"`
    Conn_type_4_available int `json:"conn_type_4_available"`
    Conn_smp_type_0_available int `json:"conn_smp_type_0_available"`
    Conn_smp_type_1_available int `json:"conn_smp_type_1_available"`
    Conn_smp_type_2_available int `json:"conn_smp_type_2_available"`
    Conn_smp_type_3_available int `json:"conn_smp_type_3_available"`
    Conn_smp_type_4_available int `json:"conn_smp_type_4_available"`
    SctpHalfOpenCounter int `json:"sctp-half-open-counter"`
    SctpEstCounter int `json:"sctp-est-counter"`
    Conn_app_smp_alloc_counter int `json:"conn_app_smp_alloc_counter"`
    Diameter_conn_counter int `json:"diameter_conn_counter"`
    Diameter_conn_freed_counter int `json:"diameter_conn_freed_counter"`
    Total_fw_conn int `json:"total_fw_conn"`
    Total_local_conn int `json:"total_local_conn"`
    Total_curr_conn int `json:"total_curr_conn"`
    Client_ssl_fatal_alert int `json:"client_ssl_fatal_alert"`
    Client_ssl_fin_rst int `json:"client_ssl_fin_rst"`
    Fp_session_fin_rst int `json:"fp_session_fin_rst"`
    Server_ssl_fatal_alert int `json:"server_ssl_fatal_alert"`
    Server_ssl_fin_rst int `json:"server_ssl_fin_rst"`
    Client_template_int_err int `json:"client_template_int_err"`
    Client_template_unknown_err int `json:"client_template_unknown_err"`
    Server_template_int_err int `json:"server_template_int_err"`
    Server_template_unknown_err int `json:"server_template_unknown_err"`
    Diameter_concurrent_user_sessions_counter int `json:"diameter_concurrent_user_sessions_counter"`
    Client_ssl_session_ticket_reuse_total int `json:"client_ssl_session_ticket_reuse_total"`
    Server_ssl_session_ticket_reuse_total int `json:"server_ssl_session_ticket_reuse_total"`
    Total_clientside_early_data_connections int `json:"total_clientside_early_data_connections"`
    Total_serverside_early_data_connections int `json:"total_serverside_early_data_connections"`
    Total_clientside_failed_early_dataConnections int `json:"total_clientside_failed_early_data-connections"`
    Total_serverside_failed_early_dataConnections int `json:"total_serverside_failed_early_data-connections"`
    Total_logging_conn int `json:"total_logging_conn"`
    Gtp_c_est_counter int `json:"gtp_c_est_counter"`
    Gtp_c_half_open_counter int `json:"gtp_c_half_open_counter"`
    Gtp_u_counter int `json:"gtp_u_counter"`
    Gtp_c_echo_counter int `json:"gtp_c_echo_counter"`
    Gtp_u_echo_counter int `json:"gtp_u_echo_counter"`
    Gtp_curr_free_conn int `json:"gtp_curr_free_conn"`
    Gtp_cum_conn_counter int `json:"gtp_cum_conn_counter"`
    Gtp_cum_conn_freed_counter int `json:"gtp_cum_conn_freed_counter"`
    Fw_blacklist_sess int `json:"fw_blacklist_sess"`
    Fw_blacklist_sess_created int `json:"fw_blacklist_sess_created"`
    Fw_blacklist_sess_freed int `json:"fw_blacklist_sess_freed"`
    Server_tcp_est_counter int `json:"server_tcp_est_counter"`
    Server_tcp_half_open_counter int `json:"server_tcp_half_open_counter"`
}

func (p *SystemSessionStats) GetId() string{
    return "1"
}

func (p *SystemSessionStats) getPath() string{
    return "system/session/stats"
}

func (p *SystemSessionStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemSessionStats,error) {
logger.Println("SystemSessionStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemSessionStats
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
