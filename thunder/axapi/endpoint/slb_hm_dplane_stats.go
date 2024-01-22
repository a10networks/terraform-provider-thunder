

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHmDplaneStats struct {
    
    Stats SlbHmDplaneStatsStats `json:"stats"`

}
type DataSlbHmDplaneStats struct {
    DtSlbHmDplaneStats SlbHmDplaneStats `json:"hm-dplane"`
}


type SlbHmDplaneStatsStats struct {
    Curr_entries int `json:"curr_entries"`
    Total_created int `json:"total_created"`
    Total_inserted int `json:"total_inserted"`
    Total_ready_to_free int `json:"total_ready_to_free"`
    Total_freed int `json:"total_freed"`
    Err_entry_create_failed int `json:"err_entry_create_failed"`
    Err_entry_create_oom int `json:"err_entry_create_oom"`
    Err_entry_insert_failed int `json:"err_entry_insert_failed"`
    Total_tcp_err int `json:"total_tcp_err"`
    Err_smart_nat_alloc int `json:"err_smart_nat_alloc"`
    Err_smart_nat_port_alloc int `json:"err_smart_nat_port_alloc"`
    Err_l4_sess_alloc int `json:"err_l4_sess_alloc"`
    Err_hm_tcp_conn_sent int `json:"err_hm_tcp_conn_sent"`
    Hm_tcp_conn_sent int `json:"hm_tcp_conn_sent"`
    Entry_deleted int `json:"entry_deleted"`
    Err_entry_create_vip_failed int `json:"err_entry_create_vip_failed"`
    Total_match_resp_code int `json:"total_match_resp_code"`
    Total_match_default_resp_code int `json:"total_match_default_resp_code"`
    Total_maintenance_received int `json:"total_maintenance_received"`
    Total_wrong_status_received int `json:"total_wrong_status_received"`
    Err_no_hm_entry int `json:"err_no_hm_entry"`
    Err_ssl_cert_name_mismatch int `json:"err_ssl_cert_name_mismatch"`
    Err_server_syn_timeout int `json:"err_server_syn_timeout"`
    Err_http2_callback int `json:"err_http2_callback"`
    Err_l7_sess_process_tcp_estab_failed int `json:"err_l7_sess_process_tcp_estab_failed"`
    Err_l7_sess_process_tcp_data_failed int `json:"err_l7_sess_process_tcp_data_failed"`
    Err_http2_ver_mismatch int `json:"err_http2_ver_mismatch"`
    Smart_nat_alloc int `json:"smart_nat_alloc"`
    Smart_nat_release int `json:"smart_nat_release"`
    Smart_nat_alloc_failed int `json:"smart_nat_alloc_failed"`
    Smart_nat_release_failed int `json:"smart_nat_release_failed"`
    Total_server_quic_conn int `json:"total_server_quic_conn"`
    Total_server_quic_conn_err int `json:"total_server_quic_conn_err"`
}

func (p *SlbHmDplaneStats) GetId() string{
    return "1"
}

func (p *SlbHmDplaneStats) getPath() string{
    return "slb/hm-dplane/stats"
}

func (p *SlbHmDplaneStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHmDplaneStats,error) {
logger.Println("SlbHmDplaneStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHmDplaneStats
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
