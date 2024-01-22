package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpACommonStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vrrp_a_common_stats`: Statistics for the object common\n\n__PLACEHOLDER__",
		ReadContext: resourceVrrpACommonStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrrp_common_dummy": {
							Type: schema.TypeInt, Optional: true, Description: "dummy counter",
						},
					},
				},
			},
		},
	}
}

func resourceVrrpACommonStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpACommonStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpACommonStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VrrpACommonStatsStats := setObjectVrrpACommonStatsStats(res)
		d.Set("stats", VrrpACommonStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVrrpACommonStatsStats(ret edpt.DataVrrpACommonStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vrrp_common_dummy": ret.DtVrrpACommonStats.Stats.Vrrp_common_dummy,
		},
	}
}

func getObjectVrrpACommonStatsStats(d []interface{}) edpt.VrrpACommonStatsStats {

	count1 := len(d)
	var ret edpt.VrrpACommonStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Vrrp_common_dummy = in["vrrp_common_dummy"].(int)
	}
	return ret
}

func dataToEndpointVrrpACommonStats(d *schema.ResourceData) edpt.VrrpACommonStats {
	var ret edpt.VrrpACommonStats

	ret.Stats = getObjectVrrpACommonStatsStats(d.Get("stats").([]interface{}))
	return ret
}
