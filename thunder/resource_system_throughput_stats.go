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
		// CreateContext: resourceSystemThroughputStatsCreate,
		// UpdateContext: resourceSystemThroughputStatsUpdate,
		ReadContext: resourceSystemThroughputStatsRead,
		// DeleteContext: resourceSystemThroughputStatsDelete,
		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"global_system_throughput_bits_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "Global System throughput in bits/sec",
						},
						"per_part_throughput_bits_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "Partition throughput in bits/sec",
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
		items := setObjectSystemThroughputStatsStats(res)
		d.SetId(obj.GetId())
		d.Set("stats", items)

		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemThroughputStatsStats(res edpt.SystemThroughputt) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"global_system_throughput_bits_per_sec": res.Throughput.Stats.GlobalSystemThroughputBitsPerSec,
			"per_part_throughput_bits_per_sec":      res.Throughput.Stats.PerPartThroughputBitsPerSec,
		},
	}
}

func getObjectSystemThroughputStatsStats(d []interface{}) edpt.SystemThroughputStatsStats {
	count := len(d)
	var ret edpt.SystemThroughputStatsStats
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.GlobalSystemThroughputBitsPerSec = in["global_system_throughput_bits_per_sec"].(int)
		ret.PerPartThroughputBitsPerSec = in["per_part_throughput_bits_per_sec"].(int)
	}
	return ret
}

func dataToEndpointSystemThroughputStats(d *schema.ResourceData) edpt.SystemThroughputStats {
	var ret edpt.SystemThroughputStats
	ret.Stats = getObjectSystemThroughputStatsStats(d.Get("stats").([]interface{}))
	return ret
}
