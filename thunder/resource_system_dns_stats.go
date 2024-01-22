package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDnsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_dns_stats`: Statistics for the object dns\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemDnsStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"slb_req": {
							Type: schema.TypeInt, Optional: true, Description: "No. of requests",
						},
						"slb_resp": {
							Type: schema.TypeInt, Optional: true, Description: "No. of responses",
						},
						"slb_no_resp": {
							Type: schema.TypeInt, Optional: true, Description: "No. of requests with no response",
						},
						"slb_req_rexmit": {
							Type: schema.TypeInt, Optional: true, Description: "No. of requests retransmit",
						},
						"slb_resp_no_match": {
							Type: schema.TypeInt, Optional: true, Description: "No. of requests and responses with no match",
						},
						"slb_no_resource": {
							Type: schema.TypeInt, Optional: true, Description: "No. of resource failures",
						},
						"nat_req": {
							Type: schema.TypeInt, Optional: true, Description: "(NAT) No. of requests",
						},
						"nat_resp": {
							Type: schema.TypeInt, Optional: true, Description: "(NAT) No. of responses",
						},
						"nat_no_resp": {
							Type: schema.TypeInt, Optional: true, Description: "(NAT) No. of resource failures",
						},
						"nat_req_rexmit": {
							Type: schema.TypeInt, Optional: true, Description: "(NAT) No. of request retransmits",
						},
						"nat_resp_no_match": {
							Type: schema.TypeInt, Optional: true, Description: "(NAT) No. of requests with no response",
						},
						"nat_no_resource": {
							Type: schema.TypeInt, Optional: true, Description: "(NAT) No. of resource failures",
						},
						"nat_xid_reused": {
							Type: schema.TypeInt, Optional: true, Description: "(NAT) No. of requests reusing a transaction id",
						},
						"filter_type_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total Query Type Drop",
						},
						"filter_class_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total Query Class Drop",
						},
						"filter_type_any_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total Query ANY Type Drop",
						},
						"slb_dns_client_ssl_succ": {
							Type: schema.TypeInt, Optional: true, Description: "No. of client ssl success",
						},
						"slb_dns_server_ssl_succ": {
							Type: schema.TypeInt, Optional: true, Description: "No. of server ssl success",
						},
						"slb_dns_udp_conn": {
							Type: schema.TypeInt, Optional: true, Description: "No. of backend udp connections",
						},
						"slb_dns_udp_conn_succ": {
							Type: schema.TypeInt, Optional: true, Description: "No. of backend udp conn established",
						},
						"slb_dns_padding_to_server_removed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"slb_dns_padding_to_client_added": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"slb_dns_edns_subnet_to_server_removed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"slb_dns_udp_retransmit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"slb_dns_udp_retransmit_fail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rpz_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "RPZ Action Drop",
						},
						"rpz_action_pass_thru": {
							Type: schema.TypeInt, Optional: true, Description: "RPZ Action Pass Through",
						},
						"rpz_action_tcp_only": {
							Type: schema.TypeInt, Optional: true, Description: "RPZ Action TCP Only",
						},
						"rpz_action_nxdomain": {
							Type: schema.TypeInt, Optional: true, Description: "RPZ Action NXDOMAIN",
						},
						"rpz_action_nodata": {
							Type: schema.TypeInt, Optional: true, Description: "RPZ Action NODATA",
						},
						"rpz_action_local_data": {
							Type: schema.TypeInt, Optional: true, Description: "RPZ Action Local Data",
						},
						"slb_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS requests drop",
						},
						"nat_slb_drop": {
							Type: schema.TypeInt, Optional: true, Description: "(NAT)DNS requests drop",
						},
						"invalid_q_len_to_udp": {
							Type: schema.TypeInt, Optional: true, Description: "invalid query length to conver to UDP",
						},
						"slb_dns_edns_ecs_received": {
							Type: schema.TypeInt, Optional: true, Description: "Number of ecs from client received",
						},
						"slb_dns_edns_ecs_inserted": {
							Type: schema.TypeInt, Optional: true, Description: "Number of ecs inserted",
						},
					},
				},
			},
		},
	}
}

func resourceSystemDnsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDnsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemDnsStatsStats := setObjectSystemDnsStatsStats(res)
		d.Set("stats", SystemDnsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemDnsStatsStats(ret edpt.DataSystemDnsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"slb_req":                               ret.DtSystemDnsStats.Stats.Slb_req,
			"slb_resp":                              ret.DtSystemDnsStats.Stats.Slb_resp,
			"slb_no_resp":                           ret.DtSystemDnsStats.Stats.Slb_no_resp,
			"slb_req_rexmit":                        ret.DtSystemDnsStats.Stats.Slb_req_rexmit,
			"slb_resp_no_match":                     ret.DtSystemDnsStats.Stats.Slb_resp_no_match,
			"slb_no_resource":                       ret.DtSystemDnsStats.Stats.Slb_no_resource,
			"nat_req":                               ret.DtSystemDnsStats.Stats.Nat_req,
			"nat_resp":                              ret.DtSystemDnsStats.Stats.Nat_resp,
			"nat_no_resp":                           ret.DtSystemDnsStats.Stats.Nat_no_resp,
			"nat_req_rexmit":                        ret.DtSystemDnsStats.Stats.Nat_req_rexmit,
			"nat_resp_no_match":                     ret.DtSystemDnsStats.Stats.Nat_resp_no_match,
			"nat_no_resource":                       ret.DtSystemDnsStats.Stats.Nat_no_resource,
			"nat_xid_reused":                        ret.DtSystemDnsStats.Stats.Nat_xid_reused,
			"filter_type_drop":                      ret.DtSystemDnsStats.Stats.Filter_type_drop,
			"filter_class_drop":                     ret.DtSystemDnsStats.Stats.Filter_class_drop,
			"filter_type_any_drop":                  ret.DtSystemDnsStats.Stats.Filter_type_any_drop,
			"slb_dns_client_ssl_succ":               ret.DtSystemDnsStats.Stats.Slb_dns_client_ssl_succ,
			"slb_dns_server_ssl_succ":               ret.DtSystemDnsStats.Stats.Slb_dns_server_ssl_succ,
			"slb_dns_udp_conn":                      ret.DtSystemDnsStats.Stats.Slb_dns_udp_conn,
			"slb_dns_udp_conn_succ":                 ret.DtSystemDnsStats.Stats.Slb_dns_udp_conn_succ,
			"slb_dns_padding_to_server_removed":     ret.DtSystemDnsStats.Stats.Slb_dns_padding_to_server_removed,
			"slb_dns_padding_to_client_added":       ret.DtSystemDnsStats.Stats.Slb_dns_padding_to_client_added,
			"slb_dns_edns_subnet_to_server_removed": ret.DtSystemDnsStats.Stats.Slb_dns_edns_subnet_to_server_removed,
			"slb_dns_udp_retransmit":                ret.DtSystemDnsStats.Stats.Slb_dns_udp_retransmit,
			"slb_dns_udp_retransmit_fail":           ret.DtSystemDnsStats.Stats.Slb_dns_udp_retransmit_fail,
			"rpz_action_drop":                       ret.DtSystemDnsStats.Stats.Rpz_action_drop,
			"rpz_action_pass_thru":                  ret.DtSystemDnsStats.Stats.Rpz_action_pass_thru,
			"rpz_action_tcp_only":                   ret.DtSystemDnsStats.Stats.Rpz_action_tcp_only,
			"rpz_action_nxdomain":                   ret.DtSystemDnsStats.Stats.Rpz_action_nxdomain,
			"rpz_action_nodata":                     ret.DtSystemDnsStats.Stats.Rpz_action_nodata,
			"rpz_action_local_data":                 ret.DtSystemDnsStats.Stats.Rpz_action_local_data,
			"slb_drop":                              ret.DtSystemDnsStats.Stats.Slb_drop,
			"nat_slb_drop":                          ret.DtSystemDnsStats.Stats.Nat_slb_drop,
			"invalid_q_len_to_udp":                  ret.DtSystemDnsStats.Stats.Invalid_q_len_to_udp,
			"slb_dns_edns_ecs_received":             ret.DtSystemDnsStats.Stats.Slb_dns_edns_ecs_received,
			"slb_dns_edns_ecs_inserted":             ret.DtSystemDnsStats.Stats.Slb_dns_edns_ecs_inserted,
		},
	}
}

func getObjectSystemDnsStatsStats(d []interface{}) edpt.SystemDnsStatsStats {

	count1 := len(d)
	var ret edpt.SystemDnsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Slb_req = in["slb_req"].(int)
		ret.Slb_resp = in["slb_resp"].(int)
		ret.Slb_no_resp = in["slb_no_resp"].(int)
		ret.Slb_req_rexmit = in["slb_req_rexmit"].(int)
		ret.Slb_resp_no_match = in["slb_resp_no_match"].(int)
		ret.Slb_no_resource = in["slb_no_resource"].(int)
		ret.Nat_req = in["nat_req"].(int)
		ret.Nat_resp = in["nat_resp"].(int)
		ret.Nat_no_resp = in["nat_no_resp"].(int)
		ret.Nat_req_rexmit = in["nat_req_rexmit"].(int)
		ret.Nat_resp_no_match = in["nat_resp_no_match"].(int)
		ret.Nat_no_resource = in["nat_no_resource"].(int)
		ret.Nat_xid_reused = in["nat_xid_reused"].(int)
		ret.Filter_type_drop = in["filter_type_drop"].(int)
		ret.Filter_class_drop = in["filter_class_drop"].(int)
		ret.Filter_type_any_drop = in["filter_type_any_drop"].(int)
		ret.Slb_dns_client_ssl_succ = in["slb_dns_client_ssl_succ"].(int)
		ret.Slb_dns_server_ssl_succ = in["slb_dns_server_ssl_succ"].(int)
		ret.Slb_dns_udp_conn = in["slb_dns_udp_conn"].(int)
		ret.Slb_dns_udp_conn_succ = in["slb_dns_udp_conn_succ"].(int)
		ret.Slb_dns_padding_to_server_removed = in["slb_dns_padding_to_server_removed"].(int)
		ret.Slb_dns_padding_to_client_added = in["slb_dns_padding_to_client_added"].(int)
		ret.Slb_dns_edns_subnet_to_server_removed = in["slb_dns_edns_subnet_to_server_removed"].(int)
		ret.Slb_dns_udp_retransmit = in["slb_dns_udp_retransmit"].(int)
		ret.Slb_dns_udp_retransmit_fail = in["slb_dns_udp_retransmit_fail"].(int)
		ret.Rpz_action_drop = in["rpz_action_drop"].(int)
		ret.Rpz_action_pass_thru = in["rpz_action_pass_thru"].(int)
		ret.Rpz_action_tcp_only = in["rpz_action_tcp_only"].(int)
		ret.Rpz_action_nxdomain = in["rpz_action_nxdomain"].(int)
		ret.Rpz_action_nodata = in["rpz_action_nodata"].(int)
		ret.Rpz_action_local_data = in["rpz_action_local_data"].(int)
		ret.Slb_drop = in["slb_drop"].(int)
		ret.Nat_slb_drop = in["nat_slb_drop"].(int)
		ret.Invalid_q_len_to_udp = in["invalid_q_len_to_udp"].(int)
		ret.Slb_dns_edns_ecs_received = in["slb_dns_edns_ecs_received"].(int)
		ret.Slb_dns_edns_ecs_inserted = in["slb_dns_edns_ecs_inserted"].(int)
	}
	return ret
}

func dataToEndpointSystemDnsStats(d *schema.ResourceData) edpt.SystemDnsStats {
	var ret edpt.SystemDnsStats

	ret.Stats = getObjectSystemDnsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
