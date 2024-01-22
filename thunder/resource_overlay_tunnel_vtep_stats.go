package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_overlay_tunnel_vtep_stats`: Statistics for the object vtep\n\n__PLACEHOLDER__",
		ReadContext: resourceOverlayTunnelVtepStatsRead,

		Schema: map[string]*schema.Schema{
			"id1": {
				Type: schema.TypeInt, Required: true, Description: "VTEP Identifier",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cfg_err_count": {
							Type: schema.TypeInt, Optional: true, Description: "Config errors",
						},
						"flooded_pkt_count": {
							Type: schema.TypeInt, Optional: true, Description: "Flooded packet count",
						},
						"encap_unresolved_count": {
							Type: schema.TypeInt, Optional: true, Description: "Encap unresolved failures",
						},
						"unknown_encap_rx_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Encap miss rx pkts",
						},
						"unknown_encap_tx_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Encap miss tx pkts",
						},
						"arp_req_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Arp request sent",
						},
						"vtep_host_learned": {
							Type: schema.TypeInt, Optional: true, Description: "Hosts learned",
						},
						"vtep_host_learn_error": {
							Type: schema.TypeInt, Optional: true, Description: "Host learn error",
						},
						"invalid_lif_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Lif pkts in",
						},
						"invalid_lif_tx": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Lif pkts out",
						},
						"unknown_vtep_tx": {
							Type: schema.TypeInt, Optional: true, Description: "Vtep unknown tx",
						},
						"unknown_vtep_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Vtep Unkown rx",
						},
						"unhandled_pkt_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Unhandled packets in",
						},
						"unhandled_pkt_tx": {
							Type: schema.TypeInt, Optional: true, Description: "Unhandled packets out",
						},
						"total_pkts_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Total packets out",
						},
						"total_bytes_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Total packet bytes in",
						},
						"unicast_pkt_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Total unicast packets in",
						},
						"bcast_pkt_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Total broadcast packets in",
						},
						"mcast_pkt_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Total multicast packets in",
						},
						"dropped_pkt_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Dropped received packets",
						},
						"encap_miss_pkts_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Encap missed in received packets",
						},
						"bad_chksum_pks_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Bad checksum in received packets",
						},
						"requeue_pkts_in": {
							Type: schema.TypeInt, Optional: true, Description: "Requeued packets in",
						},
						"pkts_out": {
							Type: schema.TypeInt, Optional: true, Description: "Packets out",
						},
						"total_bytes_tx": {
							Type: schema.TypeInt, Optional: true, Description: "Packet bytes out",
						},
						"unicast_pkt_tx": {
							Type: schema.TypeInt, Optional: true, Description: "Unicast packets out",
						},
						"bcast_pkt_tx": {
							Type: schema.TypeInt, Optional: true, Description: "Broadcast packets out",
						},
						"mcast_pkt_tx": {
							Type: schema.TypeInt, Optional: true, Description: "Multicast packets out",
						},
						"dropped_pkts_tx": {
							Type: schema.TypeInt, Optional: true, Description: "Dropped packets out",
						},
						"large_pkts_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Too large packets in",
						},
						"dot1q_pkts_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Dot1q packets in",
						},
						"frag_pkts_tx": {
							Type: schema.TypeInt, Optional: true, Description: "Frag packets out",
						},
						"reassembled_pkts_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Reassembled packets in",
						},
						"bad_inner_ipv4_len_rx": {
							Type: schema.TypeInt, Optional: true, Description: "bad inner ipv4 packet len",
						},
						"bad_inner_ipv6_len_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Bad inner ipv6 packet len",
						},
						"frag_drop_pkts_tx": {
							Type: schema.TypeInt, Optional: true, Description: "Frag dropped packets out",
						},
						"lif_un_init_rx": {
							Type: schema.TypeInt, Optional: true, Description: "Lif uninitialized packets in",
						},
					},
				},
			},
		},
	}
}

func resourceOverlayTunnelVtepStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		OverlayTunnelVtepStatsStats := setObjectOverlayTunnelVtepStatsStats(res)
		d.Set("stats", OverlayTunnelVtepStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectOverlayTunnelVtepStatsStats(ret edpt.DataOverlayTunnelVtepStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cfg_err_count":          ret.DtOverlayTunnelVtepStats.Stats.Cfg_err_count,
			"flooded_pkt_count":      ret.DtOverlayTunnelVtepStats.Stats.Flooded_pkt_count,
			"encap_unresolved_count": ret.DtOverlayTunnelVtepStats.Stats.Encap_unresolved_count,
			"unknown_encap_rx_pkt":   ret.DtOverlayTunnelVtepStats.Stats.Unknown_encap_rx_pkt,
			"unknown_encap_tx_pkt":   ret.DtOverlayTunnelVtepStats.Stats.Unknown_encap_tx_pkt,
			"arp_req_sent":           ret.DtOverlayTunnelVtepStats.Stats.Arp_req_sent,
			"vtep_host_learned":      ret.DtOverlayTunnelVtepStats.Stats.Vtep_host_learned,
			"vtep_host_learn_error":  ret.DtOverlayTunnelVtepStats.Stats.Vtep_host_learn_error,
			"invalid_lif_rx":         ret.DtOverlayTunnelVtepStats.Stats.Invalid_lif_rx,
			"invalid_lif_tx":         ret.DtOverlayTunnelVtepStats.Stats.Invalid_lif_tx,
			"unknown_vtep_tx":        ret.DtOverlayTunnelVtepStats.Stats.Unknown_vtep_tx,
			"unknown_vtep_rx":        ret.DtOverlayTunnelVtepStats.Stats.Unknown_vtep_rx,
			"unhandled_pkt_rx":       ret.DtOverlayTunnelVtepStats.Stats.Unhandled_pkt_rx,
			"unhandled_pkt_tx":       ret.DtOverlayTunnelVtepStats.Stats.Unhandled_pkt_tx,
			"total_pkts_rx":          ret.DtOverlayTunnelVtepStats.Stats.Total_pkts_rx,
			"total_bytes_rx":         ret.DtOverlayTunnelVtepStats.Stats.Total_bytes_rx,
			"unicast_pkt_rx":         ret.DtOverlayTunnelVtepStats.Stats.Unicast_pkt_rx,
			"bcast_pkt_rx":           ret.DtOverlayTunnelVtepStats.Stats.Bcast_pkt_rx,
			"mcast_pkt_rx":           ret.DtOverlayTunnelVtepStats.Stats.Mcast_pkt_rx,
			"dropped_pkt_rx":         ret.DtOverlayTunnelVtepStats.Stats.Dropped_pkt_rx,
			"encap_miss_pkts_rx":     ret.DtOverlayTunnelVtepStats.Stats.Encap_miss_pkts_rx,
			"bad_chksum_pks_rx":      ret.DtOverlayTunnelVtepStats.Stats.Bad_chksum_pks_rx,
			"requeue_pkts_in":        ret.DtOverlayTunnelVtepStats.Stats.Requeue_pkts_in,
			"pkts_out":               ret.DtOverlayTunnelVtepStats.Stats.Pkts_out,
			"total_bytes_tx":         ret.DtOverlayTunnelVtepStats.Stats.Total_bytes_tx,
			"unicast_pkt_tx":         ret.DtOverlayTunnelVtepStats.Stats.Unicast_pkt_tx,
			"bcast_pkt_tx":           ret.DtOverlayTunnelVtepStats.Stats.Bcast_pkt_tx,
			"mcast_pkt_tx":           ret.DtOverlayTunnelVtepStats.Stats.Mcast_pkt_tx,
			"dropped_pkts_tx":        ret.DtOverlayTunnelVtepStats.Stats.Dropped_pkts_tx,
			"large_pkts_rx":          ret.DtOverlayTunnelVtepStats.Stats.Large_pkts_rx,
			"dot1q_pkts_rx":          ret.DtOverlayTunnelVtepStats.Stats.Dot1q_pkts_rx,
			"frag_pkts_tx":           ret.DtOverlayTunnelVtepStats.Stats.Frag_pkts_tx,
			"reassembled_pkts_rx":    ret.DtOverlayTunnelVtepStats.Stats.Reassembled_pkts_rx,
			"bad_inner_ipv4_len_rx":  ret.DtOverlayTunnelVtepStats.Stats.Bad_inner_ipv4_len_rx,
			"bad_inner_ipv6_len_rx":  ret.DtOverlayTunnelVtepStats.Stats.Bad_inner_ipv6_len_rx,
			"frag_drop_pkts_tx":      ret.DtOverlayTunnelVtepStats.Stats.Frag_drop_pkts_tx,
			"lif_un_init_rx":         ret.DtOverlayTunnelVtepStats.Stats.Lif_un_init_rx,
		},
	}
}

func getObjectOverlayTunnelVtepStatsStats(d []interface{}) edpt.OverlayTunnelVtepStatsStats {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cfg_err_count = in["cfg_err_count"].(int)
		ret.Flooded_pkt_count = in["flooded_pkt_count"].(int)
		ret.Encap_unresolved_count = in["encap_unresolved_count"].(int)
		ret.Unknown_encap_rx_pkt = in["unknown_encap_rx_pkt"].(int)
		ret.Unknown_encap_tx_pkt = in["unknown_encap_tx_pkt"].(int)
		ret.Arp_req_sent = in["arp_req_sent"].(int)
		ret.Vtep_host_learned = in["vtep_host_learned"].(int)
		ret.Vtep_host_learn_error = in["vtep_host_learn_error"].(int)
		ret.Invalid_lif_rx = in["invalid_lif_rx"].(int)
		ret.Invalid_lif_tx = in["invalid_lif_tx"].(int)
		ret.Unknown_vtep_tx = in["unknown_vtep_tx"].(int)
		ret.Unknown_vtep_rx = in["unknown_vtep_rx"].(int)
		ret.Unhandled_pkt_rx = in["unhandled_pkt_rx"].(int)
		ret.Unhandled_pkt_tx = in["unhandled_pkt_tx"].(int)
		ret.Total_pkts_rx = in["total_pkts_rx"].(int)
		ret.Total_bytes_rx = in["total_bytes_rx"].(int)
		ret.Unicast_pkt_rx = in["unicast_pkt_rx"].(int)
		ret.Bcast_pkt_rx = in["bcast_pkt_rx"].(int)
		ret.Mcast_pkt_rx = in["mcast_pkt_rx"].(int)
		ret.Dropped_pkt_rx = in["dropped_pkt_rx"].(int)
		ret.Encap_miss_pkts_rx = in["encap_miss_pkts_rx"].(int)
		ret.Bad_chksum_pks_rx = in["bad_chksum_pks_rx"].(int)
		ret.Requeue_pkts_in = in["requeue_pkts_in"].(int)
		ret.Pkts_out = in["pkts_out"].(int)
		ret.Total_bytes_tx = in["total_bytes_tx"].(int)
		ret.Unicast_pkt_tx = in["unicast_pkt_tx"].(int)
		ret.Bcast_pkt_tx = in["bcast_pkt_tx"].(int)
		ret.Mcast_pkt_tx = in["mcast_pkt_tx"].(int)
		ret.Dropped_pkts_tx = in["dropped_pkts_tx"].(int)
		ret.Large_pkts_rx = in["large_pkts_rx"].(int)
		ret.Dot1q_pkts_rx = in["dot1q_pkts_rx"].(int)
		ret.Frag_pkts_tx = in["frag_pkts_tx"].(int)
		ret.Reassembled_pkts_rx = in["reassembled_pkts_rx"].(int)
		ret.Bad_inner_ipv4_len_rx = in["bad_inner_ipv4_len_rx"].(int)
		ret.Bad_inner_ipv6_len_rx = in["bad_inner_ipv6_len_rx"].(int)
		ret.Frag_drop_pkts_tx = in["frag_drop_pkts_tx"].(int)
		ret.Lif_un_init_rx = in["lif_un_init_rx"].(int)
	}
	return ret
}

func dataToEndpointOverlayTunnelVtepStats(d *schema.ResourceData) edpt.OverlayTunnelVtepStats {
	var ret edpt.OverlayTunnelVtepStats

	ret.Id1 = d.Get("id1").(int)

	ret.Stats = getObjectOverlayTunnelVtepStatsStats(d.Get("stats").([]interface{}))
	return ret
}
