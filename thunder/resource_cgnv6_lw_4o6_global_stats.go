package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6GlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lw_4o6_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Lw4o6GlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total Entries Configured",
						},
						"self_hairpinning_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Self-Hairpinning Drops",
						},
						"all_hairpinning_drop": {
							Type: schema.TypeInt, Optional: true, Description: "All Hairpinning Drops",
						},
						"no_match_icmpv6_sent": {
							Type: schema.TypeInt, Optional: true, Description: "No-Forward-Match ICMPv6 Sent",
						},
						"no_match_icmp_sent": {
							Type: schema.TypeInt, Optional: true, Description: "No-Reverse-Match ICMP Sent",
						},
						"icmp_inbound_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound ICMP Drops",
						},
						"fwd_lookup_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Route Lookup Failed",
						},
						"rev_lookup_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Route Lookup Failed",
						},
						"interface_not_configured": {
							Type: schema.TypeInt, Optional: true, Description: "LW-4over6 Interfaces not Configured Drops",
						},
						"no_binding_table_matches_fwd": {
							Type: schema.TypeInt, Optional: true, Description: "No Forward Binding Table Entry Match Drops",
						},
						"no_binding_table_matches_rev": {
							Type: schema.TypeInt, Optional: true, Description: "No Reverse Binding Table Entry Match Drops",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6Lw4o6GlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6GlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6GlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Lw4o6GlobalStatsStats := setObjectCgnv6Lw4o6GlobalStatsStats(res)
		d.Set("stats", Cgnv6Lw4o6GlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Lw4o6GlobalStatsStats(ret edpt.DataCgnv6Lw4o6GlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_count":                  ret.DtCgnv6Lw4o6GlobalStats.Stats.Entry_count,
			"self_hairpinning_drop":        ret.DtCgnv6Lw4o6GlobalStats.Stats.Self_hairpinning_drop,
			"all_hairpinning_drop":         ret.DtCgnv6Lw4o6GlobalStats.Stats.All_hairpinning_drop,
			"no_match_icmpv6_sent":         ret.DtCgnv6Lw4o6GlobalStats.Stats.No_match_icmpv6_sent,
			"no_match_icmp_sent":           ret.DtCgnv6Lw4o6GlobalStats.Stats.No_match_icmp_sent,
			"icmp_inbound_drop":            ret.DtCgnv6Lw4o6GlobalStats.Stats.Icmp_inbound_drop,
			"fwd_lookup_failed":            ret.DtCgnv6Lw4o6GlobalStats.Stats.Fwd_lookup_failed,
			"rev_lookup_failed":            ret.DtCgnv6Lw4o6GlobalStats.Stats.Rev_lookup_failed,
			"interface_not_configured":     ret.DtCgnv6Lw4o6GlobalStats.Stats.Interface_not_configured,
			"no_binding_table_matches_fwd": ret.DtCgnv6Lw4o6GlobalStats.Stats.No_binding_table_matches_fwd,
			"no_binding_table_matches_rev": ret.DtCgnv6Lw4o6GlobalStats.Stats.No_binding_table_matches_rev,
		},
	}
}

func getObjectCgnv6Lw4o6GlobalStatsStats(d []interface{}) edpt.Cgnv6Lw4o6GlobalStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6Lw4o6GlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Entry_count = in["entry_count"].(int)
		ret.Self_hairpinning_drop = in["self_hairpinning_drop"].(int)
		ret.All_hairpinning_drop = in["all_hairpinning_drop"].(int)
		ret.No_match_icmpv6_sent = in["no_match_icmpv6_sent"].(int)
		ret.No_match_icmp_sent = in["no_match_icmp_sent"].(int)
		ret.Icmp_inbound_drop = in["icmp_inbound_drop"].(int)
		ret.Fwd_lookup_failed = in["fwd_lookup_failed"].(int)
		ret.Rev_lookup_failed = in["rev_lookup_failed"].(int)
		ret.Interface_not_configured = in["interface_not_configured"].(int)
		ret.No_binding_table_matches_fwd = in["no_binding_table_matches_fwd"].(int)
		ret.No_binding_table_matches_rev = in["no_binding_table_matches_rev"].(int)
	}
	return ret
}

func dataToEndpointCgnv6Lw4o6GlobalStats(d *schema.ResourceData) edpt.Cgnv6Lw4o6GlobalStats {
	var ret edpt.Cgnv6Lw4o6GlobalStats

	ret.Stats = getObjectCgnv6Lw4o6GlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
