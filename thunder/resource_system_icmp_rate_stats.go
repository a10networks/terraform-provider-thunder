package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIcmpRateStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_icmp_rate_stats`: Statistics for the object icmp-rate\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemIcmpRateStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"over_limit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Over limit drops",
						},
						"limit_intf_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Interfaces rate limit drops",
						},
						"limit_vserver_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual Server rate limit drops",
						},
						"limit_total_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total rate limit drops",
						},
						"lockup_time_left": {
							Type: schema.TypeInt, Optional: true, Description: "Lockup time left",
						},
						"curr_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Current rate",
						},
						"v6_over_limit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Over limit drops (v6)",
						},
						"v6_limit_intf_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Interfaces rate limit drops (v6)",
						},
						"v6_limit_vserver_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual Server rate limit drops (v6)",
						},
						"v6_limit_total_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total rate limit drops (v6)",
						},
						"v6_lockup_time_left": {
							Type: schema.TypeInt, Optional: true, Description: "Lockup time left (v6)",
						},
						"v6_curr_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Current rate (v6)",
						},
					},
				},
			},
		},
	}
}

func resourceSystemIcmpRateStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmpRateStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmpRateStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemIcmpRateStatsStats := setObjectSystemIcmpRateStatsStats(res)
		d.Set("stats", SystemIcmpRateStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemIcmpRateStatsStats(ret edpt.DataSystemIcmpRateStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"over_limit_drop":       ret.DtSystemIcmpRateStats.Stats.Over_limit_drop,
			"limit_intf_drop":       ret.DtSystemIcmpRateStats.Stats.Limit_intf_drop,
			"limit_vserver_drop":    ret.DtSystemIcmpRateStats.Stats.Limit_vserver_drop,
			"limit_total_drop":      ret.DtSystemIcmpRateStats.Stats.Limit_total_drop,
			"lockup_time_left":      ret.DtSystemIcmpRateStats.Stats.Lockup_time_left,
			"curr_rate":             ret.DtSystemIcmpRateStats.Stats.Curr_rate,
			"v6_over_limit_drop":    ret.DtSystemIcmpRateStats.Stats.V6_over_limit_drop,
			"v6_limit_intf_drop":    ret.DtSystemIcmpRateStats.Stats.V6_limit_intf_drop,
			"v6_limit_vserver_drop": ret.DtSystemIcmpRateStats.Stats.V6_limit_vserver_drop,
			"v6_limit_total_drop":   ret.DtSystemIcmpRateStats.Stats.V6_limit_total_drop,
			"v6_lockup_time_left":   ret.DtSystemIcmpRateStats.Stats.V6_lockup_time_left,
			"v6_curr_rate":          ret.DtSystemIcmpRateStats.Stats.V6_curr_rate,
		},
	}
}

func getObjectSystemIcmpRateStatsStats(d []interface{}) edpt.SystemIcmpRateStatsStats {

	count1 := len(d)
	var ret edpt.SystemIcmpRateStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Over_limit_drop = in["over_limit_drop"].(int)
		ret.Limit_intf_drop = in["limit_intf_drop"].(int)
		ret.Limit_vserver_drop = in["limit_vserver_drop"].(int)
		ret.Limit_total_drop = in["limit_total_drop"].(int)
		ret.Lockup_time_left = in["lockup_time_left"].(int)
		ret.Curr_rate = in["curr_rate"].(int)
		ret.V6_over_limit_drop = in["v6_over_limit_drop"].(int)
		ret.V6_limit_intf_drop = in["v6_limit_intf_drop"].(int)
		ret.V6_limit_vserver_drop = in["v6_limit_vserver_drop"].(int)
		ret.V6_limit_total_drop = in["v6_limit_total_drop"].(int)
		ret.V6_lockup_time_left = in["v6_lockup_time_left"].(int)
		ret.V6_curr_rate = in["v6_curr_rate"].(int)
	}
	return ret
}

func dataToEndpointSystemIcmpRateStats(d *schema.ResourceData) edpt.SystemIcmpRateStats {
	var ret edpt.SystemIcmpRateStats

	ret.Stats = getObjectSystemIcmpRateStatsStats(d.Get("stats").([]interface{}))
	return ret
}
