package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_netflow_monitor_tmpl`: Configure template for netflow.monitor\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplDelete,

		Schema: map[string]*schema.Schema{
			"capture_config": {
				Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
			},
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nat44_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT44 Flow Records Failed",
						},
						"nat64_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT64 Flow Records Failed",
						},
						"dslite_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dslite Flow Records Failed",
						},
						"session_event_nat44_records_sent_failur": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Nat44 Session Event Records Failed",
						},
						"session_event_nat64_records_sent_failur": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Nat64 Session Event Records Falied",
						},
						"session_event_dslite_records_sent_failu": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dslite Session Event Records Failed",
						},
						"session_event_fw4_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for FW4 Session Event Records Failed",
						},
						"session_event_fw6_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for FW6 Session Event Records Failed",
						},
						"port_mapping_nat44_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Nat44 Event Records Failed",
						},
						"port_mapping_nat64_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Nat64 Event Records Failed",
						},
						"port_mapping_dslite_records_sent_failur": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Dslite Event Records failed",
						},
						"netflow_v5_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Netflow v5 Records Failed",
						},
						"netflow_v5_ext_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Netflow v5 Ext Records Failed",
						},
						"port_batching_nat44_records_sent_failur": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Nat44 Records Failed",
						},
						"port_batching_nat64_records_sent_failur": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Nat64 Records Failed",
						},
						"port_batching_dslite_records_sent_failu": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Dslite Records Failed",
						},
						"port_batching_v2_nat44_records_sent_fai": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Nat44 Records Failed",
						},
						"port_batching_v2_nat64_records_sent_fai": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Nat64 Records Failed",
						},
						"port_batching_v2_dslite_records_sent_fa": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Dslite Records Falied",
						},
						"custom_session_event_nat44_creation_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Session Creation Records Failed",
						},
						"custom_session_event_nat64_creation_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Session Creation Records Failed",
						},
						"custom_session_event_dslite_creation_re": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Session Creation Records Failed",
						},
						"custom_session_event_nat44_deletion_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Session Deletion Records Failed",
						},
						"custom_session_event_nat64_deletion_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Session Deletion Records Failed",
						},
						"custom_session_event_dslite_deletion_re": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Session Deletion Records Failed",
						},
						"custom_session_event_fw4_creation_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Session Creation Records Failed",
						},
						"custom_session_event_fw6_creation_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Session Creation Records Failed",
						},
						"custom_session_event_fw4_deletion_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Session Deletion Records Failed",
						},
						"custom_session_event_fw6_deletion_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Session Deletion Records Failed",
						},
						"custom_deny_reset_event_fw4_records_sen": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Deny/Reset Event Records Failed",
						},
						"custom_deny_reset_event_fw6_records_sen": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Deny/Reset Event Records Failed",
						},
						"custom_port_mapping_nat44_creation_reco": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Map Creation Records Failed",
						},
						"custom_port_mapping_nat64_creation_reco": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Map Creation Records Failed",
						},
						"custom_port_mapping_dslite_creation_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Map Creation Records Failed",
						},
						"custom_port_mapping_nat44_deletion_reco": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Map Deletion Records Failed",
						},
						"custom_port_mapping_nat64_deletion_reco": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Map Deletion Records Failed",
						},
						"custom_port_mapping_dslite_deletion_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Map Deletion Records Failed",
						},
						"custom_port_batching_nat44_creation_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch Creation Records Failed",
						},
						"custom_port_batching_nat64_creation_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch Creation Records Failed",
						},
						"custom_port_batching_dslite_creation_re": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch Creation Records Failed",
						},
						"custom_port_batching_nat44_deletion_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch Deletion Records Failed",
						},
						"custom_port_batching_nat64_deletion_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch Deletion Records Failed",
						},
						"custom_port_batching_dslite_deletion_re": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch Deletion Records Failed",
						},
						"custom_port_batching_v2_nat44_creation_": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch V2 Creation Records Failed",
						},
						"custom_port_batching_v2_nat64_creation_": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch V2 Creation Records Failed",
						},
						"custom_port_batching_v2_dslite_creation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch V2 Creation Records Failed",
						},
						"custom_port_batching_v2_nat44_deletion_": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch V2 Deletion Records Failed",
						},
						"custom_port_batching_v2_nat64_deletion_": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch V2 Deletion Records Failed",
						},
						"custom_port_batching_v2_dslite_deletion": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch V2 Deletion Records Failed",
						},
						"custom_gtp_c_tunnel_event_records_sent_": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP C Tunnel Records Sent Failure",
						},
						"custom_gtp_u_tunnel_event_records_sent_": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP U Tunnel Records Sent Failure",
						},
						"custom_gtp_deny_event_records_sent_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP Deny Records Sent Failure",
						},
						"custom_gtp_info_event_records_sent_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP Info Records Sent Failure",
						},
						"custom_fw_iddos_entry_created_records_s": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW iDDoS Entry Created Records Sent Failure",
						},
						"custom_fw_iddos_entry_deleted_records_s": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW iDDoS Entry Deleted Records Sent Failure",
						},
						"custom_fw_sesn_limit_exceeded_records_s": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW Session Limit Exceeded Records Sent Failure",
						},
						"custom_nat_iddos_l3_entry_created_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L3 Entry Created Records Sent Failure",
						},
						"custom_nat_iddos_l3_entry_deleted_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L3 Entry Deleted Records Sent Failure",
						},
						"custom_nat_iddos_l4_entry_created_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L4 Entry Created Records Sent Failure",
						},
						"custom_nat_iddos_l4_entry_deleted_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L4 Entry Deleted Records Sent Failure",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trigger_stats_rate": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold_exceeded_by": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
						},
						"duration": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
						},
						"nat44_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT44 Flow Records Failed",
						},
						"nat64_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT64 Flow Records Failed",
						},
						"dslite_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dslite Flow Records Failed",
						},
						"session_event_nat44_records_sent_failur": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Nat44 Session Event Records Failed",
						},
						"session_event_nat64_records_sent_failur": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Nat64 Session Event Records Falied",
						},
						"session_event_dslite_records_sent_failu": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dslite Session Event Records Failed",
						},
						"session_event_fw4_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for FW4 Session Event Records Failed",
						},
						"session_event_fw6_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for FW6 Session Event Records Failed",
						},
						"port_mapping_nat44_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Nat44 Event Records Failed",
						},
						"port_mapping_nat64_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Nat64 Event Records Failed",
						},
						"port_mapping_dslite_records_sent_failur": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Dslite Event Records failed",
						},
						"netflow_v5_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Netflow v5 Records Failed",
						},
						"netflow_v5_ext_records_sent_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Netflow v5 Ext Records Failed",
						},
						"port_batching_nat44_records_sent_failur": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Nat44 Records Failed",
						},
						"port_batching_nat64_records_sent_failur": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Nat64 Records Failed",
						},
						"port_batching_dslite_records_sent_failu": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Dslite Records Failed",
						},
						"port_batching_v2_nat44_records_sent_fai": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Nat44 Records Failed",
						},
						"port_batching_v2_nat64_records_sent_fai": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Nat64 Records Failed",
						},
						"port_batching_v2_dslite_records_sent_fa": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Dslite Records Falied",
						},
						"custom_session_event_nat44_creation_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Session Creation Records Failed",
						},
						"custom_session_event_nat64_creation_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Session Creation Records Failed",
						},
						"custom_session_event_dslite_creation_re": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Session Creation Records Failed",
						},
						"custom_session_event_nat44_deletion_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Session Deletion Records Failed",
						},
						"custom_session_event_nat64_deletion_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Session Deletion Records Failed",
						},
						"custom_session_event_dslite_deletion_re": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Session Deletion Records Failed",
						},
						"custom_session_event_fw4_creation_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Session Creation Records Failed",
						},
						"custom_session_event_fw6_creation_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Session Creation Records Failed",
						},
						"custom_session_event_fw4_deletion_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Session Deletion Records Failed",
						},
						"custom_session_event_fw6_deletion_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Session Deletion Records Failed",
						},
						"custom_deny_reset_event_fw4_records_sen": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Deny/Reset Event Records Failed",
						},
						"custom_deny_reset_event_fw6_records_sen": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Deny/Reset Event Records Failed",
						},
						"custom_port_mapping_nat44_creation_reco": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Map Creation Records Failed",
						},
						"custom_port_mapping_nat64_creation_reco": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Map Creation Records Failed",
						},
						"custom_port_mapping_dslite_creation_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Map Creation Records Failed",
						},
						"custom_port_mapping_nat44_deletion_reco": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Map Deletion Records Failed",
						},
						"custom_port_mapping_nat64_deletion_reco": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Map Deletion Records Failed",
						},
						"custom_port_mapping_dslite_deletion_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Map Deletion Records Failed",
						},
						"custom_port_batching_nat44_creation_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch Creation Records Failed",
						},
						"custom_port_batching_nat64_creation_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch Creation Records Failed",
						},
						"custom_port_batching_dslite_creation_re": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch Creation Records Failed",
						},
						"custom_port_batching_nat44_deletion_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch Deletion Records Failed",
						},
						"custom_port_batching_nat64_deletion_rec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch Deletion Records Failed",
						},
						"custom_port_batching_dslite_deletion_re": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch Deletion Records Failed",
						},
						"custom_port_batching_v2_nat44_creation_": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch V2 Creation Records Failed",
						},
						"custom_port_batching_v2_nat64_creation_": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch V2 Creation Records Failed",
						},
						"custom_port_batching_v2_dslite_creation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch V2 Creation Records Failed",
						},
						"custom_port_batching_v2_nat44_deletion_": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch V2 Deletion Records Failed",
						},
						"custom_port_batching_v2_nat64_deletion_": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch V2 Deletion Records Failed",
						},
						"custom_port_batching_v2_dslite_deletion": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch V2 Deletion Records Failed",
						},
						"custom_gtp_c_tunnel_event_records_sent_": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP C Tunnel Records Sent Failure",
						},
						"custom_gtp_u_tunnel_event_records_sent_": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP U Tunnel Records Sent Failure",
						},
						"custom_gtp_deny_event_records_sent_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP Deny Records Sent Failure",
						},
						"custom_gtp_info_event_records_sent_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP Info Records Sent Failure",
						},
						"custom_fw_iddos_entry_created_records_s": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW iDDoS Entry Created Records Sent Failure",
						},
						"custom_fw_iddos_entry_deleted_records_s": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW iDDoS Entry Deleted Records Sent Failure",
						},
						"custom_fw_sesn_limit_exceeded_records_s": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW Session Limit Exceeded Records Sent Failure",
						},
						"custom_nat_iddos_l3_entry_created_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L3 Entry Created Records Sent Failure",
						},
						"custom_nat_iddos_l3_entry_deleted_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L3 Entry Deleted Records Sent Failure",
						},
						"custom_nat_iddos_l4_entry_created_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L4 Entry Created Records Sent Failure",
						},
						"custom_nat_iddos_l4_entry_deleted_recor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L4 Entry Deleted Records Sent Failure",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trigger_stats_severity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
						},
						"error_alert": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
						},
						"error_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
						},
						"error_critical": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
						},
						"drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
						},
						"drop_alert": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
						},
						"drop_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
						},
						"drop_critical": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc2699(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc2699 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc2699
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Nat44RecordsSentFailure = in["nat44_records_sent_failure"].(int)
		ret.Nat64RecordsSentFailure = in["nat64_records_sent_failure"].(int)
		ret.DsliteRecordsSentFailure = in["dslite_records_sent_failure"].(int)
		ret.SessionEventNat44RecordsSentFailur = in["session_event_nat44_records_sent_failur"].(int)
		ret.SessionEventNat64RecordsSentFailur = in["session_event_nat64_records_sent_failur"].(int)
		ret.SessionEventDsliteRecordsSentFailu = in["session_event_dslite_records_sent_failu"].(int)
		ret.SessionEventFw4RecordsSentFailure = in["session_event_fw4_records_sent_failure"].(int)
		ret.SessionEventFw6RecordsSentFailure = in["session_event_fw6_records_sent_failure"].(int)
		ret.PortMappingNat44RecordsSentFailure = in["port_mapping_nat44_records_sent_failure"].(int)
		ret.PortMappingNat64RecordsSentFailure = in["port_mapping_nat64_records_sent_failure"].(int)
		ret.PortMappingDsliteRecordsSentFailur = in["port_mapping_dslite_records_sent_failur"].(int)
		ret.NetflowV5RecordsSentFailure = in["netflow_v5_records_sent_failure"].(int)
		ret.NetflowV5ExtRecordsSentFailure = in["netflow_v5_ext_records_sent_failure"].(int)
		ret.PortBatchingNat44RecordsSentFailur = in["port_batching_nat44_records_sent_failur"].(int)
		ret.PortBatchingNat64RecordsSentFailur = in["port_batching_nat64_records_sent_failur"].(int)
		ret.PortBatchingDsliteRecordsSentFailu = in["port_batching_dslite_records_sent_failu"].(int)
		ret.PortBatchingV2Nat44RecordsSentFai = in["port_batching_v2_nat44_records_sent_fai"].(int)
		ret.PortBatchingV2Nat64RecordsSentFai = in["port_batching_v2_nat64_records_sent_fai"].(int)
		ret.PortBatchingV2DsliteRecordsSentFa = in["port_batching_v2_dslite_records_sent_fa"].(int)
		ret.CustomSessionEventNat44CreationRec = in["custom_session_event_nat44_creation_rec"].(int)
		ret.CustomSessionEventNat64CreationRec = in["custom_session_event_nat64_creation_rec"].(int)
		ret.CustomSessionEventDsliteCreationRe = in["custom_session_event_dslite_creation_re"].(int)
		ret.CustomSessionEventNat44DeletionRec = in["custom_session_event_nat44_deletion_rec"].(int)
		ret.CustomSessionEventNat64DeletionRec = in["custom_session_event_nat64_deletion_rec"].(int)
		ret.CustomSessionEventDsliteDeletionRe = in["custom_session_event_dslite_deletion_re"].(int)
		ret.CustomSessionEventFw4CreationRecor = in["custom_session_event_fw4_creation_recor"].(int)
		ret.CustomSessionEventFw6CreationRecor = in["custom_session_event_fw6_creation_recor"].(int)
		ret.CustomSessionEventFw4DeletionRecor = in["custom_session_event_fw4_deletion_recor"].(int)
		ret.CustomSessionEventFw6DeletionRecor = in["custom_session_event_fw6_deletion_recor"].(int)
		ret.CustomDenyResetEventFw4RecordsSen = in["custom_deny_reset_event_fw4_records_sen"].(int)
		ret.CustomDenyResetEventFw6RecordsSen = in["custom_deny_reset_event_fw6_records_sen"].(int)
		ret.CustomPortMappingNat44CreationReco = in["custom_port_mapping_nat44_creation_reco"].(int)
		ret.CustomPortMappingNat64CreationReco = in["custom_port_mapping_nat64_creation_reco"].(int)
		ret.CustomPortMappingDsliteCreationRec = in["custom_port_mapping_dslite_creation_rec"].(int)
		ret.CustomPortMappingNat44DeletionReco = in["custom_port_mapping_nat44_deletion_reco"].(int)
		ret.CustomPortMappingNat64DeletionReco = in["custom_port_mapping_nat64_deletion_reco"].(int)
		ret.CustomPortMappingDsliteDeletionRec = in["custom_port_mapping_dslite_deletion_rec"].(int)
		ret.CustomPortBatchingNat44CreationRec = in["custom_port_batching_nat44_creation_rec"].(int)
		ret.CustomPortBatchingNat64CreationRec = in["custom_port_batching_nat64_creation_rec"].(int)
		ret.CustomPortBatchingDsliteCreationRe = in["custom_port_batching_dslite_creation_re"].(int)
		ret.CustomPortBatchingNat44DeletionRec = in["custom_port_batching_nat44_deletion_rec"].(int)
		ret.CustomPortBatchingNat64DeletionRec = in["custom_port_batching_nat64_deletion_rec"].(int)
		ret.CustomPortBatchingDsliteDeletionRe = in["custom_port_batching_dslite_deletion_re"].(int)
		ret.CustomPortBatchingV2Nat44Creation = in["custom_port_batching_v2_nat44_creation_"].(int)
		ret.CustomPortBatchingV2Nat64Creation = in["custom_port_batching_v2_nat64_creation_"].(int)
		ret.CustomPortBatchingV2DsliteCreation = in["custom_port_batching_v2_dslite_creation"].(int)
		ret.CustomPortBatchingV2Nat44Deletion = in["custom_port_batching_v2_nat44_deletion_"].(int)
		ret.CustomPortBatchingV2Nat64Deletion = in["custom_port_batching_v2_nat64_deletion_"].(int)
		ret.CustomPortBatchingV2DsliteDeletion = in["custom_port_batching_v2_dslite_deletion"].(int)
		ret.CustomGtpCTunnelEventRecordsSent = in["custom_gtp_c_tunnel_event_records_sent_"].(int)
		ret.CustomGtpUTunnelEventRecordsSent = in["custom_gtp_u_tunnel_event_records_sent_"].(int)
		ret.CustomGtpDenyEventRecordsSentFail = in["custom_gtp_deny_event_records_sent_fail"].(int)
		ret.CustomGtpInfoEventRecordsSentFail = in["custom_gtp_info_event_records_sent_fail"].(int)
		ret.CustomFwIddosEntryCreatedRecordsS = in["custom_fw_iddos_entry_created_records_s"].(int)
		ret.CustomFwIddosEntryDeletedRecordsS = in["custom_fw_iddos_entry_deleted_records_s"].(int)
		ret.CustomFwSesnLimitExceededRecordsS = in["custom_fw_sesn_limit_exceeded_records_s"].(int)
		ret.CustomNatIddosL3EntryCreatedRecor = in["custom_nat_iddos_l3_entry_created_recor"].(int)
		ret.CustomNatIddosL3EntryDeletedRecor = in["custom_nat_iddos_l3_entry_deleted_recor"].(int)
		ret.CustomNatIddosL4EntryCreatedRecor = in["custom_nat_iddos_l4_entry_created_recor"].(int)
		ret.CustomNatIddosL4EntryDeletedRecor = in["custom_nat_iddos_l4_entry_deleted_recor"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate2700(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate2700 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate2700
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Nat44RecordsSentFailure = in["nat44_records_sent_failure"].(int)
		ret.Nat64RecordsSentFailure = in["nat64_records_sent_failure"].(int)
		ret.DsliteRecordsSentFailure = in["dslite_records_sent_failure"].(int)
		ret.SessionEventNat44RecordsSentFailur = in["session_event_nat44_records_sent_failur"].(int)
		ret.SessionEventNat64RecordsSentFailur = in["session_event_nat64_records_sent_failur"].(int)
		ret.SessionEventDsliteRecordsSentFailu = in["session_event_dslite_records_sent_failu"].(int)
		ret.SessionEventFw4RecordsSentFailure = in["session_event_fw4_records_sent_failure"].(int)
		ret.SessionEventFw6RecordsSentFailure = in["session_event_fw6_records_sent_failure"].(int)
		ret.PortMappingNat44RecordsSentFailure = in["port_mapping_nat44_records_sent_failure"].(int)
		ret.PortMappingNat64RecordsSentFailure = in["port_mapping_nat64_records_sent_failure"].(int)
		ret.PortMappingDsliteRecordsSentFailur = in["port_mapping_dslite_records_sent_failur"].(int)
		ret.NetflowV5RecordsSentFailure = in["netflow_v5_records_sent_failure"].(int)
		ret.NetflowV5ExtRecordsSentFailure = in["netflow_v5_ext_records_sent_failure"].(int)
		ret.PortBatchingNat44RecordsSentFailur = in["port_batching_nat44_records_sent_failur"].(int)
		ret.PortBatchingNat64RecordsSentFailur = in["port_batching_nat64_records_sent_failur"].(int)
		ret.PortBatchingDsliteRecordsSentFailu = in["port_batching_dslite_records_sent_failu"].(int)
		ret.PortBatchingV2Nat44RecordsSentFai = in["port_batching_v2_nat44_records_sent_fai"].(int)
		ret.PortBatchingV2Nat64RecordsSentFai = in["port_batching_v2_nat64_records_sent_fai"].(int)
		ret.PortBatchingV2DsliteRecordsSentFa = in["port_batching_v2_dslite_records_sent_fa"].(int)
		ret.CustomSessionEventNat44CreationRec = in["custom_session_event_nat44_creation_rec"].(int)
		ret.CustomSessionEventNat64CreationRec = in["custom_session_event_nat64_creation_rec"].(int)
		ret.CustomSessionEventDsliteCreationRe = in["custom_session_event_dslite_creation_re"].(int)
		ret.CustomSessionEventNat44DeletionRec = in["custom_session_event_nat44_deletion_rec"].(int)
		ret.CustomSessionEventNat64DeletionRec = in["custom_session_event_nat64_deletion_rec"].(int)
		ret.CustomSessionEventDsliteDeletionRe = in["custom_session_event_dslite_deletion_re"].(int)
		ret.CustomSessionEventFw4CreationRecor = in["custom_session_event_fw4_creation_recor"].(int)
		ret.CustomSessionEventFw6CreationRecor = in["custom_session_event_fw6_creation_recor"].(int)
		ret.CustomSessionEventFw4DeletionRecor = in["custom_session_event_fw4_deletion_recor"].(int)
		ret.CustomSessionEventFw6DeletionRecor = in["custom_session_event_fw6_deletion_recor"].(int)
		ret.CustomDenyResetEventFw4RecordsSen = in["custom_deny_reset_event_fw4_records_sen"].(int)
		ret.CustomDenyResetEventFw6RecordsSen = in["custom_deny_reset_event_fw6_records_sen"].(int)
		ret.CustomPortMappingNat44CreationReco = in["custom_port_mapping_nat44_creation_reco"].(int)
		ret.CustomPortMappingNat64CreationReco = in["custom_port_mapping_nat64_creation_reco"].(int)
		ret.CustomPortMappingDsliteCreationRec = in["custom_port_mapping_dslite_creation_rec"].(int)
		ret.CustomPortMappingNat44DeletionReco = in["custom_port_mapping_nat44_deletion_reco"].(int)
		ret.CustomPortMappingNat64DeletionReco = in["custom_port_mapping_nat64_deletion_reco"].(int)
		ret.CustomPortMappingDsliteDeletionRec = in["custom_port_mapping_dslite_deletion_rec"].(int)
		ret.CustomPortBatchingNat44CreationRec = in["custom_port_batching_nat44_creation_rec"].(int)
		ret.CustomPortBatchingNat64CreationRec = in["custom_port_batching_nat64_creation_rec"].(int)
		ret.CustomPortBatchingDsliteCreationRe = in["custom_port_batching_dslite_creation_re"].(int)
		ret.CustomPortBatchingNat44DeletionRec = in["custom_port_batching_nat44_deletion_rec"].(int)
		ret.CustomPortBatchingNat64DeletionRec = in["custom_port_batching_nat64_deletion_rec"].(int)
		ret.CustomPortBatchingDsliteDeletionRe = in["custom_port_batching_dslite_deletion_re"].(int)
		ret.CustomPortBatchingV2Nat44Creation = in["custom_port_batching_v2_nat44_creation_"].(int)
		ret.CustomPortBatchingV2Nat64Creation = in["custom_port_batching_v2_nat64_creation_"].(int)
		ret.CustomPortBatchingV2DsliteCreation = in["custom_port_batching_v2_dslite_creation"].(int)
		ret.CustomPortBatchingV2Nat44Deletion = in["custom_port_batching_v2_nat44_deletion_"].(int)
		ret.CustomPortBatchingV2Nat64Deletion = in["custom_port_batching_v2_nat64_deletion_"].(int)
		ret.CustomPortBatchingV2DsliteDeletion = in["custom_port_batching_v2_dslite_deletion"].(int)
		ret.CustomGtpCTunnelEventRecordsSent = in["custom_gtp_c_tunnel_event_records_sent_"].(int)
		ret.CustomGtpUTunnelEventRecordsSent = in["custom_gtp_u_tunnel_event_records_sent_"].(int)
		ret.CustomGtpDenyEventRecordsSentFail = in["custom_gtp_deny_event_records_sent_fail"].(int)
		ret.CustomGtpInfoEventRecordsSentFail = in["custom_gtp_info_event_records_sent_fail"].(int)
		ret.CustomFwIddosEntryCreatedRecordsS = in["custom_fw_iddos_entry_created_records_s"].(int)
		ret.CustomFwIddosEntryDeletedRecordsS = in["custom_fw_iddos_entry_deleted_records_s"].(int)
		ret.CustomFwSesnLimitExceededRecordsS = in["custom_fw_sesn_limit_exceeded_records_s"].(int)
		ret.CustomNatIddosL3EntryCreatedRecor = in["custom_nat_iddos_l3_entry_created_recor"].(int)
		ret.CustomNatIddosL3EntryDeletedRecor = in["custom_nat_iddos_l3_entry_deleted_recor"].(int)
		ret.CustomNatIddosL4EntryCreatedRecor = in["custom_nat_iddos_l4_entry_created_recor"].(int)
		ret.CustomNatIddosL4EntryDeletedRecor = in["custom_nat_iddos_l4_entry_deleted_recor"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsSeverity2701(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsSeverity2701 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsSeverity2701
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc2699(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate2700(d.Get("trigger_stats_rate").([]interface{}))
	ret.Inst.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsSeverity2701(d.Get("trigger_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
