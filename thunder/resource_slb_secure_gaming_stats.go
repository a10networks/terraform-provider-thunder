package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSecureGamingStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_secure_gaming_stats`: Statistics for the object secure-gaming\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSecureGamingStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secure_gaming_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Secure Gaming passed",
						},
						"secure_gaming_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Secure Gaming dropped",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSecureGamingStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSecureGamingStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSecureGamingStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSecureGamingStatsStats := setObjectSlbSecureGamingStatsStats(res)
		d.Set("stats", SlbSecureGamingStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSecureGamingStatsStats(ret edpt.DataSlbSecureGamingStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"secure_gaming_pass": ret.DtSlbSecureGamingStats.Stats.Secure_gaming_pass,
			"secure_gaming_drop": ret.DtSlbSecureGamingStats.Stats.Secure_gaming_drop,
		},
	}
}

func getObjectSlbSecureGamingStatsStats(d []interface{}) edpt.SlbSecureGamingStatsStats {

	count1 := len(d)
	var ret edpt.SlbSecureGamingStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Secure_gaming_pass = in["secure_gaming_pass"].(int)
		ret.Secure_gaming_drop = in["secure_gaming_drop"].(int)
	}
	return ret
}

func dataToEndpointSlbSecureGamingStats(d *schema.ResourceData) edpt.SlbSecureGamingStats {
	var ret edpt.SlbSecureGamingStats

	ret.Stats = getObjectSlbSecureGamingStatsStats(d.Get("stats").([]interface{}))
	return ret
}
