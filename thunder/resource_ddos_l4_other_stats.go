package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosL4OtherStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_l4_other_stats`: Statistics for the object l4-other\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosL4OtherStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"other_receive": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Packets Received",
						},
						"other_frag_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Frag Received",
						},
						"other_total_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Packets Dropped",
						},
						"other_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Dst Packets Dropped",
						},
						"other_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Frag Dropped",
						},
						"other_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Src Packets Dropped",
						},
						"other_drop_black_user_cfg_src": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Src Blacklist User Packets Dropped",
						},
						"other_src_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER SrcDst Packets Dropped",
						},
						"other_drop_black_user_cfg_src_dst": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER SrcDst Blacklist User Packets Dropped",
						},
						"dst_other_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Match",
						},
						"dst_other_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter No Match",
						},
						"dst_other_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Blacklist",
						},
						"dst_other_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Drop",
						},
						"dst_other_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Default Pass",
						},
						"dst_other_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action WL",
						},
						"src_other_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Match",
						},
						"src_other_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter No Match",
						},
						"src_other_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Blacklist",
						},
						"src_other_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Drop",
						},
						"src_other_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Default Pass",
						},
						"src_other_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action WL",
						},
						"src_dst_other_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Match",
						},
						"src_dst_other_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter No Match",
						},
						"src_dst_other_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action Blacklist",
						},
						"src_dst_other_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action Drop",
						},
						"src_dst_other_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action Default Pass",
						},
						"src_dst_other_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action WL",
						},
						"other_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Exceeded",
						},
						"other_drop_bl": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Blacklisted Packets Dropped",
						},
						"other_total_bytes_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Bytes Received",
						},
						"other_total_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Bytes Dropped",
						},
					},
				},
			},
		},
	}
}

func resourceDdosL4OtherStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL4OtherStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL4OtherStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosL4OtherStatsStats := setObjectDdosL4OtherStatsStats(res)
		d.Set("stats", DdosL4OtherStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosL4OtherStatsStats(ret edpt.DataDdosL4OtherStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"other_receive":                            ret.DtDdosL4OtherStats.Stats.Other_receive,
			"other_frag_rcvd":                          ret.DtDdosL4OtherStats.Stats.Other_frag_rcvd,
			"other_total_drop":                         ret.DtDdosL4OtherStats.Stats.Other_total_drop,
			"other_dst_drop":                           ret.DtDdosL4OtherStats.Stats.Other_dst_drop,
			"other_frag_drop":                          ret.DtDdosL4OtherStats.Stats.Other_frag_drop,
			"other_src_drop":                           ret.DtDdosL4OtherStats.Stats.Other_src_drop,
			"other_drop_black_user_cfg_src":            ret.DtDdosL4OtherStats.Stats.Other_drop_black_user_cfg_src,
			"other_src_dst_drop":                       ret.DtDdosL4OtherStats.Stats.Other_src_dst_drop,
			"other_drop_black_user_cfg_src_dst":        ret.DtDdosL4OtherStats.Stats.Other_drop_black_user_cfg_src_dst,
			"dst_other_filter_match":                   ret.DtDdosL4OtherStats.Stats.Dst_other_filter_match,
			"dst_other_filter_not_match":               ret.DtDdosL4OtherStats.Stats.Dst_other_filter_not_match,
			"dst_other_filter_action_blacklist":        ret.DtDdosL4OtherStats.Stats.Dst_other_filter_action_blacklist,
			"dst_other_filter_action_drop":             ret.DtDdosL4OtherStats.Stats.Dst_other_filter_action_drop,
			"dst_other_filter_action_default_pass":     ret.DtDdosL4OtherStats.Stats.Dst_other_filter_action_default_pass,
			"dst_other_filter_action_whitelist":        ret.DtDdosL4OtherStats.Stats.Dst_other_filter_action_whitelist,
			"src_other_filter_match":                   ret.DtDdosL4OtherStats.Stats.Src_other_filter_match,
			"src_other_filter_not_match":               ret.DtDdosL4OtherStats.Stats.Src_other_filter_not_match,
			"src_other_filter_action_blacklist":        ret.DtDdosL4OtherStats.Stats.Src_other_filter_action_blacklist,
			"src_other_filter_action_drop":             ret.DtDdosL4OtherStats.Stats.Src_other_filter_action_drop,
			"src_other_filter_action_default_pass":     ret.DtDdosL4OtherStats.Stats.Src_other_filter_action_default_pass,
			"src_other_filter_action_whitelist":        ret.DtDdosL4OtherStats.Stats.Src_other_filter_action_whitelist,
			"src_dst_other_filter_match":               ret.DtDdosL4OtherStats.Stats.Src_dst_other_filter_match,
			"src_dst_other_filter_not_match":           ret.DtDdosL4OtherStats.Stats.Src_dst_other_filter_not_match,
			"src_dst_other_filter_action_blacklist":    ret.DtDdosL4OtherStats.Stats.Src_dst_other_filter_action_blacklist,
			"src_dst_other_filter_action_drop":         ret.DtDdosL4OtherStats.Stats.Src_dst_other_filter_action_drop,
			"src_dst_other_filter_action_default_pass": ret.DtDdosL4OtherStats.Stats.Src_dst_other_filter_action_default_pass,
			"src_dst_other_filter_action_whitelist":    ret.DtDdosL4OtherStats.Stats.Src_dst_other_filter_action_whitelist,
			"other_any_exceed":                         ret.DtDdosL4OtherStats.Stats.Other_any_exceed,
			"other_drop_bl":                            ret.DtDdosL4OtherStats.Stats.Other_drop_bl,
			"other_total_bytes_rcv":                    ret.DtDdosL4OtherStats.Stats.Other_total_bytes_rcv,
			"other_total_bytes_drop":                   ret.DtDdosL4OtherStats.Stats.Other_total_bytes_drop,
		},
	}
}

func getObjectDdosL4OtherStatsStats(d []interface{}) edpt.DdosL4OtherStatsStats {

	count1 := len(d)
	var ret edpt.DdosL4OtherStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Other_receive = in["other_receive"].(int)
		ret.Other_frag_rcvd = in["other_frag_rcvd"].(int)
		ret.Other_total_drop = in["other_total_drop"].(int)
		ret.Other_dst_drop = in["other_dst_drop"].(int)
		ret.Other_frag_drop = in["other_frag_drop"].(int)
		ret.Other_src_drop = in["other_src_drop"].(int)
		ret.Other_drop_black_user_cfg_src = in["other_drop_black_user_cfg_src"].(int)
		ret.Other_src_dst_drop = in["other_src_dst_drop"].(int)
		ret.Other_drop_black_user_cfg_src_dst = in["other_drop_black_user_cfg_src_dst"].(int)
		ret.Dst_other_filter_match = in["dst_other_filter_match"].(int)
		ret.Dst_other_filter_not_match = in["dst_other_filter_not_match"].(int)
		ret.Dst_other_filter_action_blacklist = in["dst_other_filter_action_blacklist"].(int)
		ret.Dst_other_filter_action_drop = in["dst_other_filter_action_drop"].(int)
		ret.Dst_other_filter_action_default_pass = in["dst_other_filter_action_default_pass"].(int)
		ret.Dst_other_filter_action_whitelist = in["dst_other_filter_action_whitelist"].(int)
		ret.Src_other_filter_match = in["src_other_filter_match"].(int)
		ret.Src_other_filter_not_match = in["src_other_filter_not_match"].(int)
		ret.Src_other_filter_action_blacklist = in["src_other_filter_action_blacklist"].(int)
		ret.Src_other_filter_action_drop = in["src_other_filter_action_drop"].(int)
		ret.Src_other_filter_action_default_pass = in["src_other_filter_action_default_pass"].(int)
		ret.Src_other_filter_action_whitelist = in["src_other_filter_action_whitelist"].(int)
		ret.Src_dst_other_filter_match = in["src_dst_other_filter_match"].(int)
		ret.Src_dst_other_filter_not_match = in["src_dst_other_filter_not_match"].(int)
		ret.Src_dst_other_filter_action_blacklist = in["src_dst_other_filter_action_blacklist"].(int)
		ret.Src_dst_other_filter_action_drop = in["src_dst_other_filter_action_drop"].(int)
		ret.Src_dst_other_filter_action_default_pass = in["src_dst_other_filter_action_default_pass"].(int)
		ret.Src_dst_other_filter_action_whitelist = in["src_dst_other_filter_action_whitelist"].(int)
		ret.Other_any_exceed = in["other_any_exceed"].(int)
		ret.Other_drop_bl = in["other_drop_bl"].(int)
		ret.Other_total_bytes_rcv = in["other_total_bytes_rcv"].(int)
		ret.Other_total_bytes_drop = in["other_total_bytes_drop"].(int)
	}
	return ret
}

func dataToEndpointDdosL4OtherStats(d *schema.ResourceData) edpt.DdosL4OtherStats {
	var ret edpt.DdosL4OtherStats

	ret.Stats = getObjectDdosL4OtherStatsStats(d.Get("stats").([]interface{}))
	return ret
}
