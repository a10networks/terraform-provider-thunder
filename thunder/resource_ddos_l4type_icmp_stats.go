package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosL4typeIcmpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_l4type_icmp_stats`: Statistics for the object l4type-icmp\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosL4typeIcmpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_type0_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 1 Exceeded",
						},
						"rate_type1_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 2 Exceeded",
						},
						"rate_type2_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 3 Exceeded",
						},
						"type_deny_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dropped",
						},
						"icmpv4_rfc_undef_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv4 RFC Undef Type Dropped",
						},
						"icmpv6_rfc_undef_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 RFC Undef Type Dropped",
						},
						"wildcard_deny_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Wildcard Dropped",
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
						"type": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type",
						},
						"type_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Blacklisted",
						},
						"wildcard": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Wildcard",
						},
						"wildcard_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Wildcard Blacklisted",
						},
						"rate_type0_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 1 Dropped",
						},
						"rate_type0_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 1 Blacklisted",
						},
						"rate_type1_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 2 Dropped",
						},
						"rate_type1_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 2 Blacklisted",
						},
						"rate_type2_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 3 Dropped",
						},
						"rate_type2_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 3 Blacklisted",
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
						"port_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Rate Exceeded",
						},
						"port_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "KiBit Rate Exceeded",
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
						"frag_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Timeout",
						},
						"src_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Fragmented Packets Dropped",
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

func resourceDdosL4typeIcmpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL4typeIcmpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL4typeIcmpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosL4typeIcmpStatsStats := setObjectDdosL4typeIcmpStatsStats(res)
		d.Set("stats", DdosL4typeIcmpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosL4typeIcmpStatsStats(ret edpt.DataDdosL4typeIcmpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rate_type0_exceed":             ret.DtDdosL4typeIcmpStats.Stats.Rate_type0_exceed,
			"rate_type1_exceed":             ret.DtDdosL4typeIcmpStats.Stats.Rate_type1_exceed,
			"rate_type2_exceed":             ret.DtDdosL4typeIcmpStats.Stats.Rate_type2_exceed,
			"type_deny_drop":                ret.DtDdosL4typeIcmpStats.Stats.Type_deny_drop,
			"icmpv4_rfc_undef_drop":         ret.DtDdosL4typeIcmpStats.Stats.Icmpv4_rfc_undef_drop,
			"icmpv6_rfc_undef_drop":         ret.DtDdosL4typeIcmpStats.Stats.Icmpv6_rfc_undef_drop,
			"wildcard_deny_drop":            ret.DtDdosL4typeIcmpStats.Stats.Wildcard_deny_drop,
			"port_rcvd":                     ret.DtDdosL4typeIcmpStats.Stats.Port_rcvd,
			"port_drop":                     ret.DtDdosL4typeIcmpStats.Stats.Port_drop,
			"port_pkt_sent":                 ret.DtDdosL4typeIcmpStats.Stats.Port_pkt_sent,
			"type":                          ret.DtDdosL4typeIcmpStats.Stats.Type,
			"type_bl":                       ret.DtDdosL4typeIcmpStats.Stats.Type_bl,
			"wildcard":                      ret.DtDdosL4typeIcmpStats.Stats.Wildcard,
			"wildcard_bl":                   ret.DtDdosL4typeIcmpStats.Stats.Wildcard_bl,
			"rate_type0_exceed_drop":        ret.DtDdosL4typeIcmpStats.Stats.Rate_type0_exceed_drop,
			"rate_type0_exceed_bl":          ret.DtDdosL4typeIcmpStats.Stats.Rate_type0_exceed_bl,
			"rate_type1_exceed_drop":        ret.DtDdosL4typeIcmpStats.Stats.Rate_type1_exceed_drop,
			"rate_type1_exceed_bl":          ret.DtDdosL4typeIcmpStats.Stats.Rate_type1_exceed_bl,
			"rate_type2_exceed_drop":        ret.DtDdosL4typeIcmpStats.Stats.Rate_type2_exceed_drop,
			"rate_type2_exceed_bl":          ret.DtDdosL4typeIcmpStats.Stats.Rate_type2_exceed_bl,
			"port_bytes":                    ret.DtDdosL4typeIcmpStats.Stats.Port_bytes,
			"outbound_port_bytes":           ret.DtDdosL4typeIcmpStats.Stats.Outbound_port_bytes,
			"outbound_port_rcvd":            ret.DtDdosL4typeIcmpStats.Stats.Outbound_port_rcvd,
			"outbound_port_pkt_sent":        ret.DtDdosL4typeIcmpStats.Stats.Outbound_port_pkt_sent,
			"port_bytes_sent":               ret.DtDdosL4typeIcmpStats.Stats.Port_bytes_sent,
			"port_bytes_drop":               ret.DtDdosL4typeIcmpStats.Stats.Port_bytes_drop,
			"port_src_bl":                   ret.DtDdosL4typeIcmpStats.Stats.Port_src_bl,
			"port_pkt_rate_exceed":          ret.DtDdosL4typeIcmpStats.Stats.Port_pkt_rate_exceed,
			"port_kbit_rate_exceed":         ret.DtDdosL4typeIcmpStats.Stats.Port_kbit_rate_exceed,
			"exceed_drop_prate_src":         ret.DtDdosL4typeIcmpStats.Stats.Exceed_drop_prate_src,
			"exceed_drop_brate_src":         ret.DtDdosL4typeIcmpStats.Stats.Exceed_drop_brate_src,
			"outbound_port_bytes_sent":      ret.DtDdosL4typeIcmpStats.Stats.Outbound_port_bytes_sent,
			"outbound_port_drop":            ret.DtDdosL4typeIcmpStats.Stats.Outbound_port_drop,
			"outbound_port_bytes_drop":      ret.DtDdosL4typeIcmpStats.Stats.Outbound_port_bytes_drop,
			"exceed_drop_brate_src_pkt":     ret.DtDdosL4typeIcmpStats.Stats.Exceed_drop_brate_src_pkt,
			"port_kbit_rate_exceed_pkt":     ret.DtDdosL4typeIcmpStats.Stats.Port_kbit_rate_exceed_pkt,
			"bl":                            ret.DtDdosL4typeIcmpStats.Stats.Bl,
			"src_drop":                      ret.DtDdosL4typeIcmpStats.Stats.Src_drop,
			"frag_rcvd":                     ret.DtDdosL4typeIcmpStats.Stats.Frag_rcvd,
			"frag_drop":                     ret.DtDdosL4typeIcmpStats.Stats.Frag_drop,
			"frag_timeout":                  ret.DtDdosL4typeIcmpStats.Stats.Frag_timeout,
			"src_frag_drop":                 ret.DtDdosL4typeIcmpStats.Stats.Src_frag_drop,
			"sflow_internal_samples_packed": ret.DtDdosL4typeIcmpStats.Stats.Sflow_internal_samples_packed,
			"sflow_external_samples_packed": ret.DtDdosL4typeIcmpStats.Stats.Sflow_external_samples_packed,
			"sflow_internal_packets_sent":   ret.DtDdosL4typeIcmpStats.Stats.Sflow_internal_packets_sent,
			"sflow_external_packets_sent":   ret.DtDdosL4typeIcmpStats.Stats.Sflow_external_packets_sent,
			"exceed_action_tunnel":          ret.DtDdosL4typeIcmpStats.Stats.Exceed_action_tunnel,
			"dst_hw_drop":                   ret.DtDdosL4typeIcmpStats.Stats.Dst_hw_drop,
			"exceed_action_drop":            ret.DtDdosL4typeIcmpStats.Stats.Exceed_action_drop,
		},
	}
}

func getObjectDdosL4typeIcmpStatsStats(d []interface{}) edpt.DdosL4typeIcmpStatsStats {

	count1 := len(d)
	var ret edpt.DdosL4typeIcmpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rate_type0_exceed = in["rate_type0_exceed"].(int)
		ret.Rate_type1_exceed = in["rate_type1_exceed"].(int)
		ret.Rate_type2_exceed = in["rate_type2_exceed"].(int)
		ret.Type_deny_drop = in["type_deny_drop"].(int)
		ret.Icmpv4_rfc_undef_drop = in["icmpv4_rfc_undef_drop"].(int)
		ret.Icmpv6_rfc_undef_drop = in["icmpv6_rfc_undef_drop"].(int)
		ret.Wildcard_deny_drop = in["wildcard_deny_drop"].(int)
		ret.Port_rcvd = in["port_rcvd"].(int)
		ret.Port_drop = in["port_drop"].(int)
		ret.Port_pkt_sent = in["port_pkt_sent"].(int)
		ret.Type = in["type"].(int)
		ret.Type_bl = in["type_bl"].(int)
		ret.Wildcard = in["wildcard"].(int)
		ret.Wildcard_bl = in["wildcard_bl"].(int)
		ret.Rate_type0_exceed_drop = in["rate_type0_exceed_drop"].(int)
		ret.Rate_type0_exceed_bl = in["rate_type0_exceed_bl"].(int)
		ret.Rate_type1_exceed_drop = in["rate_type1_exceed_drop"].(int)
		ret.Rate_type1_exceed_bl = in["rate_type1_exceed_bl"].(int)
		ret.Rate_type2_exceed_drop = in["rate_type2_exceed_drop"].(int)
		ret.Rate_type2_exceed_bl = in["rate_type2_exceed_bl"].(int)
		ret.Port_bytes = in["port_bytes"].(int)
		ret.Outbound_port_bytes = in["outbound_port_bytes"].(int)
		ret.Outbound_port_rcvd = in["outbound_port_rcvd"].(int)
		ret.Outbound_port_pkt_sent = in["outbound_port_pkt_sent"].(int)
		ret.Port_bytes_sent = in["port_bytes_sent"].(int)
		ret.Port_bytes_drop = in["port_bytes_drop"].(int)
		ret.Port_src_bl = in["port_src_bl"].(int)
		ret.Port_pkt_rate_exceed = in["port_pkt_rate_exceed"].(int)
		ret.Port_kbit_rate_exceed = in["port_kbit_rate_exceed"].(int)
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
		ret.Frag_timeout = in["frag_timeout"].(int)
		ret.Src_frag_drop = in["src_frag_drop"].(int)
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

func dataToEndpointDdosL4typeIcmpStats(d *schema.ResourceData) edpt.DdosL4typeIcmpStats {
	var ret edpt.DdosL4typeIcmpStats

	ret.Stats = getObjectDdosL4typeIcmpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
