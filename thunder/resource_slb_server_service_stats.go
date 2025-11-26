package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServerServiceStats48() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_server_service_stats`: Statistics for the object service\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbServerServiceStats48Read,

		Schema: map[string]*schema.Schema{
			"label": {
				Type: schema.TypeString, Required: true, Description: "Service Label",
			},
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
			"server_name": {
				Type: schema.TypeString, Required: true, Description: "Server_name",
			},
		},
	}
}

func resourceSlbServerServiceStats48Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerServiceStats48Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerServiceStats48(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbServerServiceStats48Stats := setObjectSlbServerServiceStats48Stats(res)
		d.Set("stats", SlbServerServiceStats48Stats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbServerServiceStats48Stats(ret edpt.DataSlbServerServiceStats48) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_conn":                ret.DtSlbServerServiceStats48.Stats.Curr_conn,
			"curr_req":                 ret.DtSlbServerServiceStats48.Stats.Curr_req,
			"total_req":                ret.DtSlbServerServiceStats48.Stats.Total_req,
			"total_req_succ":           ret.DtSlbServerServiceStats48.Stats.Total_req_succ,
			"total_fwd_bytes":          ret.DtSlbServerServiceStats48.Stats.Total_fwd_bytes,
			"total_fwd_pkts":           ret.DtSlbServerServiceStats48.Stats.Total_fwd_pkts,
			"total_rev_bytes":          ret.DtSlbServerServiceStats48.Stats.Total_rev_bytes,
			"total_rev_pkts":           ret.DtSlbServerServiceStats48.Stats.Total_rev_pkts,
			"total_conn":               ret.DtSlbServerServiceStats48.Stats.Total_conn,
			"last_total_conn":          ret.DtSlbServerServiceStats48.Stats.Last_total_conn,
			"peak_conn":                ret.DtSlbServerServiceStats48.Stats.Peak_conn,
			"es_resp_200":              ret.DtSlbServerServiceStats48.Stats.Es_resp_200,
			"es_resp_300":              ret.DtSlbServerServiceStats48.Stats.Es_resp_300,
			"es_resp_400":              ret.DtSlbServerServiceStats48.Stats.Es_resp_400,
			"es_resp_500":              ret.DtSlbServerServiceStats48.Stats.Es_resp_500,
			"es_resp_other":            ret.DtSlbServerServiceStats48.Stats.Es_resp_other,
			"es_req_count":             ret.DtSlbServerServiceStats48.Stats.Es_req_count,
			"es_resp_count":            ret.DtSlbServerServiceStats48.Stats.Es_resp_count,
			"es_resp_invalid_http":     ret.DtSlbServerServiceStats48.Stats.Es_resp_invalid_http,
			"total_rev_pkts_inspected": ret.DtSlbServerServiceStats48.Stats.Total_rev_pkts_inspected,
			"total_rev_pkts_inspected_good_status_code": ret.DtSlbServerServiceStats48.Stats.Total_rev_pkts_inspected_good_status_code,
			"response_time":    ret.DtSlbServerServiceStats48.Stats.Response_time,
			"fastest_rsp_time": ret.DtSlbServerServiceStats48.Stats.Fastest_rsp_time,
			"slowest_rsp_time": ret.DtSlbServerServiceStats48.Stats.Slowest_rsp_time,
			"curr_ssl_conn":    ret.DtSlbServerServiceStats48.Stats.Curr_ssl_conn,
			"total_ssl_conn":   ret.DtSlbServerServiceStats48.Stats.Total_ssl_conn,
			"resp_count":       ret.DtSlbServerServiceStats48.Stats.RespCount,
			"resp_1xx":         ret.DtSlbServerServiceStats48.Stats.Resp1xx,
			"resp_2xx":         ret.DtSlbServerServiceStats48.Stats.Resp2xx,
			"resp_3xx":         ret.DtSlbServerServiceStats48.Stats.Resp3xx,
			"resp_4xx":         ret.DtSlbServerServiceStats48.Stats.Resp4xx,
			"resp_5xx":         ret.DtSlbServerServiceStats48.Stats.Resp5xx,
			"resp_other":       ret.DtSlbServerServiceStats48.Stats.RespOther,
			"resp_latency":     ret.DtSlbServerServiceStats48.Stats.RespLatency,
			"curr_pconn":       ret.DtSlbServerServiceStats48.Stats.Curr_pconn,
		},
	}
}

func getObjectSlbServerServiceStats48Stats(d []interface{}) edpt.SlbServerServiceStats48Stats {

	count1 := len(d)
	var ret edpt.SlbServerServiceStats48Stats
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

func dataToEndpointSlbServerServiceStats48(d *schema.ResourceData) edpt.SlbServerServiceStats48 {
	var ret edpt.SlbServerServiceStats48

	ret.Label = d.Get("label").(string)

	ret.PortNumber = d.Get("port_number").(int)

	ret.Protocol = d.Get("protocol").(string)

	ret.Stats = getObjectSlbServerServiceStats48Stats(d.Get("stats").([]interface{}))

	ret.Server_name = d.Get("server_name").(string)
	return ret
}
