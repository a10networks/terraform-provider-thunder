

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsCacheServerStats struct {
    
    Stats DdosDnsCacheServerStatsStats `json:"stats"`

}
type DataDdosDnsCacheServerStats struct {
    DtDdosDnsCacheServerStats DdosDnsCacheServerStats `json:"dns-cache-server"`
}


type DdosDnsCacheServerStatsStats struct {
    Insert_total int `json:"insert_total"`
    Insert_success int `json:"insert_success"`
    Insert_fail_all int `json:"insert_fail_all"`
    Lookup_invalid_domain int `json:"lookup_invalid_domain"`
    Lookup_unexp_err int `json:"lookup_unexp_err"`
    Lookup_full_matched int `json:"lookup_full_matched"`
    Lookup_empty_resp int `json:"lookup_empty_resp"`
    Lookup_deleg_resp int `json:"lookup_deleg_resp"`
    Lookup_nxdomain_resp int `json:"lookup_nxdomain_resp"`
    Lookup_refuse_resp int `json:"lookup_refuse_resp"`
    Lookup_fwd_server int `json:"lookup_fwd_server"`
    Lookup_incomp_zone int `json:"lookup_incomp_zone"`
    Lookup_undefined_rtype int `json:"lookup_undefined_rtype"`
    Lookup_manual_override_action_forward int `json:"lookup_manual_override_action_forward"`
    Lookup_manual_override_action_drop int `json:"lookup_manual_override_action_drop"`
    Zt_serial_num_check_attempts int `json:"zt_serial_num_check_attempts"`
    Zt_axfr_attempts int `json:"zt_axfr_attempts"`
    Zt_completed_ok int `json:"zt_completed_ok"`
    Zt_completed_no_update int `json:"zt_completed_no_update"`
    Zt_dns_process_err int `json:"zt_dns_process_err"`
    Zt_records_processed int `json:"zt_records_processed"`
    Lookup_edns_bad_version_resp int `json:"lookup_edns_bad_version_resp"`
    Zt_tcp_conn_connect_server_fail int `json:"zt_tcp_conn_connect_server_fail"`
    Zt_tcp_conn_rst int `json:"zt_tcp_conn_rst"`
    Zt_task_no_route_retry int `json:"zt_task_no_route_retry"`
    Zt_msg_rcode_notauth int `json:"zt_msg_rcode_notauth"`
    Lookup_opcode_notimpl_resp int `json:"lookup_opcode_notimpl_resp"`
    Shard_filter_match int `json:"shard_filter_match"`
    Zt_total_fail int `json:"zt_total_fail"`
    Lookup_manual_override_action_serve int `json:"lookup_manual_override_action_serve"`
    Lookup_any_type_query_action_drop int `json:"lookup_any_type_query_action_drop"`
    Lookup_any_type_query_action_refused int `json:"lookup_any_type_query_action_refused"`
    Lookup_any_type_query_action_resp_empty int `json:"lookup_any_type_query_action_resp_empty"`
    Lookup_non_auth_zone_query_action_forward int `json:"lookup_non_auth_zone_query_action_forward"`
    Lookup_non_auth_zone_query_action_drop int `json:"lookup_non_auth_zone_query_action_drop"`
    Lookup_non_auth_zone_query_action_resp_refused int `json:"lookup_non_auth_zone_query_action_resp_refused"`
    Lookup_default_action_forward int `json:"lookup_default_action_forward"`
    Lookup_default_action_drop int `json:"lookup_default_action_drop"`
    Zt_ongoing_tasks int `json:"zt_ongoing_tasks"`
    Lookup_dnstcp_rcvd int `json:"lookup_dnstcp_rcvd"`
    Lookup_dnsudp_rcvd int `json:"lookup_dnsudp_rcvd"`
    Lookup_fwd_shard int `json:"lookup_fwd_shard"`
}

func (p *DdosDnsCacheServerStats) GetId() string{
    return "1"
}

func (p *DdosDnsCacheServerStats) getPath() string{
    return "ddos/dns-cache-server/stats"
}

func (p *DdosDnsCacheServerStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDnsCacheServerStats,error) {
logger.Println("DdosDnsCacheServerStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDnsCacheServerStats
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
