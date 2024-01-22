package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemBandwidthStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_bandwidth_stats`: Statistics for the object bandwidth\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemBandwidthStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"input_bytes_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "In Bytes per second",
						},
						"output_bytes_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "Out Bytes per second",
						},
					},
				},
			},
		},
	}
}

func resourceSystemBandwidthStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBandwidthStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBandwidthStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemBandwidthStatsStats := setObjectSystemBandwidthStatsStats(res)
		d.Set("stats", SystemBandwidthStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemBandwidthStatsStats(ret edpt.DataSystemBandwidthStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"input_bytes_per_sec":  ret.DtSystemBandwidthStats.Stats.InputBytesPerSec,
			"output_bytes_per_sec": ret.DtSystemBandwidthStats.Stats.OutputBytesPerSec,
		},
	}
}

func getObjectSystemBandwidthStatsStats(d []interface{}) edpt.SystemBandwidthStatsStats {

	count1 := len(d)
	var ret edpt.SystemBandwidthStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.InputBytesPerSec = in["input_bytes_per_sec"].(int)
		ret.OutputBytesPerSec = in["output_bytes_per_sec"].(int)
	}
	return ret
}

func dataToEndpointSystemBandwidthStats(d *schema.ResourceData) edpt.SystemBandwidthStats {
	var ret edpt.SystemBandwidthStats

	ret.Stats = getObjectSystemBandwidthStatsStats(d.Get("stats").([]interface{}))
	return ret
}
