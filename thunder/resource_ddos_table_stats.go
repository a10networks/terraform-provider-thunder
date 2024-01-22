package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTableStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_table_stats`: Statistics for the object table\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosTableStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_learn": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Learned",
						},
						"dst_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Hit",
						},
						"dst_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Missed",
						},
						"dst_entry_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Aged",
						},
						"src_learn": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Learned",
						},
						"src_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Hit",
						},
						"src_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Missed",
						},
						"src_entry_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Aged",
						},
						"src_dst_learn": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Entry Learned",
						},
						"src_dst_hit": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Entry Hit",
						},
						"src_dst_miss": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Entry Missed",
						},
						"src_dst_entry_aged": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Entry Aged",
						},
						"telem_err_misc": {
							Type: schema.TypeInt, Optional: true, Description: "From-l3-peer: Misc Error",
						},
						"telem_route_add_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "From-l3-peer: Route-add Received",
						},
						"telem_route_del_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "From-l3-peer: Route-del Received",
						},
						"telem_entry_created": {
							Type: schema.TypeInt, Optional: true, Description: "From-l3-peer: Zone Entry Created",
						},
						"telem_entry_cleared": {
							Type: schema.TypeInt, Optional: true, Description: "From-l3-peer: Zone Entry Deleted",
						},
						"telem_err_telem_entry_pre_exist": {
							Type: schema.TypeInt, Optional: true, Description: "From-l3-peer: Zone Entry Pre-exist",
						},
						"telem_err_conflict_with_static": {
							Type: schema.TypeInt, Optional: true, Description: "From-l3-peer: Conflict with Static Entry",
						},
						"telem_err_fail_to_create": {
							Type: schema.TypeInt, Optional: true, Description: "From-l3-peer: Zone Entry Create Fail",
						},
						"telem_err_fail_to_delete": {
							Type: schema.TypeInt, Optional: true, Description: "From-l3-peer: Zone Entry Delete Fail",
						},
						"src_zone_service_learn": {
							Type: schema.TypeInt, Optional: true, Description: "SrcZoneService Entry Learned",
						},
						"src_zone_service_hit": {
							Type: schema.TypeInt, Optional: true, Description: "SrcZoneService Entry Hit",
						},
						"src_zone_service_miss": {
							Type: schema.TypeInt, Optional: true, Description: "SrcZoneService Entry Missed",
						},
						"src_zone_service_entry_aged": {
							Type: schema.TypeInt, Optional: true, Description: "SrcZoneService Entry Aged",
						},
						"dst_white_list": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Whitelisted",
						},
						"src_white_list": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Whitelisted",
						},
						"src_dst_white_list": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Entry Whitelisted",
						},
						"src_zone_service_white_list": {
							Type: schema.TypeInt, Optional: true, Description: "SrcZoneService Entry Whitelisted",
						},
						"dst_black_list": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Blacklisted",
						},
						"src_black_list": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Blacklisted",
						},
						"src_dst_black_list": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Entry Blacklisted",
						},
						"src_zone_service_black_list": {
							Type: schema.TypeInt, Optional: true, Description: "SrcZoneService Entry Blacklisted",
						},
						"dst_learning_thre_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Dynamic Entry Count Overflow",
						},
						"dst_over_thre_policy_at_learning": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Overflow Policy Hit At Learning Stage",
						},
						"src_learning_thre_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src Dynamic Entry Count Overflow",
						},
						"src_over_thre_policy_at_lookup": {
							Type: schema.TypeInt, Optional: true, Description: "Src Overflow Policy Hit At Lookup Stage",
						},
						"src_over_thre_policy_at_learning": {
							Type: schema.TypeInt, Optional: true, Description: "Src Overflow Policy Hit At Learning Stage",
						},
						"src_dst_learning_thre_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Dynamic Entry Count Overflow",
						},
						"src_dst_over_thre_policy_at_lookup": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Overflow Policy Hit At Lookup Stage",
						},
						"src_dst_over_thre_policy_at_learning": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Overflow Policy Hit At Learning Stage",
						},
						"src_zone_service_learning_thre_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "SrcZoneService Dynamic Entry Count Overflow",
						},
						"src_zone_service_over_thre_policy_at_lookup": {
							Type: schema.TypeInt, Optional: true, Description: "SrcZoneService Overflow Policy Lookup Hit",
						},
						"src_zone_service_over_thre_policy_at_learning": {
							Type: schema.TypeInt, Optional: true, Description: "SrcZoneService Overflow Policy Learning Hit",
						},
						"entry_oom": {
							Type: schema.TypeInt, Optional: true, Description: "Out of Entry Memory",
						},
						"entry_ext_oom": {
							Type: schema.TypeInt, Optional: true, Description: "Out of Entry Extension Memory",
						},
						"src_dst_classlist_overflow_policy_at_learning": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Class-List Overflow Policy Hit",
						},
						"src_zone_service_classlist_overflow_policy_at_learning": {
							Type: schema.TypeInt, Optional: true, Description: "SrcZoneService Class-List Overflow Policy Hit",
						},
					},
				},
			},
		},
	}
}

func resourceDdosTableStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTableStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTableStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosTableStatsStats := setObjectDdosTableStatsStats(res)
		d.Set("stats", DdosTableStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosTableStatsStats(ret edpt.DataDdosTableStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dst_learn":                                     ret.DtDdosTableStats.Stats.Dst_learn,
			"dst_hit":                                       ret.DtDdosTableStats.Stats.Dst_hit,
			"dst_miss":                                      ret.DtDdosTableStats.Stats.Dst_miss,
			"dst_entry_aged":                                ret.DtDdosTableStats.Stats.Dst_entry_aged,
			"src_learn":                                     ret.DtDdosTableStats.Stats.Src_learn,
			"src_hit":                                       ret.DtDdosTableStats.Stats.Src_hit,
			"src_miss":                                      ret.DtDdosTableStats.Stats.Src_miss,
			"src_entry_aged":                                ret.DtDdosTableStats.Stats.Src_entry_aged,
			"src_dst_learn":                                 ret.DtDdosTableStats.Stats.Src_dst_learn,
			"src_dst_hit":                                   ret.DtDdosTableStats.Stats.Src_dst_hit,
			"src_dst_miss":                                  ret.DtDdosTableStats.Stats.Src_dst_miss,
			"src_dst_entry_aged":                            ret.DtDdosTableStats.Stats.Src_dst_entry_aged,
			"telem_err_misc":                                ret.DtDdosTableStats.Stats.Telem_err_misc,
			"telem_route_add_rcvd":                          ret.DtDdosTableStats.Stats.Telem_route_add_rcvd,
			"telem_route_del_rcvd":                          ret.DtDdosTableStats.Stats.Telem_route_del_rcvd,
			"telem_entry_created":                           ret.DtDdosTableStats.Stats.Telem_entry_created,
			"telem_entry_cleared":                           ret.DtDdosTableStats.Stats.Telem_entry_cleared,
			"telem_err_telem_entry_pre_exist":               ret.DtDdosTableStats.Stats.Telem_err_telem_entry_pre_exist,
			"telem_err_conflict_with_static":                ret.DtDdosTableStats.Stats.Telem_err_conflict_with_static,
			"telem_err_fail_to_create":                      ret.DtDdosTableStats.Stats.Telem_err_fail_to_create,
			"telem_err_fail_to_delete":                      ret.DtDdosTableStats.Stats.Telem_err_fail_to_delete,
			"src_zone_service_learn":                        ret.DtDdosTableStats.Stats.Src_zone_service_learn,
			"src_zone_service_hit":                          ret.DtDdosTableStats.Stats.Src_zone_service_hit,
			"src_zone_service_miss":                         ret.DtDdosTableStats.Stats.Src_zone_service_miss,
			"src_zone_service_entry_aged":                   ret.DtDdosTableStats.Stats.Src_zone_service_entry_aged,
			"dst_white_list":                                ret.DtDdosTableStats.Stats.Dst_white_list,
			"src_white_list":                                ret.DtDdosTableStats.Stats.Src_white_list,
			"src_dst_white_list":                            ret.DtDdosTableStats.Stats.Src_dst_white_list,
			"src_zone_service_white_list":                   ret.DtDdosTableStats.Stats.Src_zone_service_white_list,
			"dst_black_list":                                ret.DtDdosTableStats.Stats.Dst_black_list,
			"src_black_list":                                ret.DtDdosTableStats.Stats.Src_black_list,
			"src_dst_black_list":                            ret.DtDdosTableStats.Stats.Src_dst_black_list,
			"src_zone_service_black_list":                   ret.DtDdosTableStats.Stats.Src_zone_service_black_list,
			"dst_learning_thre_exceed":                      ret.DtDdosTableStats.Stats.Dst_learning_thre_exceed,
			"dst_over_thre_policy_at_learning":              ret.DtDdosTableStats.Stats.Dst_over_thre_policy_at_learning,
			"src_learning_thre_exceed":                      ret.DtDdosTableStats.Stats.Src_learning_thre_exceed,
			"src_over_thre_policy_at_lookup":                ret.DtDdosTableStats.Stats.Src_over_thre_policy_at_lookup,
			"src_over_thre_policy_at_learning":              ret.DtDdosTableStats.Stats.Src_over_thre_policy_at_learning,
			"src_dst_learning_thre_exceed":                  ret.DtDdosTableStats.Stats.Src_dst_learning_thre_exceed,
			"src_dst_over_thre_policy_at_lookup":            ret.DtDdosTableStats.Stats.Src_dst_over_thre_policy_at_lookup,
			"src_dst_over_thre_policy_at_learning":          ret.DtDdosTableStats.Stats.Src_dst_over_thre_policy_at_learning,
			"src_zone_service_learning_thre_exceed":         ret.DtDdosTableStats.Stats.Src_zone_service_learning_thre_exceed,
			"src_zone_service_over_thre_policy_at_lookup":   ret.DtDdosTableStats.Stats.Src_zone_service_over_thre_policy_at_lookup,
			"src_zone_service_over_thre_policy_at_learning": ret.DtDdosTableStats.Stats.Src_zone_service_over_thre_policy_at_learning,
			"entry_oom":                                     ret.DtDdosTableStats.Stats.Entry_oom,
			"entry_ext_oom":                                 ret.DtDdosTableStats.Stats.Entry_ext_oom,
			"src_dst_classlist_overflow_policy_at_learning": ret.DtDdosTableStats.Stats.Src_dst_classlist_overflow_policy_at_learning,
			"src_zone_service_classlist_overflow_policy_at_learning": ret.DtDdosTableStats.Stats.Src_zone_service_classlist_overflow_policy_at_learning,
		},
	}
}

func getObjectDdosTableStatsStats(d []interface{}) edpt.DdosTableStatsStats {

	count1 := len(d)
	var ret edpt.DdosTableStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dst_learn = in["dst_learn"].(int)
		ret.Dst_hit = in["dst_hit"].(int)
		ret.Dst_miss = in["dst_miss"].(int)
		ret.Dst_entry_aged = in["dst_entry_aged"].(int)
		ret.Src_learn = in["src_learn"].(int)
		ret.Src_hit = in["src_hit"].(int)
		ret.Src_miss = in["src_miss"].(int)
		ret.Src_entry_aged = in["src_entry_aged"].(int)
		ret.Src_dst_learn = in["src_dst_learn"].(int)
		ret.Src_dst_hit = in["src_dst_hit"].(int)
		ret.Src_dst_miss = in["src_dst_miss"].(int)
		ret.Src_dst_entry_aged = in["src_dst_entry_aged"].(int)
		ret.Telem_err_misc = in["telem_err_misc"].(int)
		ret.Telem_route_add_rcvd = in["telem_route_add_rcvd"].(int)
		ret.Telem_route_del_rcvd = in["telem_route_del_rcvd"].(int)
		ret.Telem_entry_created = in["telem_entry_created"].(int)
		ret.Telem_entry_cleared = in["telem_entry_cleared"].(int)
		ret.Telem_err_telem_entry_pre_exist = in["telem_err_telem_entry_pre_exist"].(int)
		ret.Telem_err_conflict_with_static = in["telem_err_conflict_with_static"].(int)
		ret.Telem_err_fail_to_create = in["telem_err_fail_to_create"].(int)
		ret.Telem_err_fail_to_delete = in["telem_err_fail_to_delete"].(int)
		ret.Src_zone_service_learn = in["src_zone_service_learn"].(int)
		ret.Src_zone_service_hit = in["src_zone_service_hit"].(int)
		ret.Src_zone_service_miss = in["src_zone_service_miss"].(int)
		ret.Src_zone_service_entry_aged = in["src_zone_service_entry_aged"].(int)
		ret.Dst_white_list = in["dst_white_list"].(int)
		ret.Src_white_list = in["src_white_list"].(int)
		ret.Src_dst_white_list = in["src_dst_white_list"].(int)
		ret.Src_zone_service_white_list = in["src_zone_service_white_list"].(int)
		ret.Dst_black_list = in["dst_black_list"].(int)
		ret.Src_black_list = in["src_black_list"].(int)
		ret.Src_dst_black_list = in["src_dst_black_list"].(int)
		ret.Src_zone_service_black_list = in["src_zone_service_black_list"].(int)
		ret.Dst_learning_thre_exceed = in["dst_learning_thre_exceed"].(int)
		ret.Dst_over_thre_policy_at_learning = in["dst_over_thre_policy_at_learning"].(int)
		ret.Src_learning_thre_exceed = in["src_learning_thre_exceed"].(int)
		ret.Src_over_thre_policy_at_lookup = in["src_over_thre_policy_at_lookup"].(int)
		ret.Src_over_thre_policy_at_learning = in["src_over_thre_policy_at_learning"].(int)
		ret.Src_dst_learning_thre_exceed = in["src_dst_learning_thre_exceed"].(int)
		ret.Src_dst_over_thre_policy_at_lookup = in["src_dst_over_thre_policy_at_lookup"].(int)
		ret.Src_dst_over_thre_policy_at_learning = in["src_dst_over_thre_policy_at_learning"].(int)
		ret.Src_zone_service_learning_thre_exceed = in["src_zone_service_learning_thre_exceed"].(int)
		ret.Src_zone_service_over_thre_policy_at_lookup = in["src_zone_service_over_thre_policy_at_lookup"].(int)
		ret.Src_zone_service_over_thre_policy_at_learning = in["src_zone_service_over_thre_policy_at_learning"].(int)
		ret.Entry_oom = in["entry_oom"].(int)
		ret.Entry_ext_oom = in["entry_ext_oom"].(int)
		ret.Src_dst_classlist_overflow_policy_at_learning = in["src_dst_classlist_overflow_policy_at_learning"].(int)
		ret.Src_zone_service_classlist_overflow_policy_at_learning = in["src_zone_service_classlist_overflow_policy_at_learning"].(int)
	}
	return ret
}

func dataToEndpointDdosTableStats(d *schema.ResourceData) edpt.DdosTableStats {
	var ret edpt.DdosTableStats

	ret.Stats = getObjectDdosTableStatsStats(d.Get("stats").([]interface{}))
	return ret
}
