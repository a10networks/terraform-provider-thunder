package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_stateful_firewall_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6StatefulFirewallGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_packet_process": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Packet Process",
						},
						"udp_packet_process": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Packet Process",
						},
						"other_packet_process": {
							Type: schema.TypeInt, Optional: true, Description: "Other Packet Process",
						},
						"packet_inbound_deny": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Packet Denied",
						},
						"packet_process_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Error Drop",
						},
						"outbound_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Session Created",
						},
						"outbound_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Session Freed",
						},
						"inbound_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Session Created",
						},
						"inbound_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound Session Freed",
						},
						"tcp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Created",
						},
						"tcp_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Freed",
						},
						"udp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Session Created",
						},
						"udp_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Session Freed",
						},
						"other_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Other Session Created",
						},
						"other_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Other Session Freed",
						},
						"session_creation_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Session Creation Failure",
						},
						"no_fwd_route": {
							Type: schema.TypeInt, Optional: true, Description: "No Forward Route",
						},
						"no_rev_route": {
							Type: schema.TypeInt, Optional: true, Description: "No Reverse Route",
						},
						"packet_standby_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Standby Drop",
						},
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
						"eif_process": {
							Type: schema.TypeInt, Optional: true, Description: "Endpnt-Independent Filter Matched",
						},
						"one_arm_drop": {
							Type: schema.TypeInt, Optional: true, Description: "One-Arm Drop",
						},
						"no_class_list_match": {
							Type: schema.TypeInt, Optional: true, Description: "No Class-List Match Drop",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6StatefulFirewallGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6StatefulFirewallGlobalStatsStats := setObjectCgnv6StatefulFirewallGlobalStatsStats(res)
		d.Set("stats", Cgnv6StatefulFirewallGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6StatefulFirewallGlobalStatsStats(ret edpt.DataCgnv6StatefulFirewallGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tcp_packet_process":        ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Tcp_packet_process,
			"udp_packet_process":        ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Udp_packet_process,
			"other_packet_process":      ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Other_packet_process,
			"packet_inbound_deny":       ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Packet_inbound_deny,
			"packet_process_failure":    ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Packet_process_failure,
			"outbound_session_created":  ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Outbound_session_created,
			"outbound_session_freed":    ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Outbound_session_freed,
			"inbound_session_created":   ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Inbound_session_created,
			"inbound_session_freed":     ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Inbound_session_freed,
			"tcp_session_created":       ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Tcp_session_created,
			"tcp_session_freed":         ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Tcp_session_freed,
			"udp_session_created":       ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Udp_session_created,
			"udp_session_freed":         ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Udp_session_freed,
			"other_session_created":     ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Other_session_created,
			"other_session_freed":       ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Other_session_freed,
			"session_creation_failure":  ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Session_creation_failure,
			"no_fwd_route":              ret.DtCgnv6StatefulFirewallGlobalStats.Stats.No_fwd_route,
			"no_rev_route":              ret.DtCgnv6StatefulFirewallGlobalStats.Stats.No_rev_route,
			"packet_standby_drop":       ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Packet_standby_drop,
			"tcp_fullcone_created":      ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Tcp_fullcone_created,
			"tcp_fullcone_freed":        ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Tcp_fullcone_freed,
			"udp_fullcone_created":      ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Udp_fullcone_created,
			"udp_fullcone_freed":        ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Udp_fullcone_freed,
			"fullcone_creation_failure": ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Fullcone_creation_failure,
			"eif_process":               ret.DtCgnv6StatefulFirewallGlobalStats.Stats.Eif_process,
			"one_arm_drop":              ret.DtCgnv6StatefulFirewallGlobalStats.Stats.One_arm_drop,
			"no_class_list_match":       ret.DtCgnv6StatefulFirewallGlobalStats.Stats.No_class_list_match,
		},
	}
}

func getObjectCgnv6StatefulFirewallGlobalStatsStats(d []interface{}) edpt.Cgnv6StatefulFirewallGlobalStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6StatefulFirewallGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp_packet_process = in["tcp_packet_process"].(int)
		ret.Udp_packet_process = in["udp_packet_process"].(int)
		ret.Other_packet_process = in["other_packet_process"].(int)
		ret.Packet_inbound_deny = in["packet_inbound_deny"].(int)
		ret.Packet_process_failure = in["packet_process_failure"].(int)
		ret.Outbound_session_created = in["outbound_session_created"].(int)
		ret.Outbound_session_freed = in["outbound_session_freed"].(int)
		ret.Inbound_session_created = in["inbound_session_created"].(int)
		ret.Inbound_session_freed = in["inbound_session_freed"].(int)
		ret.Tcp_session_created = in["tcp_session_created"].(int)
		ret.Tcp_session_freed = in["tcp_session_freed"].(int)
		ret.Udp_session_created = in["udp_session_created"].(int)
		ret.Udp_session_freed = in["udp_session_freed"].(int)
		ret.Other_session_created = in["other_session_created"].(int)
		ret.Other_session_freed = in["other_session_freed"].(int)
		ret.Session_creation_failure = in["session_creation_failure"].(int)
		ret.No_fwd_route = in["no_fwd_route"].(int)
		ret.No_rev_route = in["no_rev_route"].(int)
		ret.Packet_standby_drop = in["packet_standby_drop"].(int)
		ret.Tcp_fullcone_created = in["tcp_fullcone_created"].(int)
		ret.Tcp_fullcone_freed = in["tcp_fullcone_freed"].(int)
		ret.Udp_fullcone_created = in["udp_fullcone_created"].(int)
		ret.Udp_fullcone_freed = in["udp_fullcone_freed"].(int)
		ret.Fullcone_creation_failure = in["fullcone_creation_failure"].(int)
		ret.Eif_process = in["eif_process"].(int)
		ret.One_arm_drop = in["one_arm_drop"].(int)
		ret.No_class_list_match = in["no_class_list_match"].(int)
	}
	return ret
}

func dataToEndpointCgnv6StatefulFirewallGlobalStats(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallGlobalStats {
	var ret edpt.Cgnv6StatefulFirewallGlobalStats

	ret.Stats = getObjectCgnv6StatefulFirewallGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
