package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheServerStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dns_cache_server_stats`: Statistics for the object dns-cache-server\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDnsCacheServerStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"insert_total": {
							Type: schema.TypeInt, Optional: true, Description: "Insert Total",
						},
						"insert_success": {
							Type: schema.TypeInt, Optional: true, Description: "Insert Success",
						},
						"insert_fail_all": {
							Type: schema.TypeInt, Optional: true, Description: "Insert Fail",
						},
						"lookup_invalid_domain": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Invalid Domain",
						},
						"lookup_unexp_err": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Unexpected Error",
						},
						"lookup_full_matched": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Fully Matched",
						},
						"lookup_empty_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Empty Response",
						},
						"lookup_deleg_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Delegation Response",
						},
						"lookup_nxdomain_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Nxdomain Response",
						},
						"lookup_refuse_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Refuse Response",
						},
						"lookup_fwd_server": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Forwarded to Server",
						},
						"lookup_incomp_zone": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Incomplete Zone",
						},
						"lookup_undefined_rtype": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Undefined Record Type",
						},
						"lookup_manual_override_action_forward": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS Manual Override Action Forward",
						},
						"lookup_manual_override_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS Manual Override Action Drop",
						},
						"zt_serial_num_check_attempts": {
							Type: schema.TypeInt, Optional: true, Description: "Zone Transfer Serial Number Check Started",
						},
						"zt_axfr_attempts": {
							Type: schema.TypeInt, Optional: true, Description: "Zone Transfer AXFR Started",
						},
						"zt_completed_ok": {
							Type: schema.TypeInt, Optional: true, Description: "Zone Transfer Completed",
						},
						"zt_completed_no_update": {
							Type: schema.TypeInt, Optional: true, Description: "Zone Transfer Completed No Update",
						},
						"zt_dns_process_err": {
							Type: schema.TypeInt, Optional: true, Description: "Zone Transfer DNS Processing Errors",
						},
						"zt_records_processed": {
							Type: schema.TypeInt, Optional: true, Description: "Zone Transfer Records Processed",
						},
						"lookup_edns_bad_version_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup EDNS Bad Version Response",
						},
						"zt_tcp_conn_connect_server_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Zone Transfer TCP Connect Server Fail",
						},
						"zt_tcp_conn_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Zone Transfer TCP RST / FIN Received",
						},
						"zt_task_no_route_retry": {
							Type: schema.TypeInt, Optional: true, Description: "Zone Transfer Task No Route Fail",
						},
						"zt_msg_rcode_notauth": {
							Type: schema.TypeInt, Optional: true, Description: "Zone Transfer Server Not Auth Fail",
						},
						"lookup_opcode_notimpl_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Opcode Not Implemented Response",
						},
						"shard_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Shard Filter Matched",
						},
						"zt_total_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Zone Transfer Total Failure",
						},
						"lookup_manual_override_action_serve": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS Manual Override Action Serve",
						},
						"lookup_any_type_query_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS ANY Type Query Action Drop",
						},
						"lookup_any_type_query_action_refused": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS ANY Type Query Action Refused",
						},
						"lookup_any_type_query_action_resp_empty": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS ANY Type Query Action Response Empty",
						},
						"lookup_non_auth_zone_query_action_forward": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS Non-Authoritative Zone Query Action Forward",
						},
						"lookup_non_auth_zone_query_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS Non-Authoritative Zone Query Action Drop",
						},
						"lookup_non_auth_zone_query_action_resp_refused": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS Non-Authoritative Zone Query Action Refused",
						},
						"lookup_default_action_forward": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS Default Action Forward",
						},
						"lookup_default_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS Default Action Drop",
						},
						"zt_ongoing_tasks": {
							Type: schema.TypeInt, Optional: true, Description: "Zone Transfer Ongoing tasks",
						},
						"lookup_dnstcp_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS-TCP Request Received",
						},
						"lookup_dnsudp_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup DNS-UDP Request Received",
						},
						"lookup_fwd_shard": {
							Type: schema.TypeInt, Optional: true, Description: "Lookup Forwarded to Sharding DNS Cache",
						},
					},
				},
			},
		},
	}
}

func resourceDdosDnsCacheServerStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheServerStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheServerStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDnsCacheServerStatsStats := setObjectDdosDnsCacheServerStatsStats(res)
		d.Set("stats", DdosDnsCacheServerStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDnsCacheServerStatsStats(ret edpt.DataDdosDnsCacheServerStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"insert_total":                                   ret.DtDdosDnsCacheServerStats.Stats.Insert_total,
			"insert_success":                                 ret.DtDdosDnsCacheServerStats.Stats.Insert_success,
			"insert_fail_all":                                ret.DtDdosDnsCacheServerStats.Stats.Insert_fail_all,
			"lookup_invalid_domain":                          ret.DtDdosDnsCacheServerStats.Stats.Lookup_invalid_domain,
			"lookup_unexp_err":                               ret.DtDdosDnsCacheServerStats.Stats.Lookup_unexp_err,
			"lookup_full_matched":                            ret.DtDdosDnsCacheServerStats.Stats.Lookup_full_matched,
			"lookup_empty_resp":                              ret.DtDdosDnsCacheServerStats.Stats.Lookup_empty_resp,
			"lookup_deleg_resp":                              ret.DtDdosDnsCacheServerStats.Stats.Lookup_deleg_resp,
			"lookup_nxdomain_resp":                           ret.DtDdosDnsCacheServerStats.Stats.Lookup_nxdomain_resp,
			"lookup_refuse_resp":                             ret.DtDdosDnsCacheServerStats.Stats.Lookup_refuse_resp,
			"lookup_fwd_server":                              ret.DtDdosDnsCacheServerStats.Stats.Lookup_fwd_server,
			"lookup_incomp_zone":                             ret.DtDdosDnsCacheServerStats.Stats.Lookup_incomp_zone,
			"lookup_undefined_rtype":                         ret.DtDdosDnsCacheServerStats.Stats.Lookup_undefined_rtype,
			"lookup_manual_override_action_forward":          ret.DtDdosDnsCacheServerStats.Stats.Lookup_manual_override_action_forward,
			"lookup_manual_override_action_drop":             ret.DtDdosDnsCacheServerStats.Stats.Lookup_manual_override_action_drop,
			"zt_serial_num_check_attempts":                   ret.DtDdosDnsCacheServerStats.Stats.Zt_serial_num_check_attempts,
			"zt_axfr_attempts":                               ret.DtDdosDnsCacheServerStats.Stats.Zt_axfr_attempts,
			"zt_completed_ok":                                ret.DtDdosDnsCacheServerStats.Stats.Zt_completed_ok,
			"zt_completed_no_update":                         ret.DtDdosDnsCacheServerStats.Stats.Zt_completed_no_update,
			"zt_dns_process_err":                             ret.DtDdosDnsCacheServerStats.Stats.Zt_dns_process_err,
			"zt_records_processed":                           ret.DtDdosDnsCacheServerStats.Stats.Zt_records_processed,
			"lookup_edns_bad_version_resp":                   ret.DtDdosDnsCacheServerStats.Stats.Lookup_edns_bad_version_resp,
			"zt_tcp_conn_connect_server_fail":                ret.DtDdosDnsCacheServerStats.Stats.Zt_tcp_conn_connect_server_fail,
			"zt_tcp_conn_rst":                                ret.DtDdosDnsCacheServerStats.Stats.Zt_tcp_conn_rst,
			"zt_task_no_route_retry":                         ret.DtDdosDnsCacheServerStats.Stats.Zt_task_no_route_retry,
			"zt_msg_rcode_notauth":                           ret.DtDdosDnsCacheServerStats.Stats.Zt_msg_rcode_notauth,
			"lookup_opcode_notimpl_resp":                     ret.DtDdosDnsCacheServerStats.Stats.Lookup_opcode_notimpl_resp,
			"shard_filter_match":                             ret.DtDdosDnsCacheServerStats.Stats.Shard_filter_match,
			"zt_total_fail":                                  ret.DtDdosDnsCacheServerStats.Stats.Zt_total_fail,
			"lookup_manual_override_action_serve":            ret.DtDdosDnsCacheServerStats.Stats.Lookup_manual_override_action_serve,
			"lookup_any_type_query_action_drop":              ret.DtDdosDnsCacheServerStats.Stats.Lookup_any_type_query_action_drop,
			"lookup_any_type_query_action_refused":           ret.DtDdosDnsCacheServerStats.Stats.Lookup_any_type_query_action_refused,
			"lookup_any_type_query_action_resp_empty":        ret.DtDdosDnsCacheServerStats.Stats.Lookup_any_type_query_action_resp_empty,
			"lookup_non_auth_zone_query_action_forward":      ret.DtDdosDnsCacheServerStats.Stats.Lookup_non_auth_zone_query_action_forward,
			"lookup_non_auth_zone_query_action_drop":         ret.DtDdosDnsCacheServerStats.Stats.Lookup_non_auth_zone_query_action_drop,
			"lookup_non_auth_zone_query_action_resp_refused": ret.DtDdosDnsCacheServerStats.Stats.Lookup_non_auth_zone_query_action_resp_refused,
			"lookup_default_action_forward":                  ret.DtDdosDnsCacheServerStats.Stats.Lookup_default_action_forward,
			"lookup_default_action_drop":                     ret.DtDdosDnsCacheServerStats.Stats.Lookup_default_action_drop,
			"zt_ongoing_tasks":                               ret.DtDdosDnsCacheServerStats.Stats.Zt_ongoing_tasks,
			"lookup_dnstcp_rcvd":                             ret.DtDdosDnsCacheServerStats.Stats.Lookup_dnstcp_rcvd,
			"lookup_dnsudp_rcvd":                             ret.DtDdosDnsCacheServerStats.Stats.Lookup_dnsudp_rcvd,
			"lookup_fwd_shard":                               ret.DtDdosDnsCacheServerStats.Stats.Lookup_fwd_shard,
		},
	}
}

func getObjectDdosDnsCacheServerStatsStats(d []interface{}) edpt.DdosDnsCacheServerStatsStats {

	count1 := len(d)
	var ret edpt.DdosDnsCacheServerStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Insert_total = in["insert_total"].(int)
		ret.Insert_success = in["insert_success"].(int)
		ret.Insert_fail_all = in["insert_fail_all"].(int)
		ret.Lookup_invalid_domain = in["lookup_invalid_domain"].(int)
		ret.Lookup_unexp_err = in["lookup_unexp_err"].(int)
		ret.Lookup_full_matched = in["lookup_full_matched"].(int)
		ret.Lookup_empty_resp = in["lookup_empty_resp"].(int)
		ret.Lookup_deleg_resp = in["lookup_deleg_resp"].(int)
		ret.Lookup_nxdomain_resp = in["lookup_nxdomain_resp"].(int)
		ret.Lookup_refuse_resp = in["lookup_refuse_resp"].(int)
		ret.Lookup_fwd_server = in["lookup_fwd_server"].(int)
		ret.Lookup_incomp_zone = in["lookup_incomp_zone"].(int)
		ret.Lookup_undefined_rtype = in["lookup_undefined_rtype"].(int)
		ret.Lookup_manual_override_action_forward = in["lookup_manual_override_action_forward"].(int)
		ret.Lookup_manual_override_action_drop = in["lookup_manual_override_action_drop"].(int)
		ret.Zt_serial_num_check_attempts = in["zt_serial_num_check_attempts"].(int)
		ret.Zt_axfr_attempts = in["zt_axfr_attempts"].(int)
		ret.Zt_completed_ok = in["zt_completed_ok"].(int)
		ret.Zt_completed_no_update = in["zt_completed_no_update"].(int)
		ret.Zt_dns_process_err = in["zt_dns_process_err"].(int)
		ret.Zt_records_processed = in["zt_records_processed"].(int)
		ret.Lookup_edns_bad_version_resp = in["lookup_edns_bad_version_resp"].(int)
		ret.Zt_tcp_conn_connect_server_fail = in["zt_tcp_conn_connect_server_fail"].(int)
		ret.Zt_tcp_conn_rst = in["zt_tcp_conn_rst"].(int)
		ret.Zt_task_no_route_retry = in["zt_task_no_route_retry"].(int)
		ret.Zt_msg_rcode_notauth = in["zt_msg_rcode_notauth"].(int)
		ret.Lookup_opcode_notimpl_resp = in["lookup_opcode_notimpl_resp"].(int)
		ret.Shard_filter_match = in["shard_filter_match"].(int)
		ret.Zt_total_fail = in["zt_total_fail"].(int)
		ret.Lookup_manual_override_action_serve = in["lookup_manual_override_action_serve"].(int)
		ret.Lookup_any_type_query_action_drop = in["lookup_any_type_query_action_drop"].(int)
		ret.Lookup_any_type_query_action_refused = in["lookup_any_type_query_action_refused"].(int)
		ret.Lookup_any_type_query_action_resp_empty = in["lookup_any_type_query_action_resp_empty"].(int)
		ret.Lookup_non_auth_zone_query_action_forward = in["lookup_non_auth_zone_query_action_forward"].(int)
		ret.Lookup_non_auth_zone_query_action_drop = in["lookup_non_auth_zone_query_action_drop"].(int)
		ret.Lookup_non_auth_zone_query_action_resp_refused = in["lookup_non_auth_zone_query_action_resp_refused"].(int)
		ret.Lookup_default_action_forward = in["lookup_default_action_forward"].(int)
		ret.Lookup_default_action_drop = in["lookup_default_action_drop"].(int)
		ret.Zt_ongoing_tasks = in["zt_ongoing_tasks"].(int)
		ret.Lookup_dnstcp_rcvd = in["lookup_dnstcp_rcvd"].(int)
		ret.Lookup_dnsudp_rcvd = in["lookup_dnsudp_rcvd"].(int)
		ret.Lookup_fwd_shard = in["lookup_fwd_shard"].(int)
	}
	return ret
}

func dataToEndpointDdosDnsCacheServerStats(d *schema.ResourceData) edpt.DdosDnsCacheServerStats {
	var ret edpt.DdosDnsCacheServerStats

	ret.Stats = getObjectDdosDnsCacheServerStatsStats(d.Get("stats").([]interface{}))
	return ret
}
