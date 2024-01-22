package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6ServiceGroupMemberStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_service_group_member_stats`: Statistics for the object member\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6ServiceGroupMemberStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Member name",
			},
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Port number",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Current connections",
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
						"total_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total connections",
						},
						"total_rev_pkts_inspected": {
							Type: schema.TypeInt, Optional: true, Description: "Total reverse packets inspected",
						},
						"total_rev_pkts_inspected_status_code_2xx": {
							Type: schema.TypeInt, Optional: true, Description: "Total reverse packets inspected status code 2xx",
						},
						"total_rev_pkts_inspected_status_code_non_5xx": {
							Type: schema.TypeInt, Optional: true, Description: "Total reverse packets inspected status code non 5xx",
						},
						"curr_req": {
							Type: schema.TypeInt, Optional: true, Description: "Current requests",
						},
						"total_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total requests",
						},
						"total_req_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Total requests success",
						},
						"peak_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Peak connections",
						},
						"response_time": {
							Type: schema.TypeInt, Optional: true, Description: "Response time",
						},
						"fastest_rsp_time": {
							Type: schema.TypeInt, Optional: true, Description: "Fastest response time",
						},
						"slowest_rsp_time": {
							Type: schema.TypeInt, Optional: true, Description: "Slowest response time",
						},
						"curr_ssl_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Current SSL connections",
						},
						"total_ssl_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total SSL connections",
						},
						"curr_conn_overflow": {
							Type: schema.TypeInt, Optional: true, Description: "Current connection counter overflow count",
						},
						"state_flaps": {
							Type: schema.TypeInt, Optional: true, Description: "State flaps count",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6ServiceGroupMemberStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ServiceGroupMemberStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ServiceGroupMemberStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6ServiceGroupMemberStatsStats := setObjectCgnv6ServiceGroupMemberStatsStats(res)
		d.Set("stats", Cgnv6ServiceGroupMemberStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6ServiceGroupMemberStatsStats(ret edpt.DataCgnv6ServiceGroupMemberStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_conn":                ret.DtCgnv6ServiceGroupMemberStats.Stats.Curr_conn,
			"total_fwd_bytes":          ret.DtCgnv6ServiceGroupMemberStats.Stats.Total_fwd_bytes,
			"total_fwd_pkts":           ret.DtCgnv6ServiceGroupMemberStats.Stats.Total_fwd_pkts,
			"total_rev_bytes":          ret.DtCgnv6ServiceGroupMemberStats.Stats.Total_rev_bytes,
			"total_rev_pkts":           ret.DtCgnv6ServiceGroupMemberStats.Stats.Total_rev_pkts,
			"total_conn":               ret.DtCgnv6ServiceGroupMemberStats.Stats.Total_conn,
			"total_rev_pkts_inspected": ret.DtCgnv6ServiceGroupMemberStats.Stats.Total_rev_pkts_inspected,
			"total_rev_pkts_inspected_status_code_2xx":     ret.DtCgnv6ServiceGroupMemberStats.Stats.Total_rev_pkts_inspected_status_code_2xx,
			"total_rev_pkts_inspected_status_code_non_5xx": ret.DtCgnv6ServiceGroupMemberStats.Stats.Total_rev_pkts_inspected_status_code_non_5xx,
			"curr_req":           ret.DtCgnv6ServiceGroupMemberStats.Stats.Curr_req,
			"total_req":          ret.DtCgnv6ServiceGroupMemberStats.Stats.Total_req,
			"total_req_succ":     ret.DtCgnv6ServiceGroupMemberStats.Stats.Total_req_succ,
			"peak_conn":          ret.DtCgnv6ServiceGroupMemberStats.Stats.Peak_conn,
			"response_time":      ret.DtCgnv6ServiceGroupMemberStats.Stats.Response_time,
			"fastest_rsp_time":   ret.DtCgnv6ServiceGroupMemberStats.Stats.Fastest_rsp_time,
			"slowest_rsp_time":   ret.DtCgnv6ServiceGroupMemberStats.Stats.Slowest_rsp_time,
			"curr_ssl_conn":      ret.DtCgnv6ServiceGroupMemberStats.Stats.Curr_ssl_conn,
			"total_ssl_conn":     ret.DtCgnv6ServiceGroupMemberStats.Stats.Total_ssl_conn,
			"curr_conn_overflow": ret.DtCgnv6ServiceGroupMemberStats.Stats.Curr_conn_overflow,
			"state_flaps":        ret.DtCgnv6ServiceGroupMemberStats.Stats.State_flaps,
		},
	}
}

func getObjectCgnv6ServiceGroupMemberStatsStats(d []interface{}) edpt.Cgnv6ServiceGroupMemberStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6ServiceGroupMemberStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_conn = in["curr_conn"].(int)
		ret.Total_fwd_bytes = in["total_fwd_bytes"].(int)
		ret.Total_fwd_pkts = in["total_fwd_pkts"].(int)
		ret.Total_rev_bytes = in["total_rev_bytes"].(int)
		ret.Total_rev_pkts = in["total_rev_pkts"].(int)
		ret.Total_conn = in["total_conn"].(int)
		ret.Total_rev_pkts_inspected = in["total_rev_pkts_inspected"].(int)
		ret.Total_rev_pkts_inspected_status_code_2xx = in["total_rev_pkts_inspected_status_code_2xx"].(int)
		ret.Total_rev_pkts_inspected_status_code_non_5xx = in["total_rev_pkts_inspected_status_code_non_5xx"].(int)
		ret.Curr_req = in["curr_req"].(int)
		ret.Total_req = in["total_req"].(int)
		ret.Total_req_succ = in["total_req_succ"].(int)
		ret.Peak_conn = in["peak_conn"].(int)
		ret.Response_time = in["response_time"].(int)
		ret.Fastest_rsp_time = in["fastest_rsp_time"].(int)
		ret.Slowest_rsp_time = in["slowest_rsp_time"].(int)
		ret.Curr_ssl_conn = in["curr_ssl_conn"].(int)
		ret.Total_ssl_conn = in["total_ssl_conn"].(int)
		ret.Curr_conn_overflow = in["curr_conn_overflow"].(int)
		ret.State_flaps = in["state_flaps"].(int)
	}
	return ret
}

func dataToEndpointCgnv6ServiceGroupMemberStats(d *schema.ResourceData) edpt.Cgnv6ServiceGroupMemberStats {
	var ret edpt.Cgnv6ServiceGroupMemberStats

	ret.Name = d.Get("name").(string)

	ret.Port = d.Get("port").(int)

	ret.Stats = getObjectCgnv6ServiceGroupMemberStatsStats(d.Get("stats").([]interface{}))
	return ret
}
