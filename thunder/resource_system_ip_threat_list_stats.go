package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpThreatListStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_ip_threat_list_stats`: Statistics for the object ip-threat-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemIpThreatListStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_hit_count_in_sw": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Hit Count in SW",
						},
						"packet_hit_count_in_spe": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Hit Count in SPE",
						},
						"entries_added_in_sw": {
							Type: schema.TypeInt, Optional: true, Description: "Entries Added in SW",
						},
						"entries_removed_from_sw": {
							Type: schema.TypeInt, Optional: true, Description: "Entries Removed from SW",
						},
						"entries_added_in_spe": {
							Type: schema.TypeInt, Optional: true, Description: "Entries Added in SPE",
						},
						"entries_removed_from_spe": {
							Type: schema.TypeInt, Optional: true, Description: "Entries Removed from SPE",
						},
						"error_out_of_memory": {
							Type: schema.TypeInt, Optional: true, Description: "Out of memory Error",
						},
						"error_out_of_spe_entries": {
							Type: schema.TypeInt, Optional: true, Description: "Out of SPE Entries Error",
						},
					},
				},
			},
		},
	}
}

func resourceSystemIpThreatListStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemIpThreatListStatsStats := setObjectSystemIpThreatListStatsStats(res)
		d.Set("stats", SystemIpThreatListStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemIpThreatListStatsStats(ret edpt.DataSystemIpThreatListStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packet_hit_count_in_sw":   ret.DtSystemIpThreatListStats.Stats.Packet_hit_count_in_sw,
			"packet_hit_count_in_spe":  ret.DtSystemIpThreatListStats.Stats.Packet_hit_count_in_spe,
			"entries_added_in_sw":      ret.DtSystemIpThreatListStats.Stats.Entries_added_in_sw,
			"entries_removed_from_sw":  ret.DtSystemIpThreatListStats.Stats.Entries_removed_from_sw,
			"entries_added_in_spe":     ret.DtSystemIpThreatListStats.Stats.Entries_added_in_spe,
			"entries_removed_from_spe": ret.DtSystemIpThreatListStats.Stats.Entries_removed_from_spe,
			"error_out_of_memory":      ret.DtSystemIpThreatListStats.Stats.Error_out_of_memory,
			"error_out_of_spe_entries": ret.DtSystemIpThreatListStats.Stats.Error_out_of_spe_entries,
		},
	}
}

func getObjectSystemIpThreatListStatsStats(d []interface{}) edpt.SystemIpThreatListStatsStats {

	count1 := len(d)
	var ret edpt.SystemIpThreatListStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_hit_count_in_sw = in["packet_hit_count_in_sw"].(int)
		ret.Packet_hit_count_in_spe = in["packet_hit_count_in_spe"].(int)
		ret.Entries_added_in_sw = in["entries_added_in_sw"].(int)
		ret.Entries_removed_from_sw = in["entries_removed_from_sw"].(int)
		ret.Entries_added_in_spe = in["entries_added_in_spe"].(int)
		ret.Entries_removed_from_spe = in["entries_removed_from_spe"].(int)
		ret.Error_out_of_memory = in["error_out_of_memory"].(int)
		ret.Error_out_of_spe_entries = in["error_out_of_spe_entries"].(int)
	}
	return ret
}

func dataToEndpointSystemIpThreatListStats(d *schema.ResourceData) edpt.SystemIpThreatListStats {
	var ret edpt.SystemIpThreatListStats

	ret.Stats = getObjectSystemIpThreatListStatsStats(d.Get("stats").([]interface{}))
	return ret
}
