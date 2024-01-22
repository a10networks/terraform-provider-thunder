package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LoggingStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_logging_stats`: Statistics for the object logging\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LoggingStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Created",
						},
						"tcp_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Deleted",
						},
						"tcp_port_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Allocated",
						},
						"tcp_port_freed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Freed",
						},
						"tcp_port_batch_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Batch Allocated",
						},
						"tcp_port_batch_freed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Batch Freed",
						},
						"udp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Session Created",
						},
						"udp_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Session Deleted",
						},
						"udp_port_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Allocated",
						},
						"udp_port_freed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Freed",
						},
						"udp_port_batch_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Batch Allocated",
						},
						"udp_port_batch_freed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Batch Freed",
						},
						"icmp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Session Created",
						},
						"icmp_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Session Deleted",
						},
						"icmp_resource_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Resource Allocated",
						},
						"icmp_resource_freed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Resource Freed",
						},
						"icmpv6_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPV6 Session Created",
						},
						"icmpv6_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPV6 Session Deleted",
						},
						"icmpv6_resource_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPV6 Resource Allocated",
						},
						"icmpv6_resource_freed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPV6 Resource Freed",
						},
						"gre_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Session Created",
						},
						"gre_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Session Deleted",
						},
						"gre_resource_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Resource Allocated",
						},
						"gre_resource_freed": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Resource Freed",
						},
						"esp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "ESP Session Created",
						},
						"esp_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "ESP Session Deleted",
						},
						"esp_resource_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "ESP Resource Allocated",
						},
						"esp_resource_freed": {
							Type: schema.TypeInt, Optional: true, Description: "ESP Resource Freed",
						},
						"fixed_nat_user_ports": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT Inside User Port Mapping",
						},
						"fixed_nat_disable_config_logged": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT Periodic Configs Logged",
						},
						"fixed_nat_disable_config_logs_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT Periodic Config Logs Sent",
						},
						"fixed_nat_periodic_config_logs_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT Disabled Configs Logged",
						},
						"fixed_nat_periodic_config_logged": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT Disabled Config Logs Sent",
						},
						"fixed_nat_interim_updated": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT Interim Updated",
						},
						"enhanced_user_log": {
							Type: schema.TypeInt, Optional: true, Description: "Enhanced User Log",
						},
						"log_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Log Packets Sent",
						},
						"log_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Log Packets Dropped",
						},
						"conn_tcp_established": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Connection Established",
						},
						"conn_tcp_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Connection Lost",
						},
						"tcp_port_overloading_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Overloading Allocated",
						},
						"tcp_port_overloading_freed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Overloading Freed",
						},
						"udp_port_overloading_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Overloading Allocated",
						},
						"udp_port_overloading_freed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Overloading Freed",
						},
						"http_request_logged": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Request Logged",
						},
						"reduced_logs_by_destination": {
							Type: schema.TypeInt, Optional: true, Description: "Reduced Logs by Destination Protocol and Port",
						},
						"interim_update_scheduled": {
							Type: schema.TypeInt, Optional: true, Description: "Port Allocation Interim Update Scheduled",
						},
						"tcp_port_batch_interim_updated": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Batch Interim Updated",
						},
						"udp_port_batch_interim_updated": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Batch Interim Updated",
						},
						"iddos_l3_entry_create": {
							Type: schema.TypeInt, Optional: true, Description: "iDDoS L3 Entry Created",
						},
						"iddos_l3_entry_delete": {
							Type: schema.TypeInt, Optional: true, Description: "iDDoS L3 Entry Deleted",
						},
						"iddos_l4_entry_create": {
							Type: schema.TypeInt, Optional: true, Description: "iDDoS L4 Entry Created",
						},
						"iddos_l4_entry_delete": {
							Type: schema.TypeInt, Optional: true, Description: "iDDoS L4 Entry Deleted",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LoggingStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LoggingStatsStats := setObjectCgnv6LoggingStatsStats(res)
		d.Set("stats", Cgnv6LoggingStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LoggingStatsStats(ret edpt.DataCgnv6LoggingStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tcp_session_created":                 ret.DtCgnv6LoggingStats.Stats.TcpSessionCreated,
			"tcp_session_deleted":                 ret.DtCgnv6LoggingStats.Stats.TcpSessionDeleted,
			"tcp_port_allocated":                  ret.DtCgnv6LoggingStats.Stats.TcpPortAllocated,
			"tcp_port_freed":                      ret.DtCgnv6LoggingStats.Stats.TcpPortFreed,
			"tcp_port_batch_allocated":            ret.DtCgnv6LoggingStats.Stats.TcpPortBatchAllocated,
			"tcp_port_batch_freed":                ret.DtCgnv6LoggingStats.Stats.TcpPortBatchFreed,
			"udp_session_created":                 ret.DtCgnv6LoggingStats.Stats.UdpSessionCreated,
			"udp_session_deleted":                 ret.DtCgnv6LoggingStats.Stats.UdpSessionDeleted,
			"udp_port_allocated":                  ret.DtCgnv6LoggingStats.Stats.UdpPortAllocated,
			"udp_port_freed":                      ret.DtCgnv6LoggingStats.Stats.UdpPortFreed,
			"udp_port_batch_allocated":            ret.DtCgnv6LoggingStats.Stats.UdpPortBatchAllocated,
			"udp_port_batch_freed":                ret.DtCgnv6LoggingStats.Stats.UdpPortBatchFreed,
			"icmp_session_created":                ret.DtCgnv6LoggingStats.Stats.IcmpSessionCreated,
			"icmp_session_deleted":                ret.DtCgnv6LoggingStats.Stats.IcmpSessionDeleted,
			"icmp_resource_allocated":             ret.DtCgnv6LoggingStats.Stats.IcmpResourceAllocated,
			"icmp_resource_freed":                 ret.DtCgnv6LoggingStats.Stats.IcmpResourceFreed,
			"icmpv6_session_created":              ret.DtCgnv6LoggingStats.Stats.Icmpv6SessionCreated,
			"icmpv6_session_deleted":              ret.DtCgnv6LoggingStats.Stats.Icmpv6SessionDeleted,
			"icmpv6_resource_allocated":           ret.DtCgnv6LoggingStats.Stats.Icmpv6ResourceAllocated,
			"icmpv6_resource_freed":               ret.DtCgnv6LoggingStats.Stats.Icmpv6ResourceFreed,
			"gre_session_created":                 ret.DtCgnv6LoggingStats.Stats.GreSessionCreated,
			"gre_session_deleted":                 ret.DtCgnv6LoggingStats.Stats.GreSessionDeleted,
			"gre_resource_allocated":              ret.DtCgnv6LoggingStats.Stats.GreResourceAllocated,
			"gre_resource_freed":                  ret.DtCgnv6LoggingStats.Stats.GreResourceFreed,
			"esp_session_created":                 ret.DtCgnv6LoggingStats.Stats.EspSessionCreated,
			"esp_session_deleted":                 ret.DtCgnv6LoggingStats.Stats.EspSessionDeleted,
			"esp_resource_allocated":              ret.DtCgnv6LoggingStats.Stats.EspResourceAllocated,
			"esp_resource_freed":                  ret.DtCgnv6LoggingStats.Stats.EspResourceFreed,
			"fixed_nat_user_ports":                ret.DtCgnv6LoggingStats.Stats.FixedNatUserPorts,
			"fixed_nat_disable_config_logged":     ret.DtCgnv6LoggingStats.Stats.FixedNatDisableConfigLogged,
			"fixed_nat_disable_config_logs_sent":  ret.DtCgnv6LoggingStats.Stats.FixedNatDisableConfigLogsSent,
			"fixed_nat_periodic_config_logs_sent": ret.DtCgnv6LoggingStats.Stats.FixedNatPeriodicConfigLogsSent,
			"fixed_nat_periodic_config_logged":    ret.DtCgnv6LoggingStats.Stats.FixedNatPeriodicConfigLogged,
			"fixed_nat_interim_updated":           ret.DtCgnv6LoggingStats.Stats.FixedNatInterimUpdated,
			"enhanced_user_log":                   ret.DtCgnv6LoggingStats.Stats.EnhancedUserLog,
			"log_sent":                            ret.DtCgnv6LoggingStats.Stats.LogSent,
			"log_dropped":                         ret.DtCgnv6LoggingStats.Stats.LogDropped,
			"conn_tcp_established":                ret.DtCgnv6LoggingStats.Stats.ConnTcpEstablished,
			"conn_tcp_dropped":                    ret.DtCgnv6LoggingStats.Stats.ConnTcpDropped,
			"tcp_port_overloading_allocated":      ret.DtCgnv6LoggingStats.Stats.TcpPortOverloadingAllocated,
			"tcp_port_overloading_freed":          ret.DtCgnv6LoggingStats.Stats.TcpPortOverloadingFreed,
			"udp_port_overloading_allocated":      ret.DtCgnv6LoggingStats.Stats.UdpPortOverloadingAllocated,
			"udp_port_overloading_freed":          ret.DtCgnv6LoggingStats.Stats.UdpPortOverloadingFreed,
			"http_request_logged":                 ret.DtCgnv6LoggingStats.Stats.HttpRequestLogged,
			"reduced_logs_by_destination":         ret.DtCgnv6LoggingStats.Stats.ReducedLogsByDestination,
			"interim_update_scheduled":            ret.DtCgnv6LoggingStats.Stats.InterimUpdateScheduled,
			"tcp_port_batch_interim_updated":      ret.DtCgnv6LoggingStats.Stats.TcpPortBatchInterimUpdated,
			"udp_port_batch_interim_updated":      ret.DtCgnv6LoggingStats.Stats.UdpPortBatchInterimUpdated,
			"iddos_l3_entry_create":               ret.DtCgnv6LoggingStats.Stats.IddosL3EntryCreate,
			"iddos_l3_entry_delete":               ret.DtCgnv6LoggingStats.Stats.IddosL3EntryDelete,
			"iddos_l4_entry_create":               ret.DtCgnv6LoggingStats.Stats.IddosL4EntryCreate,
			"iddos_l4_entry_delete":               ret.DtCgnv6LoggingStats.Stats.IddosL4EntryDelete,
		},
	}
}

func getObjectCgnv6LoggingStatsStats(d []interface{}) edpt.Cgnv6LoggingStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6LoggingStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpSessionCreated = in["tcp_session_created"].(int)
		ret.TcpSessionDeleted = in["tcp_session_deleted"].(int)
		ret.TcpPortAllocated = in["tcp_port_allocated"].(int)
		ret.TcpPortFreed = in["tcp_port_freed"].(int)
		ret.TcpPortBatchAllocated = in["tcp_port_batch_allocated"].(int)
		ret.TcpPortBatchFreed = in["tcp_port_batch_freed"].(int)
		ret.UdpSessionCreated = in["udp_session_created"].(int)
		ret.UdpSessionDeleted = in["udp_session_deleted"].(int)
		ret.UdpPortAllocated = in["udp_port_allocated"].(int)
		ret.UdpPortFreed = in["udp_port_freed"].(int)
		ret.UdpPortBatchAllocated = in["udp_port_batch_allocated"].(int)
		ret.UdpPortBatchFreed = in["udp_port_batch_freed"].(int)
		ret.IcmpSessionCreated = in["icmp_session_created"].(int)
		ret.IcmpSessionDeleted = in["icmp_session_deleted"].(int)
		ret.IcmpResourceAllocated = in["icmp_resource_allocated"].(int)
		ret.IcmpResourceFreed = in["icmp_resource_freed"].(int)
		ret.Icmpv6SessionCreated = in["icmpv6_session_created"].(int)
		ret.Icmpv6SessionDeleted = in["icmpv6_session_deleted"].(int)
		ret.Icmpv6ResourceAllocated = in["icmpv6_resource_allocated"].(int)
		ret.Icmpv6ResourceFreed = in["icmpv6_resource_freed"].(int)
		ret.GreSessionCreated = in["gre_session_created"].(int)
		ret.GreSessionDeleted = in["gre_session_deleted"].(int)
		ret.GreResourceAllocated = in["gre_resource_allocated"].(int)
		ret.GreResourceFreed = in["gre_resource_freed"].(int)
		ret.EspSessionCreated = in["esp_session_created"].(int)
		ret.EspSessionDeleted = in["esp_session_deleted"].(int)
		ret.EspResourceAllocated = in["esp_resource_allocated"].(int)
		ret.EspResourceFreed = in["esp_resource_freed"].(int)
		ret.FixedNatUserPorts = in["fixed_nat_user_ports"].(int)
		ret.FixedNatDisableConfigLogged = in["fixed_nat_disable_config_logged"].(int)
		ret.FixedNatDisableConfigLogsSent = in["fixed_nat_disable_config_logs_sent"].(int)
		ret.FixedNatPeriodicConfigLogsSent = in["fixed_nat_periodic_config_logs_sent"].(int)
		ret.FixedNatPeriodicConfigLogged = in["fixed_nat_periodic_config_logged"].(int)
		ret.FixedNatInterimUpdated = in["fixed_nat_interim_updated"].(int)
		ret.EnhancedUserLog = in["enhanced_user_log"].(int)
		ret.LogSent = in["log_sent"].(int)
		ret.LogDropped = in["log_dropped"].(int)
		ret.ConnTcpEstablished = in["conn_tcp_established"].(int)
		ret.ConnTcpDropped = in["conn_tcp_dropped"].(int)
		ret.TcpPortOverloadingAllocated = in["tcp_port_overloading_allocated"].(int)
		ret.TcpPortOverloadingFreed = in["tcp_port_overloading_freed"].(int)
		ret.UdpPortOverloadingAllocated = in["udp_port_overloading_allocated"].(int)
		ret.UdpPortOverloadingFreed = in["udp_port_overloading_freed"].(int)
		ret.HttpRequestLogged = in["http_request_logged"].(int)
		ret.ReducedLogsByDestination = in["reduced_logs_by_destination"].(int)
		ret.InterimUpdateScheduled = in["interim_update_scheduled"].(int)
		ret.TcpPortBatchInterimUpdated = in["tcp_port_batch_interim_updated"].(int)
		ret.UdpPortBatchInterimUpdated = in["udp_port_batch_interim_updated"].(int)
		ret.IddosL3EntryCreate = in["iddos_l3_entry_create"].(int)
		ret.IddosL3EntryDelete = in["iddos_l3_entry_delete"].(int)
		ret.IddosL4EntryCreate = in["iddos_l4_entry_create"].(int)
		ret.IddosL4EntryDelete = in["iddos_l4_entry_delete"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LoggingStats(d *schema.ResourceData) edpt.Cgnv6LoggingStats {
	var ret edpt.Cgnv6LoggingStats

	ret.Stats = getObjectCgnv6LoggingStatsStats(d.Get("stats").([]interface{}))
	return ret
}
