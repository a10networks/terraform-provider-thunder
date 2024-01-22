package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServiceGroupMemberStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_service_group_member_stats`: Statistics for the object member\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServiceGroupMemberStatsRead,

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
							Type: schema.TypeInt, Optional: true, Description: "Current established connections",
						},
						"total_fwd_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes processed in forward direction",
						},
						"total_fwd_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Packets processed in forward direction",
						},
						"total_rev_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes processed in reverse direction",
						},
						"total_rev_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Packets processed in reverse direction",
						},
						"total_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total established connections",
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
							Type: schema.TypeInt, Optional: true, Description: "Total requests successful",
						},
						"peak_conn": {
							Type: schema.TypeInt, Optional: true, Description: "",
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
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationServiceGroupMemberStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServiceGroupMemberStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServiceGroupMemberStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServiceGroupMemberStatsStats := setObjectAamAuthenticationServiceGroupMemberStatsStats(res)
		d.Set("stats", AamAuthenticationServiceGroupMemberStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationServiceGroupMemberStatsStats(ret edpt.DataAamAuthenticationServiceGroupMemberStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_conn":                ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Curr_conn,
			"total_fwd_bytes":          ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Total_fwd_bytes,
			"total_fwd_pkts":           ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Total_fwd_pkts,
			"total_rev_bytes":          ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Total_rev_bytes,
			"total_rev_pkts":           ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Total_rev_pkts,
			"total_conn":               ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Total_conn,
			"total_rev_pkts_inspected": ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Total_rev_pkts_inspected,
			"total_rev_pkts_inspected_status_code_2xx":     ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Total_rev_pkts_inspected_status_code_2xx,
			"total_rev_pkts_inspected_status_code_non_5xx": ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Total_rev_pkts_inspected_status_code_non_5xx,
			"curr_req":           ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Curr_req,
			"total_req":          ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Total_req,
			"total_req_succ":     ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Total_req_succ,
			"peak_conn":          ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Peak_conn,
			"response_time":      ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Response_time,
			"fastest_rsp_time":   ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Fastest_rsp_time,
			"slowest_rsp_time":   ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Slowest_rsp_time,
			"curr_ssl_conn":      ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Curr_ssl_conn,
			"total_ssl_conn":     ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Total_ssl_conn,
			"curr_conn_overflow": ret.DtAamAuthenticationServiceGroupMemberStats.Stats.Curr_conn_overflow,
		},
	}
}

func getObjectAamAuthenticationServiceGroupMemberStatsStats(d []interface{}) edpt.AamAuthenticationServiceGroupMemberStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServiceGroupMemberStatsStats
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
	}
	return ret
}

func dataToEndpointAamAuthenticationServiceGroupMemberStats(d *schema.ResourceData) edpt.AamAuthenticationServiceGroupMemberStats {
	var ret edpt.AamAuthenticationServiceGroupMemberStats

	ret.Name = d.Get("name").(string)

	ret.Port = d.Get("port").(int)

	ret.Stats = getObjectAamAuthenticationServiceGroupMemberStatsStats(d.Get("stats").([]interface{}))
	return ret
}
