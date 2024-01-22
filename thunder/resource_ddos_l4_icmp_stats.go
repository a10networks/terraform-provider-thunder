package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosL4IcmpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_l4_icmp_stats`: Statistics for the object l4-icmp\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosL4IcmpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Src Packets Dropped",
						},
						"icmp_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Dst Packets Dropped",
						},
						"icmp_wildcard_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Wildcard Blacklisted",
						},
						"src_icmp_bl_user_config": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Src Blacklisted User Packets Dropped",
						},
						"icmp_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Packets Received",
						},
						"icmp_echo_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Echo Received",
						},
						"icmp_other_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP other Received",
						},
						"icmp_src_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP SrcDst Packets Dropped",
						},
						"icmp_src_dst_bl_drop_user_cfg": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP SrcDst Blacklisted User Packets Dropped",
						},
						"icmp_type_deny_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dropped",
						},
						"icmp_v4_rfc_undef_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv4 RFC Undef Type Dropped",
						},
						"icmp_v6_rfc_undef_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 RFC Undef Type Dropped",
						},
						"icmp_wildcard_deny_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Wildcard Dropped",
						},
						"icmp_rate_exceed0": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Rate 1 Exceeded",
						},
						"icmp_rate_exceed1": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Rate 2 Exceeded",
						},
						"icmp_rate_exceed2": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Rate 3 Exceeded",
						},
						"icmp_type": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type",
						},
						"icmp_type_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Blacklisted",
						},
						"icmp_rate_exceed0_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Rate 1 Dropped",
						},
						"icmp_rate_exceed0_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Rate 1 Blacklisted",
						},
						"icmp_rate_exceed1_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Rate 2 Dropped",
						},
						"icmp_rate_exceed1_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Rate 2 Blacklisted",
						},
						"icmp_rate_exceed2_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Rate 3 Dropped",
						},
						"icmp_rate_exceed2_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Rate 3 Blacklisted",
						},
						"icmp_wildcard": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Wildcard",
						},
						"icmp_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Exceeded",
						},
						"icmp_drop_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Blacklisted Packets Dropped",
						},
						"icmp_frag_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Frag Received",
						},
						"icmp_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Frag Dropped",
						},
						"icmp_total_bytes_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Bytes Received",
						},
						"icmp_total_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Packets Dropped",
						},
						"icmp_total_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Bytes Dropped",
						},
					},
				},
			},
		},
	}
}

func resourceDdosL4IcmpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL4IcmpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL4IcmpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosL4IcmpStatsStats := setObjectDdosL4IcmpStatsStats(res)
		d.Set("stats", DdosL4IcmpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosL4IcmpStatsStats(ret edpt.DataDdosL4IcmpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"icmp_src_drop":                 ret.DtDdosL4IcmpStats.Stats.Icmp_src_drop,
			"icmp_dst_drop":                 ret.DtDdosL4IcmpStats.Stats.Icmp_dst_drop,
			"icmp_wildcard_bl":              ret.DtDdosL4IcmpStats.Stats.Icmp_wildcard_bl,
			"src_icmp_bl_user_config":       ret.DtDdosL4IcmpStats.Stats.Src_icmp_bl_user_config,
			"icmp_rcvd":                     ret.DtDdosL4IcmpStats.Stats.Icmp_rcvd,
			"icmp_echo_rcvd":                ret.DtDdosL4IcmpStats.Stats.Icmp_echo_rcvd,
			"icmp_other_rcvd":               ret.DtDdosL4IcmpStats.Stats.Icmp_other_rcvd,
			"icmp_src_dst_drop":             ret.DtDdosL4IcmpStats.Stats.Icmp_src_dst_drop,
			"icmp_src_dst_bl_drop_user_cfg": ret.DtDdosL4IcmpStats.Stats.Icmp_src_dst_bl_drop_user_cfg,
			"icmp_type_deny_drop":           ret.DtDdosL4IcmpStats.Stats.Icmp_type_deny_drop,
			"icmp_v4_rfc_undef_drop":        ret.DtDdosL4IcmpStats.Stats.Icmp_v4_rfc_undef_drop,
			"icmp_v6_rfc_undef_drop":        ret.DtDdosL4IcmpStats.Stats.Icmp_v6_rfc_undef_drop,
			"icmp_wildcard_deny_drop":       ret.DtDdosL4IcmpStats.Stats.Icmp_wildcard_deny_drop,
			"icmp_rate_exceed0":             ret.DtDdosL4IcmpStats.Stats.Icmp_rate_exceed0,
			"icmp_rate_exceed1":             ret.DtDdosL4IcmpStats.Stats.Icmp_rate_exceed1,
			"icmp_rate_exceed2":             ret.DtDdosL4IcmpStats.Stats.Icmp_rate_exceed2,
			"icmp_type":                     ret.DtDdosL4IcmpStats.Stats.Icmp_type,
			"icmp_type_bl":                  ret.DtDdosL4IcmpStats.Stats.Icmp_type_bl,
			"icmp_rate_exceed0_drop":        ret.DtDdosL4IcmpStats.Stats.Icmp_rate_exceed0_drop,
			"icmp_rate_exceed0_bl":          ret.DtDdosL4IcmpStats.Stats.Icmp_rate_exceed0_bl,
			"icmp_rate_exceed1_drop":        ret.DtDdosL4IcmpStats.Stats.Icmp_rate_exceed1_drop,
			"icmp_rate_exceed1_bl":          ret.DtDdosL4IcmpStats.Stats.Icmp_rate_exceed1_bl,
			"icmp_rate_exceed2_drop":        ret.DtDdosL4IcmpStats.Stats.Icmp_rate_exceed2_drop,
			"icmp_rate_exceed2_bl":          ret.DtDdosL4IcmpStats.Stats.Icmp_rate_exceed2_bl,
			"icmp_wildcard":                 ret.DtDdosL4IcmpStats.Stats.Icmp_wildcard,
			"icmp_any_exceed":               ret.DtDdosL4IcmpStats.Stats.Icmp_any_exceed,
			"icmp_drop_bl":                  ret.DtDdosL4IcmpStats.Stats.Icmp_drop_bl,
			"icmp_frag_rcvd":                ret.DtDdosL4IcmpStats.Stats.Icmp_frag_rcvd,
			"icmp_frag_drop":                ret.DtDdosL4IcmpStats.Stats.Icmp_frag_drop,
			"icmp_total_bytes_rcv":          ret.DtDdosL4IcmpStats.Stats.Icmp_total_bytes_rcv,
			"icmp_total_drop":               ret.DtDdosL4IcmpStats.Stats.Icmp_total_drop,
			"icmp_total_bytes_drop":         ret.DtDdosL4IcmpStats.Stats.Icmp_total_bytes_drop,
		},
	}
}

func getObjectDdosL4IcmpStatsStats(d []interface{}) edpt.DdosL4IcmpStatsStats {

	count1 := len(d)
	var ret edpt.DdosL4IcmpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Icmp_src_drop = in["icmp_src_drop"].(int)
		ret.Icmp_dst_drop = in["icmp_dst_drop"].(int)
		ret.Icmp_wildcard_bl = in["icmp_wildcard_bl"].(int)
		ret.Src_icmp_bl_user_config = in["src_icmp_bl_user_config"].(int)
		ret.Icmp_rcvd = in["icmp_rcvd"].(int)
		ret.Icmp_echo_rcvd = in["icmp_echo_rcvd"].(int)
		ret.Icmp_other_rcvd = in["icmp_other_rcvd"].(int)
		ret.Icmp_src_dst_drop = in["icmp_src_dst_drop"].(int)
		ret.Icmp_src_dst_bl_drop_user_cfg = in["icmp_src_dst_bl_drop_user_cfg"].(int)
		ret.Icmp_type_deny_drop = in["icmp_type_deny_drop"].(int)
		ret.Icmp_v4_rfc_undef_drop = in["icmp_v4_rfc_undef_drop"].(int)
		ret.Icmp_v6_rfc_undef_drop = in["icmp_v6_rfc_undef_drop"].(int)
		ret.Icmp_wildcard_deny_drop = in["icmp_wildcard_deny_drop"].(int)
		ret.Icmp_rate_exceed0 = in["icmp_rate_exceed0"].(int)
		ret.Icmp_rate_exceed1 = in["icmp_rate_exceed1"].(int)
		ret.Icmp_rate_exceed2 = in["icmp_rate_exceed2"].(int)
		ret.Icmp_type = in["icmp_type"].(int)
		ret.Icmp_type_bl = in["icmp_type_bl"].(int)
		ret.Icmp_rate_exceed0_drop = in["icmp_rate_exceed0_drop"].(int)
		ret.Icmp_rate_exceed0_bl = in["icmp_rate_exceed0_bl"].(int)
		ret.Icmp_rate_exceed1_drop = in["icmp_rate_exceed1_drop"].(int)
		ret.Icmp_rate_exceed1_bl = in["icmp_rate_exceed1_bl"].(int)
		ret.Icmp_rate_exceed2_drop = in["icmp_rate_exceed2_drop"].(int)
		ret.Icmp_rate_exceed2_bl = in["icmp_rate_exceed2_bl"].(int)
		ret.Icmp_wildcard = in["icmp_wildcard"].(int)
		ret.Icmp_any_exceed = in["icmp_any_exceed"].(int)
		ret.Icmp_drop_bl = in["icmp_drop_bl"].(int)
		ret.Icmp_frag_rcvd = in["icmp_frag_rcvd"].(int)
		ret.Icmp_frag_drop = in["icmp_frag_drop"].(int)
		ret.Icmp_total_bytes_rcv = in["icmp_total_bytes_rcv"].(int)
		ret.Icmp_total_drop = in["icmp_total_drop"].(int)
		ret.Icmp_total_bytes_drop = in["icmp_total_bytes_drop"].(int)
	}
	return ret
}

func dataToEndpointDdosL4IcmpStats(d *schema.ResourceData) edpt.DdosL4IcmpStats {
	var ret edpt.DdosL4IcmpStats

	ret.Stats = getObjectDdosL4IcmpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
