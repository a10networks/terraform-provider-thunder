

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbVirtualServerPortStats56 struct {
    
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Stats SlbVirtualServerPortStats56Stats `json:"stats"`
    Name string 

}
type DataSlbVirtualServerPortStats56 struct {
    DtSlbVirtualServerPortStats56 SlbVirtualServerPortStats56 `json:"port"`
}


type SlbVirtualServerPortStats56Stats struct {
    Curr_conn int `json:"curr_conn"`
    Total_l4_conn int `json:"total_l4_conn"`
    Total_l7_conn int `json:"total_l7_conn"`
    Total_tcp_conn int `json:"total_tcp_conn"`
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
    Loc_permit int `json:"loc_permit"`
    Loc_deny int `json:"loc_deny"`
    Loc_conn int `json:"loc_conn"`
    Curr_ssl_conn int `json:"curr_ssl_conn"`
    Total_ssl_conn int `json:"total_ssl_conn"`
    BackendTimeToFirstByte int `json:"backend-time-to-first-byte"`
    BackendTimeToLastByte int `json:"backend-time-to-last-byte"`
    InLatency int `json:"in-latency"`
    OutLatency int `json:"out-latency"`
    Total_fwd_bytes_out int `json:"total_fwd_bytes_out"`
    Total_fwd_pkts_out int `json:"total_fwd_pkts_out"`
    Total_rev_bytes_out int `json:"total_rev_bytes_out"`
    Total_rev_pkts_out int `json:"total_rev_pkts_out"`
    Curr_req_rate int `json:"curr_req_rate"`
    Curr_resp int `json:"curr_resp"`
    Total_resp int `json:"total_resp"`
    Total_resp_succ int `json:"total_resp_succ"`
    Curr_resp_rate int `json:"curr_resp_rate"`
    Dnsrrl_total_allowed int `json:"dnsrrl_total_allowed"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Dnsrrl_total_slipped int `json:"dnsrrl_total_slipped"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    ThroughputBitsPerSec int `json:"throughput-bits-per-sec"`
    DynamicMemory int `json:"dynamic-memory"`
    Ip_only_lb_fwd_bytes int `json:"ip_only_lb_fwd_bytes"`
    Ip_only_lb_rev_bytes int `json:"ip_only_lb_rev_bytes"`
    Ip_only_lb_fwd_pkts int `json:"ip_only_lb_fwd_pkts"`
    Ip_only_lb_rev_pkts int `json:"ip_only_lb_rev_pkts"`
    Total_dns_filter_type_drop int `json:"total_dns_filter_type_drop"`
    Total_dns_filter_class_drop int `json:"total_dns_filter_class_drop"`
    Dns_filter_type_a_drop int `json:"dns_filter_type_a_drop"`
    Dns_filter_type_aaaa_drop int `json:"dns_filter_type_aaaa_drop"`
    Dns_filter_type_cname_drop int `json:"dns_filter_type_cname_drop"`
    Dns_filter_type_mx_drop int `json:"dns_filter_type_mx_drop"`
    Dns_filter_type_ns_drop int `json:"dns_filter_type_ns_drop"`
    Dns_filter_type_srv_drop int `json:"dns_filter_type_srv_drop"`
    Dns_filter_type_ptr_drop int `json:"dns_filter_type_ptr_drop"`
    Dns_filter_type_soa_drop int `json:"dns_filter_type_soa_drop"`
    Dns_filter_type_txt_drop int `json:"dns_filter_type_txt_drop"`
    Dns_filter_type_any_drop int `json:"dns_filter_type_any_drop"`
    Dns_filter_type_others_drop int `json:"dns_filter_type_others_drop"`
    Dns_filter_class_internet_drop int `json:"dns_filter_class_internet_drop"`
    Dns_filter_class_chaos_drop int `json:"dns_filter_class_chaos_drop"`
    Dns_filter_class_hesiod_drop int `json:"dns_filter_class_hesiod_drop"`
    Dns_filter_class_none_drop int `json:"dns_filter_class_none_drop"`
    Dns_filter_class_any_drop int `json:"dns_filter_class_any_drop"`
    Dns_filter_class_others_drop int `json:"dns_filter_class_others_drop"`
    Dns_rpz_action_drop int `json:"dns_rpz_action_drop"`
    Dns_rpz_action_pass_thru int `json:"dns_rpz_action_pass_thru"`
    Dns_rpz_action_tcp_only int `json:"dns_rpz_action_tcp_only"`
    Dns_rpz_action_nxdomain int `json:"dns_rpz_action_nxdomain"`
    Dns_rpz_action_nodata int `json:"dns_rpz_action_nodata"`
    Dns_rpz_action_local_data int `json:"dns_rpz_action_local_data"`
    Dns_rpz_trigger_client_ip int `json:"dns_rpz_trigger_client_ip"`
    Dns_rpz_trigger_resp_ip int `json:"dns_rpz_trigger_resp_ip"`
    Dns_rpz_trigger_ns_ip int `json:"dns_rpz_trigger_ns_ip"`
    Dns_rpz_trigger_qname int `json:"dns_rpz_trigger_qname"`
    Dns_rpz_trigger_ns_name int `json:"dns_rpz_trigger_ns_name"`
    Compression_bytes_before_br int `json:"compression_bytes_before_br"`
    Compression_bytes_after_br int `json:"compression_bytes_after_br"`
    Compression_bytes_before_total int `json:"compression_bytes_before_total"`
    Compression_bytes_after_total int `json:"compression_bytes_after_total"`
    Compression_hit_br int `json:"compression_hit_br"`
    Compression_miss_br int `json:"compression_miss_br"`
    Compression_hit_total int `json:"compression_hit_total"`
    Compression_miss_total int `json:"compression_miss_total"`
    Dnsrrl_total_tc int `json:"dnsrrl_total_tc"`
    Http1_client_idle_timeout int `json:"http1_client_idle_timeout"`
    Http2_client_idle_timeout int `json:"http2_client_idle_timeout"`
}

func (p *SlbVirtualServerPortStats56) GetId() string{
    return "1"
}

func (p *SlbVirtualServerPortStats56) getPath() string{
    
    return "slb/virtual-server/" +p.Name + "/port/" +strconv.Itoa(p.PortNumber)+"+"+p.Protocol+"/stats"
}

func (p *SlbVirtualServerPortStats56) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbVirtualServerPortStats56,error) {
logger.Println("SlbVirtualServerPortStats56::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbVirtualServerPortStats56
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
