package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAutomaticUpdateChecknowStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_automatic_update_checknow_stats`: Statistics for the object checknow\n\n__PLACEHOLDER__",
		ReadContext: resourceAutomaticUpdateChecknowStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dummy": {
							Type: schema.TypeInt, Optional: true, Description: "Entry for a10countergen",
						},
					},
				},
			},
		},
	}
}

func resourceAutomaticUpdateChecknowStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateChecknowStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateChecknowStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AutomaticUpdateChecknowStatsStats := setObjectAutomaticUpdateChecknowStatsStats(res)
		d.Set("stats", AutomaticUpdateChecknowStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAutomaticUpdateChecknowStatsStats(ret edpt.DataAutomaticUpdateChecknowStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dummy": ret.DtAutomaticUpdateChecknowStats.Stats.Dummy,
		},
	}
}

func getObjectAutomaticUpdateChecknowStatsStats(d []interface{}) edpt.AutomaticUpdateChecknowStatsStats {

	count1 := len(d)
	var ret edpt.AutomaticUpdateChecknowStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dummy = in["dummy"].(int)
	}
	return ret
}

func dataToEndpointAutomaticUpdateChecknowStats(d *schema.ResourceData) edpt.AutomaticUpdateChecknowStats {
	var ret edpt.AutomaticUpdateChecknowStats

	ret.Stats = getObjectAutomaticUpdateChecknowStatsStats(d.Get("stats").([]interface{}))
	return ret
}
