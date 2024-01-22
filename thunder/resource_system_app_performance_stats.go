package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAppPerformanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_app_performance_stats`: Statistics for the object app-performance\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemAppPerformanceStatsRead,

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
					},
				},
			},
		},
	}
}

func resourceSystemAppPerformanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAppPerformanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAppPerformanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemAppPerformanceStatsStats := setObjectSystemAppPerformanceStatsStats(res)
		d.Set("stats", SystemAppPerformanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemAppPerformanceStatsStats(ret edpt.DataSystemAppPerformanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_throughput_bits_per_sec": ret.DtSystemAppPerformanceStats.Stats.TotalThroughputBitsPerSec,
			"l4_conns_per_sec":              ret.DtSystemAppPerformanceStats.Stats.L4ConnsPerSec,
			"l7_conns_per_sec":              ret.DtSystemAppPerformanceStats.Stats.L7ConnsPerSec,
			"l7_trans_per_sec":              ret.DtSystemAppPerformanceStats.Stats.L7TransPerSec,
			"ssl_conns_per_sec":             ret.DtSystemAppPerformanceStats.Stats.SslConnsPerSec,
			"ip_nat_conns_per_sec":          ret.DtSystemAppPerformanceStats.Stats.IpNatConnsPerSec,
			"total_new_conns_per_sec":       ret.DtSystemAppPerformanceStats.Stats.TotalNewConnsPerSec,
			"total_curr_conns":              ret.DtSystemAppPerformanceStats.Stats.TotalCurrConns,
			"l4_bandwidth":                  ret.DtSystemAppPerformanceStats.Stats.L4Bandwidth,
			"l7_bandwidth":                  ret.DtSystemAppPerformanceStats.Stats.L7Bandwidth,
			"serv_ssl_conns_per_sec":        ret.DtSystemAppPerformanceStats.Stats.ServSslConnsPerSec,
			"fw_conns_per_sec":              ret.DtSystemAppPerformanceStats.Stats.FwConnsPerSec,
			"gifw_conns_per_sec":            ret.DtSystemAppPerformanceStats.Stats.GifwConnsPerSec,
		},
	}
}

func getObjectSystemAppPerformanceStatsStats(d []interface{}) edpt.SystemAppPerformanceStatsStats {

	count1 := len(d)
	var ret edpt.SystemAppPerformanceStatsStats
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
	}
	return ret
}

func dataToEndpointSystemAppPerformanceStats(d *schema.ResourceData) edpt.SystemAppPerformanceStats {
	var ret edpt.SystemAppPerformanceStats

	ret.Stats = getObjectSystemAppPerformanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
