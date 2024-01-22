package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServiceGroupStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_service_group_stats`: Statistics for the object service-group\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbServiceGroupStatsRead,

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
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "SLB Service Name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_selection_fail_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Drops due to Service selection failure",
						},
						"server_selection_fail_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Resets sent out for Service selection failure",
						},
						"service_peak_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Peak connection count for the Service Group",
						},
						"service_healthy_host": {
							Type: schema.TypeInt, Optional: true, Description: "Service Group healthy host count",
						},
						"service_unhealthy_host": {
							Type: schema.TypeInt, Optional: true, Description: "Service Group unhealthy host count",
						},
						"service_req_count": {
							Type: schema.TypeInt, Optional: true, Description: "Service Group request count",
						},
						"service_resp_count": {
							Type: schema.TypeInt, Optional: true, Description: "Service Group response count",
						},
						"service_resp_2xx": {
							Type: schema.TypeInt, Optional: true, Description: "Service Group response 2xx count",
						},
						"service_resp_3xx": {
							Type: schema.TypeInt, Optional: true, Description: "Service Group response 3xx count",
						},
						"service_resp_4xx": {
							Type: schema.TypeInt, Optional: true, Description: "Service Group response 4xx count",
						},
						"service_resp_5xx": {
							Type: schema.TypeInt, Optional: true, Description: "Service Group response 5xx count",
						},
						"service_curr_conn_overflow": {
							Type: schema.TypeInt, Optional: true, Description: "Current connection counter overflow count",
						},
					},
				},
			},
		},
	}
}

func resourceSlbServiceGroupStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServiceGroupStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServiceGroupStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbServiceGroupStatsMemberList := setSliceSlbServiceGroupStatsMemberList(res)
		d.Set("member_list", SlbServiceGroupStatsMemberList)
		SlbServiceGroupStatsStats := setObjectSlbServiceGroupStatsStats(res)
		d.Set("stats", SlbServiceGroupStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceSlbServiceGroupStatsMemberList(d edpt.DataSlbServiceGroupStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtSlbServiceGroupStats.MemberList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["port"] = item.Port
		in["stats"] = setObjectSlbServiceGroupStatsMemberListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectSlbServiceGroupStatsMemberListStats(d edpt.SlbServiceGroupStatsMemberListStats) []map[string]interface{} {
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

func setObjectSlbServiceGroupStatsStats(ret edpt.DataSlbServiceGroupStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"server_selection_fail_drop":  ret.DtSlbServiceGroupStats.Stats.Server_selection_fail_drop,
			"server_selection_fail_reset": ret.DtSlbServiceGroupStats.Stats.Server_selection_fail_reset,
			"service_peak_conn":           ret.DtSlbServiceGroupStats.Stats.Service_peak_conn,
			"service_healthy_host":        ret.DtSlbServiceGroupStats.Stats.Service_healthy_host,
			"service_unhealthy_host":      ret.DtSlbServiceGroupStats.Stats.Service_unhealthy_host,
			"service_req_count":           ret.DtSlbServiceGroupStats.Stats.Service_req_count,
			"service_resp_count":          ret.DtSlbServiceGroupStats.Stats.Service_resp_count,
			"service_resp_2xx":            ret.DtSlbServiceGroupStats.Stats.Service_resp_2xx,
			"service_resp_3xx":            ret.DtSlbServiceGroupStats.Stats.Service_resp_3xx,
			"service_resp_4xx":            ret.DtSlbServiceGroupStats.Stats.Service_resp_4xx,
			"service_resp_5xx":            ret.DtSlbServiceGroupStats.Stats.Service_resp_5xx,
			"service_curr_conn_overflow":  ret.DtSlbServiceGroupStats.Stats.Service_curr_conn_overflow,
		},
	}
}

func getSliceSlbServiceGroupStatsMemberList(d []interface{}) []edpt.SlbServiceGroupStatsMemberList {

	count1 := len(d)
	ret := make([]edpt.SlbServiceGroupStatsMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServiceGroupStatsMemberList
		oi.Name = in["name"].(string)
		oi.Port = in["port"].(int)
		oi.Stats = getObjectSlbServiceGroupStatsMemberListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbServiceGroupStatsMemberListStats(d []interface{}) edpt.SlbServiceGroupStatsMemberListStats {

	count1 := len(d)
	var ret edpt.SlbServiceGroupStatsMemberListStats
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

func getObjectSlbServiceGroupStatsStats(d []interface{}) edpt.SlbServiceGroupStatsStats {

	count1 := len(d)
	var ret edpt.SlbServiceGroupStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Server_selection_fail_drop = in["server_selection_fail_drop"].(int)
		ret.Server_selection_fail_reset = in["server_selection_fail_reset"].(int)
		ret.Service_peak_conn = in["service_peak_conn"].(int)
		ret.Service_healthy_host = in["service_healthy_host"].(int)
		ret.Service_unhealthy_host = in["service_unhealthy_host"].(int)
		ret.Service_req_count = in["service_req_count"].(int)
		ret.Service_resp_count = in["service_resp_count"].(int)
		ret.Service_resp_2xx = in["service_resp_2xx"].(int)
		ret.Service_resp_3xx = in["service_resp_3xx"].(int)
		ret.Service_resp_4xx = in["service_resp_4xx"].(int)
		ret.Service_resp_5xx = in["service_resp_5xx"].(int)
		ret.Service_curr_conn_overflow = in["service_curr_conn_overflow"].(int)
	}
	return ret
}

func dataToEndpointSlbServiceGroupStats(d *schema.ResourceData) edpt.SlbServiceGroupStats {
	var ret edpt.SlbServiceGroupStats

	ret.MemberList = getSliceSlbServiceGroupStatsMemberList(d.Get("member_list").([]interface{}))

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSlbServiceGroupStatsStats(d.Get("stats").([]interface{}))
	return ret
}
