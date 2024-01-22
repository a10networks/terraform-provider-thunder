package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAutomaticUpdateInfoStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_automatic_update_info_stats`: Statistics for the object info\n\n__PLACEHOLDER__",
		ReadContext: resourceAutomaticUpdateInfoStatsRead,

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

func resourceAutomaticUpdateInfoStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateInfoStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdateInfoStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AutomaticUpdateInfoStatsStats := setObjectAutomaticUpdateInfoStatsStats(res)
		d.Set("stats", AutomaticUpdateInfoStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAutomaticUpdateInfoStatsStats(ret edpt.DataAutomaticUpdateInfoStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dummy": ret.DtAutomaticUpdateInfoStats.Stats.Dummy,
		},
	}
}

func getObjectAutomaticUpdateInfoStatsStats(d []interface{}) edpt.AutomaticUpdateInfoStatsStats {

	count1 := len(d)
	var ret edpt.AutomaticUpdateInfoStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dummy = in["dummy"].(int)
	}
	return ret
}

func dataToEndpointAutomaticUpdateInfoStats(d *schema.ResourceData) edpt.AutomaticUpdateInfoStats {
	var ret edpt.AutomaticUpdateInfoStats

	ret.Stats = getObjectAutomaticUpdateInfoStatsStats(d.Get("stats").([]interface{}))
	return ret
}
