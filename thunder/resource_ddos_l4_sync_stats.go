package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosL4SyncStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_l4_sync_stats`: Statistics for the object l4-sync\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosL4SyncStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sync_hello_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Hello Received",
						},
						"sync_hello_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Hello Sent",
						},
						"sync_dst_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Dst Sent",
						},
						"sync_dst_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Dst Received",
						},
						"sync_src_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Src Sent",
						},
						"sync_src_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Src Received",
						},
						"sync_src_dst_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync SrcDst Sent",
						},
						"sync_src_dst_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "Sync SrcDst Received",
						},
						"sync_szp_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync SrcZoneService Sent",
						},
						"sync_szp_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "Sync SrcZoneService Received",
						},
						"sync_sent_no_peer": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent No Peer",
						},
						"sync_sent_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent Failed",
						},
						"sync_rcv_hello_unk_peer": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Hello Unknown Peer",
						},
						"sync_src_dst_no_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Sync SrcDst No-Dst",
						},
						"sync_szp_no_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Sync SrcZoneService No-Dst",
						},
						"sync_rcv_entry_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Entry Create Failed",
						},
						"sync_rcv_entry_conflict_static": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Entry Conflict Static",
						},
						"sync_rcv_unk_msg": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Unknown Msg",
						},
						"sync_rcv_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Failed",
						},
						"sync_rcv_hello_unk_subtype": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Hello Unknown Subtype",
						},
						"sync_rcv_entry_unk_subtype": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Entry Unknown Subtype",
						},
						"sync_rcv_entry_unk": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Unknown Entry",
						},
						"sync_src_udp_wl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent Src UDP Authenticated",
						},
						"sync_src_tcp_wl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent Src TCP Authenticated",
						},
						"sync_src_icmp_wl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent Src ICMP Authenticated",
						},
						"sync_src_other_wl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent Src OTHER Authenticated",
						},
						"sync_src_dns_udp_auth_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent Src DNS UDP Authenticated",
						},
						"sync_src_dns_tcp_auth_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent Src DNS Force-TCP Authenticated",
						},
						"sync_src_bl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent Src Blacklisted",
						},
						"sync_src_dst_udp_wl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcDst UDP Authenticated",
						},
						"sync_src_dst_tcp_wl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcDst TCP Authenticated",
						},
						"sync_src_dst_icmp_wl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcDst ICMP Authenticated",
						},
						"sync_src_dst_other_wl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcDst OTHER Authenticated",
						},
						"sync_src_dst_dns_udp_auth_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcDst DNS UDP Authenticated",
						},
						"sync_src_dst_dns_tcp_auth_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcDst DNS Force-TCP Authenticated",
						},
						"sync_src_dst_bl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcDst Blacklisted",
						},
						"sync_szp_tcp_auth_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcZoneService TCP Authenticated",
						},
						"sync_szp_udp_auth_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcZoneService UDP Authenticated",
						},
						"sync_szp_icmp_auth_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcZoneService ICMP Authenticated",
						},
						"sync_szp_other_auth_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcZoneService OTHER Authenticated",
						},
						"sync_szp_dns_udp_auth_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcZoneService DNS UDP Authenticated",
						},
						"sync_szp_dns_tcp_auth_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcZoneService DNS Force-TCP Authenticated",
						},
						"sync_szp_bl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Sent SrcZoneService Blacklist",
						},
						"sync_src_udp_wl_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Src UDP Authenticated",
						},
						"sync_src_tcp_wl_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Src TCP Authenticated",
						},
						"sync_src_icmp_wl_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Src ICMP Authenticated",
						},
						"sync_src_other_wl_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Src OTHER Authenticated",
						},
						"sync_src_dns_udp_auth_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Src DNS UDP Authenticated",
						},
						"sync_src_dns_tcp_auth_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Src DNS Force-TCP Authenticated",
						},
						"sync_src_bl_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received Src Blacklisted",
						},
						"sync_src_dst_udp_wl_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcDst UDP Authenticated",
						},
						"sync_src_dst_tcp_wl_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcDst TCP Authenticated",
						},
						"sync_src_dst_icmp_wl_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcDst ICMP Authenticated",
						},
						"sync_src_dst_other_wl_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcDst OTHER Authenticated",
						},
						"sync_src_dst_dns_udp_auth_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcDst DNS UDP Authenticated",
						},
						"sync_src_dst_dns_tcp_auth_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcDst DNS Force-TCP Authenticated",
						},
						"sync_src_dst_bl_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcDst Blacklisted",
						},
						"sync_szp_tcp_auth_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcZoneService TCP Authenticated",
						},
						"sync_szp_udp_auth_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcZoneService UDP Authenticated",
						},
						"sync_szp_icmp_auth_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcZoneService ICMP Authenticated",
						},
						"sync_szp_other_auth_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcZoneService OTHER Authenticated",
						},
						"sync_szp_dns_udp_auth_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcZoneService DNS UDP Authenticated",
						},
						"sync_szp_dns_tcp_auth_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcZoneService DNS Force-TCP Authenticated",
						},
						"sync_szp_bl_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Received SrcZoneService Blacklist",
						},
					},
				},
			},
		},
	}
}

func resourceDdosL4SyncStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL4SyncStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL4SyncStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosL4SyncStatsStats := setObjectDdosL4SyncStatsStats(res)
		d.Set("stats", DdosL4SyncStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosL4SyncStatsStats(ret edpt.DataDdosL4SyncStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sync_hello_rcv":                 ret.DtDdosL4SyncStats.Stats.Sync_hello_rcv,
			"sync_hello_sent":                ret.DtDdosL4SyncStats.Stats.Sync_hello_sent,
			"sync_dst_sent":                  ret.DtDdosL4SyncStats.Stats.Sync_dst_sent,
			"sync_dst_rcv":                   ret.DtDdosL4SyncStats.Stats.Sync_dst_rcv,
			"sync_src_sent":                  ret.DtDdosL4SyncStats.Stats.Sync_src_sent,
			"sync_src_rcv":                   ret.DtDdosL4SyncStats.Stats.Sync_src_rcv,
			"sync_src_dst_sent":              ret.DtDdosL4SyncStats.Stats.Sync_src_dst_sent,
			"sync_src_dst_rcv":               ret.DtDdosL4SyncStats.Stats.Sync_src_dst_rcv,
			"sync_szp_sent":                  ret.DtDdosL4SyncStats.Stats.Sync_szp_sent,
			"sync_szp_rcv":                   ret.DtDdosL4SyncStats.Stats.Sync_szp_rcv,
			"sync_sent_no_peer":              ret.DtDdosL4SyncStats.Stats.Sync_sent_no_peer,
			"sync_sent_fail":                 ret.DtDdosL4SyncStats.Stats.Sync_sent_fail,
			"sync_rcv_hello_unk_peer":        ret.DtDdosL4SyncStats.Stats.Sync_rcv_hello_unk_peer,
			"sync_src_dst_no_dst_drop":       ret.DtDdosL4SyncStats.Stats.Sync_src_dst_no_dst_drop,
			"sync_szp_no_dst_drop":           ret.DtDdosL4SyncStats.Stats.Sync_szp_no_dst_drop,
			"sync_rcv_entry_create_fail":     ret.DtDdosL4SyncStats.Stats.Sync_rcv_entry_create_fail,
			"sync_rcv_entry_conflict_static": ret.DtDdosL4SyncStats.Stats.Sync_rcv_entry_conflict_static,
			"sync_rcv_unk_msg":               ret.DtDdosL4SyncStats.Stats.Sync_rcv_unk_msg,
			"sync_rcv_fail":                  ret.DtDdosL4SyncStats.Stats.Sync_rcv_fail,
			"sync_rcv_hello_unk_subtype":     ret.DtDdosL4SyncStats.Stats.Sync_rcv_hello_unk_subtype,
			"sync_rcv_entry_unk_subtype":     ret.DtDdosL4SyncStats.Stats.Sync_rcv_entry_unk_subtype,
			"sync_rcv_entry_unk":             ret.DtDdosL4SyncStats.Stats.Sync_rcv_entry_unk,
			"sync_src_udp_wl_sent":           ret.DtDdosL4SyncStats.Stats.Sync_src_udp_wl_sent,
			"sync_src_tcp_wl_sent":           ret.DtDdosL4SyncStats.Stats.Sync_src_tcp_wl_sent,
			"sync_src_icmp_wl_sent":          ret.DtDdosL4SyncStats.Stats.Sync_src_icmp_wl_sent,
			"sync_src_other_wl_sent":         ret.DtDdosL4SyncStats.Stats.Sync_src_other_wl_sent,
			"sync_src_dns_udp_auth_sent":     ret.DtDdosL4SyncStats.Stats.Sync_src_dns_udp_auth_sent,
			"sync_src_dns_tcp_auth_sent":     ret.DtDdosL4SyncStats.Stats.Sync_src_dns_tcp_auth_sent,
			"sync_src_bl_sent":               ret.DtDdosL4SyncStats.Stats.Sync_src_bl_sent,
			"sync_src_dst_udp_wl_sent":       ret.DtDdosL4SyncStats.Stats.Sync_src_dst_udp_wl_sent,
			"sync_src_dst_tcp_wl_sent":       ret.DtDdosL4SyncStats.Stats.Sync_src_dst_tcp_wl_sent,
			"sync_src_dst_icmp_wl_sent":      ret.DtDdosL4SyncStats.Stats.Sync_src_dst_icmp_wl_sent,
			"sync_src_dst_other_wl_sent":     ret.DtDdosL4SyncStats.Stats.Sync_src_dst_other_wl_sent,
			"sync_src_dst_dns_udp_auth_sent": ret.DtDdosL4SyncStats.Stats.Sync_src_dst_dns_udp_auth_sent,
			"sync_src_dst_dns_tcp_auth_sent": ret.DtDdosL4SyncStats.Stats.Sync_src_dst_dns_tcp_auth_sent,
			"sync_src_dst_bl_sent":           ret.DtDdosL4SyncStats.Stats.Sync_src_dst_bl_sent,
			"sync_szp_tcp_auth_sent":         ret.DtDdosL4SyncStats.Stats.Sync_szp_tcp_auth_sent,
			"sync_szp_udp_auth_sent":         ret.DtDdosL4SyncStats.Stats.Sync_szp_udp_auth_sent,
			"sync_szp_icmp_auth_sent":        ret.DtDdosL4SyncStats.Stats.Sync_szp_icmp_auth_sent,
			"sync_szp_other_auth_sent":       ret.DtDdosL4SyncStats.Stats.Sync_szp_other_auth_sent,
			"sync_szp_dns_udp_auth_sent":     ret.DtDdosL4SyncStats.Stats.Sync_szp_dns_udp_auth_sent,
			"sync_szp_dns_tcp_auth_sent":     ret.DtDdosL4SyncStats.Stats.Sync_szp_dns_tcp_auth_sent,
			"sync_szp_bl_sent":               ret.DtDdosL4SyncStats.Stats.Sync_szp_bl_sent,
			"sync_src_udp_wl_rcvd":           ret.DtDdosL4SyncStats.Stats.Sync_src_udp_wl_rcvd,
			"sync_src_tcp_wl_rcvd":           ret.DtDdosL4SyncStats.Stats.Sync_src_tcp_wl_rcvd,
			"sync_src_icmp_wl_rcvd":          ret.DtDdosL4SyncStats.Stats.Sync_src_icmp_wl_rcvd,
			"sync_src_other_wl_rcvd":         ret.DtDdosL4SyncStats.Stats.Sync_src_other_wl_rcvd,
			"sync_src_dns_udp_auth_rcvd":     ret.DtDdosL4SyncStats.Stats.Sync_src_dns_udp_auth_rcvd,
			"sync_src_dns_tcp_auth_rcvd":     ret.DtDdosL4SyncStats.Stats.Sync_src_dns_tcp_auth_rcvd,
			"sync_src_bl_rcvd":               ret.DtDdosL4SyncStats.Stats.Sync_src_bl_rcvd,
			"sync_src_dst_udp_wl_rcvd":       ret.DtDdosL4SyncStats.Stats.Sync_src_dst_udp_wl_rcvd,
			"sync_src_dst_tcp_wl_rcvd":       ret.DtDdosL4SyncStats.Stats.Sync_src_dst_tcp_wl_rcvd,
			"sync_src_dst_icmp_wl_rcvd":      ret.DtDdosL4SyncStats.Stats.Sync_src_dst_icmp_wl_rcvd,
			"sync_src_dst_other_wl_rcvd":     ret.DtDdosL4SyncStats.Stats.Sync_src_dst_other_wl_rcvd,
			"sync_src_dst_dns_udp_auth_rcvd": ret.DtDdosL4SyncStats.Stats.Sync_src_dst_dns_udp_auth_rcvd,
			"sync_src_dst_dns_tcp_auth_rcvd": ret.DtDdosL4SyncStats.Stats.Sync_src_dst_dns_tcp_auth_rcvd,
			"sync_src_dst_bl_rcvd":           ret.DtDdosL4SyncStats.Stats.Sync_src_dst_bl_rcvd,
			"sync_szp_tcp_auth_rcvd":         ret.DtDdosL4SyncStats.Stats.Sync_szp_tcp_auth_rcvd,
			"sync_szp_udp_auth_rcvd":         ret.DtDdosL4SyncStats.Stats.Sync_szp_udp_auth_rcvd,
			"sync_szp_icmp_auth_rcvd":        ret.DtDdosL4SyncStats.Stats.Sync_szp_icmp_auth_rcvd,
			"sync_szp_other_auth_rcvd":       ret.DtDdosL4SyncStats.Stats.Sync_szp_other_auth_rcvd,
			"sync_szp_dns_udp_auth_rcvd":     ret.DtDdosL4SyncStats.Stats.Sync_szp_dns_udp_auth_rcvd,
			"sync_szp_dns_tcp_auth_rcvd":     ret.DtDdosL4SyncStats.Stats.Sync_szp_dns_tcp_auth_rcvd,
			"sync_szp_bl_rcvd":               ret.DtDdosL4SyncStats.Stats.Sync_szp_bl_rcvd,
		},
	}
}

func getObjectDdosL4SyncStatsStats(d []interface{}) edpt.DdosL4SyncStatsStats {

	count1 := len(d)
	var ret edpt.DdosL4SyncStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Sync_hello_rcv = in["sync_hello_rcv"].(int)
		ret.Sync_hello_sent = in["sync_hello_sent"].(int)
		ret.Sync_dst_sent = in["sync_dst_sent"].(int)
		ret.Sync_dst_rcv = in["sync_dst_rcv"].(int)
		ret.Sync_src_sent = in["sync_src_sent"].(int)
		ret.Sync_src_rcv = in["sync_src_rcv"].(int)
		ret.Sync_src_dst_sent = in["sync_src_dst_sent"].(int)
		ret.Sync_src_dst_rcv = in["sync_src_dst_rcv"].(int)
		ret.Sync_szp_sent = in["sync_szp_sent"].(int)
		ret.Sync_szp_rcv = in["sync_szp_rcv"].(int)
		ret.Sync_sent_no_peer = in["sync_sent_no_peer"].(int)
		ret.Sync_sent_fail = in["sync_sent_fail"].(int)
		ret.Sync_rcv_hello_unk_peer = in["sync_rcv_hello_unk_peer"].(int)
		ret.Sync_src_dst_no_dst_drop = in["sync_src_dst_no_dst_drop"].(int)
		ret.Sync_szp_no_dst_drop = in["sync_szp_no_dst_drop"].(int)
		ret.Sync_rcv_entry_create_fail = in["sync_rcv_entry_create_fail"].(int)
		ret.Sync_rcv_entry_conflict_static = in["sync_rcv_entry_conflict_static"].(int)
		ret.Sync_rcv_unk_msg = in["sync_rcv_unk_msg"].(int)
		ret.Sync_rcv_fail = in["sync_rcv_fail"].(int)
		ret.Sync_rcv_hello_unk_subtype = in["sync_rcv_hello_unk_subtype"].(int)
		ret.Sync_rcv_entry_unk_subtype = in["sync_rcv_entry_unk_subtype"].(int)
		ret.Sync_rcv_entry_unk = in["sync_rcv_entry_unk"].(int)
		ret.Sync_src_udp_wl_sent = in["sync_src_udp_wl_sent"].(int)
		ret.Sync_src_tcp_wl_sent = in["sync_src_tcp_wl_sent"].(int)
		ret.Sync_src_icmp_wl_sent = in["sync_src_icmp_wl_sent"].(int)
		ret.Sync_src_other_wl_sent = in["sync_src_other_wl_sent"].(int)
		ret.Sync_src_dns_udp_auth_sent = in["sync_src_dns_udp_auth_sent"].(int)
		ret.Sync_src_dns_tcp_auth_sent = in["sync_src_dns_tcp_auth_sent"].(int)
		ret.Sync_src_bl_sent = in["sync_src_bl_sent"].(int)
		ret.Sync_src_dst_udp_wl_sent = in["sync_src_dst_udp_wl_sent"].(int)
		ret.Sync_src_dst_tcp_wl_sent = in["sync_src_dst_tcp_wl_sent"].(int)
		ret.Sync_src_dst_icmp_wl_sent = in["sync_src_dst_icmp_wl_sent"].(int)
		ret.Sync_src_dst_other_wl_sent = in["sync_src_dst_other_wl_sent"].(int)
		ret.Sync_src_dst_dns_udp_auth_sent = in["sync_src_dst_dns_udp_auth_sent"].(int)
		ret.Sync_src_dst_dns_tcp_auth_sent = in["sync_src_dst_dns_tcp_auth_sent"].(int)
		ret.Sync_src_dst_bl_sent = in["sync_src_dst_bl_sent"].(int)
		ret.Sync_szp_tcp_auth_sent = in["sync_szp_tcp_auth_sent"].(int)
		ret.Sync_szp_udp_auth_sent = in["sync_szp_udp_auth_sent"].(int)
		ret.Sync_szp_icmp_auth_sent = in["sync_szp_icmp_auth_sent"].(int)
		ret.Sync_szp_other_auth_sent = in["sync_szp_other_auth_sent"].(int)
		ret.Sync_szp_dns_udp_auth_sent = in["sync_szp_dns_udp_auth_sent"].(int)
		ret.Sync_szp_dns_tcp_auth_sent = in["sync_szp_dns_tcp_auth_sent"].(int)
		ret.Sync_szp_bl_sent = in["sync_szp_bl_sent"].(int)
		ret.Sync_src_udp_wl_rcvd = in["sync_src_udp_wl_rcvd"].(int)
		ret.Sync_src_tcp_wl_rcvd = in["sync_src_tcp_wl_rcvd"].(int)
		ret.Sync_src_icmp_wl_rcvd = in["sync_src_icmp_wl_rcvd"].(int)
		ret.Sync_src_other_wl_rcvd = in["sync_src_other_wl_rcvd"].(int)
		ret.Sync_src_dns_udp_auth_rcvd = in["sync_src_dns_udp_auth_rcvd"].(int)
		ret.Sync_src_dns_tcp_auth_rcvd = in["sync_src_dns_tcp_auth_rcvd"].(int)
		ret.Sync_src_bl_rcvd = in["sync_src_bl_rcvd"].(int)
		ret.Sync_src_dst_udp_wl_rcvd = in["sync_src_dst_udp_wl_rcvd"].(int)
		ret.Sync_src_dst_tcp_wl_rcvd = in["sync_src_dst_tcp_wl_rcvd"].(int)
		ret.Sync_src_dst_icmp_wl_rcvd = in["sync_src_dst_icmp_wl_rcvd"].(int)
		ret.Sync_src_dst_other_wl_rcvd = in["sync_src_dst_other_wl_rcvd"].(int)
		ret.Sync_src_dst_dns_udp_auth_rcvd = in["sync_src_dst_dns_udp_auth_rcvd"].(int)
		ret.Sync_src_dst_dns_tcp_auth_rcvd = in["sync_src_dst_dns_tcp_auth_rcvd"].(int)
		ret.Sync_src_dst_bl_rcvd = in["sync_src_dst_bl_rcvd"].(int)
		ret.Sync_szp_tcp_auth_rcvd = in["sync_szp_tcp_auth_rcvd"].(int)
		ret.Sync_szp_udp_auth_rcvd = in["sync_szp_udp_auth_rcvd"].(int)
		ret.Sync_szp_icmp_auth_rcvd = in["sync_szp_icmp_auth_rcvd"].(int)
		ret.Sync_szp_other_auth_rcvd = in["sync_szp_other_auth_rcvd"].(int)
		ret.Sync_szp_dns_udp_auth_rcvd = in["sync_szp_dns_udp_auth_rcvd"].(int)
		ret.Sync_szp_dns_tcp_auth_rcvd = in["sync_szp_dns_tcp_auth_rcvd"].(int)
		ret.Sync_szp_bl_rcvd = in["sync_szp_bl_rcvd"].(int)
	}
	return ret
}

func dataToEndpointDdosL4SyncStats(d *schema.ResourceData) edpt.DdosL4SyncStats {
	var ret edpt.DdosL4SyncStats

	ret.Stats = getObjectDdosL4SyncStatsStats(d.Get("stats").([]interface{}))
	return ret
}
