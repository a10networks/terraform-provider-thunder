package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwServerStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_server_stats`: Statistics for the object server\n\n__PLACEHOLDER__",
		ReadContext: resourceFwServerStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Server Name",
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_number": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"curr_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Current connections",
									},
									"curr_req": {
										Type: schema.TypeInt, Optional: true, Description: "Current requests",
									},
									"total_req": {
										Type: schema.TypeInt, Optional: true, Description: "Total requests",
									},
									"total_req_succ": {
										Type: schema.TypeInt, Optional: true, Description: "Total request success",
									},
									"total_fwd_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Forward bytes",
									},
									"total_fwd_pkts": {
										Type: schema.TypeInt, Optional: true, Description: "Forward packets",
									},
									"total_rev_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Reverse bytes",
									},
									"total_rev_pkts": {
										Type: schema.TypeInt, Optional: true, Description: "Reverse packets",
									},
									"total_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Total connections",
									},
									"last_total_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Last total connections",
									},
									"peak_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Peak connections",
									},
									"es_resp_200": {
										Type: schema.TypeInt, Optional: true, Description: "Response status 200",
									},
									"es_resp_300": {
										Type: schema.TypeInt, Optional: true, Description: "Response status 300",
									},
									"es_resp_400": {
										Type: schema.TypeInt, Optional: true, Description: "Response status 400",
									},
									"es_resp_500": {
										Type: schema.TypeInt, Optional: true, Description: "Response status 500",
									},
									"es_resp_other": {
										Type: schema.TypeInt, Optional: true, Description: "Response status other",
									},
									"es_req_count": {
										Type: schema.TypeInt, Optional: true, Description: "Total proxy request",
									},
									"es_resp_count": {
										Type: schema.TypeInt, Optional: true, Description: "Total proxy Response",
									},
									"es_resp_invalid_http": {
										Type: schema.TypeInt, Optional: true, Description: "Total non-http response",
									},
									"total_rev_pkts_inspected": {
										Type: schema.TypeInt, Optional: true, Description: "Total reverse packets inspected",
									},
									"total_rev_pkts_inspected_good_status_code": {
										Type: schema.TypeInt, Optional: true, Description: "Total reverse packets with good status code inspected",
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
								},
							},
						},
					},
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Current connections",
						},
						"total_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total connections",
						},
						"fwd_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Forward packets",
						},
						"rev_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Packets",
						},
						"peak_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Peak connections",
						},
					},
				},
			},
		},
	}
}

func resourceFwServerStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwServerStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwServerStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwServerStatsPortList := setSliceFwServerStatsPortList(res)
		d.Set("port_list", FwServerStatsPortList)
		FwServerStatsStats := setObjectFwServerStatsStats(res)
		d.Set("stats", FwServerStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceFwServerStatsPortList(d edpt.DataFwServerStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtFwServerStats.PortList {
		in := make(map[string]interface{})
		in["port_number"] = item.PortNumber
		in["protocol"] = item.Protocol
		in["stats"] = setObjectFwServerStatsPortListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectFwServerStatsPortListStats(d edpt.FwServerStatsPortListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["curr_conn"] = d.Curr_conn

	in["curr_req"] = d.Curr_req

	in["total_req"] = d.Total_req

	in["total_req_succ"] = d.Total_req_succ

	in["total_fwd_bytes"] = d.Total_fwd_bytes

	in["total_fwd_pkts"] = d.Total_fwd_pkts

	in["total_rev_bytes"] = d.Total_rev_bytes

	in["total_rev_pkts"] = d.Total_rev_pkts

	in["total_conn"] = d.Total_conn

	in["last_total_conn"] = d.Last_total_conn

	in["peak_conn"] = d.Peak_conn

	in["es_resp_200"] = d.Es_resp_200

	in["es_resp_300"] = d.Es_resp_300

	in["es_resp_400"] = d.Es_resp_400

	in["es_resp_500"] = d.Es_resp_500

	in["es_resp_other"] = d.Es_resp_other

	in["es_req_count"] = d.Es_req_count

	in["es_resp_count"] = d.Es_resp_count

	in["es_resp_invalid_http"] = d.Es_resp_invalid_http

	in["total_rev_pkts_inspected"] = d.Total_rev_pkts_inspected

	in["total_rev_pkts_inspected_good_status_code"] = d.Total_rev_pkts_inspected_good_status_code

	in["response_time"] = d.Response_time

	in["fastest_rsp_time"] = d.Fastest_rsp_time

	in["slowest_rsp_time"] = d.Slowest_rsp_time
	result = append(result, in)
	return result
}

func setObjectFwServerStatsStats(ret edpt.DataFwServerStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_conn":  ret.DtFwServerStats.Stats.CurrConn,
			"total_conn": ret.DtFwServerStats.Stats.TotalConn,
			"fwd_pkt":    ret.DtFwServerStats.Stats.FwdPkt,
			"rev_pkt":    ret.DtFwServerStats.Stats.RevPkt,
			"peak_conn":  ret.DtFwServerStats.Stats.PeakConn,
		},
	}
}

func getSliceFwServerStatsPortList(d []interface{}) []edpt.FwServerStatsPortList {

	count1 := len(d)
	ret := make([]edpt.FwServerStatsPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwServerStatsPortList
		oi.PortNumber = in["port_number"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Stats = getObjectFwServerStatsPortListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwServerStatsPortListStats(d []interface{}) edpt.FwServerStatsPortListStats {

	count1 := len(d)
	var ret edpt.FwServerStatsPortListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_conn = in["curr_conn"].(int)
		ret.Curr_req = in["curr_req"].(int)
		ret.Total_req = in["total_req"].(int)
		ret.Total_req_succ = in["total_req_succ"].(int)
		ret.Total_fwd_bytes = in["total_fwd_bytes"].(int)
		ret.Total_fwd_pkts = in["total_fwd_pkts"].(int)
		ret.Total_rev_bytes = in["total_rev_bytes"].(int)
		ret.Total_rev_pkts = in["total_rev_pkts"].(int)
		ret.Total_conn = in["total_conn"].(int)
		ret.Last_total_conn = in["last_total_conn"].(int)
		ret.Peak_conn = in["peak_conn"].(int)
		ret.Es_resp_200 = in["es_resp_200"].(int)
		ret.Es_resp_300 = in["es_resp_300"].(int)
		ret.Es_resp_400 = in["es_resp_400"].(int)
		ret.Es_resp_500 = in["es_resp_500"].(int)
		ret.Es_resp_other = in["es_resp_other"].(int)
		ret.Es_req_count = in["es_req_count"].(int)
		ret.Es_resp_count = in["es_resp_count"].(int)
		ret.Es_resp_invalid_http = in["es_resp_invalid_http"].(int)
		ret.Total_rev_pkts_inspected = in["total_rev_pkts_inspected"].(int)
		ret.Total_rev_pkts_inspected_good_status_code = in["total_rev_pkts_inspected_good_status_code"].(int)
		ret.Response_time = in["response_time"].(int)
		ret.Fastest_rsp_time = in["fastest_rsp_time"].(int)
		ret.Slowest_rsp_time = in["slowest_rsp_time"].(int)
	}
	return ret
}

func getObjectFwServerStatsStats(d []interface{}) edpt.FwServerStatsStats {

	count1 := len(d)
	var ret edpt.FwServerStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CurrConn = in["curr_conn"].(int)
		ret.TotalConn = in["total_conn"].(int)
		ret.FwdPkt = in["fwd_pkt"].(int)
		ret.RevPkt = in["rev_pkt"].(int)
		ret.PeakConn = in["peak_conn"].(int)
	}
	return ret
}

func dataToEndpointFwServerStats(d *schema.ResourceData) edpt.FwServerStats {
	var ret edpt.FwServerStats

	ret.Name = d.Get("name").(string)

	ret.PortList = getSliceFwServerStatsPortList(d.Get("port_list").([]interface{}))

	ret.Stats = getObjectFwServerStatsStats(d.Get("stats").([]interface{}))
	return ret
}
