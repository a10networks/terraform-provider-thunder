package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPlayerIdGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_player_id_global_stats`: Statistics for the object player-id-global\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbPlayerIdGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_playerids_created": {
							Type: schema.TypeInt, Optional: true, Description: "Playerid records created",
						},
						"total_playerids_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "Playerid records deleted",
						},
						"total_abs_max_age_outs": {
							Type: schema.TypeInt, Optional: true, Description: "Playerid records max time aged out",
						},
						"total_pkt_activity_age_outs": {
							Type: schema.TypeInt, Optional: true, Description: "Playerid records idle timeout",
						},
						"total_invalid_playerid_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid playerid packets",
						},
						"total_invalid_playerid_drops": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid playerid packet drops",
						},
						"total_valid_playerid_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Valid playerid packets",
						},
					},
				},
			},
		},
	}
}

func resourceSlbPlayerIdGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPlayerIdGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPlayerIdGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbPlayerIdGlobalStatsStats := setObjectSlbPlayerIdGlobalStatsStats(res)
		d.Set("stats", SlbPlayerIdGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbPlayerIdGlobalStatsStats(ret edpt.DataSlbPlayerIdGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_playerids_created":      ret.DtSlbPlayerIdGlobalStats.Stats.Total_playerids_created,
			"total_playerids_deleted":      ret.DtSlbPlayerIdGlobalStats.Stats.Total_playerids_deleted,
			"total_abs_max_age_outs":       ret.DtSlbPlayerIdGlobalStats.Stats.Total_abs_max_age_outs,
			"total_pkt_activity_age_outs":  ret.DtSlbPlayerIdGlobalStats.Stats.Total_pkt_activity_age_outs,
			"total_invalid_playerid_pkts":  ret.DtSlbPlayerIdGlobalStats.Stats.Total_invalid_playerid_pkts,
			"total_invalid_playerid_drops": ret.DtSlbPlayerIdGlobalStats.Stats.Total_invalid_playerid_drops,
			"total_valid_playerid_pkts":    ret.DtSlbPlayerIdGlobalStats.Stats.Total_valid_playerid_pkts,
		},
	}
}

func getObjectSlbPlayerIdGlobalStatsStats(d []interface{}) edpt.SlbPlayerIdGlobalStatsStats {

	count1 := len(d)
	var ret edpt.SlbPlayerIdGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_playerids_created = in["total_playerids_created"].(int)
		ret.Total_playerids_deleted = in["total_playerids_deleted"].(int)
		ret.Total_abs_max_age_outs = in["total_abs_max_age_outs"].(int)
		ret.Total_pkt_activity_age_outs = in["total_pkt_activity_age_outs"].(int)
		ret.Total_invalid_playerid_pkts = in["total_invalid_playerid_pkts"].(int)
		ret.Total_invalid_playerid_drops = in["total_invalid_playerid_drops"].(int)
		ret.Total_valid_playerid_pkts = in["total_valid_playerid_pkts"].(int)
	}
	return ret
}

func dataToEndpointSlbPlayerIdGlobalStats(d *schema.ResourceData) edpt.SlbPlayerIdGlobalStats {
	var ret edpt.SlbPlayerIdGlobalStats

	ret.Stats = getObjectSlbPlayerIdGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
