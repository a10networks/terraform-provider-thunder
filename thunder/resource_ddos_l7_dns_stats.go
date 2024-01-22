package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosL7DnsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_l7_dns_stats`: Statistics for the object l7-dns\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosL7DnsStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"force_tcp_auth": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth Force-TCP",
						},
						"dns_auth_udp": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth UDP",
						},
						"dns_malform_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Malform Query Dropped",
						},
						"dns_qry_any_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query ANY Dropped",
						},
						"dst_rate_limit0": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 1 Exceeded",
						},
						"dst_rate_limit1": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 2 Exceeded",
						},
						"dst_rate_limit2": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 3 Exceeded",
						},
						"dst_rate_limit3": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 4 Exceeded",
						},
						"dst_rate_limit4": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 5 Exceeded",
						},
						"src_rate_limit0": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 1 Exceeded",
						},
						"src_rate_limit1": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 2 Exceeded",
						},
						"src_rate_limit2": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 3 Exceeded",
						},
						"src_rate_limit3": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 4 Exceeded",
						},
						"src_rate_limit4": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 5 Exceeded",
						},
						"dns_auth_udp_pass": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth UDP Passed",
						},
						"dns_fqdn_stage2_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Rate Exceeded",
						},
						"dns_is_nx": {
							Type: schema.TypeInt, Optional: true, Description: "NXDOMAIN Response",
						},
						"dns_nx_drop": {
							Type: schema.TypeInt, Optional: true, Description: "NXDOMAIN Query Dropped",
						},
						"dns_nx_bl": {
							Type: schema.TypeInt, Optional: true, Description: "NXDOMAIN Query Blacklisted",
						},
						"dns_tcp_auth_pass": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth Force-TCP Passed",
						},
						"dns_auth_udp_fail": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth UDP Failed",
						},
						"dns_auth_udp_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth UDP Timeout",
						},
						"dns_fqdn_label_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Label Length Exceeded",
						},
						"dns_pkt_processed": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Packets Processed",
						},
						"dns_query_type_a": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Type A",
						},
						"dns_query_type_aaaa": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Type AAAA",
						},
						"dns_query_type_ns": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Type NS",
						},
						"dns_query_type_cname": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Type CNAME",
						},
						"dns_query_type_any": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Type ANY",
						},
						"dns_query_type_srv": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Type SRV",
						},
						"dns_query_type_mx": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Type MX",
						},
						"dns_query_type_soa": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Type SOA",
						},
						"dns_query_type_opt": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Type OPT",
						},
						"dns_dg_action_permit": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Domain Group Action Permit",
						},
						"dns_dg_action_deny": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Domain Group Action Deny",
						},
						"dns_fqdn_rate_by_label_count_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Rate by Label Count Exceeded",
						},
						"dns_udp_auth_retry_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth UDP Retry-Gap Drop",
						},
						"dns_policy_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Policy Dropped",
						},
						"dns_fqdn_label_count_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Label Count Exceeded",
						},
						"dns_rrtype_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Record Type Dropped",
						},
						"force_tcp_auth_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth Force-TCP With UDP Auth Timeout",
						},
						"dns_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth Dropped",
						},
						"dns_auth_resp": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth Responded",
						},
						"force_tcp_auth_conn_hit": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth Force-TCP With UDP Auth Connection Hit",
						},
						"dns_auth_udp_fail_bl": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth UDP Fail Blacklisted",
						},
						"dns_nx_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "NXDOMAIN Response Rate Exceeded",
						},
						"dns_query_class_whitelist_miss": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Class Whitelist Miss",
						},
						"dns_query_class_in": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Class INTERNET",
						},
						"dns_query_class_csnet": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Class CSNET",
						},
						"dns_query_class_chaos": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Class CHAOS",
						},
						"dns_query_class_hs": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Class HESIOD",
						},
						"dns_query_class_none": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Class NONE",
						},
						"dns_query_class_any": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Class ANY",
						},
						"dns_dg_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Domain Group Domain Query Rate Exceeded",
						},
						"dns_outbound_query_response_size_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Outbound Query Resp Size Exceeded",
						},
						"dns_outbound_query_sess_timed_out": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Outbound Query Session Timed Out",
						},
						"non_query_opcode_pass_through": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Non Query Opcode Pass Through",
						},
					},
				},
			},
		},
	}
}

func resourceDdosL7DnsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL7DnsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL7DnsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosL7DnsStatsStats := setObjectDdosL7DnsStatsStats(res)
		d.Set("stats", DdosL7DnsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosL7DnsStatsStats(ret edpt.DataDdosL7DnsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"force_tcp_auth":                          ret.DtDdosL7DnsStats.Stats.Force_tcp_auth,
			"dns_auth_udp":                            ret.DtDdosL7DnsStats.Stats.Dns_auth_udp,
			"dns_malform_drop":                        ret.DtDdosL7DnsStats.Stats.Dns_malform_drop,
			"dns_qry_any_drop":                        ret.DtDdosL7DnsStats.Stats.Dns_qry_any_drop,
			"dst_rate_limit0":                         ret.DtDdosL7DnsStats.Stats.Dst_rate_limit0,
			"dst_rate_limit1":                         ret.DtDdosL7DnsStats.Stats.Dst_rate_limit1,
			"dst_rate_limit2":                         ret.DtDdosL7DnsStats.Stats.Dst_rate_limit2,
			"dst_rate_limit3":                         ret.DtDdosL7DnsStats.Stats.Dst_rate_limit3,
			"dst_rate_limit4":                         ret.DtDdosL7DnsStats.Stats.Dst_rate_limit4,
			"src_rate_limit0":                         ret.DtDdosL7DnsStats.Stats.Src_rate_limit0,
			"src_rate_limit1":                         ret.DtDdosL7DnsStats.Stats.Src_rate_limit1,
			"src_rate_limit2":                         ret.DtDdosL7DnsStats.Stats.Src_rate_limit2,
			"src_rate_limit3":                         ret.DtDdosL7DnsStats.Stats.Src_rate_limit3,
			"src_rate_limit4":                         ret.DtDdosL7DnsStats.Stats.Src_rate_limit4,
			"dns_auth_udp_pass":                       ret.DtDdosL7DnsStats.Stats.Dns_auth_udp_pass,
			"dns_fqdn_stage2_exceed":                  ret.DtDdosL7DnsStats.Stats.Dns_fqdn_stage2_exceed,
			"dns_is_nx":                               ret.DtDdosL7DnsStats.Stats.Dns_is_nx,
			"dns_nx_drop":                             ret.DtDdosL7DnsStats.Stats.Dns_nx_drop,
			"dns_nx_bl":                               ret.DtDdosL7DnsStats.Stats.Dns_nx_bl,
			"dns_tcp_auth_pass":                       ret.DtDdosL7DnsStats.Stats.Dns_tcp_auth_pass,
			"dns_auth_udp_fail":                       ret.DtDdosL7DnsStats.Stats.Dns_auth_udp_fail,
			"dns_auth_udp_timeout":                    ret.DtDdosL7DnsStats.Stats.Dns_auth_udp_timeout,
			"dns_fqdn_label_len_exceed":               ret.DtDdosL7DnsStats.Stats.Dns_fqdn_label_len_exceed,
			"dns_pkt_processed":                       ret.DtDdosL7DnsStats.Stats.Dns_pkt_processed,
			"dns_query_type_a":                        ret.DtDdosL7DnsStats.Stats.Dns_query_type_a,
			"dns_query_type_aaaa":                     ret.DtDdosL7DnsStats.Stats.Dns_query_type_aaaa,
			"dns_query_type_ns":                       ret.DtDdosL7DnsStats.Stats.Dns_query_type_ns,
			"dns_query_type_cname":                    ret.DtDdosL7DnsStats.Stats.Dns_query_type_cname,
			"dns_query_type_any":                      ret.DtDdosL7DnsStats.Stats.Dns_query_type_any,
			"dns_query_type_srv":                      ret.DtDdosL7DnsStats.Stats.Dns_query_type_srv,
			"dns_query_type_mx":                       ret.DtDdosL7DnsStats.Stats.Dns_query_type_mx,
			"dns_query_type_soa":                      ret.DtDdosL7DnsStats.Stats.Dns_query_type_soa,
			"dns_query_type_opt":                      ret.DtDdosL7DnsStats.Stats.Dns_query_type_opt,
			"dns_dg_action_permit":                    ret.DtDdosL7DnsStats.Stats.Dns_dg_action_permit,
			"dns_dg_action_deny":                      ret.DtDdosL7DnsStats.Stats.Dns_dg_action_deny,
			"dns_fqdn_rate_by_label_count_exceed":     ret.DtDdosL7DnsStats.Stats.Dns_fqdn_rate_by_label_count_exceed,
			"dns_udp_auth_retry_gap_drop":             ret.DtDdosL7DnsStats.Stats.Dns_udp_auth_retry_gap_drop,
			"dns_policy_drop":                         ret.DtDdosL7DnsStats.Stats.Dns_policy_drop,
			"dns_fqdn_label_count_exceed":             ret.DtDdosL7DnsStats.Stats.Dns_fqdn_label_count_exceed,
			"dns_rrtype_drop":                         ret.DtDdosL7DnsStats.Stats.Dns_rrtype_drop,
			"force_tcp_auth_timeout":                  ret.DtDdosL7DnsStats.Stats.Force_tcp_auth_timeout,
			"dns_auth_drop":                           ret.DtDdosL7DnsStats.Stats.Dns_auth_drop,
			"dns_auth_resp":                           ret.DtDdosL7DnsStats.Stats.Dns_auth_resp,
			"force_tcp_auth_conn_hit":                 ret.DtDdosL7DnsStats.Stats.Force_tcp_auth_conn_hit,
			"dns_auth_udp_fail_bl":                    ret.DtDdosL7DnsStats.Stats.Dns_auth_udp_fail_bl,
			"dns_nx_exceed":                           ret.DtDdosL7DnsStats.Stats.Dns_nx_exceed,
			"dns_query_class_whitelist_miss":          ret.DtDdosL7DnsStats.Stats.Dns_query_class_whitelist_miss,
			"dns_query_class_in":                      ret.DtDdosL7DnsStats.Stats.Dns_query_class_in,
			"dns_query_class_csnet":                   ret.DtDdosL7DnsStats.Stats.Dns_query_class_csnet,
			"dns_query_class_chaos":                   ret.DtDdosL7DnsStats.Stats.Dns_query_class_chaos,
			"dns_query_class_hs":                      ret.DtDdosL7DnsStats.Stats.Dns_query_class_hs,
			"dns_query_class_none":                    ret.DtDdosL7DnsStats.Stats.Dns_query_class_none,
			"dns_query_class_any":                     ret.DtDdosL7DnsStats.Stats.Dns_query_class_any,
			"dns_dg_rate_exceed":                      ret.DtDdosL7DnsStats.Stats.Dns_dg_rate_exceed,
			"dns_outbound_query_response_size_exceed": ret.DtDdosL7DnsStats.Stats.Dns_outbound_query_response_size_exceed,
			"dns_outbound_query_sess_timed_out":       ret.DtDdosL7DnsStats.Stats.Dns_outbound_query_sess_timed_out,
			"non_query_opcode_pass_through":           ret.DtDdosL7DnsStats.Stats.Non_query_opcode_pass_through,
		},
	}
}

func getObjectDdosL7DnsStatsStats(d []interface{}) edpt.DdosL7DnsStatsStats {

	count1 := len(d)
	var ret edpt.DdosL7DnsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Force_tcp_auth = in["force_tcp_auth"].(int)
		ret.Dns_auth_udp = in["dns_auth_udp"].(int)
		ret.Dns_malform_drop = in["dns_malform_drop"].(int)
		ret.Dns_qry_any_drop = in["dns_qry_any_drop"].(int)
		ret.Dst_rate_limit0 = in["dst_rate_limit0"].(int)
		ret.Dst_rate_limit1 = in["dst_rate_limit1"].(int)
		ret.Dst_rate_limit2 = in["dst_rate_limit2"].(int)
		ret.Dst_rate_limit3 = in["dst_rate_limit3"].(int)
		ret.Dst_rate_limit4 = in["dst_rate_limit4"].(int)
		ret.Src_rate_limit0 = in["src_rate_limit0"].(int)
		ret.Src_rate_limit1 = in["src_rate_limit1"].(int)
		ret.Src_rate_limit2 = in["src_rate_limit2"].(int)
		ret.Src_rate_limit3 = in["src_rate_limit3"].(int)
		ret.Src_rate_limit4 = in["src_rate_limit4"].(int)
		ret.Dns_auth_udp_pass = in["dns_auth_udp_pass"].(int)
		ret.Dns_fqdn_stage2_exceed = in["dns_fqdn_stage2_exceed"].(int)
		ret.Dns_is_nx = in["dns_is_nx"].(int)
		ret.Dns_nx_drop = in["dns_nx_drop"].(int)
		ret.Dns_nx_bl = in["dns_nx_bl"].(int)
		ret.Dns_tcp_auth_pass = in["dns_tcp_auth_pass"].(int)
		ret.Dns_auth_udp_fail = in["dns_auth_udp_fail"].(int)
		ret.Dns_auth_udp_timeout = in["dns_auth_udp_timeout"].(int)
		ret.Dns_fqdn_label_len_exceed = in["dns_fqdn_label_len_exceed"].(int)
		ret.Dns_pkt_processed = in["dns_pkt_processed"].(int)
		ret.Dns_query_type_a = in["dns_query_type_a"].(int)
		ret.Dns_query_type_aaaa = in["dns_query_type_aaaa"].(int)
		ret.Dns_query_type_ns = in["dns_query_type_ns"].(int)
		ret.Dns_query_type_cname = in["dns_query_type_cname"].(int)
		ret.Dns_query_type_any = in["dns_query_type_any"].(int)
		ret.Dns_query_type_srv = in["dns_query_type_srv"].(int)
		ret.Dns_query_type_mx = in["dns_query_type_mx"].(int)
		ret.Dns_query_type_soa = in["dns_query_type_soa"].(int)
		ret.Dns_query_type_opt = in["dns_query_type_opt"].(int)
		ret.Dns_dg_action_permit = in["dns_dg_action_permit"].(int)
		ret.Dns_dg_action_deny = in["dns_dg_action_deny"].(int)
		ret.Dns_fqdn_rate_by_label_count_exceed = in["dns_fqdn_rate_by_label_count_exceed"].(int)
		ret.Dns_udp_auth_retry_gap_drop = in["dns_udp_auth_retry_gap_drop"].(int)
		ret.Dns_policy_drop = in["dns_policy_drop"].(int)
		ret.Dns_fqdn_label_count_exceed = in["dns_fqdn_label_count_exceed"].(int)
		ret.Dns_rrtype_drop = in["dns_rrtype_drop"].(int)
		ret.Force_tcp_auth_timeout = in["force_tcp_auth_timeout"].(int)
		ret.Dns_auth_drop = in["dns_auth_drop"].(int)
		ret.Dns_auth_resp = in["dns_auth_resp"].(int)
		ret.Force_tcp_auth_conn_hit = in["force_tcp_auth_conn_hit"].(int)
		ret.Dns_auth_udp_fail_bl = in["dns_auth_udp_fail_bl"].(int)
		ret.Dns_nx_exceed = in["dns_nx_exceed"].(int)
		ret.Dns_query_class_whitelist_miss = in["dns_query_class_whitelist_miss"].(int)
		ret.Dns_query_class_in = in["dns_query_class_in"].(int)
		ret.Dns_query_class_csnet = in["dns_query_class_csnet"].(int)
		ret.Dns_query_class_chaos = in["dns_query_class_chaos"].(int)
		ret.Dns_query_class_hs = in["dns_query_class_hs"].(int)
		ret.Dns_query_class_none = in["dns_query_class_none"].(int)
		ret.Dns_query_class_any = in["dns_query_class_any"].(int)
		ret.Dns_dg_rate_exceed = in["dns_dg_rate_exceed"].(int)
		ret.Dns_outbound_query_response_size_exceed = in["dns_outbound_query_response_size_exceed"].(int)
		ret.Dns_outbound_query_sess_timed_out = in["dns_outbound_query_sess_timed_out"].(int)
		ret.Non_query_opcode_pass_through = in["non_query_opcode_pass_through"].(int)
	}
	return ret
}

func dataToEndpointDdosL7DnsStats(d *schema.ResourceData) edpt.DdosL7DnsStats {
	var ret edpt.DdosL7DnsStats

	ret.Stats = getObjectDdosL7DnsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
