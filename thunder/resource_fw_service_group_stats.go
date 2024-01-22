package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwServiceGroupStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_service_group_stats`: Statistics for the object service-group\n\n__PLACEHOLDER__",
		ReadContext: resourceFwServiceGroupStatsRead,

		Schema: map[string]*schema.Schema{
			"member_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "FW Service Name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_selection_fail_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Service selection fail drop",
						},
						"server_selection_fail_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Service selection fail reset",
						},
						"service_peak_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Service peak connection",
						},
					},
				},
			},
		},
	}
}

func resourceFwServiceGroupStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwServiceGroupStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwServiceGroupStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwServiceGroupStatsMemberList := setSliceFwServiceGroupStatsMemberList(res)
		d.Set("member_list", FwServiceGroupStatsMemberList)
		FwServiceGroupStatsStats := setObjectFwServiceGroupStatsStats(res)
		d.Set("stats", FwServiceGroupStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceFwServiceGroupStatsMemberList(d edpt.DataFwServiceGroupStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtFwServiceGroupStats.MemberList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["port"] = item.Port
		in["stats"] = setObjectFwServiceGroupStatsMemberListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectFwServiceGroupStatsMemberListStats(d edpt.FwServiceGroupStatsMemberListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["curr_conn"] = d.Curr_conn

	in["total_fwd_bytes"] = d.Total_fwd_bytes

	in["total_fwd_pkts"] = d.Total_fwd_pkts

	in["total_rev_bytes"] = d.Total_rev_bytes

	in["total_rev_pkts"] = d.Total_rev_pkts

	in["total_conn"] = d.Total_conn

	in["total_rev_pkts_inspected"] = d.Total_rev_pkts_inspected

	in["total_rev_pkts_inspected_status_code_2xx"] = d.Total_rev_pkts_inspected_status_code_2xx

	in["total_rev_pkts_inspected_status_code_non_5xx"] = d.Total_rev_pkts_inspected_status_code_non_5xx

	in["curr_req"] = d.Curr_req

	in["total_req"] = d.Total_req

	in["total_req_succ"] = d.Total_req_succ

	in["peak_conn"] = d.Peak_conn

	in["response_time"] = d.Response_time

	in["fastest_rsp_time"] = d.Fastest_rsp_time

	in["slowest_rsp_time"] = d.Slowest_rsp_time

	in["curr_ssl_conn"] = d.Curr_ssl_conn

	in["total_ssl_conn"] = d.Total_ssl_conn

	in["curr_conn_overflow"] = d.Curr_conn_overflow

	in["state_flaps"] = d.State_flaps
	result = append(result, in)
	return result
}

func setObjectFwServiceGroupStatsStats(ret edpt.DataFwServiceGroupStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"server_selection_fail_drop":  ret.DtFwServiceGroupStats.Stats.Server_selection_fail_drop,
			"server_selection_fail_reset": ret.DtFwServiceGroupStats.Stats.Server_selection_fail_reset,
			"service_peak_conn":           ret.DtFwServiceGroupStats.Stats.Service_peak_conn,
		},
	}
}

func getSliceFwServiceGroupStatsMemberList(d []interface{}) []edpt.FwServiceGroupStatsMemberList {

	count1 := len(d)
	ret := make([]edpt.FwServiceGroupStatsMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwServiceGroupStatsMemberList
		oi.Name = in["name"].(string)
		oi.Port = in["port"].(int)
		oi.Stats = getObjectFwServiceGroupStatsMemberListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwServiceGroupStatsMemberListStats(d []interface{}) edpt.FwServiceGroupStatsMemberListStats {

	count1 := len(d)
	var ret edpt.FwServiceGroupStatsMemberListStats
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

func getObjectFwServiceGroupStatsStats(d []interface{}) edpt.FwServiceGroupStatsStats {

	count1 := len(d)
	var ret edpt.FwServiceGroupStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Server_selection_fail_drop = in["server_selection_fail_drop"].(int)
		ret.Server_selection_fail_reset = in["server_selection_fail_reset"].(int)
		ret.Service_peak_conn = in["service_peak_conn"].(int)
	}
	return ret
}

func dataToEndpointFwServiceGroupStats(d *schema.ResourceData) edpt.FwServiceGroupStats {
	var ret edpt.FwServiceGroupStats

	ret.MemberList = getSliceFwServiceGroupStatsMemberList(d.Get("member_list").([]interface{}))

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectFwServiceGroupStatsStats(d.Get("stats").([]interface{}))
	return ret
}
