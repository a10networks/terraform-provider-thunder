package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwServerPortStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_server_port_stats`: Statistics for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceFwServerPortStatsRead,

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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceFwServerPortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwServerPortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwServerPortStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwServerPortStatsStats := setObjectFwServerPortStatsStats(res)
		d.Set("stats", FwServerPortStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwServerPortStatsStats(ret edpt.DataFwServerPortStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_conn":                ret.DtFwServerPortStats.Stats.Curr_conn,
			"curr_req":                 ret.DtFwServerPortStats.Stats.Curr_req,
			"total_req":                ret.DtFwServerPortStats.Stats.Total_req,
			"total_req_succ":           ret.DtFwServerPortStats.Stats.Total_req_succ,
			"total_fwd_bytes":          ret.DtFwServerPortStats.Stats.Total_fwd_bytes,
			"total_fwd_pkts":           ret.DtFwServerPortStats.Stats.Total_fwd_pkts,
			"total_rev_bytes":          ret.DtFwServerPortStats.Stats.Total_rev_bytes,
			"total_rev_pkts":           ret.DtFwServerPortStats.Stats.Total_rev_pkts,
			"total_conn":               ret.DtFwServerPortStats.Stats.Total_conn,
			"last_total_conn":          ret.DtFwServerPortStats.Stats.Last_total_conn,
			"peak_conn":                ret.DtFwServerPortStats.Stats.Peak_conn,
			"es_resp_200":              ret.DtFwServerPortStats.Stats.Es_resp_200,
			"es_resp_300":              ret.DtFwServerPortStats.Stats.Es_resp_300,
			"es_resp_400":              ret.DtFwServerPortStats.Stats.Es_resp_400,
			"es_resp_500":              ret.DtFwServerPortStats.Stats.Es_resp_500,
			"es_resp_other":            ret.DtFwServerPortStats.Stats.Es_resp_other,
			"es_req_count":             ret.DtFwServerPortStats.Stats.Es_req_count,
			"es_resp_count":            ret.DtFwServerPortStats.Stats.Es_resp_count,
			"es_resp_invalid_http":     ret.DtFwServerPortStats.Stats.Es_resp_invalid_http,
			"total_rev_pkts_inspected": ret.DtFwServerPortStats.Stats.Total_rev_pkts_inspected,
			"total_rev_pkts_inspected_good_status_code": ret.DtFwServerPortStats.Stats.Total_rev_pkts_inspected_good_status_code,
			"response_time":    ret.DtFwServerPortStats.Stats.Response_time,
			"fastest_rsp_time": ret.DtFwServerPortStats.Stats.Fastest_rsp_time,
			"slowest_rsp_time": ret.DtFwServerPortStats.Stats.Slowest_rsp_time,
		},
	}
}

func getObjectFwServerPortStatsStats(d []interface{}) edpt.FwServerPortStatsStats {

	count1 := len(d)
	var ret edpt.FwServerPortStatsStats
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

func dataToEndpointFwServerPortStats(d *schema.ResourceData) edpt.FwServerPortStats {
	var ret edpt.FwServerPortStats

	ret.PortNumber = d.Get("port_number").(int)

	ret.Protocol = d.Get("protocol").(string)

	ret.Stats = getObjectFwServerPortStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
