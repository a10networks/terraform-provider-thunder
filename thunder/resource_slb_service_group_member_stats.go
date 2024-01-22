package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServiceGroupMemberStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_service_group_member_stats`: Statistics for the object member\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbServiceGroupMemberStatsRead,

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

func resourceSlbServiceGroupMemberStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServiceGroupMemberStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServiceGroupMemberStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbServiceGroupMemberStatsStats := setObjectSlbServiceGroupMemberStatsStats(res)
		d.Set("stats", SlbServiceGroupMemberStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbServiceGroupMemberStatsStats(ret edpt.DataSlbServiceGroupMemberStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_conn":                ret.DtSlbServiceGroupMemberStats.Stats.Curr_conn,
			"total_fwd_bytes":          ret.DtSlbServiceGroupMemberStats.Stats.Total_fwd_bytes,
			"total_fwd_pkts":           ret.DtSlbServiceGroupMemberStats.Stats.Total_fwd_pkts,
			"total_rev_bytes":          ret.DtSlbServiceGroupMemberStats.Stats.Total_rev_bytes,
			"total_rev_pkts":           ret.DtSlbServiceGroupMemberStats.Stats.Total_rev_pkts,
			"total_conn":               ret.DtSlbServiceGroupMemberStats.Stats.Total_conn,
			"total_rev_pkts_inspected": ret.DtSlbServiceGroupMemberStats.Stats.Total_rev_pkts_inspected,
			"total_rev_pkts_inspected_status_code_2xx":     ret.DtSlbServiceGroupMemberStats.Stats.Total_rev_pkts_inspected_status_code_2xx,
			"total_rev_pkts_inspected_status_code_non_5xx": ret.DtSlbServiceGroupMemberStats.Stats.Total_rev_pkts_inspected_status_code_non_5xx,
			"curr_req":           ret.DtSlbServiceGroupMemberStats.Stats.Curr_req,
			"total_req":          ret.DtSlbServiceGroupMemberStats.Stats.Total_req,
			"total_req_succ":     ret.DtSlbServiceGroupMemberStats.Stats.Total_req_succ,
			"peak_conn":          ret.DtSlbServiceGroupMemberStats.Stats.Peak_conn,
			"response_time":      ret.DtSlbServiceGroupMemberStats.Stats.Response_time,
			"fastest_rsp_time":   ret.DtSlbServiceGroupMemberStats.Stats.Fastest_rsp_time,
			"slowest_rsp_time":   ret.DtSlbServiceGroupMemberStats.Stats.Slowest_rsp_time,
			"curr_ssl_conn":      ret.DtSlbServiceGroupMemberStats.Stats.Curr_ssl_conn,
			"total_ssl_conn":     ret.DtSlbServiceGroupMemberStats.Stats.Total_ssl_conn,
			"curr_conn_overflow": ret.DtSlbServiceGroupMemberStats.Stats.Curr_conn_overflow,
			"state_flaps":        ret.DtSlbServiceGroupMemberStats.Stats.State_flaps,
		},
	}
}

func getObjectSlbServiceGroupMemberStatsStats(d []interface{}) edpt.SlbServiceGroupMemberStatsStats {

	count1 := len(d)
	var ret edpt.SlbServiceGroupMemberStatsStats
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

func dataToEndpointSlbServiceGroupMemberStats(d *schema.ResourceData) edpt.SlbServiceGroupMemberStats {
	var ret edpt.SlbServiceGroupMemberStats

	ret.Name = d.Get("name").(string)

	ret.Port = d.Get("port").(int)

	ret.Stats = getObjectSlbServiceGroupMemberStatsStats(d.Get("stats").([]interface{}))
	return ret
}
