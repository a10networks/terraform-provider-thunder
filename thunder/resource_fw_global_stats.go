package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceFwGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Full-cone Created",
						},
						"tcp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Full-cone Freed",
						},
						"udp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Full-cone Created",
						},
						"udp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Full-cone Freed",
						},
						"fullcone_creation_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Full-Cone Creation Failure",
						},
						"data_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Created",
						},
						"data_session_created_local": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Created Local",
						},
						"data_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Freed",
						},
						"data_session_freed_local": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Freed Local",
						},
						"dyn_blist_sess_created": {
							Type: schema.TypeInt, Optional: true, Description: "Dynamic Blacklist Session Created",
						},
						"dyn_blist_sess_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Dynamic Blacklist Freed",
						},
						"dyn_blist_pkt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dynamic Blacklist - Packet Drop",
						},
						"active_fullcone_session": {
							Type: schema.TypeInt, Optional: true, Description: "Total Active Full-cone sessions",
						},
						"limit_entry_created": {
							Type: schema.TypeInt, Optional: true, Description: "Limit Entry Created",
						},
						"limit_entry_marked_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "Limit Entry Marked Deleted",
						},
						"undetermined_rule_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Undetermined rule detected",
						},
						"non_syn_pkt_fwd_allowed": {
							Type: schema.TypeInt, Optional: true, Description: "Non-SYN pkt forward allowed",
						},
						"fwd_ingress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packets TCP",
						},
						"fwd_egress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packets TCP",
						},
						"rev_ingress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packets TCP",
						},
						"rev_egress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packets TCP",
						},
						"fwd_ingress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Bytes TCP",
						},
						"fwd_egress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Bytes TCP",
						},
						"rev_ingress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Bytes TCP",
						},
						"rev_egress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Bytes TCP",
						},
						"fwd_ingress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packets UDP",
						},
						"fwd_egress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packets UDP",
						},
						"rev_ingress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packets UDP",
						},
						"rev_egress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packets UDP",
						},
						"fwd_ingress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Bytes UDP",
						},
						"fwd_egress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Bytes UDP",
						},
						"rev_ingress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Bytes UDP",
						},
						"rev_egress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Bytes UDP",
						},
						"fwd_ingress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packets ICMP",
						},
						"fwd_egress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packets ICMP",
						},
						"rev_ingress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packets ICMP",
						},
						"rev_egress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packets ICMP",
						},
						"fwd_ingress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Bytes ICMP",
						},
						"fwd_egress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Bytes ICMP",
						},
						"rev_ingress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Bytes ICMP",
						},
						"rev_egress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Bytes ICMP",
						},
						"fwd_ingress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packets OTHERS",
						},
						"fwd_egress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packets OTHERS",
						},
						"rev_ingress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packets OTHERS",
						},
						"rev_egress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packets OTHERS",
						},
						"fwd_ingress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Bytes OTHERS",
						},
						"fwd_egress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Bytes OTHERS",
						},
						"rev_ingress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Bytes OTHERS",
						},
						"rev_egress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Bytes OTHERS",
						},
						"fwd_ingress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packet size between 0 and 200",
						},
						"fwd_ingress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packet size between 201 and 800",
						},
						"fwd_ingress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packet size between 801 and 1550",
						},
						"fwd_ingress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packet size between 1551 and 9000",
						},
						"fwd_egress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packet size between 0 and 200",
						},
						"fwd_egress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packet size between 201 and 800",
						},
						"fwd_egress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packet size between 801 and 1550",
						},
						"fwd_egress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packet size between 1551 and 9000",
						},
						"rev_ingress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packet size between 0 and 200",
						},
						"rev_ingress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packet size between 201 and 800",
						},
						"rev_ingress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packet size between 801 and 1550",
						},
						"rev_ingress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packet size between 1551 and 9000",
						},
						"rev_egress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packet size between 0 and 200",
						},
						"rev_egress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packet size between 201 and 800",
						},
						"rev_egress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packet size between 801 and 1550",
						},
						"rev_egress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packet size between 1551 and 9000",
						},
					},
				},
			},
		},
	}
}

func resourceFwGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwGlobalStatsStats := setObjectFwGlobalStatsStats(res)
		d.Set("stats", FwGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwGlobalStatsStats(ret edpt.DataFwGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tcp_fullcone_created":        ret.DtFwGlobalStats.Stats.Tcp_fullcone_created,
			"tcp_fullcone_freed":          ret.DtFwGlobalStats.Stats.Tcp_fullcone_freed,
			"udp_fullcone_created":        ret.DtFwGlobalStats.Stats.Udp_fullcone_created,
			"udp_fullcone_freed":          ret.DtFwGlobalStats.Stats.Udp_fullcone_freed,
			"fullcone_creation_failure":   ret.DtFwGlobalStats.Stats.Fullcone_creation_failure,
			"data_session_created":        ret.DtFwGlobalStats.Stats.Data_session_created,
			"data_session_created_local":  ret.DtFwGlobalStats.Stats.Data_session_created_local,
			"data_session_freed":          ret.DtFwGlobalStats.Stats.Data_session_freed,
			"data_session_freed_local":    ret.DtFwGlobalStats.Stats.Data_session_freed_local,
			"dyn_blist_sess_created":      ret.DtFwGlobalStats.Stats.Dyn_blist_sess_created,
			"dyn_blist_sess_freed":        ret.DtFwGlobalStats.Stats.Dyn_blist_sess_freed,
			"dyn_blist_pkt_drop":          ret.DtFwGlobalStats.Stats.Dyn_blist_pkt_drop,
			"active_fullcone_session":     ret.DtFwGlobalStats.Stats.Active_fullcone_session,
			"limit_entry_created":         ret.DtFwGlobalStats.Stats.LimitEntryCreated,
			"limit_entry_marked_deleted":  ret.DtFwGlobalStats.Stats.LimitEntryMarkedDeleted,
			"undetermined_rule_counter":   ret.DtFwGlobalStats.Stats.UndeterminedRuleCounter,
			"non_syn_pkt_fwd_allowed":     ret.DtFwGlobalStats.Stats.Non_syn_pkt_fwd_allowed,
			"fwd_ingress_packets_tcp":     ret.DtFwGlobalStats.Stats.Fwd_ingress_packets_tcp,
			"fwd_egress_packets_tcp":      ret.DtFwGlobalStats.Stats.Fwd_egress_packets_tcp,
			"rev_ingress_packets_tcp":     ret.DtFwGlobalStats.Stats.Rev_ingress_packets_tcp,
			"rev_egress_packets_tcp":      ret.DtFwGlobalStats.Stats.Rev_egress_packets_tcp,
			"fwd_ingress_bytes_tcp":       ret.DtFwGlobalStats.Stats.Fwd_ingress_bytes_tcp,
			"fwd_egress_bytes_tcp":        ret.DtFwGlobalStats.Stats.Fwd_egress_bytes_tcp,
			"rev_ingress_bytes_tcp":       ret.DtFwGlobalStats.Stats.Rev_ingress_bytes_tcp,
			"rev_egress_bytes_tcp":        ret.DtFwGlobalStats.Stats.Rev_egress_bytes_tcp,
			"fwd_ingress_packets_udp":     ret.DtFwGlobalStats.Stats.Fwd_ingress_packets_udp,
			"fwd_egress_packets_udp":      ret.DtFwGlobalStats.Stats.Fwd_egress_packets_udp,
			"rev_ingress_packets_udp":     ret.DtFwGlobalStats.Stats.Rev_ingress_packets_udp,
			"rev_egress_packets_udp":      ret.DtFwGlobalStats.Stats.Rev_egress_packets_udp,
			"fwd_ingress_bytes_udp":       ret.DtFwGlobalStats.Stats.Fwd_ingress_bytes_udp,
			"fwd_egress_bytes_udp":        ret.DtFwGlobalStats.Stats.Fwd_egress_bytes_udp,
			"rev_ingress_bytes_udp":       ret.DtFwGlobalStats.Stats.Rev_ingress_bytes_udp,
			"rev_egress_bytes_udp":        ret.DtFwGlobalStats.Stats.Rev_egress_bytes_udp,
			"fwd_ingress_packets_icmp":    ret.DtFwGlobalStats.Stats.Fwd_ingress_packets_icmp,
			"fwd_egress_packets_icmp":     ret.DtFwGlobalStats.Stats.Fwd_egress_packets_icmp,
			"rev_ingress_packets_icmp":    ret.DtFwGlobalStats.Stats.Rev_ingress_packets_icmp,
			"rev_egress_packets_icmp":     ret.DtFwGlobalStats.Stats.Rev_egress_packets_icmp,
			"fwd_ingress_bytes_icmp":      ret.DtFwGlobalStats.Stats.Fwd_ingress_bytes_icmp,
			"fwd_egress_bytes_icmp":       ret.DtFwGlobalStats.Stats.Fwd_egress_bytes_icmp,
			"rev_ingress_bytes_icmp":      ret.DtFwGlobalStats.Stats.Rev_ingress_bytes_icmp,
			"rev_egress_bytes_icmp":       ret.DtFwGlobalStats.Stats.Rev_egress_bytes_icmp,
			"fwd_ingress_packets_others":  ret.DtFwGlobalStats.Stats.Fwd_ingress_packets_others,
			"fwd_egress_packets_others":   ret.DtFwGlobalStats.Stats.Fwd_egress_packets_others,
			"rev_ingress_packets_others":  ret.DtFwGlobalStats.Stats.Rev_ingress_packets_others,
			"rev_egress_packets_others":   ret.DtFwGlobalStats.Stats.Rev_egress_packets_others,
			"fwd_ingress_bytes_others":    ret.DtFwGlobalStats.Stats.Fwd_ingress_bytes_others,
			"fwd_egress_bytes_others":     ret.DtFwGlobalStats.Stats.Fwd_egress_bytes_others,
			"rev_ingress_bytes_others":    ret.DtFwGlobalStats.Stats.Rev_ingress_bytes_others,
			"rev_egress_bytes_others":     ret.DtFwGlobalStats.Stats.Rev_egress_bytes_others,
			"fwd_ingress_pkt_size_range1": ret.DtFwGlobalStats.Stats.Fwd_ingress_pkt_size_range1,
			"fwd_ingress_pkt_size_range2": ret.DtFwGlobalStats.Stats.Fwd_ingress_pkt_size_range2,
			"fwd_ingress_pkt_size_range3": ret.DtFwGlobalStats.Stats.Fwd_ingress_pkt_size_range3,
			"fwd_ingress_pkt_size_range4": ret.DtFwGlobalStats.Stats.Fwd_ingress_pkt_size_range4,
			"fwd_egress_pkt_size_range1":  ret.DtFwGlobalStats.Stats.Fwd_egress_pkt_size_range1,
			"fwd_egress_pkt_size_range2":  ret.DtFwGlobalStats.Stats.Fwd_egress_pkt_size_range2,
			"fwd_egress_pkt_size_range3":  ret.DtFwGlobalStats.Stats.Fwd_egress_pkt_size_range3,
			"fwd_egress_pkt_size_range4":  ret.DtFwGlobalStats.Stats.Fwd_egress_pkt_size_range4,
			"rev_ingress_pkt_size_range1": ret.DtFwGlobalStats.Stats.Rev_ingress_pkt_size_range1,
			"rev_ingress_pkt_size_range2": ret.DtFwGlobalStats.Stats.Rev_ingress_pkt_size_range2,
			"rev_ingress_pkt_size_range3": ret.DtFwGlobalStats.Stats.Rev_ingress_pkt_size_range3,
			"rev_ingress_pkt_size_range4": ret.DtFwGlobalStats.Stats.Rev_ingress_pkt_size_range4,
			"rev_egress_pkt_size_range1":  ret.DtFwGlobalStats.Stats.Rev_egress_pkt_size_range1,
			"rev_egress_pkt_size_range2":  ret.DtFwGlobalStats.Stats.Rev_egress_pkt_size_range2,
			"rev_egress_pkt_size_range3":  ret.DtFwGlobalStats.Stats.Rev_egress_pkt_size_range3,
			"rev_egress_pkt_size_range4":  ret.DtFwGlobalStats.Stats.Rev_egress_pkt_size_range4,
		},
	}
}

func getObjectFwGlobalStatsStats(d []interface{}) edpt.FwGlobalStatsStats {

	count1 := len(d)
	var ret edpt.FwGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp_fullcone_created = in["tcp_fullcone_created"].(int)
		ret.Tcp_fullcone_freed = in["tcp_fullcone_freed"].(int)
		ret.Udp_fullcone_created = in["udp_fullcone_created"].(int)
		ret.Udp_fullcone_freed = in["udp_fullcone_freed"].(int)
		ret.Fullcone_creation_failure = in["fullcone_creation_failure"].(int)
		ret.Data_session_created = in["data_session_created"].(int)
		ret.Data_session_created_local = in["data_session_created_local"].(int)
		ret.Data_session_freed = in["data_session_freed"].(int)
		ret.Data_session_freed_local = in["data_session_freed_local"].(int)
		ret.Dyn_blist_sess_created = in["dyn_blist_sess_created"].(int)
		ret.Dyn_blist_sess_freed = in["dyn_blist_sess_freed"].(int)
		ret.Dyn_blist_pkt_drop = in["dyn_blist_pkt_drop"].(int)
		ret.Active_fullcone_session = in["active_fullcone_session"].(int)
		ret.LimitEntryCreated = in["limit_entry_created"].(int)
		ret.LimitEntryMarkedDeleted = in["limit_entry_marked_deleted"].(int)
		ret.UndeterminedRuleCounter = in["undetermined_rule_counter"].(int)
		ret.Non_syn_pkt_fwd_allowed = in["non_syn_pkt_fwd_allowed"].(int)
		ret.Fwd_ingress_packets_tcp = in["fwd_ingress_packets_tcp"].(int)
		ret.Fwd_egress_packets_tcp = in["fwd_egress_packets_tcp"].(int)
		ret.Rev_ingress_packets_tcp = in["rev_ingress_packets_tcp"].(int)
		ret.Rev_egress_packets_tcp = in["rev_egress_packets_tcp"].(int)
		ret.Fwd_ingress_bytes_tcp = in["fwd_ingress_bytes_tcp"].(int)
		ret.Fwd_egress_bytes_tcp = in["fwd_egress_bytes_tcp"].(int)
		ret.Rev_ingress_bytes_tcp = in["rev_ingress_bytes_tcp"].(int)
		ret.Rev_egress_bytes_tcp = in["rev_egress_bytes_tcp"].(int)
		ret.Fwd_ingress_packets_udp = in["fwd_ingress_packets_udp"].(int)
		ret.Fwd_egress_packets_udp = in["fwd_egress_packets_udp"].(int)
		ret.Rev_ingress_packets_udp = in["rev_ingress_packets_udp"].(int)
		ret.Rev_egress_packets_udp = in["rev_egress_packets_udp"].(int)
		ret.Fwd_ingress_bytes_udp = in["fwd_ingress_bytes_udp"].(int)
		ret.Fwd_egress_bytes_udp = in["fwd_egress_bytes_udp"].(int)
		ret.Rev_ingress_bytes_udp = in["rev_ingress_bytes_udp"].(int)
		ret.Rev_egress_bytes_udp = in["rev_egress_bytes_udp"].(int)
		ret.Fwd_ingress_packets_icmp = in["fwd_ingress_packets_icmp"].(int)
		ret.Fwd_egress_packets_icmp = in["fwd_egress_packets_icmp"].(int)
		ret.Rev_ingress_packets_icmp = in["rev_ingress_packets_icmp"].(int)
		ret.Rev_egress_packets_icmp = in["rev_egress_packets_icmp"].(int)
		ret.Fwd_ingress_bytes_icmp = in["fwd_ingress_bytes_icmp"].(int)
		ret.Fwd_egress_bytes_icmp = in["fwd_egress_bytes_icmp"].(int)
		ret.Rev_ingress_bytes_icmp = in["rev_ingress_bytes_icmp"].(int)
		ret.Rev_egress_bytes_icmp = in["rev_egress_bytes_icmp"].(int)
		ret.Fwd_ingress_packets_others = in["fwd_ingress_packets_others"].(int)
		ret.Fwd_egress_packets_others = in["fwd_egress_packets_others"].(int)
		ret.Rev_ingress_packets_others = in["rev_ingress_packets_others"].(int)
		ret.Rev_egress_packets_others = in["rev_egress_packets_others"].(int)
		ret.Fwd_ingress_bytes_others = in["fwd_ingress_bytes_others"].(int)
		ret.Fwd_egress_bytes_others = in["fwd_egress_bytes_others"].(int)
		ret.Rev_ingress_bytes_others = in["rev_ingress_bytes_others"].(int)
		ret.Rev_egress_bytes_others = in["rev_egress_bytes_others"].(int)
		ret.Fwd_ingress_pkt_size_range1 = in["fwd_ingress_pkt_size_range1"].(int)
		ret.Fwd_ingress_pkt_size_range2 = in["fwd_ingress_pkt_size_range2"].(int)
		ret.Fwd_ingress_pkt_size_range3 = in["fwd_ingress_pkt_size_range3"].(int)
		ret.Fwd_ingress_pkt_size_range4 = in["fwd_ingress_pkt_size_range4"].(int)
		ret.Fwd_egress_pkt_size_range1 = in["fwd_egress_pkt_size_range1"].(int)
		ret.Fwd_egress_pkt_size_range2 = in["fwd_egress_pkt_size_range2"].(int)
		ret.Fwd_egress_pkt_size_range3 = in["fwd_egress_pkt_size_range3"].(int)
		ret.Fwd_egress_pkt_size_range4 = in["fwd_egress_pkt_size_range4"].(int)
		ret.Rev_ingress_pkt_size_range1 = in["rev_ingress_pkt_size_range1"].(int)
		ret.Rev_ingress_pkt_size_range2 = in["rev_ingress_pkt_size_range2"].(int)
		ret.Rev_ingress_pkt_size_range3 = in["rev_ingress_pkt_size_range3"].(int)
		ret.Rev_ingress_pkt_size_range4 = in["rev_ingress_pkt_size_range4"].(int)
		ret.Rev_egress_pkt_size_range1 = in["rev_egress_pkt_size_range1"].(int)
		ret.Rev_egress_pkt_size_range2 = in["rev_egress_pkt_size_range2"].(int)
		ret.Rev_egress_pkt_size_range3 = in["rev_egress_pkt_size_range3"].(int)
		ret.Rev_egress_pkt_size_range4 = in["rev_egress_pkt_size_range4"].(int)
	}
	return ret
}

func dataToEndpointFwGlobalStats(d *schema.ResourceData) edpt.FwGlobalStats {
	var ret edpt.FwGlobalStats

	ret.Stats = getObjectFwGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
