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
						"ppsl_drop_egr": {
							Type: schema.TypeInt, Optional: true, Description: "Packet-Per-Sec Limit Drop at egress",
						},
						"ppsl_drop_ing": {
							Type: schema.TypeInt, Optional: true, Description: "Packet-Per-Sec Limit Drop at ingress",
						},
						"ppsl_ignore_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Packet-Per-Sec Limit ignored packets count",
						},
						"licexpire_drop": {
							Type: schema.TypeInt, Optional: true, Description: "License Expire Drop",
						},
						"bwl_drop": {
							Type: schema.TypeInt, Optional: true, Description: "BW Limit Drop",
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
			"ppsl_drop_egr":        ret.DtSystemBandwidthStats.Stats.Ppsl_drop_egr,
			"ppsl_drop_ing":        ret.DtSystemBandwidthStats.Stats.Ppsl_drop_ing,
			"ppsl_ignore_limit":    ret.DtSystemBandwidthStats.Stats.Ppsl_ignore_limit,
			"licexpire_drop":       ret.DtSystemBandwidthStats.Stats.Licexpire_drop,
			"bwl_drop":             ret.DtSystemBandwidthStats.Stats.Bwl_drop,
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
		ret.Ppsl_drop_egr = in["ppsl_drop_egr"].(int)
		ret.Ppsl_drop_ing = in["ppsl_drop_ing"].(int)
		ret.Ppsl_ignore_limit = in["ppsl_ignore_limit"].(int)
		ret.Licexpire_drop = in["licexpire_drop"].(int)
		ret.Bwl_drop = in["bwl_drop"].(int)
	}
	return ret
}

func dataToEndpointSystemBandwidthStats(d *schema.ResourceData) edpt.SystemBandwidthStats {
	var ret edpt.SystemBandwidthStats

	ret.Stats = getObjectSystemBandwidthStatsStats(d.Get("stats").([]interface{}))
	return ret
}
