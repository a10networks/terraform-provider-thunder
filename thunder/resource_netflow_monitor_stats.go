package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowMonitorStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_netflow_monitor_stats`: Statistics for the object monitor\n\n__PLACEHOLDER__",
		ReadContext: resourceNetflowMonitorStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name of netflow monitor",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packets_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sent Packets Count",
						},
						"bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sent Bytes Count",
						},
						"nat44_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 Flow Records Sent",
						},
						"nat44_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 Flow Records Failed",
						},
						"nat64_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Flow Records Sent",
						},
						"nat64_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Flow Records Failed",
						},
						"dslite_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Dslite Flow Records Sent",
						},
						"dslite_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Dslite Flow Records Failed",
						},
						"session_event_nat44_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Nat44 Session Event Records Sent",
						},
						"session_event_nat44_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Nat44 Session Event Records Failed",
						},
						"session_event_nat64_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Nat64 Session Event Records Sent",
						},
						"session_event_nat64_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Nat64 Session Event Records Falied",
						},
						"session_event_dslite_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Dslite Session Event Records Sent",
						},
						"session_event_dslite_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Dslite Session Event Records Failed",
						},
						"session_event_fw4_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "FW4 Session Event Records Sent",
						},
						"session_event_fw4_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "FW4 Session Event Records Failed",
						},
						"session_event_fw6_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "FW6 Session Event Records Sent",
						},
						"session_event_fw6_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "FW6 Session Event Records Failed",
						},
						"port_mapping_nat44_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Port Mapping Nat44 Event Records Sent",
						},
						"port_mapping_nat44_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Port Mapping Nat44 Event Records Failed",
						},
						"port_mapping_nat64_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Port Mapping Nat64 Event Records Sent",
						},
						"port_mapping_nat64_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Port Mapping Nat64 Event Records Failed",
						},
						"port_mapping_dslite_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Port Mapping Dslite Event Records Sent",
						},
						"port_mapping_dslite_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Port Mapping Dslite Event Records failed",
						},
						"netflow_v5_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v5 Records Sent",
						},
						"netflow_v5_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v5 Records Failed",
						},
						"netflow_v5_ext_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v5 Ext Records Sent",
						},
						"netflow_v5_ext_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v5 Ext Records Failed",
						},
						"port_batching_nat44_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Port Batching Nat44 Records Sent",
						},
						"port_batching_nat44_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Port Batching Nat44 Records Failed",
						},
						"port_batching_nat64_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Port Batching Nat64 Records Sent",
						},
						"port_batching_nat64_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Port Batching Nat64 Records Failed",
						},
						"port_batching_dslite_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Port Batching Dslite Records Sent",
						},
						"port_batching_dslite_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Port Batching Dslite Records Failed",
						},
						"port_batching_v2_nat44_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Port Batching V2 Nat44 Records Sent",
						},
						"port_batching_v2_nat44_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Port Batching V2 Nat44 Records Failed",
						},
						"port_batching_v2_nat64_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Port Batching V2 Nat64 Records Sent",
						},
						"port_batching_v2_nat64_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Port Batching V2 Nat64 Records Failed",
						},
						"port_batching_v2_dslite_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Port Batching V2 Dslite Records Sent",
						},
						"port_batching_v2_dslite_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Port Batching V2 Dslite Records Falied",
						},
						"custom_session_event_nat44_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Session Creation Records Sent",
						},
						"custom_session_event_nat44_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Session Creation Records Failed",
						},
						"custom_session_event_nat64_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Session Creation Records Sent",
						},
						"custom_session_event_nat64_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Session Creation Records Failed",
						},
						"custom_session_event_dslite_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Session Creation Records Sent",
						},
						"custom_session_event_dslite_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Session Creation Records Failed",
						},
						"custom_session_event_nat44_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Session Deletion Records Sent",
						},
						"custom_session_event_nat44_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Session Deletion Records Failed",
						},
						"custom_session_event_nat64_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Session Deletion Records Sent",
						},
						"custom_session_event_nat64_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Session Deletion Records Failed",
						},
						"custom_session_event_dslite_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Session Deletion Records Sent",
						},
						"custom_session_event_dslite_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Session Deletion Records Failed",
						},
						"custom_session_event_fw4_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW4 Session Creation Records Sent",
						},
						"custom_session_event_fw4_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW4 Session Creation Records Failed",
						},
						"custom_session_event_fw6_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW6 Session Creation Records Sent",
						},
						"custom_session_event_fw6_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW6 Session Creation Records Failed",
						},
						"custom_session_event_fw4_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW4 Session Deletion Records Sent",
						},
						"custom_session_event_fw4_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW4 Session Deletion Records Failed",
						},
						"custom_session_event_fw6_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW6 Session Deletion Records Sent",
						},
						"custom_session_event_fw6_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW6 Session Deletion Records Failed",
						},
						"custom_deny_reset_event_fw4_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW4 Deny/Reset Event Records Sent",
						},
						"custom_deny_reset_event_fw4_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW4 Deny/Reset Event Records Failed",
						},
						"custom_deny_reset_event_fw6_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW6 Deny/Reset Event Records Sent",
						},
						"custom_deny_reset_event_fw6_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW6 Deny/Reset Event Records Failed",
						},
						"custom_port_mapping_nat44_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Port Map Creation Records Sent",
						},
						"custom_port_mapping_nat44_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Port Map Creation Records Failed",
						},
						"custom_port_mapping_nat64_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Port Map Creation Records Sent",
						},
						"custom_port_mapping_nat64_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Port Map Creation Records Failed",
						},
						"custom_port_mapping_dslite_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Port Map Creation Records Sent",
						},
						"custom_port_mapping_dslite_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Port Map Creation Records Failed",
						},
						"custom_port_mapping_nat44_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Port Map Deletion Records Sent",
						},
						"custom_port_mapping_nat44_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Port Map Deletion Records Failed",
						},
						"custom_port_mapping_nat64_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Port Map Deletion Records Sent",
						},
						"custom_port_mapping_nat64_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Port Map Deletion Records Failed",
						},
						"custom_port_mapping_dslite_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Port Map Deletion Records Sent",
						},
						"custom_port_mapping_dslite_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Port Map Deletion Records Failed",
						},
						"custom_port_batching_nat44_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Port Batch Creation Records Sent",
						},
						"custom_port_batching_nat44_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Port Batch Creation Records Failed",
						},
						"custom_port_batching_nat64_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Port Batch Creation Records Sent",
						},
						"custom_port_batching_nat64_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Port Batch Creation Records Failed",
						},
						"custom_port_batching_dslite_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Port Batch Creation Records Sent",
						},
						"custom_port_batching_dslite_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Port Batch Creation Records Failed",
						},
						"custom_port_batching_nat44_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Port Batch Deletion Records Sent",
						},
						"custom_port_batching_nat44_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Port Batch Deletion Records Failed",
						},
						"custom_port_batching_nat64_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Port Batch Deletion Records Sent",
						},
						"custom_port_batching_nat64_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Port Batch Deletion Records Failed",
						},
						"custom_port_batching_dslite_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Port Batch Deletion Records Sent",
						},
						"custom_port_batching_dslite_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Port Batch Deletion Records Failed",
						},
						"custom_port_batching_v2_nat44_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Port Batch V2 Creation Records Sent",
						},
						"custom_port_batching_v2_nat44_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Port Batch V2 Creation Records Failed",
						},
						"custom_port_batching_v2_nat64_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Port Batch V2 Creation Records Sent",
						},
						"custom_port_batching_v2_nat64_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Port Batch V2 Creation Records Failed",
						},
						"custom_port_batching_v2_dslite_creation_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Port Batch V2 Creation Records Sent",
						},
						"custom_port_batching_v2_dslite_creation_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Port Batch V2 Creation Records Failed",
						},
						"custom_port_batching_v2_nat44_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Port Batch V2 Deletion Records Sent",
						},
						"custom_port_batching_v2_nat44_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat44 Port Batch V2 Deletion Records Failed",
						},
						"custom_port_batching_v2_nat64_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Port Batch V2 Deletion Records Sent",
						},
						"custom_port_batching_v2_nat64_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Nat64 Port Batch V2 Deletion Records Failed",
						},
						"custom_port_batching_v2_dslite_deletion_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Port Batch V2 Deletion Records Sent",
						},
						"custom_port_batching_v2_dslite_deletion_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom Dslite Port Batch V2 Deletion Records Failed",
						},
						"reduced_logs_by_destination": {
							Type: schema.TypeInt, Optional: true, Description: "Reduced Logs by Destination Protocol and Port",
						},
						"custom_gtp_c_tunnel_event_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom GTP C Tunnel Records Sent",
						},
						"custom_gtp_c_tunnel_event_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom GTP C Tunnel Records Sent Failure",
						},
						"custom_gtp_u_tunnel_event_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom GTP U Tunnel Records Sent",
						},
						"custom_gtp_u_tunnel_event_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom GTP U Tunnel Records Sent Failure",
						},
						"custom_gtp_deny_event_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom GTP Deny Records Sent",
						},
						"custom_gtp_deny_event_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom GTP Deny Records Sent Failure",
						},
						"custom_gtp_info_event_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom GTP Info Records Sent",
						},
						"custom_gtp_info_event_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom GTP Info Records Sent Failure",
						},
						"custom_fw_iddos_entry_created_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW iDDoS Entry Created Records Sent",
						},
						"custom_fw_iddos_entry_created_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW iDDoS Entry Created Records Sent Failure",
						},
						"custom_fw_iddos_entry_deleted_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW iDDoS Entry Deleted Records Sent",
						},
						"custom_fw_iddos_entry_deleted_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW iDDoS Entry Deleted Records Sent Failure",
						},
						"custom_fw_sesn_limit_exceeded_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW Session Limit Exceeded Records Sent",
						},
						"custom_fw_sesn_limit_exceeded_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom FW Session Limit Exceeded Records Sent Failure",
						},
						"custom_nat_iddos_l3_entry_created_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom NAT iDDoS L3 Entry Created Records Sent",
						},
						"custom_nat_iddos_l3_entry_created_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom NAT iDDoS L3 Entry Created Records Sent Failure",
						},
						"custom_nat_iddos_l3_entry_deleted_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom NAT iDDoS L3 Entry Deleted Records Sent",
						},
						"custom_nat_iddos_l3_entry_deleted_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom NAT iDDoS L3 Entry Deleted Records Sent Failure",
						},
						"custom_nat_iddos_l4_entry_created_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom NAT iDDoS L4 Entry Created Records Sent",
						},
						"custom_nat_iddos_l4_entry_created_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom NAT iDDoS L4 Entry Created Records Sent Failure",
						},
						"custom_nat_iddos_l4_entry_deleted_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom NAT iDDoS L4 Entry Deleted Records Sent",
						},
						"custom_nat_iddos_l4_entry_deleted_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom NAT iDDoS L4 Entry Deleted Records Sent Failure",
						},
						"custom_gtp_rate_limit_periodic_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Custom GTP Rate Limit Periodic Records Sent",
						},
						"custom_gtp_rate_limit_periodic_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Custom GTP Rate Limit Periodic Records Sent Failure",
						},
						"ddos_general_stat_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ddos_general_stat_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ddos_http_stat_records_sent": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ddos_http_stat_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceNetflowMonitorStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetflowMonitorStatsStats := setObjectNetflowMonitorStatsStats(res)
		d.Set("stats", NetflowMonitorStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetflowMonitorStatsStats(ret edpt.DataNetflowMonitorStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packets_sent":                                                 ret.DtNetflowMonitorStats.Stats.PacketsSent,
			"bytes_sent":                                                   ret.DtNetflowMonitorStats.Stats.BytesSent,
			"nat44_records_sent":                                           ret.DtNetflowMonitorStats.Stats.Nat44RecordsSent,
			"nat44_records_sent_failure":                                   ret.DtNetflowMonitorStats.Stats.Nat44RecordsSentFailure,
			"nat64_records_sent":                                           ret.DtNetflowMonitorStats.Stats.Nat64RecordsSent,
			"nat64_records_sent_failure":                                   ret.DtNetflowMonitorStats.Stats.Nat64RecordsSentFailure,
			"dslite_records_sent":                                          ret.DtNetflowMonitorStats.Stats.DsliteRecordsSent,
			"dslite_records_sent_failure":                                  ret.DtNetflowMonitorStats.Stats.DsliteRecordsSentFailure,
			"session_event_nat44_records_sent":                             ret.DtNetflowMonitorStats.Stats.SessionEventNat44RecordsSent,
			"session_event_nat44_records_sent_failure":                     ret.DtNetflowMonitorStats.Stats.SessionEventNat44RecordsSentFailure,
			"session_event_nat64_records_sent":                             ret.DtNetflowMonitorStats.Stats.SessionEventNat64RecordsSent,
			"session_event_nat64_records_sent_failure":                     ret.DtNetflowMonitorStats.Stats.SessionEventNat64RecordsSentFailure,
			"session_event_dslite_records_sent":                            ret.DtNetflowMonitorStats.Stats.SessionEventDsliteRecordsSent,
			"session_event_dslite_records_sent_failure":                    ret.DtNetflowMonitorStats.Stats.SessionEventDsliteRecordsSentFailure,
			"session_event_fw4_records_sent":                               ret.DtNetflowMonitorStats.Stats.SessionEventFw4RecordsSent,
			"session_event_fw4_records_sent_failure":                       ret.DtNetflowMonitorStats.Stats.SessionEventFw4RecordsSentFailure,
			"session_event_fw6_records_sent":                               ret.DtNetflowMonitorStats.Stats.SessionEventFw6RecordsSent,
			"session_event_fw6_records_sent_failure":                       ret.DtNetflowMonitorStats.Stats.SessionEventFw6RecordsSentFailure,
			"port_mapping_nat44_records_sent":                              ret.DtNetflowMonitorStats.Stats.PortMappingNat44RecordsSent,
			"port_mapping_nat44_records_sent_failure":                      ret.DtNetflowMonitorStats.Stats.PortMappingNat44RecordsSentFailure,
			"port_mapping_nat64_records_sent":                              ret.DtNetflowMonitorStats.Stats.PortMappingNat64RecordsSent,
			"port_mapping_nat64_records_sent_failure":                      ret.DtNetflowMonitorStats.Stats.PortMappingNat64RecordsSentFailure,
			"port_mapping_dslite_records_sent":                             ret.DtNetflowMonitorStats.Stats.PortMappingDsliteRecordsSent,
			"port_mapping_dslite_records_sent_failure":                     ret.DtNetflowMonitorStats.Stats.PortMappingDsliteRecordsSentFailure,
			"netflow_v5_records_sent":                                      ret.DtNetflowMonitorStats.Stats.NetflowV5RecordsSent,
			"netflow_v5_records_sent_failure":                              ret.DtNetflowMonitorStats.Stats.NetflowV5RecordsSentFailure,
			"netflow_v5_ext_records_sent":                                  ret.DtNetflowMonitorStats.Stats.NetflowV5ExtRecordsSent,
			"netflow_v5_ext_records_sent_failure":                          ret.DtNetflowMonitorStats.Stats.NetflowV5ExtRecordsSentFailure,
			"port_batching_nat44_records_sent":                             ret.DtNetflowMonitorStats.Stats.PortBatchingNat44RecordsSent,
			"port_batching_nat44_records_sent_failure":                     ret.DtNetflowMonitorStats.Stats.PortBatchingNat44RecordsSentFailure,
			"port_batching_nat64_records_sent":                             ret.DtNetflowMonitorStats.Stats.PortBatchingNat64RecordsSent,
			"port_batching_nat64_records_sent_failure":                     ret.DtNetflowMonitorStats.Stats.PortBatchingNat64RecordsSentFailure,
			"port_batching_dslite_records_sent":                            ret.DtNetflowMonitorStats.Stats.PortBatchingDsliteRecordsSent,
			"port_batching_dslite_records_sent_failure":                    ret.DtNetflowMonitorStats.Stats.PortBatchingDsliteRecordsSentFailure,
			"port_batching_v2_nat44_records_sent":                          ret.DtNetflowMonitorStats.Stats.PortBatchingV2Nat44RecordsSent,
			"port_batching_v2_nat44_records_sent_failure":                  ret.DtNetflowMonitorStats.Stats.PortBatchingV2Nat44RecordsSentFailure,
			"port_batching_v2_nat64_records_sent":                          ret.DtNetflowMonitorStats.Stats.PortBatchingV2Nat64RecordsSent,
			"port_batching_v2_nat64_records_sent_failure":                  ret.DtNetflowMonitorStats.Stats.PortBatchingV2Nat64RecordsSentFailure,
			"port_batching_v2_dslite_records_sent":                         ret.DtNetflowMonitorStats.Stats.PortBatchingV2DsliteRecordsSent,
			"port_batching_v2_dslite_records_sent_failure":                 ret.DtNetflowMonitorStats.Stats.PortBatchingV2DsliteRecordsSentFailure,
			"custom_session_event_nat44_creation_records_sent":             ret.DtNetflowMonitorStats.Stats.CustomSessionEventNat44CreationRecordsSent,
			"custom_session_event_nat44_creation_records_sent_failure":     ret.DtNetflowMonitorStats.Stats.CustomSessionEventNat44CreationRecordsSentFailure,
			"custom_session_event_nat64_creation_records_sent":             ret.DtNetflowMonitorStats.Stats.CustomSessionEventNat64CreationRecordsSent,
			"custom_session_event_nat64_creation_records_sent_failure":     ret.DtNetflowMonitorStats.Stats.CustomSessionEventNat64CreationRecordsSentFailure,
			"custom_session_event_dslite_creation_records_sent":            ret.DtNetflowMonitorStats.Stats.CustomSessionEventDsliteCreationRecordsSent,
			"custom_session_event_dslite_creation_records_sent_failure":    ret.DtNetflowMonitorStats.Stats.CustomSessionEventDsliteCreationRecordsSentFailure,
			"custom_session_event_nat44_deletion_records_sent":             ret.DtNetflowMonitorStats.Stats.CustomSessionEventNat44DeletionRecordsSent,
			"custom_session_event_nat44_deletion_records_sent_failure":     ret.DtNetflowMonitorStats.Stats.CustomSessionEventNat44DeletionRecordsSentFailure,
			"custom_session_event_nat64_deletion_records_sent":             ret.DtNetflowMonitorStats.Stats.CustomSessionEventNat64DeletionRecordsSent,
			"custom_session_event_nat64_deletion_records_sent_failure":     ret.DtNetflowMonitorStats.Stats.CustomSessionEventNat64DeletionRecordsSentFailure,
			"custom_session_event_dslite_deletion_records_sent":            ret.DtNetflowMonitorStats.Stats.CustomSessionEventDsliteDeletionRecordsSent,
			"custom_session_event_dslite_deletion_records_sent_failure":    ret.DtNetflowMonitorStats.Stats.CustomSessionEventDsliteDeletionRecordsSentFailure,
			"custom_session_event_fw4_creation_records_sent":               ret.DtNetflowMonitorStats.Stats.CustomSessionEventFw4CreationRecordsSent,
			"custom_session_event_fw4_creation_records_sent_failure":       ret.DtNetflowMonitorStats.Stats.CustomSessionEventFw4CreationRecordsSentFailure,
			"custom_session_event_fw6_creation_records_sent":               ret.DtNetflowMonitorStats.Stats.CustomSessionEventFw6CreationRecordsSent,
			"custom_session_event_fw6_creation_records_sent_failure":       ret.DtNetflowMonitorStats.Stats.CustomSessionEventFw6CreationRecordsSentFailure,
			"custom_session_event_fw4_deletion_records_sent":               ret.DtNetflowMonitorStats.Stats.CustomSessionEventFw4DeletionRecordsSent,
			"custom_session_event_fw4_deletion_records_sent_failure":       ret.DtNetflowMonitorStats.Stats.CustomSessionEventFw4DeletionRecordsSentFailure,
			"custom_session_event_fw6_deletion_records_sent":               ret.DtNetflowMonitorStats.Stats.CustomSessionEventFw6DeletionRecordsSent,
			"custom_session_event_fw6_deletion_records_sent_failure":       ret.DtNetflowMonitorStats.Stats.CustomSessionEventFw6DeletionRecordsSentFailure,
			"custom_deny_reset_event_fw4_records_sent":                     ret.DtNetflowMonitorStats.Stats.CustomDenyResetEventFw4RecordsSent,
			"custom_deny_reset_event_fw4_records_sent_failure":             ret.DtNetflowMonitorStats.Stats.CustomDenyResetEventFw4RecordsSentFailure,
			"custom_deny_reset_event_fw6_records_sent":                     ret.DtNetflowMonitorStats.Stats.CustomDenyResetEventFw6RecordsSent,
			"custom_deny_reset_event_fw6_records_sent_failure":             ret.DtNetflowMonitorStats.Stats.CustomDenyResetEventFw6RecordsSentFailure,
			"custom_port_mapping_nat44_creation_records_sent":              ret.DtNetflowMonitorStats.Stats.CustomPortMappingNat44CreationRecordsSent,
			"custom_port_mapping_nat44_creation_records_sent_failure":      ret.DtNetflowMonitorStats.Stats.CustomPortMappingNat44CreationRecordsSentFailure,
			"custom_port_mapping_nat64_creation_records_sent":              ret.DtNetflowMonitorStats.Stats.CustomPortMappingNat64CreationRecordsSent,
			"custom_port_mapping_nat64_creation_records_sent_failure":      ret.DtNetflowMonitorStats.Stats.CustomPortMappingNat64CreationRecordsSentFailure,
			"custom_port_mapping_dslite_creation_records_sent":             ret.DtNetflowMonitorStats.Stats.CustomPortMappingDsliteCreationRecordsSent,
			"custom_port_mapping_dslite_creation_records_sent_failure":     ret.DtNetflowMonitorStats.Stats.CustomPortMappingDsliteCreationRecordsSentFailure,
			"custom_port_mapping_nat44_deletion_records_sent":              ret.DtNetflowMonitorStats.Stats.CustomPortMappingNat44DeletionRecordsSent,
			"custom_port_mapping_nat44_deletion_records_sent_failure":      ret.DtNetflowMonitorStats.Stats.CustomPortMappingNat44DeletionRecordsSentFailure,
			"custom_port_mapping_nat64_deletion_records_sent":              ret.DtNetflowMonitorStats.Stats.CustomPortMappingNat64DeletionRecordsSent,
			"custom_port_mapping_nat64_deletion_records_sent_failure":      ret.DtNetflowMonitorStats.Stats.CustomPortMappingNat64DeletionRecordsSentFailure,
			"custom_port_mapping_dslite_deletion_records_sent":             ret.DtNetflowMonitorStats.Stats.CustomPortMappingDsliteDeletionRecordsSent,
			"custom_port_mapping_dslite_deletion_records_sent_failure":     ret.DtNetflowMonitorStats.Stats.CustomPortMappingDsliteDeletionRecordsSentFailure,
			"custom_port_batching_nat44_creation_records_sent":             ret.DtNetflowMonitorStats.Stats.CustomPortBatchingNat44CreationRecordsSent,
			"custom_port_batching_nat44_creation_records_sent_failure":     ret.DtNetflowMonitorStats.Stats.CustomPortBatchingNat44CreationRecordsSentFailure,
			"custom_port_batching_nat64_creation_records_sent":             ret.DtNetflowMonitorStats.Stats.CustomPortBatchingNat64CreationRecordsSent,
			"custom_port_batching_nat64_creation_records_sent_failure":     ret.DtNetflowMonitorStats.Stats.CustomPortBatchingNat64CreationRecordsSentFailure,
			"custom_port_batching_dslite_creation_records_sent":            ret.DtNetflowMonitorStats.Stats.CustomPortBatchingDsliteCreationRecordsSent,
			"custom_port_batching_dslite_creation_records_sent_failure":    ret.DtNetflowMonitorStats.Stats.CustomPortBatchingDsliteCreationRecordsSentFailure,
			"custom_port_batching_nat44_deletion_records_sent":             ret.DtNetflowMonitorStats.Stats.CustomPortBatchingNat44DeletionRecordsSent,
			"custom_port_batching_nat44_deletion_records_sent_failure":     ret.DtNetflowMonitorStats.Stats.CustomPortBatchingNat44DeletionRecordsSentFailure,
			"custom_port_batching_nat64_deletion_records_sent":             ret.DtNetflowMonitorStats.Stats.CustomPortBatchingNat64DeletionRecordsSent,
			"custom_port_batching_nat64_deletion_records_sent_failure":     ret.DtNetflowMonitorStats.Stats.CustomPortBatchingNat64DeletionRecordsSentFailure,
			"custom_port_batching_dslite_deletion_records_sent":            ret.DtNetflowMonitorStats.Stats.CustomPortBatchingDsliteDeletionRecordsSent,
			"custom_port_batching_dslite_deletion_records_sent_failure":    ret.DtNetflowMonitorStats.Stats.CustomPortBatchingDsliteDeletionRecordsSentFailure,
			"custom_port_batching_v2_nat44_creation_records_sent":          ret.DtNetflowMonitorStats.Stats.CustomPortBatchingV2Nat44CreationRecordsSent,
			"custom_port_batching_v2_nat44_creation_records_sent_failure":  ret.DtNetflowMonitorStats.Stats.CustomPortBatchingV2Nat44CreationRecordsSentFailure,
			"custom_port_batching_v2_nat64_creation_records_sent":          ret.DtNetflowMonitorStats.Stats.CustomPortBatchingV2Nat64CreationRecordsSent,
			"custom_port_batching_v2_nat64_creation_records_sent_failure":  ret.DtNetflowMonitorStats.Stats.CustomPortBatchingV2Nat64CreationRecordsSentFailure,
			"custom_port_batching_v2_dslite_creation_records_sent":         ret.DtNetflowMonitorStats.Stats.CustomPortBatchingV2DsliteCreationRecordsSent,
			"custom_port_batching_v2_dslite_creation_records_sent_failure": ret.DtNetflowMonitorStats.Stats.CustomPortBatchingV2DsliteCreationRecordsSentFailure,
			"custom_port_batching_v2_nat44_deletion_records_sent":          ret.DtNetflowMonitorStats.Stats.CustomPortBatchingV2Nat44DeletionRecordsSent,
			"custom_port_batching_v2_nat44_deletion_records_sent_failure":  ret.DtNetflowMonitorStats.Stats.CustomPortBatchingV2Nat44DeletionRecordsSentFailure,
			"custom_port_batching_v2_nat64_deletion_records_sent":          ret.DtNetflowMonitorStats.Stats.CustomPortBatchingV2Nat64DeletionRecordsSent,
			"custom_port_batching_v2_nat64_deletion_records_sent_failure":  ret.DtNetflowMonitorStats.Stats.CustomPortBatchingV2Nat64DeletionRecordsSentFailure,
			"custom_port_batching_v2_dslite_deletion_records_sent":         ret.DtNetflowMonitorStats.Stats.CustomPortBatchingV2DsliteDeletionRecordsSent,
			"custom_port_batching_v2_dslite_deletion_records_sent_failure": ret.DtNetflowMonitorStats.Stats.CustomPortBatchingV2DsliteDeletionRecordsSentFailure,
			"reduced_logs_by_destination":                                  ret.DtNetflowMonitorStats.Stats.ReducedLogsByDestination,
			"custom_gtp_c_tunnel_event_records_sent":                       ret.DtNetflowMonitorStats.Stats.CustomGtpCTunnelEventRecordsSent,
			"custom_gtp_c_tunnel_event_records_sent_failure":               ret.DtNetflowMonitorStats.Stats.CustomGtpCTunnelEventRecordsSentFailure,
			"custom_gtp_u_tunnel_event_records_sent":                       ret.DtNetflowMonitorStats.Stats.CustomGtpUTunnelEventRecordsSent,
			"custom_gtp_u_tunnel_event_records_sent_failure":               ret.DtNetflowMonitorStats.Stats.CustomGtpUTunnelEventRecordsSentFailure,
			"custom_gtp_deny_event_records_sent":                           ret.DtNetflowMonitorStats.Stats.CustomGtpDenyEventRecordsSent,
			"custom_gtp_deny_event_records_sent_failure":                   ret.DtNetflowMonitorStats.Stats.CustomGtpDenyEventRecordsSentFailure,
			"custom_gtp_info_event_records_sent":                           ret.DtNetflowMonitorStats.Stats.CustomGtpInfoEventRecordsSent,
			"custom_gtp_info_event_records_sent_failure":                   ret.DtNetflowMonitorStats.Stats.CustomGtpInfoEventRecordsSentFailure,
			"custom_fw_iddos_entry_created_records_sent":                   ret.DtNetflowMonitorStats.Stats.CustomFwIddosEntryCreatedRecordsSent,
			"custom_fw_iddos_entry_created_records_sent_failure":           ret.DtNetflowMonitorStats.Stats.CustomFwIddosEntryCreatedRecordsSentFailure,
			"custom_fw_iddos_entry_deleted_records_sent":                   ret.DtNetflowMonitorStats.Stats.CustomFwIddosEntryDeletedRecordsSent,
			"custom_fw_iddos_entry_deleted_records_sent_failure":           ret.DtNetflowMonitorStats.Stats.CustomFwIddosEntryDeletedRecordsSentFailure,
			"custom_fw_sesn_limit_exceeded_records_sent":                   ret.DtNetflowMonitorStats.Stats.CustomFwSesnLimitExceededRecordsSent,
			"custom_fw_sesn_limit_exceeded_records_sent_failure":           ret.DtNetflowMonitorStats.Stats.CustomFwSesnLimitExceededRecordsSentFailure,
			"custom_nat_iddos_l3_entry_created_records_sent":               ret.DtNetflowMonitorStats.Stats.CustomNatIddosL3EntryCreatedRecordsSent,
			"custom_nat_iddos_l3_entry_created_records_sent_failure":       ret.DtNetflowMonitorStats.Stats.CustomNatIddosL3EntryCreatedRecordsSentFailure,
			"custom_nat_iddos_l3_entry_deleted_records_sent":               ret.DtNetflowMonitorStats.Stats.CustomNatIddosL3EntryDeletedRecordsSent,
			"custom_nat_iddos_l3_entry_deleted_records_sent_failure":       ret.DtNetflowMonitorStats.Stats.CustomNatIddosL3EntryDeletedRecordsSentFailure,
			"custom_nat_iddos_l4_entry_created_records_sent":               ret.DtNetflowMonitorStats.Stats.CustomNatIddosL4EntryCreatedRecordsSent,
			"custom_nat_iddos_l4_entry_created_records_sent_failure":       ret.DtNetflowMonitorStats.Stats.CustomNatIddosL4EntryCreatedRecordsSentFailure,
			"custom_nat_iddos_l4_entry_deleted_records_sent":               ret.DtNetflowMonitorStats.Stats.CustomNatIddosL4EntryDeletedRecordsSent,
			"custom_nat_iddos_l4_entry_deleted_records_sent_failure":       ret.DtNetflowMonitorStats.Stats.CustomNatIddosL4EntryDeletedRecordsSentFailure,
			"custom_gtp_rate_limit_periodic_records_sent":                  ret.DtNetflowMonitorStats.Stats.CustomGtpRateLimitPeriodicRecordsSent,
			"custom_gtp_rate_limit_periodic_records_sent_failure":          ret.DtNetflowMonitorStats.Stats.CustomGtpRateLimitPeriodicRecordsSentFailure,
			"ddos_general_stat_records_sent":                               ret.DtNetflowMonitorStats.Stats.DdosGeneralStatRecordsSent,
			"ddos_general_stat_records_sent_failure":                       ret.DtNetflowMonitorStats.Stats.DdosGeneralStatRecordsSentFailure,
			"ddos_http_stat_records_sent":                                  ret.DtNetflowMonitorStats.Stats.DdosHttpStatRecordsSent,
			"ddos_http_stat_records_sent_failure":                          ret.DtNetflowMonitorStats.Stats.DdosHttpStatRecordsSentFailure,
		},
	}
}

func getObjectNetflowMonitorStatsStats(d []interface{}) edpt.NetflowMonitorStatsStats {

	count1 := len(d)
	var ret edpt.NetflowMonitorStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketsSent = in["packets_sent"].(int)
		ret.BytesSent = in["bytes_sent"].(int)
		ret.Nat44RecordsSent = in["nat44_records_sent"].(int)
		ret.Nat44RecordsSentFailure = in["nat44_records_sent_failure"].(int)
		ret.Nat64RecordsSent = in["nat64_records_sent"].(int)
		ret.Nat64RecordsSentFailure = in["nat64_records_sent_failure"].(int)
		ret.DsliteRecordsSent = in["dslite_records_sent"].(int)
		ret.DsliteRecordsSentFailure = in["dslite_records_sent_failure"].(int)
		ret.SessionEventNat44RecordsSent = in["session_event_nat44_records_sent"].(int)
		ret.SessionEventNat44RecordsSentFailure = in["session_event_nat44_records_sent_failure"].(int)
		ret.SessionEventNat64RecordsSent = in["session_event_nat64_records_sent"].(int)
		ret.SessionEventNat64RecordsSentFailure = in["session_event_nat64_records_sent_failure"].(int)
		ret.SessionEventDsliteRecordsSent = in["session_event_dslite_records_sent"].(int)
		ret.SessionEventDsliteRecordsSentFailure = in["session_event_dslite_records_sent_failure"].(int)
		ret.SessionEventFw4RecordsSent = in["session_event_fw4_records_sent"].(int)
		ret.SessionEventFw4RecordsSentFailure = in["session_event_fw4_records_sent_failure"].(int)
		ret.SessionEventFw6RecordsSent = in["session_event_fw6_records_sent"].(int)
		ret.SessionEventFw6RecordsSentFailure = in["session_event_fw6_records_sent_failure"].(int)
		ret.PortMappingNat44RecordsSent = in["port_mapping_nat44_records_sent"].(int)
		ret.PortMappingNat44RecordsSentFailure = in["port_mapping_nat44_records_sent_failure"].(int)
		ret.PortMappingNat64RecordsSent = in["port_mapping_nat64_records_sent"].(int)
		ret.PortMappingNat64RecordsSentFailure = in["port_mapping_nat64_records_sent_failure"].(int)
		ret.PortMappingDsliteRecordsSent = in["port_mapping_dslite_records_sent"].(int)
		ret.PortMappingDsliteRecordsSentFailure = in["port_mapping_dslite_records_sent_failure"].(int)
		ret.NetflowV5RecordsSent = in["netflow_v5_records_sent"].(int)
		ret.NetflowV5RecordsSentFailure = in["netflow_v5_records_sent_failure"].(int)
		ret.NetflowV5ExtRecordsSent = in["netflow_v5_ext_records_sent"].(int)
		ret.NetflowV5ExtRecordsSentFailure = in["netflow_v5_ext_records_sent_failure"].(int)
		ret.PortBatchingNat44RecordsSent = in["port_batching_nat44_records_sent"].(int)
		ret.PortBatchingNat44RecordsSentFailure = in["port_batching_nat44_records_sent_failure"].(int)
		ret.PortBatchingNat64RecordsSent = in["port_batching_nat64_records_sent"].(int)
		ret.PortBatchingNat64RecordsSentFailure = in["port_batching_nat64_records_sent_failure"].(int)
		ret.PortBatchingDsliteRecordsSent = in["port_batching_dslite_records_sent"].(int)
		ret.PortBatchingDsliteRecordsSentFailure = in["port_batching_dslite_records_sent_failure"].(int)
		ret.PortBatchingV2Nat44RecordsSent = in["port_batching_v2_nat44_records_sent"].(int)
		ret.PortBatchingV2Nat44RecordsSentFailure = in["port_batching_v2_nat44_records_sent_failure"].(int)
		ret.PortBatchingV2Nat64RecordsSent = in["port_batching_v2_nat64_records_sent"].(int)
		ret.PortBatchingV2Nat64RecordsSentFailure = in["port_batching_v2_nat64_records_sent_failure"].(int)
		ret.PortBatchingV2DsliteRecordsSent = in["port_batching_v2_dslite_records_sent"].(int)
		ret.PortBatchingV2DsliteRecordsSentFailure = in["port_batching_v2_dslite_records_sent_failure"].(int)
		ret.CustomSessionEventNat44CreationRecordsSent = in["custom_session_event_nat44_creation_records_sent"].(int)
		ret.CustomSessionEventNat44CreationRecordsSentFailure = in["custom_session_event_nat44_creation_records_sent_failure"].(int)
		ret.CustomSessionEventNat64CreationRecordsSent = in["custom_session_event_nat64_creation_records_sent"].(int)
		ret.CustomSessionEventNat64CreationRecordsSentFailure = in["custom_session_event_nat64_creation_records_sent_failure"].(int)
		ret.CustomSessionEventDsliteCreationRecordsSent = in["custom_session_event_dslite_creation_records_sent"].(int)
		ret.CustomSessionEventDsliteCreationRecordsSentFailure = in["custom_session_event_dslite_creation_records_sent_failure"].(int)
		ret.CustomSessionEventNat44DeletionRecordsSent = in["custom_session_event_nat44_deletion_records_sent"].(int)
		ret.CustomSessionEventNat44DeletionRecordsSentFailure = in["custom_session_event_nat44_deletion_records_sent_failure"].(int)
		ret.CustomSessionEventNat64DeletionRecordsSent = in["custom_session_event_nat64_deletion_records_sent"].(int)
		ret.CustomSessionEventNat64DeletionRecordsSentFailure = in["custom_session_event_nat64_deletion_records_sent_failure"].(int)
		ret.CustomSessionEventDsliteDeletionRecordsSent = in["custom_session_event_dslite_deletion_records_sent"].(int)
		ret.CustomSessionEventDsliteDeletionRecordsSentFailure = in["custom_session_event_dslite_deletion_records_sent_failure"].(int)
		ret.CustomSessionEventFw4CreationRecordsSent = in["custom_session_event_fw4_creation_records_sent"].(int)
		ret.CustomSessionEventFw4CreationRecordsSentFailure = in["custom_session_event_fw4_creation_records_sent_failure"].(int)
		ret.CustomSessionEventFw6CreationRecordsSent = in["custom_session_event_fw6_creation_records_sent"].(int)
		ret.CustomSessionEventFw6CreationRecordsSentFailure = in["custom_session_event_fw6_creation_records_sent_failure"].(int)
		ret.CustomSessionEventFw4DeletionRecordsSent = in["custom_session_event_fw4_deletion_records_sent"].(int)
		ret.CustomSessionEventFw4DeletionRecordsSentFailure = in["custom_session_event_fw4_deletion_records_sent_failure"].(int)
		ret.CustomSessionEventFw6DeletionRecordsSent = in["custom_session_event_fw6_deletion_records_sent"].(int)
		ret.CustomSessionEventFw6DeletionRecordsSentFailure = in["custom_session_event_fw6_deletion_records_sent_failure"].(int)
		ret.CustomDenyResetEventFw4RecordsSent = in["custom_deny_reset_event_fw4_records_sent"].(int)
		ret.CustomDenyResetEventFw4RecordsSentFailure = in["custom_deny_reset_event_fw4_records_sent_failure"].(int)
		ret.CustomDenyResetEventFw6RecordsSent = in["custom_deny_reset_event_fw6_records_sent"].(int)
		ret.CustomDenyResetEventFw6RecordsSentFailure = in["custom_deny_reset_event_fw6_records_sent_failure"].(int)
		ret.CustomPortMappingNat44CreationRecordsSent = in["custom_port_mapping_nat44_creation_records_sent"].(int)
		ret.CustomPortMappingNat44CreationRecordsSentFailure = in["custom_port_mapping_nat44_creation_records_sent_failure"].(int)
		ret.CustomPortMappingNat64CreationRecordsSent = in["custom_port_mapping_nat64_creation_records_sent"].(int)
		ret.CustomPortMappingNat64CreationRecordsSentFailure = in["custom_port_mapping_nat64_creation_records_sent_failure"].(int)
		ret.CustomPortMappingDsliteCreationRecordsSent = in["custom_port_mapping_dslite_creation_records_sent"].(int)
		ret.CustomPortMappingDsliteCreationRecordsSentFailure = in["custom_port_mapping_dslite_creation_records_sent_failure"].(int)
		ret.CustomPortMappingNat44DeletionRecordsSent = in["custom_port_mapping_nat44_deletion_records_sent"].(int)
		ret.CustomPortMappingNat44DeletionRecordsSentFailure = in["custom_port_mapping_nat44_deletion_records_sent_failure"].(int)
		ret.CustomPortMappingNat64DeletionRecordsSent = in["custom_port_mapping_nat64_deletion_records_sent"].(int)
		ret.CustomPortMappingNat64DeletionRecordsSentFailure = in["custom_port_mapping_nat64_deletion_records_sent_failure"].(int)
		ret.CustomPortMappingDsliteDeletionRecordsSent = in["custom_port_mapping_dslite_deletion_records_sent"].(int)
		ret.CustomPortMappingDsliteDeletionRecordsSentFailure = in["custom_port_mapping_dslite_deletion_records_sent_failure"].(int)
		ret.CustomPortBatchingNat44CreationRecordsSent = in["custom_port_batching_nat44_creation_records_sent"].(int)
		ret.CustomPortBatchingNat44CreationRecordsSentFailure = in["custom_port_batching_nat44_creation_records_sent_failure"].(int)
		ret.CustomPortBatchingNat64CreationRecordsSent = in["custom_port_batching_nat64_creation_records_sent"].(int)
		ret.CustomPortBatchingNat64CreationRecordsSentFailure = in["custom_port_batching_nat64_creation_records_sent_failure"].(int)
		ret.CustomPortBatchingDsliteCreationRecordsSent = in["custom_port_batching_dslite_creation_records_sent"].(int)
		ret.CustomPortBatchingDsliteCreationRecordsSentFailure = in["custom_port_batching_dslite_creation_records_sent_failure"].(int)
		ret.CustomPortBatchingNat44DeletionRecordsSent = in["custom_port_batching_nat44_deletion_records_sent"].(int)
		ret.CustomPortBatchingNat44DeletionRecordsSentFailure = in["custom_port_batching_nat44_deletion_records_sent_failure"].(int)
		ret.CustomPortBatchingNat64DeletionRecordsSent = in["custom_port_batching_nat64_deletion_records_sent"].(int)
		ret.CustomPortBatchingNat64DeletionRecordsSentFailure = in["custom_port_batching_nat64_deletion_records_sent_failure"].(int)
		ret.CustomPortBatchingDsliteDeletionRecordsSent = in["custom_port_batching_dslite_deletion_records_sent"].(int)
		ret.CustomPortBatchingDsliteDeletionRecordsSentFailure = in["custom_port_batching_dslite_deletion_records_sent_failure"].(int)
		ret.CustomPortBatchingV2Nat44CreationRecordsSent = in["custom_port_batching_v2_nat44_creation_records_sent"].(int)
		ret.CustomPortBatchingV2Nat44CreationRecordsSentFailure = in["custom_port_batching_v2_nat44_creation_records_sent_failure"].(int)
		ret.CustomPortBatchingV2Nat64CreationRecordsSent = in["custom_port_batching_v2_nat64_creation_records_sent"].(int)
		ret.CustomPortBatchingV2Nat64CreationRecordsSentFailure = in["custom_port_batching_v2_nat64_creation_records_sent_failure"].(int)
		ret.CustomPortBatchingV2DsliteCreationRecordsSent = in["custom_port_batching_v2_dslite_creation_records_sent"].(int)
		ret.CustomPortBatchingV2DsliteCreationRecordsSentFailure = in["custom_port_batching_v2_dslite_creation_records_sent_failure"].(int)
		ret.CustomPortBatchingV2Nat44DeletionRecordsSent = in["custom_port_batching_v2_nat44_deletion_records_sent"].(int)
		ret.CustomPortBatchingV2Nat44DeletionRecordsSentFailure = in["custom_port_batching_v2_nat44_deletion_records_sent_failure"].(int)
		ret.CustomPortBatchingV2Nat64DeletionRecordsSent = in["custom_port_batching_v2_nat64_deletion_records_sent"].(int)
		ret.CustomPortBatchingV2Nat64DeletionRecordsSentFailure = in["custom_port_batching_v2_nat64_deletion_records_sent_failure"].(int)
		ret.CustomPortBatchingV2DsliteDeletionRecordsSent = in["custom_port_batching_v2_dslite_deletion_records_sent"].(int)
		ret.CustomPortBatchingV2DsliteDeletionRecordsSentFailure = in["custom_port_batching_v2_dslite_deletion_records_sent_failure"].(int)
		ret.ReducedLogsByDestination = in["reduced_logs_by_destination"].(int)
		ret.CustomGtpCTunnelEventRecordsSent = in["custom_gtp_c_tunnel_event_records_sent"].(int)
		ret.CustomGtpCTunnelEventRecordsSentFailure = in["custom_gtp_c_tunnel_event_records_sent_failure"].(int)
		ret.CustomGtpUTunnelEventRecordsSent = in["custom_gtp_u_tunnel_event_records_sent"].(int)
		ret.CustomGtpUTunnelEventRecordsSentFailure = in["custom_gtp_u_tunnel_event_records_sent_failure"].(int)
		ret.CustomGtpDenyEventRecordsSent = in["custom_gtp_deny_event_records_sent"].(int)
		ret.CustomGtpDenyEventRecordsSentFailure = in["custom_gtp_deny_event_records_sent_failure"].(int)
		ret.CustomGtpInfoEventRecordsSent = in["custom_gtp_info_event_records_sent"].(int)
		ret.CustomGtpInfoEventRecordsSentFailure = in["custom_gtp_info_event_records_sent_failure"].(int)
		ret.CustomFwIddosEntryCreatedRecordsSent = in["custom_fw_iddos_entry_created_records_sent"].(int)
		ret.CustomFwIddosEntryCreatedRecordsSentFailure = in["custom_fw_iddos_entry_created_records_sent_failure"].(int)
		ret.CustomFwIddosEntryDeletedRecordsSent = in["custom_fw_iddos_entry_deleted_records_sent"].(int)
		ret.CustomFwIddosEntryDeletedRecordsSentFailure = in["custom_fw_iddos_entry_deleted_records_sent_failure"].(int)
		ret.CustomFwSesnLimitExceededRecordsSent = in["custom_fw_sesn_limit_exceeded_records_sent"].(int)
		ret.CustomFwSesnLimitExceededRecordsSentFailure = in["custom_fw_sesn_limit_exceeded_records_sent_failure"].(int)
		ret.CustomNatIddosL3EntryCreatedRecordsSent = in["custom_nat_iddos_l3_entry_created_records_sent"].(int)
		ret.CustomNatIddosL3EntryCreatedRecordsSentFailure = in["custom_nat_iddos_l3_entry_created_records_sent_failure"].(int)
		ret.CustomNatIddosL3EntryDeletedRecordsSent = in["custom_nat_iddos_l3_entry_deleted_records_sent"].(int)
		ret.CustomNatIddosL3EntryDeletedRecordsSentFailure = in["custom_nat_iddos_l3_entry_deleted_records_sent_failure"].(int)
		ret.CustomNatIddosL4EntryCreatedRecordsSent = in["custom_nat_iddos_l4_entry_created_records_sent"].(int)
		ret.CustomNatIddosL4EntryCreatedRecordsSentFailure = in["custom_nat_iddos_l4_entry_created_records_sent_failure"].(int)
		ret.CustomNatIddosL4EntryDeletedRecordsSent = in["custom_nat_iddos_l4_entry_deleted_records_sent"].(int)
		ret.CustomNatIddosL4EntryDeletedRecordsSentFailure = in["custom_nat_iddos_l4_entry_deleted_records_sent_failure"].(int)
		ret.CustomGtpRateLimitPeriodicRecordsSent = in["custom_gtp_rate_limit_periodic_records_sent"].(int)
		ret.CustomGtpRateLimitPeriodicRecordsSentFailure = in["custom_gtp_rate_limit_periodic_records_sent_failure"].(int)
		ret.DdosGeneralStatRecordsSent = in["ddos_general_stat_records_sent"].(int)
		ret.DdosGeneralStatRecordsSentFailure = in["ddos_general_stat_records_sent_failure"].(int)
		ret.DdosHttpStatRecordsSent = in["ddos_http_stat_records_sent"].(int)
		ret.DdosHttpStatRecordsSentFailure = in["ddos_http_stat_records_sent_failure"].(int)
	}
	return ret
}

func dataToEndpointNetflowMonitorStats(d *schema.ResourceData) edpt.NetflowMonitorStats {
	var ret edpt.NetflowMonitorStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectNetflowMonitorStatsStats(d.Get("stats").([]interface{}))
	return ret
}
