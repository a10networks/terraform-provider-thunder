package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAppStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_app_stats`: Statistics for the object app\n\n__PLACEHOLDER__",
		ReadContext: resourceFwAppStatsRead,

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

func resourceFwAppStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAppStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAppStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwAppStatsStats := setObjectFwAppStatsStats(res)
		d.Set("stats", FwAppStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwAppStatsStats(ret edpt.DataFwAppStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dummy": ret.DtFwAppStats.Stats.Dummy,
		},
	}
}

func getObjectFwAppStatsStats(d []interface{}) edpt.FwAppStatsStats {

	count1 := len(d)
	var ret edpt.FwAppStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dummy = in["dummy"].(int)
	}
	return ret
}

func dataToEndpointFwAppStats(d *schema.ResourceData) edpt.FwAppStats {
	var ret edpt.FwAppStats

	ret.Stats = getObjectFwAppStatsStats(d.Get("stats").([]interface{}))
	return ret
}
