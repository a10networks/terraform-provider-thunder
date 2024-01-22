package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFastHttpProxyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_fast_http_proxy_stats`: Statistics for the object fast-http-proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbFastHttpProxyStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Curr Proxy Conns",
						},
						"total_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Total Proxy Conns",
						},
						"req": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP requests",
						},
						"req_succ": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP requests(succ)",
						},
						"noproxy": {
							Type: schema.TypeInt, Optional: true, Description: "No proxy error",
						},
						"client_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Client RST",
						},
						"server_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Server RST",
						},
						"notuple": {
							Type: schema.TypeInt, Optional: true, Description: "No tuple error",
						},
						"parsereq_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parse req fail",
						},
						"svrsel_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Server selection fail",
						},
						"fwdreq_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Fwd req fail",
						},
						"fwdreqdata_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Fwd req data fail",
						},
						"req_retran": {
							Type: schema.TypeInt, Optional: true, Description: "Packets retrans",
						},
						"req_ofo": {
							Type: schema.TypeInt, Optional: true, Description: "Packets ofo",
						},
						"server_resel": {
							Type: schema.TypeInt, Optional: true, Description: "Server reselection",
						},
						"svr_prem_close": {
							Type: schema.TypeInt, Optional: true, Description: "Server premature close",
						},
						"new_svrconn": {
							Type: schema.TypeInt, Optional: true, Description: "Server conn made",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Source NAT failure",
						},
						"tcpoutrst": {
							Type: schema.TypeInt, Optional: true, Description: "Out RSTs",
						},
						"full_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Full proxy tot",
						},
						"full_proxy_post": {
							Type: schema.TypeInt, Optional: true, Description: "Full proxy POST",
						},
						"full_proxy_pipeline": {
							Type: schema.TypeInt, Optional: true, Description: "Full proxy pipeline",
						},
						"full_proxy_fpga_err": {
							Type: schema.TypeInt, Optional: true, Description: "Full proxy fpga err",
						},
						"req_over_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Request over limit",
						},
						"req_rate_over_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Request rate over limit",
						},
						"close_on_ddos": {
							Type: schema.TypeInt, Optional: true, Description: "Close on DDoS",
						},
						"full_proxy_put": {
							Type: schema.TypeInt, Optional: true, Description: "Full proxy PUT",
						},
					},
				},
			},
		},
	}
}

func resourceSlbFastHttpProxyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFastHttpProxyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFastHttpProxyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbFastHttpProxyStatsStats := setObjectSlbFastHttpProxyStatsStats(res)
		d.Set("stats", SlbFastHttpProxyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbFastHttpProxyStatsStats(ret edpt.DataSlbFastHttpProxyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_proxy":          ret.DtSlbFastHttpProxyStats.Stats.Curr_proxy,
			"total_proxy":         ret.DtSlbFastHttpProxyStats.Stats.Total_proxy,
			"req":                 ret.DtSlbFastHttpProxyStats.Stats.Req,
			"req_succ":            ret.DtSlbFastHttpProxyStats.Stats.Req_succ,
			"noproxy":             ret.DtSlbFastHttpProxyStats.Stats.Noproxy,
			"client_rst":          ret.DtSlbFastHttpProxyStats.Stats.Client_rst,
			"server_rst":          ret.DtSlbFastHttpProxyStats.Stats.Server_rst,
			"notuple":             ret.DtSlbFastHttpProxyStats.Stats.Notuple,
			"parsereq_fail":       ret.DtSlbFastHttpProxyStats.Stats.Parsereq_fail,
			"svrsel_fail":         ret.DtSlbFastHttpProxyStats.Stats.Svrsel_fail,
			"fwdreq_fail":         ret.DtSlbFastHttpProxyStats.Stats.Fwdreq_fail,
			"fwdreqdata_fail":     ret.DtSlbFastHttpProxyStats.Stats.Fwdreqdata_fail,
			"req_retran":          ret.DtSlbFastHttpProxyStats.Stats.Req_retran,
			"req_ofo":             ret.DtSlbFastHttpProxyStats.Stats.Req_ofo,
			"server_resel":        ret.DtSlbFastHttpProxyStats.Stats.Server_resel,
			"svr_prem_close":      ret.DtSlbFastHttpProxyStats.Stats.Svr_prem_close,
			"new_svrconn":         ret.DtSlbFastHttpProxyStats.Stats.New_svrconn,
			"snat_fail":           ret.DtSlbFastHttpProxyStats.Stats.Snat_fail,
			"tcpoutrst":           ret.DtSlbFastHttpProxyStats.Stats.Tcpoutrst,
			"full_proxy":          ret.DtSlbFastHttpProxyStats.Stats.Full_proxy,
			"full_proxy_post":     ret.DtSlbFastHttpProxyStats.Stats.Full_proxy_post,
			"full_proxy_pipeline": ret.DtSlbFastHttpProxyStats.Stats.Full_proxy_pipeline,
			"full_proxy_fpga_err": ret.DtSlbFastHttpProxyStats.Stats.Full_proxy_fpga_err,
			"req_over_limit":      ret.DtSlbFastHttpProxyStats.Stats.Req_over_limit,
			"req_rate_over_limit": ret.DtSlbFastHttpProxyStats.Stats.Req_rate_over_limit,
			"close_on_ddos":       ret.DtSlbFastHttpProxyStats.Stats.Close_on_ddos,
			"full_proxy_put":      ret.DtSlbFastHttpProxyStats.Stats.Full_proxy_put,
		},
	}
}

func getObjectSlbFastHttpProxyStatsStats(d []interface{}) edpt.SlbFastHttpProxyStatsStats {

	count1 := len(d)
	var ret edpt.SlbFastHttpProxyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_proxy = in["curr_proxy"].(int)
		ret.Total_proxy = in["total_proxy"].(int)
		ret.Req = in["req"].(int)
		ret.Req_succ = in["req_succ"].(int)
		ret.Noproxy = in["noproxy"].(int)
		ret.Client_rst = in["client_rst"].(int)
		ret.Server_rst = in["server_rst"].(int)
		ret.Notuple = in["notuple"].(int)
		ret.Parsereq_fail = in["parsereq_fail"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.Fwdreq_fail = in["fwdreq_fail"].(int)
		ret.Fwdreqdata_fail = in["fwdreqdata_fail"].(int)
		ret.Req_retran = in["req_retran"].(int)
		ret.Req_ofo = in["req_ofo"].(int)
		ret.Server_resel = in["server_resel"].(int)
		ret.Svr_prem_close = in["svr_prem_close"].(int)
		ret.New_svrconn = in["new_svrconn"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Tcpoutrst = in["tcpoutrst"].(int)
		ret.Full_proxy = in["full_proxy"].(int)
		ret.Full_proxy_post = in["full_proxy_post"].(int)
		ret.Full_proxy_pipeline = in["full_proxy_pipeline"].(int)
		ret.Full_proxy_fpga_err = in["full_proxy_fpga_err"].(int)
		ret.Req_over_limit = in["req_over_limit"].(int)
		ret.Req_rate_over_limit = in["req_rate_over_limit"].(int)
		ret.Close_on_ddos = in["close_on_ddos"].(int)
		ret.Full_proxy_put = in["full_proxy_put"].(int)
	}
	return ret
}

func dataToEndpointSlbFastHttpProxyStats(d *schema.ResourceData) edpt.SlbFastHttpProxyStats {
	var ret edpt.SlbFastHttpProxyStats

	ret.Stats = getObjectSlbFastHttpProxyStatsStats(d.Get("stats").([]interface{}))
	return ret
}
