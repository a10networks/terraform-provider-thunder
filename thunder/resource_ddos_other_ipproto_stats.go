package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosOtherIpprotoStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_other_ipproto_stats`: Statistics for the object other-ipproto\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosOtherIpprotoStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter1_match": {
							Type: schema.TypeInt, Optional: true, Description: "Filter1 Match",
						},
						"filter2_match": {
							Type: schema.TypeInt, Optional: true, Description: "Filter2 Match",
						},
						"filter3_match": {
							Type: schema.TypeInt, Optional: true, Description: "Filter3 Match",
						},
						"filter4_match": {
							Type: schema.TypeInt, Optional: true, Description: "Filter4 Match",
						},
						"filter5_match": {
							Type: schema.TypeInt, Optional: true, Description: "Filter5 Match",
						},
						"filter_none_match": {
							Type: schema.TypeInt, Optional: true, Description: "Filter No Match",
						},
						"port_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Packets Received",
						},
						"port_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Packets Dropped",
						},
						"port_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Packets Forwarded",
						},
						"port_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Rate Exceeded",
						},
						"port_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "KiBit Rate Exceeded",
						},
						"filter_auth_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Auth Failed",
						},
						"port_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Bytes Received",
						},
						"outbound_port_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Bytes Received",
						},
						"outbound_port_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Packets Received",
						},
						"outbound_port_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Packets Forwarded",
						},
						"port_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Bytes Forwarded",
						},
						"port_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Bytes Dropped",
						},
						"port_src_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Src Blacklisted",
						},
						"filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Action Blacklist",
						},
						"filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Action Drop",
						},
						"filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Action Default Pass",
						},
						"filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Action Whitelist",
						},
						"exceed_drop_prate_src": {
							Type: schema.TypeInt, Optional: true, Description: "Src Pkt Rate Exceeded",
						},
						"exceed_drop_brate_src": {
							Type: schema.TypeInt, Optional: true, Description: "Src KiBit Rate Exceeded",
						},
						"outbound_port_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Bytes Forwarded",
						},
						"outbound_port_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Packets Dropped",
						},
						"outbound_port_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Bytes Dropped",
						},
						"exceed_drop_brate_src_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Src KiBit Rate Exceeded Count",
						},
						"port_kbit_rate_exceed_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "KiBit Rate Exceeded Count",
						},
						"bl": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Blacklisted",
						},
						"src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Packets Dropped",
						},
						"frag_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Received",
						},
						"frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Dropped",
						},
						"filter_total_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Not Matched on Pkt",
						},
						"src_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Blacklist",
						},
						"src_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Drop",
						},
						"src_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Default Pass",
						},
						"src_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Whitelist",
						},
						"frag_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Timeout",
						},
						"src_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Fragmented Packets Dropped",
						},
						"src_filter1_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter1 Match",
						},
						"src_filter2_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter2 Match",
						},
						"src_filter3_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter3 Match",
						},
						"src_filter4_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter4 Match",
						},
						"src_filter5_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter5 Match",
						},
						"src_filter_none_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter No Match",
						},
						"src_filter_total_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Not Matched on Pkt",
						},
						"src_filter_auth_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Auth Failed",
						},
						"sflow_internal_samples_packed": {
							Type: schema.TypeInt, Optional: true, Description: "Sflow Internal Samples Packed",
						},
						"sflow_external_samples_packed": {
							Type: schema.TypeInt, Optional: true, Description: "Sflow External Samples Packed",
						},
						"sflow_internal_packets_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sflow Internal Packets Sent",
						},
						"sflow_external_packets_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sflow External Packets Sent",
						},
						"exceed_action_tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Exceed Action: Tunnel",
						},
						"dst_hw_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Hardware Packets Dropped",
						},
						"exceed_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Exceed Action: Dropped",
						},
					},
				},
			},
		},
	}
}

func resourceDdosOtherIpprotoStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOtherIpprotoStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOtherIpprotoStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosOtherIpprotoStatsStats := setObjectDdosOtherIpprotoStatsStats(res)
		d.Set("stats", DdosOtherIpprotoStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosOtherIpprotoStatsStats(ret edpt.DataDdosOtherIpprotoStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"filter1_match":                  ret.DtDdosOtherIpprotoStats.Stats.Filter1_match,
			"filter2_match":                  ret.DtDdosOtherIpprotoStats.Stats.Filter2_match,
			"filter3_match":                  ret.DtDdosOtherIpprotoStats.Stats.Filter3_match,
			"filter4_match":                  ret.DtDdosOtherIpprotoStats.Stats.Filter4_match,
			"filter5_match":                  ret.DtDdosOtherIpprotoStats.Stats.Filter5_match,
			"filter_none_match":              ret.DtDdosOtherIpprotoStats.Stats.Filter_none_match,
			"port_rcvd":                      ret.DtDdosOtherIpprotoStats.Stats.Port_rcvd,
			"port_drop":                      ret.DtDdosOtherIpprotoStats.Stats.Port_drop,
			"port_pkt_sent":                  ret.DtDdosOtherIpprotoStats.Stats.Port_pkt_sent,
			"port_pkt_rate_exceed":           ret.DtDdosOtherIpprotoStats.Stats.Port_pkt_rate_exceed,
			"port_kbit_rate_exceed":          ret.DtDdosOtherIpprotoStats.Stats.Port_kbit_rate_exceed,
			"filter_auth_fail":               ret.DtDdosOtherIpprotoStats.Stats.Filter_auth_fail,
			"port_bytes":                     ret.DtDdosOtherIpprotoStats.Stats.Port_bytes,
			"outbound_port_bytes":            ret.DtDdosOtherIpprotoStats.Stats.Outbound_port_bytes,
			"outbound_port_rcvd":             ret.DtDdosOtherIpprotoStats.Stats.Outbound_port_rcvd,
			"outbound_port_pkt_sent":         ret.DtDdosOtherIpprotoStats.Stats.Outbound_port_pkt_sent,
			"port_bytes_sent":                ret.DtDdosOtherIpprotoStats.Stats.Port_bytes_sent,
			"port_bytes_drop":                ret.DtDdosOtherIpprotoStats.Stats.Port_bytes_drop,
			"port_src_bl":                    ret.DtDdosOtherIpprotoStats.Stats.Port_src_bl,
			"filter_action_blacklist":        ret.DtDdosOtherIpprotoStats.Stats.Filter_action_blacklist,
			"filter_action_drop":             ret.DtDdosOtherIpprotoStats.Stats.Filter_action_drop,
			"filter_action_default_pass":     ret.DtDdosOtherIpprotoStats.Stats.Filter_action_default_pass,
			"filter_action_whitelist":        ret.DtDdosOtherIpprotoStats.Stats.Filter_action_whitelist,
			"exceed_drop_prate_src":          ret.DtDdosOtherIpprotoStats.Stats.Exceed_drop_prate_src,
			"exceed_drop_brate_src":          ret.DtDdosOtherIpprotoStats.Stats.Exceed_drop_brate_src,
			"outbound_port_bytes_sent":       ret.DtDdosOtherIpprotoStats.Stats.Outbound_port_bytes_sent,
			"outbound_port_drop":             ret.DtDdosOtherIpprotoStats.Stats.Outbound_port_drop,
			"outbound_port_bytes_drop":       ret.DtDdosOtherIpprotoStats.Stats.Outbound_port_bytes_drop,
			"exceed_drop_brate_src_pkt":      ret.DtDdosOtherIpprotoStats.Stats.Exceed_drop_brate_src_pkt,
			"port_kbit_rate_exceed_pkt":      ret.DtDdosOtherIpprotoStats.Stats.Port_kbit_rate_exceed_pkt,
			"bl":                             ret.DtDdosOtherIpprotoStats.Stats.Bl,
			"src_drop":                       ret.DtDdosOtherIpprotoStats.Stats.Src_drop,
			"frag_rcvd":                      ret.DtDdosOtherIpprotoStats.Stats.Frag_rcvd,
			"frag_drop":                      ret.DtDdosOtherIpprotoStats.Stats.Frag_drop,
			"filter_total_not_match":         ret.DtDdosOtherIpprotoStats.Stats.Filter_total_not_match,
			"src_filter_action_blacklist":    ret.DtDdosOtherIpprotoStats.Stats.Src_filter_action_blacklist,
			"src_filter_action_drop":         ret.DtDdosOtherIpprotoStats.Stats.Src_filter_action_drop,
			"src_filter_action_default_pass": ret.DtDdosOtherIpprotoStats.Stats.Src_filter_action_default_pass,
			"src_filter_action_whitelist":    ret.DtDdosOtherIpprotoStats.Stats.Src_filter_action_whitelist,
			"frag_timeout":                   ret.DtDdosOtherIpprotoStats.Stats.Frag_timeout,
			"src_frag_drop":                  ret.DtDdosOtherIpprotoStats.Stats.Src_frag_drop,
			"src_filter1_match":              ret.DtDdosOtherIpprotoStats.Stats.Src_filter1_match,
			"src_filter2_match":              ret.DtDdosOtherIpprotoStats.Stats.Src_filter2_match,
			"src_filter3_match":              ret.DtDdosOtherIpprotoStats.Stats.Src_filter3_match,
			"src_filter4_match":              ret.DtDdosOtherIpprotoStats.Stats.Src_filter4_match,
			"src_filter5_match":              ret.DtDdosOtherIpprotoStats.Stats.Src_filter5_match,
			"src_filter_none_match":          ret.DtDdosOtherIpprotoStats.Stats.Src_filter_none_match,
			"src_filter_total_not_match":     ret.DtDdosOtherIpprotoStats.Stats.Src_filter_total_not_match,
			"src_filter_auth_fail":           ret.DtDdosOtherIpprotoStats.Stats.Src_filter_auth_fail,
			"sflow_internal_samples_packed":  ret.DtDdosOtherIpprotoStats.Stats.Sflow_internal_samples_packed,
			"sflow_external_samples_packed":  ret.DtDdosOtherIpprotoStats.Stats.Sflow_external_samples_packed,
			"sflow_internal_packets_sent":    ret.DtDdosOtherIpprotoStats.Stats.Sflow_internal_packets_sent,
			"sflow_external_packets_sent":    ret.DtDdosOtherIpprotoStats.Stats.Sflow_external_packets_sent,
			"exceed_action_tunnel":           ret.DtDdosOtherIpprotoStats.Stats.Exceed_action_tunnel,
			"dst_hw_drop":                    ret.DtDdosOtherIpprotoStats.Stats.Dst_hw_drop,
			"exceed_action_drop":             ret.DtDdosOtherIpprotoStats.Stats.Exceed_action_drop,
		},
	}
}

func getObjectDdosOtherIpprotoStatsStats(d []interface{}) edpt.DdosOtherIpprotoStatsStats {

	count1 := len(d)
	var ret edpt.DdosOtherIpprotoStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Filter1_match = in["filter1_match"].(int)
		ret.Filter2_match = in["filter2_match"].(int)
		ret.Filter3_match = in["filter3_match"].(int)
		ret.Filter4_match = in["filter4_match"].(int)
		ret.Filter5_match = in["filter5_match"].(int)
		ret.Filter_none_match = in["filter_none_match"].(int)
		ret.Port_rcvd = in["port_rcvd"].(int)
		ret.Port_drop = in["port_drop"].(int)
		ret.Port_pkt_sent = in["port_pkt_sent"].(int)
		ret.Port_pkt_rate_exceed = in["port_pkt_rate_exceed"].(int)
		ret.Port_kbit_rate_exceed = in["port_kbit_rate_exceed"].(int)
		ret.Filter_auth_fail = in["filter_auth_fail"].(int)
		ret.Port_bytes = in["port_bytes"].(int)
		ret.Outbound_port_bytes = in["outbound_port_bytes"].(int)
		ret.Outbound_port_rcvd = in["outbound_port_rcvd"].(int)
		ret.Outbound_port_pkt_sent = in["outbound_port_pkt_sent"].(int)
		ret.Port_bytes_sent = in["port_bytes_sent"].(int)
		ret.Port_bytes_drop = in["port_bytes_drop"].(int)
		ret.Port_src_bl = in["port_src_bl"].(int)
		ret.Filter_action_blacklist = in["filter_action_blacklist"].(int)
		ret.Filter_action_drop = in["filter_action_drop"].(int)
		ret.Filter_action_default_pass = in["filter_action_default_pass"].(int)
		ret.Filter_action_whitelist = in["filter_action_whitelist"].(int)
		ret.Exceed_drop_prate_src = in["exceed_drop_prate_src"].(int)
		ret.Exceed_drop_brate_src = in["exceed_drop_brate_src"].(int)
		ret.Outbound_port_bytes_sent = in["outbound_port_bytes_sent"].(int)
		ret.Outbound_port_drop = in["outbound_port_drop"].(int)
		ret.Outbound_port_bytes_drop = in["outbound_port_bytes_drop"].(int)
		ret.Exceed_drop_brate_src_pkt = in["exceed_drop_brate_src_pkt"].(int)
		ret.Port_kbit_rate_exceed_pkt = in["port_kbit_rate_exceed_pkt"].(int)
		ret.Bl = in["bl"].(int)
		ret.Src_drop = in["src_drop"].(int)
		ret.Frag_rcvd = in["frag_rcvd"].(int)
		ret.Frag_drop = in["frag_drop"].(int)
		ret.Filter_total_not_match = in["filter_total_not_match"].(int)
		ret.Src_filter_action_blacklist = in["src_filter_action_blacklist"].(int)
		ret.Src_filter_action_drop = in["src_filter_action_drop"].(int)
		ret.Src_filter_action_default_pass = in["src_filter_action_default_pass"].(int)
		ret.Src_filter_action_whitelist = in["src_filter_action_whitelist"].(int)
		ret.Frag_timeout = in["frag_timeout"].(int)
		ret.Src_frag_drop = in["src_frag_drop"].(int)
		ret.Src_filter1_match = in["src_filter1_match"].(int)
		ret.Src_filter2_match = in["src_filter2_match"].(int)
		ret.Src_filter3_match = in["src_filter3_match"].(int)
		ret.Src_filter4_match = in["src_filter4_match"].(int)
		ret.Src_filter5_match = in["src_filter5_match"].(int)
		ret.Src_filter_none_match = in["src_filter_none_match"].(int)
		ret.Src_filter_total_not_match = in["src_filter_total_not_match"].(int)
		ret.Src_filter_auth_fail = in["src_filter_auth_fail"].(int)
		ret.Sflow_internal_samples_packed = in["sflow_internal_samples_packed"].(int)
		ret.Sflow_external_samples_packed = in["sflow_external_samples_packed"].(int)
		ret.Sflow_internal_packets_sent = in["sflow_internal_packets_sent"].(int)
		ret.Sflow_external_packets_sent = in["sflow_external_packets_sent"].(int)
		ret.Exceed_action_tunnel = in["exceed_action_tunnel"].(int)
		ret.Dst_hw_drop = in["dst_hw_drop"].(int)
		ret.Exceed_action_drop = in["exceed_action_drop"].(int)
	}
	return ret
}

func dataToEndpointDdosOtherIpprotoStats(d *schema.ResourceData) edpt.DdosOtherIpprotoStats {
	var ret edpt.DdosOtherIpprotoStats

	ret.Stats = getObjectDdosOtherIpprotoStatsStats(d.Get("stats").([]interface{}))
	return ret
}
