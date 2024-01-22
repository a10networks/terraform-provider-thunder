package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwDdosProtectionStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_ddos_protection_stats`: Statistics for the object ddos-protection\n\n__PLACEHOLDER__",
		ReadContext: resourceFwDdosProtectionStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ddos_entries_too_many": {
							Type: schema.TypeInt, Optional: true, Description: "Too many DDOS entries",
						},
						"ddos_entry_added": {
							Type: schema.TypeInt, Optional: true, Description: "DDOS entry added",
						},
						"ddos_entry_removed": {
							Type: schema.TypeInt, Optional: true, Description: "DDOS entry removed",
						},
						"ddos_entry_added_to_bgp": {
							Type: schema.TypeInt, Optional: true, Description: "DDoS Entry added to BGP",
						},
						"ddos_entry_removed_from_bgp": {
							Type: schema.TypeInt, Optional: true, Description: "DDoS Entry Removed from BGP",
						},
						"ddos_entry_add_to_bgp_failure": {
							Type: schema.TypeInt, Optional: true, Description: "DDoS Entry BGP add failures",
						},
						"ddos_entry_remove_from_bgp_failure": {
							Type: schema.TypeInt, Optional: true, Description: "DDOS entry BGP remove failures",
						},
						"ddos_packet_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "DDOS Packet Drop",
						},
					},
				},
			},
		},
	}
}

func resourceFwDdosProtectionStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwDdosProtectionStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwDdosProtectionStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwDdosProtectionStatsStats := setObjectFwDdosProtectionStatsStats(res)
		d.Set("stats", FwDdosProtectionStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwDdosProtectionStatsStats(ret edpt.DataFwDdosProtectionStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entries_too_many":              ret.DtFwDdosProtectionStats.Stats.Ddos_entries_too_many,
			"ddos_entry_added":                   ret.DtFwDdosProtectionStats.Stats.Ddos_entry_added,
			"ddos_entry_removed":                 ret.DtFwDdosProtectionStats.Stats.Ddos_entry_removed,
			"ddos_entry_added_to_bgp":            ret.DtFwDdosProtectionStats.Stats.Ddos_entry_added_to_bgp,
			"ddos_entry_removed_from_bgp":        ret.DtFwDdosProtectionStats.Stats.Ddos_entry_removed_from_bgp,
			"ddos_entry_add_to_bgp_failure":      ret.DtFwDdosProtectionStats.Stats.Ddos_entry_add_to_bgp_failure,
			"ddos_entry_remove_from_bgp_failure": ret.DtFwDdosProtectionStats.Stats.Ddos_entry_remove_from_bgp_failure,
			"ddos_packet_dropped":                ret.DtFwDdosProtectionStats.Stats.Ddos_packet_dropped,
		},
	}
}

func getObjectFwDdosProtectionStatsStats(d []interface{}) edpt.FwDdosProtectionStatsStats {

	count1 := len(d)
	var ret edpt.FwDdosProtectionStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entries_too_many = in["ddos_entries_too_many"].(int)
		ret.Ddos_entry_added = in["ddos_entry_added"].(int)
		ret.Ddos_entry_removed = in["ddos_entry_removed"].(int)
		ret.Ddos_entry_added_to_bgp = in["ddos_entry_added_to_bgp"].(int)
		ret.Ddos_entry_removed_from_bgp = in["ddos_entry_removed_from_bgp"].(int)
		ret.Ddos_entry_add_to_bgp_failure = in["ddos_entry_add_to_bgp_failure"].(int)
		ret.Ddos_entry_remove_from_bgp_failure = in["ddos_entry_remove_from_bgp_failure"].(int)
		ret.Ddos_packet_dropped = in["ddos_packet_dropped"].(int)
	}
	return ret
}

func dataToEndpointFwDdosProtectionStats(d *schema.ResourceData) edpt.FwDdosProtectionStats {
	var ret edpt.FwDdosProtectionStats

	ret.Stats = getObjectFwDdosProtectionStatsStats(d.Get("stats").([]interface{}))
	return ret
}
