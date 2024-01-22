package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSportRateLimitStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_sport_rate_limit_stats`: Statistics for the object sport-rate-limit\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSportRateLimitStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alloc_sport": {
							Type: schema.TypeInt, Optional: true, Description: "Alloc'd src port entry",
						},
						"alloc_sportip": {
							Type: schema.TypeInt, Optional: true, Description: "Alloc'd src port-ip entry",
						},
						"freed_sport": {
							Type: schema.TypeInt, Optional: true, Description: "Freed src port entry",
						},
						"freed_sportip": {
							Type: schema.TypeInt, Optional: true, Description: "Freed src port-ip entry",
						},
						"total_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total rate exceed drop",
						},
						"total_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Total rate exceed reset",
						},
						"total_log": {
							Type: schema.TypeInt, Optional: true, Description: "Total log sent",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSportRateLimitStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSportRateLimitStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSportRateLimitStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSportRateLimitStatsStats := setObjectSlbSportRateLimitStatsStats(res)
		d.Set("stats", SlbSportRateLimitStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSportRateLimitStatsStats(ret edpt.DataSlbSportRateLimitStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"alloc_sport":   ret.DtSlbSportRateLimitStats.Stats.Alloc_sport,
			"alloc_sportip": ret.DtSlbSportRateLimitStats.Stats.Alloc_sportip,
			"freed_sport":   ret.DtSlbSportRateLimitStats.Stats.Freed_sport,
			"freed_sportip": ret.DtSlbSportRateLimitStats.Stats.Freed_sportip,
			"total_drop":    ret.DtSlbSportRateLimitStats.Stats.Total_drop,
			"total_reset":   ret.DtSlbSportRateLimitStats.Stats.Total_reset,
			"total_log":     ret.DtSlbSportRateLimitStats.Stats.Total_log,
		},
	}
}

func getObjectSlbSportRateLimitStatsStats(d []interface{}) edpt.SlbSportRateLimitStatsStats {

	count1 := len(d)
	var ret edpt.SlbSportRateLimitStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Alloc_sport = in["alloc_sport"].(int)
		ret.Alloc_sportip = in["alloc_sportip"].(int)
		ret.Freed_sport = in["freed_sport"].(int)
		ret.Freed_sportip = in["freed_sportip"].(int)
		ret.Total_drop = in["total_drop"].(int)
		ret.Total_reset = in["total_reset"].(int)
		ret.Total_log = in["total_log"].(int)
	}
	return ret
}

func dataToEndpointSlbSportRateLimitStats(d *schema.ResourceData) edpt.SlbSportRateLimitStats {
	var ret edpt.SlbSportRateLimitStats

	ret.Stats = getObjectSlbSportRateLimitStatsStats(d.Get("stats").([]interface{}))
	return ret
}
