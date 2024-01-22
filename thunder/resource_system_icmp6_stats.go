package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIcmp6Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_icmp6_stats`: Statistics for the object icmp6\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemIcmp6StatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"in_msgs": {
							Type: schema.TypeInt, Optional: true, Description: "In messages",
						},
						"in_errors": {
							Type: schema.TypeInt, Optional: true, Description: "In Errors",
						},
						"in_dest_un_reach": {
							Type: schema.TypeInt, Optional: true, Description: "In Destunation Unreachable",
						},
						"in_pkt_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "In Packet too big",
						},
						"in_time_exceeds": {
							Type: schema.TypeInt, Optional: true, Description: "In TTL Exceeds",
						},
						"in_param_prob": {
							Type: schema.TypeInt, Optional: true, Description: "In Parameter Problem",
						},
						"in_echoes": {
							Type: schema.TypeInt, Optional: true, Description: "In Echo requests",
						},
						"in_exho_reply": {
							Type: schema.TypeInt, Optional: true, Description: "In Echo replies",
						},
						"in_grp_mem_query": {
							Type: schema.TypeInt, Optional: true, Description: "In Group member query",
						},
						"in_grp_mem_resp": {
							Type: schema.TypeInt, Optional: true, Description: "In Group member reply",
						},
						"in_grp_mem_reduction": {
							Type: schema.TypeInt, Optional: true, Description: "In Group member reduction",
						},
						"in_router_sol": {
							Type: schema.TypeInt, Optional: true, Description: "In Router solicitation",
						},
						"in_ra": {
							Type: schema.TypeInt, Optional: true, Description: "In Router advertisement",
						},
						"in_ns": {
							Type: schema.TypeInt, Optional: true, Description: "In neighbor solicitation",
						},
						"in_na": {
							Type: schema.TypeInt, Optional: true, Description: "In neighbor advertisement",
						},
						"in_redirect": {
							Type: schema.TypeInt, Optional: true, Description: "In Redirects",
						},
						"out_msg": {
							Type: schema.TypeInt, Optional: true, Description: "Out Messages",
						},
						"out_dst_un_reach": {
							Type: schema.TypeInt, Optional: true, Description: "Out Destination Unreachable",
						},
						"out_pkt_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "Out Packet too big",
						},
						"out_time_exceeds": {
							Type: schema.TypeInt, Optional: true, Description: "Out TTL Exceeds",
						},
						"out_param_prob": {
							Type: schema.TypeInt, Optional: true, Description: "Out Parameter Problem",
						},
						"out_echo_req": {
							Type: schema.TypeInt, Optional: true, Description: "Out Echo requests",
						},
						"out_echo_replies": {
							Type: schema.TypeInt, Optional: true, Description: "Out Echo replies",
						},
						"out_rs": {
							Type: schema.TypeInt, Optional: true, Description: "Out Router solicitation",
						},
						"out_ra": {
							Type: schema.TypeInt, Optional: true, Description: "Out Router advertisement",
						},
						"out_ns": {
							Type: schema.TypeInt, Optional: true, Description: "Out neighbor solicitation",
						},
						"out_na": {
							Type: schema.TypeInt, Optional: true, Description: "Out neighbor advertisement",
						},
						"out_redirects": {
							Type: schema.TypeInt, Optional: true, Description: "Out Redirects",
						},
						"out_mem_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Out Group member reply",
						},
						"out_mem_reductions": {
							Type: schema.TypeInt, Optional: true, Description: "Out Group member reduction",
						},
						"err_rs": {
							Type: schema.TypeInt, Optional: true, Description: "Error Router solicitation",
						},
						"err_ra": {
							Type: schema.TypeInt, Optional: true, Description: "Error Router advertisement",
						},
						"err_ns": {
							Type: schema.TypeInt, Optional: true, Description: "Error Neighbor solicitation",
						},
						"err_na": {
							Type: schema.TypeInt, Optional: true, Description: "Error Neighbor advertisement",
						},
						"err_redirects": {
							Type: schema.TypeInt, Optional: true, Description: "Error Redirects",
						},
						"err_echoes": {
							Type: schema.TypeInt, Optional: true, Description: "Error Echo requests",
						},
						"err_echo_replies": {
							Type: schema.TypeInt, Optional: true, Description: "Error Echo replies",
						},
					},
				},
			},
		},
	}
}

func resourceSystemIcmp6StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmp6StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmp6Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemIcmp6StatsStats := setObjectSystemIcmp6StatsStats(res)
		d.Set("stats", SystemIcmp6StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemIcmp6StatsStats(ret edpt.DataSystemIcmp6Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"in_msgs":              ret.DtSystemIcmp6Stats.Stats.In_msgs,
			"in_errors":            ret.DtSystemIcmp6Stats.Stats.In_errors,
			"in_dest_un_reach":     ret.DtSystemIcmp6Stats.Stats.In_dest_un_reach,
			"in_pkt_too_big":       ret.DtSystemIcmp6Stats.Stats.In_pkt_too_big,
			"in_time_exceeds":      ret.DtSystemIcmp6Stats.Stats.In_time_exceeds,
			"in_param_prob":        ret.DtSystemIcmp6Stats.Stats.In_param_prob,
			"in_echoes":            ret.DtSystemIcmp6Stats.Stats.In_echoes,
			"in_exho_reply":        ret.DtSystemIcmp6Stats.Stats.In_exho_reply,
			"in_grp_mem_query":     ret.DtSystemIcmp6Stats.Stats.In_grp_mem_query,
			"in_grp_mem_resp":      ret.DtSystemIcmp6Stats.Stats.In_grp_mem_resp,
			"in_grp_mem_reduction": ret.DtSystemIcmp6Stats.Stats.In_grp_mem_reduction,
			"in_router_sol":        ret.DtSystemIcmp6Stats.Stats.In_router_sol,
			"in_ra":                ret.DtSystemIcmp6Stats.Stats.In_ra,
			"in_ns":                ret.DtSystemIcmp6Stats.Stats.In_ns,
			"in_na":                ret.DtSystemIcmp6Stats.Stats.In_na,
			"in_redirect":          ret.DtSystemIcmp6Stats.Stats.In_redirect,
			"out_msg":              ret.DtSystemIcmp6Stats.Stats.Out_msg,
			"out_dst_un_reach":     ret.DtSystemIcmp6Stats.Stats.Out_dst_un_reach,
			"out_pkt_too_big":      ret.DtSystemIcmp6Stats.Stats.Out_pkt_too_big,
			"out_time_exceeds":     ret.DtSystemIcmp6Stats.Stats.Out_time_exceeds,
			"out_param_prob":       ret.DtSystemIcmp6Stats.Stats.Out_param_prob,
			"out_echo_req":         ret.DtSystemIcmp6Stats.Stats.Out_echo_req,
			"out_echo_replies":     ret.DtSystemIcmp6Stats.Stats.Out_echo_replies,
			"out_rs":               ret.DtSystemIcmp6Stats.Stats.Out_rs,
			"out_ra":               ret.DtSystemIcmp6Stats.Stats.Out_ra,
			"out_ns":               ret.DtSystemIcmp6Stats.Stats.Out_ns,
			"out_na":               ret.DtSystemIcmp6Stats.Stats.Out_na,
			"out_redirects":        ret.DtSystemIcmp6Stats.Stats.Out_redirects,
			"out_mem_resp":         ret.DtSystemIcmp6Stats.Stats.Out_mem_resp,
			"out_mem_reductions":   ret.DtSystemIcmp6Stats.Stats.Out_mem_reductions,
			"err_rs":               ret.DtSystemIcmp6Stats.Stats.Err_rs,
			"err_ra":               ret.DtSystemIcmp6Stats.Stats.Err_ra,
			"err_ns":               ret.DtSystemIcmp6Stats.Stats.Err_ns,
			"err_na":               ret.DtSystemIcmp6Stats.Stats.Err_na,
			"err_redirects":        ret.DtSystemIcmp6Stats.Stats.Err_redirects,
			"err_echoes":           ret.DtSystemIcmp6Stats.Stats.Err_echoes,
			"err_echo_replies":     ret.DtSystemIcmp6Stats.Stats.Err_echo_replies,
		},
	}
}

func getObjectSystemIcmp6StatsStats(d []interface{}) edpt.SystemIcmp6StatsStats {

	count1 := len(d)
	var ret edpt.SystemIcmp6StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.In_msgs = in["in_msgs"].(int)
		ret.In_errors = in["in_errors"].(int)
		ret.In_dest_un_reach = in["in_dest_un_reach"].(int)
		ret.In_pkt_too_big = in["in_pkt_too_big"].(int)
		ret.In_time_exceeds = in["in_time_exceeds"].(int)
		ret.In_param_prob = in["in_param_prob"].(int)
		ret.In_echoes = in["in_echoes"].(int)
		ret.In_exho_reply = in["in_exho_reply"].(int)
		ret.In_grp_mem_query = in["in_grp_mem_query"].(int)
		ret.In_grp_mem_resp = in["in_grp_mem_resp"].(int)
		ret.In_grp_mem_reduction = in["in_grp_mem_reduction"].(int)
		ret.In_router_sol = in["in_router_sol"].(int)
		ret.In_ra = in["in_ra"].(int)
		ret.In_ns = in["in_ns"].(int)
		ret.In_na = in["in_na"].(int)
		ret.In_redirect = in["in_redirect"].(int)
		ret.Out_msg = in["out_msg"].(int)
		ret.Out_dst_un_reach = in["out_dst_un_reach"].(int)
		ret.Out_pkt_too_big = in["out_pkt_too_big"].(int)
		ret.Out_time_exceeds = in["out_time_exceeds"].(int)
		ret.Out_param_prob = in["out_param_prob"].(int)
		ret.Out_echo_req = in["out_echo_req"].(int)
		ret.Out_echo_replies = in["out_echo_replies"].(int)
		ret.Out_rs = in["out_rs"].(int)
		ret.Out_ra = in["out_ra"].(int)
		ret.Out_ns = in["out_ns"].(int)
		ret.Out_na = in["out_na"].(int)
		ret.Out_redirects = in["out_redirects"].(int)
		ret.Out_mem_resp = in["out_mem_resp"].(int)
		ret.Out_mem_reductions = in["out_mem_reductions"].(int)
		ret.Err_rs = in["err_rs"].(int)
		ret.Err_ra = in["err_ra"].(int)
		ret.Err_ns = in["err_ns"].(int)
		ret.Err_na = in["err_na"].(int)
		ret.Err_redirects = in["err_redirects"].(int)
		ret.Err_echoes = in["err_echoes"].(int)
		ret.Err_echo_replies = in["err_echo_replies"].(int)
	}
	return ret
}

func dataToEndpointSystemIcmp6Stats(d *schema.ResourceData) edpt.SystemIcmp6Stats {
	var ret edpt.SystemIcmp6Stats

	ret.Stats = getObjectSystemIcmp6StatsStats(d.Get("stats").([]interface{}))
	return ret
}
