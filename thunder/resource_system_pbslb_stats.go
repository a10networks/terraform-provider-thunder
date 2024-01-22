package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPbslbStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_pbslb_stats`: Statistics for the object pbslb\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemPbslbStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_entries": {
							Type: schema.TypeInt, Optional: true, Description: "Current PBSLB Entry Count",
						},
						"total_v4_entries_created": {
							Type: schema.TypeInt, Optional: true, Description: "Total V4 Entry Count Created",
						},
						"total_v4_entries_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total V4 Entry Count Freed",
						},
						"total_v6_entries_created": {
							Type: schema.TypeInt, Optional: true, Description: "Total V6 Entry Count Created",
						},
						"total_v6_entries_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total V6 Entry Count Freed",
						},
						"total_domain_entries_created": {
							Type: schema.TypeInt, Optional: true, Description: "Total Domain Entry Count Created",
						},
						"total_domain_entries_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total Domain Entry Count Freed",
						},
						"total_direct_action_entries_created": {
							Type: schema.TypeInt, Optional: true, Description: "Total Direct Action Entry Count Created",
						},
						"total_direct_action_entries_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total Direct Action Entry Count Freed",
						},
						"curr_entries_target_global": {
							Type: schema.TypeInt, Optional: true, Description: "Current Entry Target Global",
						},
						"curr_entries_target_vserver": {
							Type: schema.TypeInt, Optional: true, Description: "Current Entry Target Vserver",
						},
						"curr_entries_target_vport": {
							Type: schema.TypeInt, Optional: true, Description: "Current Entry Target Vport",
						},
						"curr_entries_target_loc": {
							Type: schema.TypeInt, Optional: true, Description: "Current Entry Target LOC",
						},
						"curr_entries_target_rserver": {
							Type: schema.TypeInt, Optional: true, Description: "Current Entry Target Rserver",
						},
						"curr_entries_target_rport": {
							Type: schema.TypeInt, Optional: true, Description: "Current Entry Target Rport",
						},
						"curr_entries_target_service": {
							Type: schema.TypeInt, Optional: true, Description: "Current Entry Target Service",
						},
						"curr_entries_stats": {
							Type: schema.TypeInt, Optional: true, Description: "Current Entry Stats Count",
						},
					},
				},
			},
		},
	}
}

func resourceSystemPbslbStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPbslbStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPbslbStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemPbslbStatsStats := setObjectSystemPbslbStatsStats(res)
		d.Set("stats", SystemPbslbStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemPbslbStatsStats(ret edpt.DataSystemPbslbStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_entries":                        ret.DtSystemPbslbStats.Stats.Curr_entries,
			"total_v4_entries_created":            ret.DtSystemPbslbStats.Stats.Total_v4_entries_created,
			"total_v4_entries_freed":              ret.DtSystemPbslbStats.Stats.Total_v4_entries_freed,
			"total_v6_entries_created":            ret.DtSystemPbslbStats.Stats.Total_v6_entries_created,
			"total_v6_entries_freed":              ret.DtSystemPbslbStats.Stats.Total_v6_entries_freed,
			"total_domain_entries_created":        ret.DtSystemPbslbStats.Stats.Total_domain_entries_created,
			"total_domain_entries_freed":          ret.DtSystemPbslbStats.Stats.Total_domain_entries_freed,
			"total_direct_action_entries_created": ret.DtSystemPbslbStats.Stats.Total_direct_action_entries_created,
			"total_direct_action_entries_freed":   ret.DtSystemPbslbStats.Stats.Total_direct_action_entries_freed,
			"curr_entries_target_global":          ret.DtSystemPbslbStats.Stats.Curr_entries_target_global,
			"curr_entries_target_vserver":         ret.DtSystemPbslbStats.Stats.Curr_entries_target_vserver,
			"curr_entries_target_vport":           ret.DtSystemPbslbStats.Stats.Curr_entries_target_vport,
			"curr_entries_target_loc":             ret.DtSystemPbslbStats.Stats.Curr_entries_target_loc,
			"curr_entries_target_rserver":         ret.DtSystemPbslbStats.Stats.Curr_entries_target_rserver,
			"curr_entries_target_rport":           ret.DtSystemPbslbStats.Stats.Curr_entries_target_rport,
			"curr_entries_target_service":         ret.DtSystemPbslbStats.Stats.Curr_entries_target_service,
			"curr_entries_stats":                  ret.DtSystemPbslbStats.Stats.Curr_entries_stats,
		},
	}
}

func getObjectSystemPbslbStatsStats(d []interface{}) edpt.SystemPbslbStatsStats {

	count1 := len(d)
	var ret edpt.SystemPbslbStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_entries = in["curr_entries"].(int)
		ret.Total_v4_entries_created = in["total_v4_entries_created"].(int)
		ret.Total_v4_entries_freed = in["total_v4_entries_freed"].(int)
		ret.Total_v6_entries_created = in["total_v6_entries_created"].(int)
		ret.Total_v6_entries_freed = in["total_v6_entries_freed"].(int)
		ret.Total_domain_entries_created = in["total_domain_entries_created"].(int)
		ret.Total_domain_entries_freed = in["total_domain_entries_freed"].(int)
		ret.Total_direct_action_entries_created = in["total_direct_action_entries_created"].(int)
		ret.Total_direct_action_entries_freed = in["total_direct_action_entries_freed"].(int)
		ret.Curr_entries_target_global = in["curr_entries_target_global"].(int)
		ret.Curr_entries_target_vserver = in["curr_entries_target_vserver"].(int)
		ret.Curr_entries_target_vport = in["curr_entries_target_vport"].(int)
		ret.Curr_entries_target_loc = in["curr_entries_target_loc"].(int)
		ret.Curr_entries_target_rserver = in["curr_entries_target_rserver"].(int)
		ret.Curr_entries_target_rport = in["curr_entries_target_rport"].(int)
		ret.Curr_entries_target_service = in["curr_entries_target_service"].(int)
		ret.Curr_entries_stats = in["curr_entries_stats"].(int)
	}
	return ret
}

func dataToEndpointSystemPbslbStats(d *schema.ResourceData) edpt.SystemPbslbStats {
	var ret edpt.SystemPbslbStats

	ret.Stats = getObjectSystemPbslbStatsStats(d.Get("stats").([]interface{}))
	return ret
}
