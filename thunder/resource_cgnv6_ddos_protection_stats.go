package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DdosProtectionStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_ddos_protection_stats`: Statistics for the object ddos-protection\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6DdosProtectionStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l3_entry_added": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Entry Added",
						},
						"l3_entry_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Entry Deleted",
						},
						"l3_entry_added_to_bgp": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Entry added to BGP",
						},
						"l3_entry_removed_from_bgp": {
							Type: schema.TypeInt, Optional: true, Description: "Entry removed from BGP",
						},
						"l3_entry_added_to_hw": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Entry added to HW",
						},
						"l3_entry_removed_from_hw": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Entry removed from HW",
						},
						"l3_entry_too_many": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Too many entries",
						},
						"l3_entry_match_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Entry match drop",
						},
						"l3_entry_match_drop_hw": {
							Type: schema.TypeInt, Optional: true, Description: "L3 HW entry match drop",
						},
						"l3_entry_drop_max_hw_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Entry Drop due to HW Limit Exceeded",
						},
						"l4_entry_added": {
							Type: schema.TypeInt, Optional: true, Description: "L4 Entry added",
						},
						"l4_entry_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "L4 Entry deleted",
						},
						"l4_entry_added_to_hw": {
							Type: schema.TypeInt, Optional: true, Description: "L4 Entry added to HW",
						},
						"l4_entry_removed_from_hw": {
							Type: schema.TypeInt, Optional: true, Description: "L4 Entry removed from HW",
						},
						"l4_hw_out_of_entries": {
							Type: schema.TypeInt, Optional: true, Description: "HW out of L4 entries",
						},
						"l4_entry_match_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L4 Entry match drop",
						},
						"l4_entry_match_drop_hw": {
							Type: schema.TypeInt, Optional: true, Description: "L4 HW Entry match drop",
						},
						"l4_entry_drop_max_hw_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "L4 Entry Drop due to HW Limit Exceeded",
						},
						"l4_entry_list_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "L4 Entry list alloc",
						},
						"l4_entry_list_free": {
							Type: schema.TypeInt, Optional: true, Description: "L4 Entry list free",
						},
						"l4_entry_list_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Description: "L4 Entry list alloc failures",
						},
						"ip_node_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Node alloc",
						},
						"ip_node_free": {
							Type: schema.TypeInt, Optional: true, Description: "Node free",
						},
						"ip_node_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Node alloc failures",
						},
						"ip_port_block_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Port block alloc",
						},
						"ip_port_block_free": {
							Type: schema.TypeInt, Optional: true, Description: "Port block free",
						},
						"ip_port_block_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Port block alloc failure",
						},
						"ip_other_block_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Other block alloc",
						},
						"ip_other_block_free": {
							Type: schema.TypeInt, Optional: true, Description: "Other block free",
						},
						"ip_other_block_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Other block alloc failure",
						},
						"entry_added_shadow": {
							Type: schema.TypeInt, Optional: true, Description: "Entry added shadow",
						},
						"entry_invalidated": {
							Type: schema.TypeInt, Optional: true, Description: "Entry invalidated",
						},
						"l3_entry_add_to_bgp_failure": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Entry BGP add failures",
						},
						"l3_entry_remove_from_bgp_failure": {
							Type: schema.TypeInt, Optional: true, Description: "L3 entry BGP remove failures",
						},
						"l3_entry_add_to_hw_failure": {
							Type: schema.TypeInt, Optional: true, Description: "L3 entry HW add failure",
						},
						"syn_cookie_syn_ack_sent": {
							Type: schema.TypeInt, Optional: true, Description: "SYN cookie SYN ACK sent",
						},
						"syn_cookie_verification_passed": {
							Type: schema.TypeInt, Optional: true, Description: "SYN cookie verification passed",
						},
						"syn_cookie_verification_failed": {
							Type: schema.TypeInt, Optional: true, Description: "SYN cookie verification failed",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6DdosProtectionStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DdosProtectionStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DdosProtectionStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6DdosProtectionStatsStats := setObjectCgnv6DdosProtectionStatsStats(res)
		d.Set("stats", Cgnv6DdosProtectionStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6DdosProtectionStatsStats(ret edpt.DataCgnv6DdosProtectionStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"l3_entry_added":                   ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_added,
			"l3_entry_deleted":                 ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_deleted,
			"l3_entry_added_to_bgp":            ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_added_to_bgp,
			"l3_entry_removed_from_bgp":        ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_removed_from_bgp,
			"l3_entry_added_to_hw":             ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_added_to_hw,
			"l3_entry_removed_from_hw":         ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_removed_from_hw,
			"l3_entry_too_many":                ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_too_many,
			"l3_entry_match_drop":              ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_match_drop,
			"l3_entry_match_drop_hw":           ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_match_drop_hw,
			"l3_entry_drop_max_hw_exceeded":    ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_drop_max_hw_exceeded,
			"l4_entry_added":                   ret.DtCgnv6DdosProtectionStats.Stats.L4_entry_added,
			"l4_entry_deleted":                 ret.DtCgnv6DdosProtectionStats.Stats.L4_entry_deleted,
			"l4_entry_added_to_hw":             ret.DtCgnv6DdosProtectionStats.Stats.L4_entry_added_to_hw,
			"l4_entry_removed_from_hw":         ret.DtCgnv6DdosProtectionStats.Stats.L4_entry_removed_from_hw,
			"l4_hw_out_of_entries":             ret.DtCgnv6DdosProtectionStats.Stats.L4_hw_out_of_entries,
			"l4_entry_match_drop":              ret.DtCgnv6DdosProtectionStats.Stats.L4_entry_match_drop,
			"l4_entry_match_drop_hw":           ret.DtCgnv6DdosProtectionStats.Stats.L4_entry_match_drop_hw,
			"l4_entry_drop_max_hw_exceeded":    ret.DtCgnv6DdosProtectionStats.Stats.L4_entry_drop_max_hw_exceeded,
			"l4_entry_list_alloc":              ret.DtCgnv6DdosProtectionStats.Stats.L4_entry_list_alloc,
			"l4_entry_list_free":               ret.DtCgnv6DdosProtectionStats.Stats.L4_entry_list_free,
			"l4_entry_list_alloc_failure":      ret.DtCgnv6DdosProtectionStats.Stats.L4_entry_list_alloc_failure,
			"ip_node_alloc":                    ret.DtCgnv6DdosProtectionStats.Stats.Ip_node_alloc,
			"ip_node_free":                     ret.DtCgnv6DdosProtectionStats.Stats.Ip_node_free,
			"ip_node_alloc_failure":            ret.DtCgnv6DdosProtectionStats.Stats.Ip_node_alloc_failure,
			"ip_port_block_alloc":              ret.DtCgnv6DdosProtectionStats.Stats.Ip_port_block_alloc,
			"ip_port_block_free":               ret.DtCgnv6DdosProtectionStats.Stats.Ip_port_block_free,
			"ip_port_block_alloc_failure":      ret.DtCgnv6DdosProtectionStats.Stats.Ip_port_block_alloc_failure,
			"ip_other_block_alloc":             ret.DtCgnv6DdosProtectionStats.Stats.Ip_other_block_alloc,
			"ip_other_block_free":              ret.DtCgnv6DdosProtectionStats.Stats.Ip_other_block_free,
			"ip_other_block_alloc_failure":     ret.DtCgnv6DdosProtectionStats.Stats.Ip_other_block_alloc_failure,
			"entry_added_shadow":               ret.DtCgnv6DdosProtectionStats.Stats.Entry_added_shadow,
			"entry_invalidated":                ret.DtCgnv6DdosProtectionStats.Stats.Entry_invalidated,
			"l3_entry_add_to_bgp_failure":      ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_add_to_bgp_failure,
			"l3_entry_remove_from_bgp_failure": ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_remove_from_bgp_failure,
			"l3_entry_add_to_hw_failure":       ret.DtCgnv6DdosProtectionStats.Stats.L3_entry_add_to_hw_failure,
			"syn_cookie_syn_ack_sent":          ret.DtCgnv6DdosProtectionStats.Stats.Syn_cookie_syn_ack_sent,
			"syn_cookie_verification_passed":   ret.DtCgnv6DdosProtectionStats.Stats.Syn_cookie_verification_passed,
			"syn_cookie_verification_failed":   ret.DtCgnv6DdosProtectionStats.Stats.Syn_cookie_verification_failed,
		},
	}
}

func getObjectCgnv6DdosProtectionStatsStats(d []interface{}) edpt.Cgnv6DdosProtectionStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6DdosProtectionStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L3_entry_added = in["l3_entry_added"].(int)
		ret.L3_entry_deleted = in["l3_entry_deleted"].(int)
		ret.L3_entry_added_to_bgp = in["l3_entry_added_to_bgp"].(int)
		ret.L3_entry_removed_from_bgp = in["l3_entry_removed_from_bgp"].(int)
		ret.L3_entry_added_to_hw = in["l3_entry_added_to_hw"].(int)
		ret.L3_entry_removed_from_hw = in["l3_entry_removed_from_hw"].(int)
		ret.L3_entry_too_many = in["l3_entry_too_many"].(int)
		ret.L3_entry_match_drop = in["l3_entry_match_drop"].(int)
		ret.L3_entry_match_drop_hw = in["l3_entry_match_drop_hw"].(int)
		ret.L3_entry_drop_max_hw_exceeded = in["l3_entry_drop_max_hw_exceeded"].(int)
		ret.L4_entry_added = in["l4_entry_added"].(int)
		ret.L4_entry_deleted = in["l4_entry_deleted"].(int)
		ret.L4_entry_added_to_hw = in["l4_entry_added_to_hw"].(int)
		ret.L4_entry_removed_from_hw = in["l4_entry_removed_from_hw"].(int)
		ret.L4_hw_out_of_entries = in["l4_hw_out_of_entries"].(int)
		ret.L4_entry_match_drop = in["l4_entry_match_drop"].(int)
		ret.L4_entry_match_drop_hw = in["l4_entry_match_drop_hw"].(int)
		ret.L4_entry_drop_max_hw_exceeded = in["l4_entry_drop_max_hw_exceeded"].(int)
		ret.L4_entry_list_alloc = in["l4_entry_list_alloc"].(int)
		ret.L4_entry_list_free = in["l4_entry_list_free"].(int)
		ret.L4_entry_list_alloc_failure = in["l4_entry_list_alloc_failure"].(int)
		ret.Ip_node_alloc = in["ip_node_alloc"].(int)
		ret.Ip_node_free = in["ip_node_free"].(int)
		ret.Ip_node_alloc_failure = in["ip_node_alloc_failure"].(int)
		ret.Ip_port_block_alloc = in["ip_port_block_alloc"].(int)
		ret.Ip_port_block_free = in["ip_port_block_free"].(int)
		ret.Ip_port_block_alloc_failure = in["ip_port_block_alloc_failure"].(int)
		ret.Ip_other_block_alloc = in["ip_other_block_alloc"].(int)
		ret.Ip_other_block_free = in["ip_other_block_free"].(int)
		ret.Ip_other_block_alloc_failure = in["ip_other_block_alloc_failure"].(int)
		ret.Entry_added_shadow = in["entry_added_shadow"].(int)
		ret.Entry_invalidated = in["entry_invalidated"].(int)
		ret.L3_entry_add_to_bgp_failure = in["l3_entry_add_to_bgp_failure"].(int)
		ret.L3_entry_remove_from_bgp_failure = in["l3_entry_remove_from_bgp_failure"].(int)
		ret.L3_entry_add_to_hw_failure = in["l3_entry_add_to_hw_failure"].(int)
		ret.Syn_cookie_syn_ack_sent = in["syn_cookie_syn_ack_sent"].(int)
		ret.Syn_cookie_verification_passed = in["syn_cookie_verification_passed"].(int)
		ret.Syn_cookie_verification_failed = in["syn_cookie_verification_failed"].(int)
	}
	return ret
}

func dataToEndpointCgnv6DdosProtectionStats(d *schema.ResourceData) edpt.Cgnv6DdosProtectionStats {
	var ret edpt.Cgnv6DdosProtectionStats

	ret.Stats = getObjectCgnv6DdosProtectionStatsStats(d.Get("stats").([]interface{}))
	return ret
}
