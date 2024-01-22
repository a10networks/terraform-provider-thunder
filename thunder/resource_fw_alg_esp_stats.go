package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgEspStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_alg_esp_stats`: Statistics for the object esp\n\n__PLACEHOLDER__",
		ReadContext: resourceFwAlgEspStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_created": {
							Type: schema.TypeInt, Optional: true, Description: "ESP Sessions Created",
						},
						"helper_created": {
							Type: schema.TypeInt, Optional: true, Description: "ESP Helper Sessions Created",
						},
						"helper_freed": {
							Type: schema.TypeInt, Optional: true, Description: "ESP Helper Sessions Freed",
						},
					},
				},
			},
		},
	}
}

func resourceFwAlgEspStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgEspStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgEspStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwAlgEspStatsStats := setObjectFwAlgEspStatsStats(res)
		d.Set("stats", FwAlgEspStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwAlgEspStatsStats(ret edpt.DataFwAlgEspStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_created": ret.DtFwAlgEspStats.Stats.SessionCreated,
			"helper_created":  ret.DtFwAlgEspStats.Stats.HelperCreated,
			"helper_freed":    ret.DtFwAlgEspStats.Stats.HelperFreed,
		},
	}
}

func getObjectFwAlgEspStatsStats(d []interface{}) edpt.FwAlgEspStatsStats {

	count1 := len(d)
	var ret edpt.FwAlgEspStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionCreated = in["session_created"].(int)
		ret.HelperCreated = in["helper_created"].(int)
		ret.HelperFreed = in["helper_freed"].(int)
	}
	return ret
}

func dataToEndpointFwAlgEspStats(d *schema.ResourceData) edpt.FwAlgEspStats {
	var ret edpt.FwAlgEspStats

	ret.Stats = getObjectFwAlgEspStatsStats(d.Get("stats").([]interface{}))
	return ret
}
