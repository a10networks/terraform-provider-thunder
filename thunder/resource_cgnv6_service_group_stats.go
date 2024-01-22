package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6ServiceGroupStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_service_group_stats`: Statistics for the object service-group\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6ServiceGroupStatsRead,

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
				Type: schema.TypeString, Required: true, Description: "CGNV6 Service Name",
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

func resourceCgnv6ServiceGroupStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ServiceGroupStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ServiceGroupStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6ServiceGroupStatsMemberList := setSliceCgnv6ServiceGroupStatsMemberList(res)
		d.Set("member_list", Cgnv6ServiceGroupStatsMemberList)
		Cgnv6ServiceGroupStatsStats := setObjectCgnv6ServiceGroupStatsStats(res)
		d.Set("stats", Cgnv6ServiceGroupStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceCgnv6ServiceGroupStatsMemberList(d edpt.DataCgnv6ServiceGroupStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtCgnv6ServiceGroupStats.MemberList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["port"] = item.Port
		in["stats"] = setObjectCgnv6ServiceGroupStatsMemberListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectCgnv6ServiceGroupStatsMemberListStats(d edpt.Cgnv6ServiceGroupStatsMemberListStats) []map[string]interface{} {
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

func setObjectCgnv6ServiceGroupStatsStats(ret edpt.DataCgnv6ServiceGroupStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"server_selection_fail_drop":  ret.DtCgnv6ServiceGroupStats.Stats.Server_selection_fail_drop,
			"server_selection_fail_reset": ret.DtCgnv6ServiceGroupStats.Stats.Server_selection_fail_reset,
			"service_peak_conn":           ret.DtCgnv6ServiceGroupStats.Stats.Service_peak_conn,
		},
	}
}

func getSliceCgnv6ServiceGroupStatsMemberList(d []interface{}) []edpt.Cgnv6ServiceGroupStatsMemberList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ServiceGroupStatsMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ServiceGroupStatsMemberList
		oi.Name = in["name"].(string)
		oi.Port = in["port"].(int)
		oi.Stats = getObjectCgnv6ServiceGroupStatsMemberListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6ServiceGroupStatsMemberListStats(d []interface{}) edpt.Cgnv6ServiceGroupStatsMemberListStats {

	count1 := len(d)
	var ret edpt.Cgnv6ServiceGroupStatsMemberListStats
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

func getObjectCgnv6ServiceGroupStatsStats(d []interface{}) edpt.Cgnv6ServiceGroupStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6ServiceGroupStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Server_selection_fail_drop = in["server_selection_fail_drop"].(int)
		ret.Server_selection_fail_reset = in["server_selection_fail_reset"].(int)
		ret.Service_peak_conn = in["service_peak_conn"].(int)
	}
	return ret
}

func dataToEndpointCgnv6ServiceGroupStats(d *schema.ResourceData) edpt.Cgnv6ServiceGroupStats {
	var ret edpt.Cgnv6ServiceGroupStats

	ret.MemberList = getSliceCgnv6ServiceGroupStatsMemberList(d.Get("member_list").([]interface{}))

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectCgnv6ServiceGroupStatsStats(d.Get("stats").([]interface{}))
	return ret
}
