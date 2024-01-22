package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPop3ProxyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_pop3_proxy_stats`: Statistics for the object pop3-proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbPop3ProxyStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num": {
							Type: schema.TypeInt, Optional: true, Description: "Num",
						},
						"curr": {
							Type: schema.TypeInt, Optional: true, Description: "Current proxy conns",
						},
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "Total proxy conns",
						},
						"svrsel_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Server selection failure",
						},
						"no_route": {
							Type: schema.TypeInt, Optional: true, Description: "no route failure",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Description: "source nat failure",
						},
						"line_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "line too long",
						},
						"line_mem_freed": {
							Type: schema.TypeInt, Optional: true, Description: "request line freed",
						},
						"invalid_start_line": {
							Type: schema.TypeInt, Optional: true, Description: "invalid start line",
						},
						"stls": {
							Type: schema.TypeInt, Optional: true, Description: "stls cmd",
						},
						"request_dont_care": {
							Type: schema.TypeInt, Optional: true, Description: "other cmd",
						},
						"unsupported_command": {
							Type: schema.TypeInt, Optional: true, Description: "Unsupported cmd",
						},
						"bad_sequence": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Sequence",
						},
						"rsv_persist_conn_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Serv Sel Persist fail",
						},
						"smp_v6_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Serv Sel SMPv6 fail",
						},
						"smp_v4_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Serv Sel SMPv4 fail",
						},
						"insert_tuple_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Serv Sel insert tuple fail",
						},
						"cl_est_err": {
							Type: schema.TypeInt, Optional: true, Description: "Client EST state erro",
						},
						"ser_connecting_err": {
							Type: schema.TypeInt, Optional: true, Description: "Serv CTNG state error",
						},
						"server_response_err": {
							Type: schema.TypeInt, Optional: true, Description: "Serv RESP state error",
						},
						"cl_request_err": {
							Type: schema.TypeInt, Optional: true, Description: "Client RQ state error",
						},
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "Total POP3 Request",
						},
						"control_to_ssl": {
							Type: schema.TypeInt, Optional: true, Description: "Control chn ssl",
						},
					},
				},
			},
		},
	}
}

func resourceSlbPop3ProxyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPop3ProxyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPop3ProxyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbPop3ProxyStatsStats := setObjectSlbPop3ProxyStatsStats(res)
		d.Set("stats", SlbPop3ProxyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbPop3ProxyStatsStats(ret edpt.DataSlbPop3ProxyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"num":                   ret.DtSlbPop3ProxyStats.Stats.Num,
			"curr":                  ret.DtSlbPop3ProxyStats.Stats.Curr,
			"total":                 ret.DtSlbPop3ProxyStats.Stats.Total,
			"svrsel_fail":           ret.DtSlbPop3ProxyStats.Stats.Svrsel_fail,
			"no_route":              ret.DtSlbPop3ProxyStats.Stats.No_route,
			"snat_fail":             ret.DtSlbPop3ProxyStats.Stats.Snat_fail,
			"line_too_long":         ret.DtSlbPop3ProxyStats.Stats.Line_too_long,
			"line_mem_freed":        ret.DtSlbPop3ProxyStats.Stats.Line_mem_freed,
			"invalid_start_line":    ret.DtSlbPop3ProxyStats.Stats.Invalid_start_line,
			"stls":                  ret.DtSlbPop3ProxyStats.Stats.Stls,
			"request_dont_care":     ret.DtSlbPop3ProxyStats.Stats.Request_dont_care,
			"unsupported_command":   ret.DtSlbPop3ProxyStats.Stats.Unsupported_command,
			"bad_sequence":          ret.DtSlbPop3ProxyStats.Stats.Bad_sequence,
			"rsv_persist_conn_fail": ret.DtSlbPop3ProxyStats.Stats.Rsv_persist_conn_fail,
			"smp_v6_fail":           ret.DtSlbPop3ProxyStats.Stats.Smp_v6_fail,
			"smp_v4_fail":           ret.DtSlbPop3ProxyStats.Stats.Smp_v4_fail,
			"insert_tuple_fail":     ret.DtSlbPop3ProxyStats.Stats.Insert_tuple_fail,
			"cl_est_err":            ret.DtSlbPop3ProxyStats.Stats.Cl_est_err,
			"ser_connecting_err":    ret.DtSlbPop3ProxyStats.Stats.Ser_connecting_err,
			"server_response_err":   ret.DtSlbPop3ProxyStats.Stats.Server_response_err,
			"cl_request_err":        ret.DtSlbPop3ProxyStats.Stats.Cl_request_err,
			"request":               ret.DtSlbPop3ProxyStats.Stats.Request,
			"control_to_ssl":        ret.DtSlbPop3ProxyStats.Stats.Control_to_ssl,
		},
	}
}

func getObjectSlbPop3ProxyStatsStats(d []interface{}) edpt.SlbPop3ProxyStatsStats {

	count1 := len(d)
	var ret edpt.SlbPop3ProxyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Num = in["num"].(int)
		ret.Curr = in["curr"].(int)
		ret.Total = in["total"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Line_mem_freed = in["line_mem_freed"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Stls = in["stls"].(int)
		ret.Request_dont_care = in["request_dont_care"].(int)
		ret.Unsupported_command = in["unsupported_command"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		ret.Cl_request_err = in["cl_request_err"].(int)
		ret.Request = in["request"].(int)
		ret.Control_to_ssl = in["control_to_ssl"].(int)
	}
	return ret
}

func dataToEndpointSlbPop3ProxyStats(d *schema.ResourceData) edpt.SlbPop3ProxyStats {
	var ret edpt.SlbPop3ProxyStats

	ret.Stats = getObjectSlbPop3ProxyStatsStats(d.Get("stats").([]interface{}))
	return ret
}
