package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServerStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_server_stats`: Statistics for the object server\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbServerStatsRead,

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
										Type: schema.TypeInt, Optional: true, Description: "Total Requests",
									},
									"total_req_succ": {
										Type: schema.TypeInt, Optional: true, Description: "Total requests succ",
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
										Type: schema.TypeInt, Optional: true, Description: "Total proxy requests",
									},
									"es_resp_count": {
										Type: schema.TypeInt, Optional: true, Description: "Total proxy response",
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
									"curr_ssl_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Current SSL connections",
									},
									"total_ssl_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Total SSL connections",
									},
									"resp_count": {
										Type: schema.TypeInt, Optional: true, Description: "Total Response Count",
									},
									"resp_1xx": {
										Type: schema.TypeInt, Optional: true, Description: "Response status 1xx",
									},
									"resp_2xx": {
										Type: schema.TypeInt, Optional: true, Description: "Response status 2xx",
									},
									"resp_3xx": {
										Type: schema.TypeInt, Optional: true, Description: "Response status 3xx",
									},
									"resp_4xx": {
										Type: schema.TypeInt, Optional: true, Description: "Response status 4xx",
									},
									"resp_5xx": {
										Type: schema.TypeInt, Optional: true, Description: "Response status 5xx",
									},
									"resp_other": {
										Type: schema.TypeInt, Optional: true, Description: "Response status Other",
									},
									"resp_latency": {
										Type: schema.TypeInt, Optional: true, Description: "Time to First Response Byte",
									},
									"curr_pconn": {
										Type: schema.TypeInt, Optional: true, Description: "Current persistent connections",
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
							Type: schema.TypeInt, Optional: true, Description: "Current established connections",
						},
						"total_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total established connections",
						},
						"fwd_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Packets Processed",
						},
						"rev_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Packets Processed",
						},
						"peak_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Peak number of established connections",
						},
						"total_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total Requests processed",
						},
						"total_req_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Total Requests succeeded",
						},
						"curr_ssl_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Current SSL connections established",
						},
						"total_ssl_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total SSL connections established",
						},
						"total_fwd_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes processed in forward direction",
						},
						"total_rev_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes processed in reverse direction",
						},
						"total_fwd_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Packets processed in forward direction",
						},
						"total_rev_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Packets processed in reverse direction",
						},
						"ip_only_lb_fwd_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "IP-Only-LB Bytes processed in forward direction",
						},
						"ip_only_lb_rev_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "IP-Only-LB Bytes processed in reverse direction",
						},
						"ip_only_lb_fwd_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "IP-Only-LB Packets processed in forward direction",
						},
						"ip_only_lb_rev_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "IP-Only-LB Packets processed in reverse direction",
						},
					},
				},
			},
		},
	}
}

func resourceSlbServerStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbServerStatsPortList := setSliceSlbServerStatsPortList(res)
		d.Set("port_list", SlbServerStatsPortList)
		SlbServerStatsStats := setObjectSlbServerStatsStats(res)
		d.Set("stats", SlbServerStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceSlbServerStatsPortList(d edpt.DataSlbServerStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtSlbServerStats.PortList {
		in := make(map[string]interface{})
		in["port_number"] = item.PortNumber
		in["protocol"] = item.Protocol
		in["stats"] = setObjectSlbServerStatsPortListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectSlbServerStatsPortListStats(d edpt.SlbServerStatsPortListStats) []map[string]interface{} {
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

	in["curr_ssl_conn"] = d.Curr_ssl_conn

	in["total_ssl_conn"] = d.Total_ssl_conn

	in["resp_count"] = d.RespCount

	in["resp_1xx"] = d.Resp1xx

	in["resp_2xx"] = d.Resp2xx

	in["resp_3xx"] = d.Resp3xx

	in["resp_4xx"] = d.Resp4xx

	in["resp_5xx"] = d.Resp5xx

	in["resp_other"] = d.RespOther

	in["resp_latency"] = d.RespLatency

	in["curr_pconn"] = d.Curr_pconn
	result = append(result, in)
	return result
}

func setObjectSlbServerStatsStats(ret edpt.DataSlbServerStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_conn":            ret.DtSlbServerStats.Stats.CurrConn,
			"total_conn":           ret.DtSlbServerStats.Stats.TotalConn,
			"fwd_pkt":              ret.DtSlbServerStats.Stats.FwdPkt,
			"rev_pkt":              ret.DtSlbServerStats.Stats.RevPkt,
			"peak_conn":            ret.DtSlbServerStats.Stats.PeakConn,
			"total_req":            ret.DtSlbServerStats.Stats.Total_req,
			"total_req_succ":       ret.DtSlbServerStats.Stats.Total_req_succ,
			"curr_ssl_conn":        ret.DtSlbServerStats.Stats.Curr_ssl_conn,
			"total_ssl_conn":       ret.DtSlbServerStats.Stats.Total_ssl_conn,
			"total_fwd_bytes":      ret.DtSlbServerStats.Stats.Total_fwd_bytes,
			"total_rev_bytes":      ret.DtSlbServerStats.Stats.Total_rev_bytes,
			"total_fwd_pkts":       ret.DtSlbServerStats.Stats.Total_fwd_pkts,
			"total_rev_pkts":       ret.DtSlbServerStats.Stats.Total_rev_pkts,
			"ip_only_lb_fwd_bytes": ret.DtSlbServerStats.Stats.Ip_only_lb_fwd_bytes,
			"ip_only_lb_rev_bytes": ret.DtSlbServerStats.Stats.Ip_only_lb_rev_bytes,
			"ip_only_lb_fwd_pkts":  ret.DtSlbServerStats.Stats.Ip_only_lb_fwd_pkts,
			"ip_only_lb_rev_pkts":  ret.DtSlbServerStats.Stats.Ip_only_lb_rev_pkts,
		},
	}
}

func getSliceSlbServerStatsPortList(d []interface{}) []edpt.SlbServerStatsPortList {

	count1 := len(d)
	ret := make([]edpt.SlbServerStatsPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerStatsPortList
		oi.PortNumber = in["port_number"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Stats = getObjectSlbServerStatsPortListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbServerStatsPortListStats(d []interface{}) edpt.SlbServerStatsPortListStats {

	count1 := len(d)
	var ret edpt.SlbServerStatsPortListStats
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
		ret.Curr_ssl_conn = in["curr_ssl_conn"].(int)
		ret.Total_ssl_conn = in["total_ssl_conn"].(int)
		ret.RespCount = in["resp_count"].(int)
		ret.Resp1xx = in["resp_1xx"].(int)
		ret.Resp2xx = in["resp_2xx"].(int)
		ret.Resp3xx = in["resp_3xx"].(int)
		ret.Resp4xx = in["resp_4xx"].(int)
		ret.Resp5xx = in["resp_5xx"].(int)
		ret.RespOther = in["resp_other"].(int)
		ret.RespLatency = in["resp_latency"].(int)
		ret.Curr_pconn = in["curr_pconn"].(int)
	}
	return ret
}

func getObjectSlbServerStatsStats(d []interface{}) edpt.SlbServerStatsStats {

	count1 := len(d)
	var ret edpt.SlbServerStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CurrConn = in["curr_conn"].(int)
		ret.TotalConn = in["total_conn"].(int)
		ret.FwdPkt = in["fwd_pkt"].(int)
		ret.RevPkt = in["rev_pkt"].(int)
		ret.PeakConn = in["peak_conn"].(int)
		ret.Total_req = in["total_req"].(int)
		ret.Total_req_succ = in["total_req_succ"].(int)
		ret.Curr_ssl_conn = in["curr_ssl_conn"].(int)
		ret.Total_ssl_conn = in["total_ssl_conn"].(int)
		ret.Total_fwd_bytes = in["total_fwd_bytes"].(int)
		ret.Total_rev_bytes = in["total_rev_bytes"].(int)
		ret.Total_fwd_pkts = in["total_fwd_pkts"].(int)
		ret.Total_rev_pkts = in["total_rev_pkts"].(int)
		ret.Ip_only_lb_fwd_bytes = in["ip_only_lb_fwd_bytes"].(int)
		ret.Ip_only_lb_rev_bytes = in["ip_only_lb_rev_bytes"].(int)
		ret.Ip_only_lb_fwd_pkts = in["ip_only_lb_fwd_pkts"].(int)
		ret.Ip_only_lb_rev_pkts = in["ip_only_lb_rev_pkts"].(int)
	}
	return ret
}

func dataToEndpointSlbServerStats(d *schema.ResourceData) edpt.SlbServerStats {
	var ret edpt.SlbServerStats

	ret.Name = d.Get("name").(string)

	ret.PortList = getSliceSlbServerStatsPortList(d.Get("port_list").([]interface{}))

	ret.Stats = getObjectSlbServerStatsStats(d.Get("stats").([]interface{}))
	return ret
}
