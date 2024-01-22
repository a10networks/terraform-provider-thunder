package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSwitchStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_switch_stats`: Statistics for the object switch\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosSwitchStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Received",
						},
						"ip_sent": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Sent",
						},
						"ipv6_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Received",
						},
						"ipv6_sent": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Sent",
						},
						"in_stateless_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Stateless Packets Received",
						},
						"noroute": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4/v6 Out No Route",
						},
						"pkt_not_for_ddos": {
							Type: schema.TypeInt, Optional: true, Description: "Not For DDOS",
						},
						"ingress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Bytes Received",
						},
						"egress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Bytes Received",
						},
						"ingress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Packets Received",
						},
						"egress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Packets Received",
						},
						"src_ip_bypass": {
							Type: schema.TypeInt, Optional: true, Description: "Src IP Bypass",
						},
						"dst_ip_bypass": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP Bypass",
						},
						"dst_blackhole_inject": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Blackhole Injected",
						},
						"dst_blackhole_withdraw": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Blackhole Withdrawn",
						},
						"mpls_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "MPLS Received",
						},
						"mpls_drop": {
							Type: schema.TypeInt, Optional: true, Description: "MPLS Dropped",
						},
						"mpls_malformed": {
							Type: schema.TypeInt, Optional: true, Description: "MPLS Malformed",
						},
						"jumbo_frag_drop_by_filter": {
							Type: schema.TypeInt, Optional: true, Description: "Jumbo Fragment Filter Miss Drop",
						},
						"jumbo_frag_drop_before_slb": {
							Type: schema.TypeInt, Optional: true, Description: "Jumbo Fragment Non Data Plane Drop",
						},
						"jumbo_outgoing_mtu_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Jumbo Outgoing MTU Exceed Drop",
						},
					},
				},
			},
		},
	}
}

func resourceDdosSwitchStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSwitchStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSwitchStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosSwitchStatsStats := setObjectDdosSwitchStatsStats(res)
		d.Set("stats", DdosSwitchStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosSwitchStatsStats(ret edpt.DataDdosSwitchStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ip_rcvd":                        ret.DtDdosSwitchStats.Stats.Ip_rcvd,
			"ip_sent":                        ret.DtDdosSwitchStats.Stats.Ip_sent,
			"ipv6_rcvd":                      ret.DtDdosSwitchStats.Stats.Ipv6_rcvd,
			"ipv6_sent":                      ret.DtDdosSwitchStats.Stats.Ipv6_sent,
			"in_stateless_pkt":               ret.DtDdosSwitchStats.Stats.In_stateless_pkt,
			"noroute":                        ret.DtDdosSwitchStats.Stats.Noroute,
			"pkt_not_for_ddos":               ret.DtDdosSwitchStats.Stats.Pkt_not_for_ddos,
			"ingress_bytes":                  ret.DtDdosSwitchStats.Stats.Ingress_bytes,
			"egress_bytes":                   ret.DtDdosSwitchStats.Stats.Egress_bytes,
			"ingress_packets":                ret.DtDdosSwitchStats.Stats.Ingress_packets,
			"egress_packets":                 ret.DtDdosSwitchStats.Stats.Egress_packets,
			"src_ip_bypass":                  ret.DtDdosSwitchStats.Stats.Src_ip_bypass,
			"dst_ip_bypass":                  ret.DtDdosSwitchStats.Stats.Dst_ip_bypass,
			"dst_blackhole_inject":           ret.DtDdosSwitchStats.Stats.Dst_blackhole_inject,
			"dst_blackhole_withdraw":         ret.DtDdosSwitchStats.Stats.Dst_blackhole_withdraw,
			"mpls_rcvd":                      ret.DtDdosSwitchStats.Stats.Mpls_rcvd,
			"mpls_drop":                      ret.DtDdosSwitchStats.Stats.Mpls_drop,
			"mpls_malformed":                 ret.DtDdosSwitchStats.Stats.Mpls_malformed,
			"jumbo_frag_drop_by_filter":      ret.DtDdosSwitchStats.Stats.Jumbo_frag_drop_by_filter,
			"jumbo_frag_drop_before_slb":     ret.DtDdosSwitchStats.Stats.Jumbo_frag_drop_before_slb,
			"jumbo_outgoing_mtu_exceed_drop": ret.DtDdosSwitchStats.Stats.Jumbo_outgoing_mtu_exceed_drop,
		},
	}
}

func getObjectDdosSwitchStatsStats(d []interface{}) edpt.DdosSwitchStatsStats {

	count1 := len(d)
	var ret edpt.DdosSwitchStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ip_rcvd = in["ip_rcvd"].(int)
		ret.Ip_sent = in["ip_sent"].(int)
		ret.Ipv6_rcvd = in["ipv6_rcvd"].(int)
		ret.Ipv6_sent = in["ipv6_sent"].(int)
		ret.In_stateless_pkt = in["in_stateless_pkt"].(int)
		ret.Noroute = in["noroute"].(int)
		ret.Pkt_not_for_ddos = in["pkt_not_for_ddos"].(int)
		ret.Ingress_bytes = in["ingress_bytes"].(int)
		ret.Egress_bytes = in["egress_bytes"].(int)
		ret.Ingress_packets = in["ingress_packets"].(int)
		ret.Egress_packets = in["egress_packets"].(int)
		ret.Src_ip_bypass = in["src_ip_bypass"].(int)
		ret.Dst_ip_bypass = in["dst_ip_bypass"].(int)
		ret.Dst_blackhole_inject = in["dst_blackhole_inject"].(int)
		ret.Dst_blackhole_withdraw = in["dst_blackhole_withdraw"].(int)
		ret.Mpls_rcvd = in["mpls_rcvd"].(int)
		ret.Mpls_drop = in["mpls_drop"].(int)
		ret.Mpls_malformed = in["mpls_malformed"].(int)
		ret.Jumbo_frag_drop_by_filter = in["jumbo_frag_drop_by_filter"].(int)
		ret.Jumbo_frag_drop_before_slb = in["jumbo_frag_drop_before_slb"].(int)
		ret.Jumbo_outgoing_mtu_exceed_drop = in["jumbo_outgoing_mtu_exceed_drop"].(int)
	}
	return ret
}

func dataToEndpointDdosSwitchStats(d *schema.ResourceData) edpt.DdosSwitchStats {
	var ret edpt.DdosSwitchStats

	ret.Stats = getObjectDdosSwitchStatsStats(d.Get("stats").([]interface{}))
	return ret
}
