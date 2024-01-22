package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatAlgEspStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_alg_esp_stats`: Statistics for the object esp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatAlgEspStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_created": {
							Type: schema.TypeInt, Optional: true, Description: "ESP Sessions Created",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6FixedNatAlgEspStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgEspStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgEspStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatAlgEspStatsStats := setObjectCgnv6FixedNatAlgEspStatsStats(res)
		d.Set("stats", Cgnv6FixedNatAlgEspStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatAlgEspStatsStats(ret edpt.DataCgnv6FixedNatAlgEspStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_created": ret.DtCgnv6FixedNatAlgEspStats.Stats.SessionCreated,
		},
	}
}

func getObjectCgnv6FixedNatAlgEspStatsStats(d []interface{}) edpt.Cgnv6FixedNatAlgEspStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatAlgEspStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionCreated = in["session_created"].(int)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatAlgEspStats(d *schema.ResourceData) edpt.Cgnv6FixedNatAlgEspStats {
	var ret edpt.Cgnv6FixedNatAlgEspStats

	ret.Stats = getObjectCgnv6FixedNatAlgEspStatsStats(d.Get("stats").([]interface{}))
	return ret
}
