package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPerfStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_perf_stats`: Statistics for the object perf\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbPerfStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_throughput_bits_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "Total Throughput in bits/sec",
						},
						"l4_conns_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "L4 Connections/sec",
						},
						"l7_conns_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "L7 Connections/sec",
						},
						"l7_trans_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "L7 Transactions/sec",
						},
						"ssl_conns_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "SSL Connections/sec",
						},
						"ip_nat_conns_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "IP NAT Connections/sec",
						},
						"total_new_conns_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "Total New Connections Established/sec",
						},
						"total_curr_conns": {
							Type: schema.TypeInt, Optional: true, Description: "Total Current Established Connections",
						},
						"l4_bandwidth": {
							Type: schema.TypeInt, Optional: true, Description: "L4 Bandwidth in bits/sec",
						},
						"l7_bandwidth": {
							Type: schema.TypeInt, Optional: true, Description: "L7 Bandwidth in bits/sec",
						},
						"serv_ssl_conns_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "Server SSL Connections/sec",
						},
						"fw_conns_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "FW Connections/sec",
						},
						"gifw_conns_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "GiFW Connections/sec",
						},
						"l7_proxy_conns_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "L7 Proxy Connections/sec",
						},
						"l7_proxy_trans_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "L7 Proxy Transactions/sec",
						},
					},
				},
			},
		},
	}
}

func resourceSlbPerfStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPerfStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPerfStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbPerfStatsStats := setObjectSlbPerfStatsStats(res)
		d.Set("stats", SlbPerfStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbPerfStatsStats(ret edpt.DataSlbPerfStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_throughput_bits_per_sec": ret.DtSlbPerfStats.Stats.TotalThroughputBitsPerSec,
			"l4_conns_per_sec":              ret.DtSlbPerfStats.Stats.L4ConnsPerSec,
			"l7_conns_per_sec":              ret.DtSlbPerfStats.Stats.L7ConnsPerSec,
			"l7_trans_per_sec":              ret.DtSlbPerfStats.Stats.L7TransPerSec,
			"ssl_conns_per_sec":             ret.DtSlbPerfStats.Stats.SslConnsPerSec,
			"ip_nat_conns_per_sec":          ret.DtSlbPerfStats.Stats.IpNatConnsPerSec,
			"total_new_conns_per_sec":       ret.DtSlbPerfStats.Stats.TotalNewConnsPerSec,
			"total_curr_conns":              ret.DtSlbPerfStats.Stats.TotalCurrConns,
			"l4_bandwidth":                  ret.DtSlbPerfStats.Stats.L4Bandwidth,
			"l7_bandwidth":                  ret.DtSlbPerfStats.Stats.L7Bandwidth,
			"serv_ssl_conns_per_sec":        ret.DtSlbPerfStats.Stats.ServSslConnsPerSec,
			"fw_conns_per_sec":              ret.DtSlbPerfStats.Stats.FwConnsPerSec,
			"gifw_conns_per_sec":            ret.DtSlbPerfStats.Stats.GifwConnsPerSec,
			"l7_proxy_conns_per_sec":        ret.DtSlbPerfStats.Stats.L7ProxyConnsPerSec,
			"l7_proxy_trans_per_sec":        ret.DtSlbPerfStats.Stats.L7ProxyTransPerSec,
		},
	}
}

func getObjectSlbPerfStatsStats(d []interface{}) edpt.SlbPerfStatsStats {

	count1 := len(d)
	var ret edpt.SlbPerfStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalThroughputBitsPerSec = in["total_throughput_bits_per_sec"].(int)
		ret.L4ConnsPerSec = in["l4_conns_per_sec"].(int)
		ret.L7ConnsPerSec = in["l7_conns_per_sec"].(int)
		ret.L7TransPerSec = in["l7_trans_per_sec"].(int)
		ret.SslConnsPerSec = in["ssl_conns_per_sec"].(int)
		ret.IpNatConnsPerSec = in["ip_nat_conns_per_sec"].(int)
		ret.TotalNewConnsPerSec = in["total_new_conns_per_sec"].(int)
		ret.TotalCurrConns = in["total_curr_conns"].(int)
		ret.L4Bandwidth = in["l4_bandwidth"].(int)
		ret.L7Bandwidth = in["l7_bandwidth"].(int)
		ret.ServSslConnsPerSec = in["serv_ssl_conns_per_sec"].(int)
		ret.FwConnsPerSec = in["fw_conns_per_sec"].(int)
		ret.GifwConnsPerSec = in["gifw_conns_per_sec"].(int)
		ret.L7ProxyConnsPerSec = in["l7_proxy_conns_per_sec"].(int)
		ret.L7ProxyTransPerSec = in["l7_proxy_trans_per_sec"].(int)
	}
	return ret
}

func dataToEndpointSlbPerfStats(d *schema.ResourceData) edpt.SlbPerfStats {
	var ret edpt.SlbPerfStats

	ret.Stats = getObjectSlbPerfStatsStats(d.Get("stats").([]interface{}))
	return ret
}
