package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Dns64VirtualserverPortStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_dns64_virtualserver_port_stats`: Statistics for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Dns64VirtualserverPortStatsRead,

		Schema: map[string]*schema.Schema{
			"port_number": {
				Type: schema.TypeInt, Required: true, Description: "Port",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'dns-udp': DNS service over UDP;",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Current connection",
						},
						"total_l4_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total L4 connections",
						},
						"total_l7_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total L7 connections",
						},
						"toatal_tcp_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP connections",
						},
						"total_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total connections",
						},
						"total_fwd_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Total forward bytes",
						},
						"total_fwd_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Total forward packets",
						},
						"total_rev_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Total reverse bytes",
						},
						"total_rev_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Total reverse packets",
						},
						"total_dns_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Total DNS packets",
						},
						"total_mf_dns_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Total MF DNS packets",
						},
						"es_total_failure_actions": {
							Type: schema.TypeInt, Optional: true, Description: "Total failure actions",
						},
						"compression_bytes_before": {
							Type: schema.TypeInt, Optional: true, Description: "Data into compression engine",
						},
						"compression_bytes_after": {
							Type: schema.TypeInt, Optional: true, Description: "Data out of compression engine",
						},
						"compression_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests compressed",
						},
						"compression_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests NOT compressed",
						},
						"compression_miss_no_client": {
							Type: schema.TypeInt, Optional: true, Description: "Compression miss no client",
						},
						"compression_miss_template_exclusion": {
							Type: schema.TypeInt, Optional: true, Description: "Compression miss template exclusion",
						},
						"curr_req": {
							Type: schema.TypeInt, Optional: true, Description: "Current requests",
						},
						"total_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total requests",
						},
						"total_req_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Total successful requests",
						},
						"peak_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Peak connections",
						},
						"curr_conn_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Current connection rate",
						},
						"last_rsp_time": {
							Type: schema.TypeInt, Optional: true, Description: "Last response time",
						},
						"fastest_rsp_time": {
							Type: schema.TypeInt, Optional: true, Description: "Fastest response time",
						},
						"slowest_rsp_time": {
							Type: schema.TypeInt, Optional: true, Description: "Slowest response time",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceCgnv6Dns64VirtualserverPortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Dns64VirtualserverPortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Dns64VirtualserverPortStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Dns64VirtualserverPortStatsStats := setObjectCgnv6Dns64VirtualserverPortStatsStats(res)
		d.Set("stats", Cgnv6Dns64VirtualserverPortStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Dns64VirtualserverPortStatsStats(ret edpt.DataCgnv6Dns64VirtualserverPortStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_conn":                           ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Curr_conn,
			"total_l4_conn":                       ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Total_l4_conn,
			"total_l7_conn":                       ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Total_l7_conn,
			"toatal_tcp_conn":                     ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Toatal_tcp_conn,
			"total_conn":                          ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Total_conn,
			"total_fwd_bytes":                     ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Total_fwd_bytes,
			"total_fwd_pkts":                      ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Total_fwd_pkts,
			"total_rev_bytes":                     ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Total_rev_bytes,
			"total_rev_pkts":                      ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Total_rev_pkts,
			"total_dns_pkts":                      ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Total_dns_pkts,
			"total_mf_dns_pkts":                   ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Total_mf_dns_pkts,
			"es_total_failure_actions":            ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Es_total_failure_actions,
			"compression_bytes_before":            ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Compression_bytes_before,
			"compression_bytes_after":             ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Compression_bytes_after,
			"compression_hit":                     ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Compression_hit,
			"compression_miss":                    ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Compression_miss,
			"compression_miss_no_client":          ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Compression_miss_no_client,
			"compression_miss_template_exclusion": ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Compression_miss_template_exclusion,
			"curr_req":                            ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Curr_req,
			"total_req":                           ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Total_req,
			"total_req_succ":                      ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Total_req_succ,
			"peak_conn":                           ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Peak_conn,
			"curr_conn_rate":                      ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Curr_conn_rate,
			"last_rsp_time":                       ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Last_rsp_time,
			"fastest_rsp_time":                    ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Fastest_rsp_time,
			"slowest_rsp_time":                    ret.DtCgnv6Dns64VirtualserverPortStats.Stats.Slowest_rsp_time,
		},
	}
}

func getObjectCgnv6Dns64VirtualserverPortStatsStats(d []interface{}) edpt.Cgnv6Dns64VirtualserverPortStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6Dns64VirtualserverPortStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_conn = in["curr_conn"].(int)
		ret.Total_l4_conn = in["total_l4_conn"].(int)
		ret.Total_l7_conn = in["total_l7_conn"].(int)
		ret.Toatal_tcp_conn = in["toatal_tcp_conn"].(int)
		ret.Total_conn = in["total_conn"].(int)
		ret.Total_fwd_bytes = in["total_fwd_bytes"].(int)
		ret.Total_fwd_pkts = in["total_fwd_pkts"].(int)
		ret.Total_rev_bytes = in["total_rev_bytes"].(int)
		ret.Total_rev_pkts = in["total_rev_pkts"].(int)
		ret.Total_dns_pkts = in["total_dns_pkts"].(int)
		ret.Total_mf_dns_pkts = in["total_mf_dns_pkts"].(int)
		ret.Es_total_failure_actions = in["es_total_failure_actions"].(int)
		ret.Compression_bytes_before = in["compression_bytes_before"].(int)
		ret.Compression_bytes_after = in["compression_bytes_after"].(int)
		ret.Compression_hit = in["compression_hit"].(int)
		ret.Compression_miss = in["compression_miss"].(int)
		ret.Compression_miss_no_client = in["compression_miss_no_client"].(int)
		ret.Compression_miss_template_exclusion = in["compression_miss_template_exclusion"].(int)
		ret.Curr_req = in["curr_req"].(int)
		ret.Total_req = in["total_req"].(int)
		ret.Total_req_succ = in["total_req_succ"].(int)
		ret.Peak_conn = in["peak_conn"].(int)
		ret.Curr_conn_rate = in["curr_conn_rate"].(int)
		ret.Last_rsp_time = in["last_rsp_time"].(int)
		ret.Fastest_rsp_time = in["fastest_rsp_time"].(int)
		ret.Slowest_rsp_time = in["slowest_rsp_time"].(int)
	}
	return ret
}

func dataToEndpointCgnv6Dns64VirtualserverPortStats(d *schema.ResourceData) edpt.Cgnv6Dns64VirtualserverPortStats {
	var ret edpt.Cgnv6Dns64VirtualserverPortStats

	ret.PortNumber = d.Get("port_number").(int)

	ret.Protocol = d.Get("protocol").(string)

	ret.Stats = getObjectCgnv6Dns64VirtualserverPortStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
