package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemViewMemoryViewStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_view_memory_view_stats`: Statistics for the object memory-view\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemViewMemoryViewStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"usage_percentage": {
							Type: schema.TypeInt, Optional: true, Description: "Usage percentage",
						},
					},
				},
			},
		},
	}
}

func resourceSystemViewMemoryViewStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemViewMemoryViewStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemViewMemoryViewStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemViewMemoryViewStatsStats := setObjectSystemViewMemoryViewStatsStats(res)
		d.Set("stats", SystemViewMemoryViewStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemViewMemoryViewStatsStats(ret edpt.DataSystemViewMemoryViewStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"usage_percentage": ret.DtSystemViewMemoryViewStats.Stats.UsagePercentage,
		},
	}
}

func getObjectSystemViewMemoryViewStatsStats(d []interface{}) edpt.SystemViewMemoryViewStatsStats {

	count1 := len(d)
	var ret edpt.SystemViewMemoryViewStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UsagePercentage = in["usage_percentage"].(int)
	}
	return ret
}

func dataToEndpointSystemViewMemoryViewStats(d *schema.ResourceData) edpt.SystemViewMemoryViewStats {
	var ret edpt.SystemViewMemoryViewStats

	ret.Stats = getObjectSystemViewMemoryViewStatsStats(d.Get("stats").([]interface{}))
	return ret
}
