package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMemoryStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_memory_stats`: Statistics for the object memory\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemMemoryStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"usage_percentage": {
							Type: schema.TypeInt, Optional: true, Description: "Memory Usage percentage",
						},
					},
				},
			},
		},
	}
}

func resourceSystemMemoryStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMemoryStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMemoryStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemMemoryStatsStats := setObjectSystemMemoryStatsStats(res)
		d.Set("stats", SystemMemoryStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemMemoryStatsStats(ret edpt.DataSystemMemoryStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"usage_percentage": ret.DtSystemMemoryStats.Stats.UsagePercentage,
		},
	}
}

func getObjectSystemMemoryStatsStats(d []interface{}) edpt.SystemMemoryStatsStats {

	count1 := len(d)
	var ret edpt.SystemMemoryStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UsagePercentage = in["usage_percentage"].(int)
	}
	return ret
}

func dataToEndpointSystemMemoryStats(d *schema.ResourceData) edpt.SystemMemoryStats {
	var ret edpt.SystemMemoryStats

	ret.Stats = getObjectSystemMemoryStatsStats(d.Get("stats").([]interface{}))
	return ret
}
