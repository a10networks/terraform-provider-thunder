

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosL7DnsStats struct {
    
    Stats DdosL7DnsStatsStats `json:"stats"`

}
type DataDdosL7DnsStats struct {
    DtDdosL7DnsStats DdosL7DnsStats `json:"l7-dns"`
}


type DdosL7DnsStatsStats struct {
    Force_tcp_auth int `json:"force_tcp_auth"`
    Dns_auth_udp int `json:"dns_auth_udp"`
    Dns_malform_drop int `json:"dns_malform_drop"`
    Dns_qry_any_drop int `json:"dns_qry_any_drop"`
    Dst_rate_limit0 int `json:"dst_rate_limit0"`
    Dst_rate_limit1 int `json:"dst_rate_limit1"`
    Dst_rate_limit2 int `json:"dst_rate_limit2"`
    Dst_rate_limit3 int `json:"dst_rate_limit3"`
    Dst_rate_limit4 int `json:"dst_rate_limit4"`
    Src_rate_limit0 int `json:"src_rate_limit0"`
    Src_rate_limit1 int `json:"src_rate_limit1"`
    Src_rate_limit2 int `json:"src_rate_limit2"`
    Src_rate_limit3 int `json:"src_rate_limit3"`
    Src_rate_limit4 int `json:"src_rate_limit4"`
    Dns_auth_udp_pass int `json:"dns_auth_udp_pass"`
    Dns_fqdn_stage2_exceed int `json:"dns_fqdn_stage2_exceed"`
    Dns_is_nx int `json:"dns_is_nx"`
    Dns_nx_drop int `json:"dns_nx_drop"`
    Dns_nx_bl int `json:"dns_nx_bl"`
    Dns_tcp_auth_pass int `json:"dns_tcp_auth_pass"`
    Dns_auth_udp_fail int `json:"dns_auth_udp_fail"`
    Dns_auth_udp_timeout int `json:"dns_auth_udp_timeout"`
    Dns_fqdn_label_len_exceed int `json:"dns_fqdn_label_len_exceed"`
    Dns_pkt_processed int `json:"dns_pkt_processed"`
    Dns_query_type_a int `json:"dns_query_type_a"`
    Dns_query_type_aaaa int `json:"dns_query_type_aaaa"`
    Dns_query_type_ns int `json:"dns_query_type_ns"`
    Dns_query_type_cname int `json:"dns_query_type_cname"`
    Dns_query_type_any int `json:"dns_query_type_any"`
    Dns_query_type_srv int `json:"dns_query_type_srv"`
    Dns_query_type_mx int `json:"dns_query_type_mx"`
    Dns_query_type_soa int `json:"dns_query_type_soa"`
    Dns_query_type_opt int `json:"dns_query_type_opt"`
    Dns_dg_action_permit int `json:"dns_dg_action_permit"`
    Dns_dg_action_deny int `json:"dns_dg_action_deny"`
    Dns_fqdn_rate_by_label_count_exceed int `json:"dns_fqdn_rate_by_label_count_exceed"`
    Dns_udp_auth_retry_gap_drop int `json:"dns_udp_auth_retry_gap_drop"`
    Dns_policy_drop int `json:"dns_policy_drop"`
    Dns_fqdn_label_count_exceed int `json:"dns_fqdn_label_count_exceed"`
    Dns_rrtype_drop int `json:"dns_rrtype_drop"`
    Force_tcp_auth_timeout int `json:"force_tcp_auth_timeout"`
    Dns_auth_drop int `json:"dns_auth_drop"`
    Dns_auth_resp int `json:"dns_auth_resp"`
    Force_tcp_auth_conn_hit int `json:"force_tcp_auth_conn_hit"`
    Dns_auth_udp_fail_bl int `json:"dns_auth_udp_fail_bl"`
    Dns_nx_exceed int `json:"dns_nx_exceed"`
    Dns_query_class_whitelist_miss int `json:"dns_query_class_whitelist_miss"`
    Dns_query_class_in int `json:"dns_query_class_in"`
    Dns_query_class_csnet int `json:"dns_query_class_csnet"`
    Dns_query_class_chaos int `json:"dns_query_class_chaos"`
    Dns_query_class_hs int `json:"dns_query_class_hs"`
    Dns_query_class_none int `json:"dns_query_class_none"`
    Dns_query_class_any int `json:"dns_query_class_any"`
    Dns_dg_rate_exceed int `json:"dns_dg_rate_exceed"`
    Dns_outbound_query_response_size_exceed int `json:"dns_outbound_query_response_size_exceed"`
    Dns_outbound_query_sess_timed_out int `json:"dns_outbound_query_sess_timed_out"`
    Non_query_opcode_pass_through int `json:"non_query_opcode_pass_through"`
}

func (p *DdosL7DnsStats) GetId() string{
    return "1"
}

func (p *DdosL7DnsStats) getPath() string{
    return "ddos/l7-dns/stats"
}

func (p *DdosL7DnsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosL7DnsStats,error) {
logger.Println("DdosL7DnsStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosL7DnsStats
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
