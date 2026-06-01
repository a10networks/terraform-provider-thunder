package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemThroughputStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_throughput_stats`: Statistics for the object throughput\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemThroughputStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"global_system_throughput_bits_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "Global System egress throughput in bits/sec",
						},
						"global_system_ingress_throughput_bits_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "Global System ingress throughput in bits/sec",
						},
						"per_part_throughput_bits_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "Partition throughput in bits/sec",
						},
						"global_client_ssl_count": {
							Type: schema.TypeInt, Optional: true, Description: "global ssl count",
						},
						"global_server_ssl_count": {
							Type: schema.TypeInt, Optional: true, Description: "global server ssl count",
						},
						"global_client_ssl_connections_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "global ssl conneciton per sec",
						},
						"global_server_ssl_connections_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "global server ssl conneciton",
						},
					},
				},
			},
		},
	}
}

func resourceSystemThroughputStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemThroughputStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemThroughputStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemThroughputStatsStats := setObjectSystemThroughputStatsStats(res)
		d.Set("stats", SystemThroughputStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemThroughputStatsStats(ret edpt.DataSystemThroughputStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"global_system_throughput_bits_per_sec":         ret.DtSystemThroughputStats.Stats.GlobalSystemThroughputBitsPerSec,
			"global_system_ingress_throughput_bits_per_sec": ret.DtSystemThroughputStats.Stats.GlobalSystemIngressThroughputBitsPerSec,
			"per_part_throughput_bits_per_sec":              ret.DtSystemThroughputStats.Stats.PerPartThroughputBitsPerSec,
			"global_client_ssl_count":                       ret.DtSystemThroughputStats.Stats.GlobalClientSslCount,
			"global_server_ssl_count":                       ret.DtSystemThroughputStats.Stats.GlobalServerSslCount,
			"global_client_ssl_connections_per_sec":         ret.DtSystemThroughputStats.Stats.GlobalClientSslConnectionsPerSec,
			"global_server_ssl_connections_per_sec":         ret.DtSystemThroughputStats.Stats.GlobalServerSslConnectionsPerSec,
		},
	}
}

func getObjectSystemThroughputStatsStats(d []interface{}) edpt.SystemThroughputStatsStats {

	count1 := len(d)
	var ret edpt.SystemThroughputStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GlobalSystemThroughputBitsPerSec = in["global_system_throughput_bits_per_sec"].(int)
		ret.GlobalSystemIngressThroughputBitsPerSec = in["global_system_ingress_throughput_bits_per_sec"].(int)
		ret.PerPartThroughputBitsPerSec = in["per_part_throughput_bits_per_sec"].(int)
		ret.GlobalClientSslCount = in["global_client_ssl_count"].(int)
		ret.GlobalServerSslCount = in["global_server_ssl_count"].(int)
		ret.GlobalClientSslConnectionsPerSec = in["global_client_ssl_connections_per_sec"].(int)
		ret.GlobalServerSslConnectionsPerSec = in["global_server_ssl_connections_per_sec"].(int)
	}
	return ret
}

func dataToEndpointSystemThroughputStats(d *schema.ResourceData) edpt.SystemThroughputStats {
	var ret edpt.SystemThroughputStats

	ret.Stats = getObjectSystemThroughputStatsStats(d.Get("stats").([]interface{}))
	return ret
}
