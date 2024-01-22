package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_netflow_monitor_tmpl_trigger_stats_inc`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"custom_deny_reset_event_fw4_records_sen": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Deny/Reset Event Records Failed",
			},
			"custom_deny_reset_event_fw6_records_sen": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Deny/Reset Event Records Failed",
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
			"custom_gtp_c_tunnel_event_records_sent_": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP C Tunnel Records Sent Failure",
			},
			"custom_gtp_deny_event_records_sent_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP Deny Records Sent Failure",
			},
			"custom_gtp_info_event_records_sent_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP Info Records Sent Failure",
			},
			"custom_gtp_u_tunnel_event_records_sent_": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP U Tunnel Records Sent Failure",
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
			"custom_port_batching_dslite_creation_re": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch Creation Records Failed",
			},
			"custom_port_batching_dslite_deletion_re": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch Deletion Records Failed",
			},
			"custom_port_batching_nat44_creation_rec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch Creation Records Failed",
			},
			"custom_port_batching_nat44_deletion_rec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch Deletion Records Failed",
			},
			"custom_port_batching_nat64_creation_rec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch Creation Records Failed",
			},
			"custom_port_batching_nat64_deletion_rec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch Deletion Records Failed",
			},
			"custom_port_batching_v2_dslite_creation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch V2 Creation Records Failed",
			},
			"custom_port_batching_v2_dslite_deletion": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch V2 Deletion Records Failed",
			},
			"custom_port_batching_v2_nat44_creation_": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch V2 Creation Records Failed",
			},
			"custom_port_batching_v2_nat44_deletion_": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch V2 Deletion Records Failed",
			},
			"custom_port_batching_v2_nat64_creation_": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch V2 Creation Records Failed",
			},
			"custom_port_batching_v2_nat64_deletion_": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch V2 Deletion Records Failed",
			},
			"custom_port_mapping_dslite_creation_rec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Map Creation Records Failed",
			},
			"custom_port_mapping_dslite_deletion_rec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Map Deletion Records Failed",
			},
			"custom_port_mapping_nat44_creation_reco": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Map Creation Records Failed",
			},
			"custom_port_mapping_nat44_deletion_reco": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Map Deletion Records Failed",
			},
			"custom_port_mapping_nat64_creation_reco": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Map Creation Records Failed",
			},
			"custom_port_mapping_nat64_deletion_reco": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Map Deletion Records Failed",
			},
			"custom_session_event_dslite_creation_re": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Session Creation Records Failed",
			},
			"custom_session_event_dslite_deletion_re": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Session Deletion Records Failed",
			},
			"custom_session_event_fw4_creation_recor": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Session Creation Records Failed",
			},
			"custom_session_event_fw4_deletion_recor": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Session Deletion Records Failed",
			},
			"custom_session_event_fw6_creation_recor": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Session Creation Records Failed",
			},
			"custom_session_event_fw6_deletion_recor": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Session Deletion Records Failed",
			},
			"custom_session_event_nat44_creation_rec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Session Creation Records Failed",
			},
			"custom_session_event_nat44_deletion_rec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Session Deletion Records Failed",
			},
			"custom_session_event_nat64_creation_rec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Session Creation Records Failed",
			},
			"custom_session_event_nat64_deletion_rec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Session Deletion Records Failed",
			},
			"dslite_records_sent_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dslite Flow Records Failed",
			},
			"nat44_records_sent_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT44 Flow Records Failed",
			},
			"nat64_records_sent_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT64 Flow Records Failed",
			},
			"netflow_v5_ext_records_sent_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Netflow v5 Ext Records Failed",
			},
			"netflow_v5_records_sent_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Netflow v5 Records Failed",
			},
			"port_batching_dslite_records_sent_failu": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Dslite Records Failed",
			},
			"port_batching_nat44_records_sent_failur": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Nat44 Records Failed",
			},
			"port_batching_nat64_records_sent_failur": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Nat64 Records Failed",
			},
			"port_batching_v2_dslite_records_sent_fa": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Dslite Records Falied",
			},
			"port_batching_v2_nat44_records_sent_fai": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Nat44 Records Failed",
			},
			"port_batching_v2_nat64_records_sent_fai": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Nat64 Records Failed",
			},
			"port_mapping_dslite_records_sent_failur": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Dslite Event Records failed",
			},
			"port_mapping_nat44_records_sent_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Nat44 Event Records Failed",
			},
			"port_mapping_nat64_records_sent_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Nat64 Event Records Failed",
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
			"session_event_nat44_records_sent_failur": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Nat44 Session Event Records Failed",
			},
			"session_event_nat64_records_sent_failur": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Nat64 Session Event Records Falied",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc
	ret.Inst.CustomDenyResetEventFw4RecordsSen = d.Get("custom_deny_reset_event_fw4_records_sen").(int)
	ret.Inst.CustomDenyResetEventFw6RecordsSen = d.Get("custom_deny_reset_event_fw6_records_sen").(int)
	ret.Inst.CustomFwIddosEntryCreatedRecordsS = d.Get("custom_fw_iddos_entry_created_records_s").(int)
	ret.Inst.CustomFwIddosEntryDeletedRecordsS = d.Get("custom_fw_iddos_entry_deleted_records_s").(int)
	ret.Inst.CustomFwSesnLimitExceededRecordsS = d.Get("custom_fw_sesn_limit_exceeded_records_s").(int)
	ret.Inst.CustomGtpCTunnelEventRecordsSent = d.Get("custom_gtp_c_tunnel_event_records_sent_").(int)
	ret.Inst.CustomGtpDenyEventRecordsSentFail = d.Get("custom_gtp_deny_event_records_sent_fail").(int)
	ret.Inst.CustomGtpInfoEventRecordsSentFail = d.Get("custom_gtp_info_event_records_sent_fail").(int)
	ret.Inst.CustomGtpUTunnelEventRecordsSent = d.Get("custom_gtp_u_tunnel_event_records_sent_").(int)
	ret.Inst.CustomNatIddosL3EntryCreatedRecor = d.Get("custom_nat_iddos_l3_entry_created_recor").(int)
	ret.Inst.CustomNatIddosL3EntryDeletedRecor = d.Get("custom_nat_iddos_l3_entry_deleted_recor").(int)
	ret.Inst.CustomNatIddosL4EntryCreatedRecor = d.Get("custom_nat_iddos_l4_entry_created_recor").(int)
	ret.Inst.CustomNatIddosL4EntryDeletedRecor = d.Get("custom_nat_iddos_l4_entry_deleted_recor").(int)
	ret.Inst.CustomPortBatchingDsliteCreationRe = d.Get("custom_port_batching_dslite_creation_re").(int)
	ret.Inst.CustomPortBatchingDsliteDeletionRe = d.Get("custom_port_batching_dslite_deletion_re").(int)
	ret.Inst.CustomPortBatchingNat44CreationRec = d.Get("custom_port_batching_nat44_creation_rec").(int)
	ret.Inst.CustomPortBatchingNat44DeletionRec = d.Get("custom_port_batching_nat44_deletion_rec").(int)
	ret.Inst.CustomPortBatchingNat64CreationRec = d.Get("custom_port_batching_nat64_creation_rec").(int)
	ret.Inst.CustomPortBatchingNat64DeletionRec = d.Get("custom_port_batching_nat64_deletion_rec").(int)
	ret.Inst.CustomPortBatchingV2DsliteCreation = d.Get("custom_port_batching_v2_dslite_creation").(int)
	ret.Inst.CustomPortBatchingV2DsliteDeletion = d.Get("custom_port_batching_v2_dslite_deletion").(int)
	ret.Inst.CustomPortBatchingV2Nat44Creation = d.Get("custom_port_batching_v2_nat44_creation_").(int)
	ret.Inst.CustomPortBatchingV2Nat44Deletion = d.Get("custom_port_batching_v2_nat44_deletion_").(int)
	ret.Inst.CustomPortBatchingV2Nat64Creation = d.Get("custom_port_batching_v2_nat64_creation_").(int)
	ret.Inst.CustomPortBatchingV2Nat64Deletion = d.Get("custom_port_batching_v2_nat64_deletion_").(int)
	ret.Inst.CustomPortMappingDsliteCreationRec = d.Get("custom_port_mapping_dslite_creation_rec").(int)
	ret.Inst.CustomPortMappingDsliteDeletionRec = d.Get("custom_port_mapping_dslite_deletion_rec").(int)
	ret.Inst.CustomPortMappingNat44CreationReco = d.Get("custom_port_mapping_nat44_creation_reco").(int)
	ret.Inst.CustomPortMappingNat44DeletionReco = d.Get("custom_port_mapping_nat44_deletion_reco").(int)
	ret.Inst.CustomPortMappingNat64CreationReco = d.Get("custom_port_mapping_nat64_creation_reco").(int)
	ret.Inst.CustomPortMappingNat64DeletionReco = d.Get("custom_port_mapping_nat64_deletion_reco").(int)
	ret.Inst.CustomSessionEventDsliteCreationRe = d.Get("custom_session_event_dslite_creation_re").(int)
	ret.Inst.CustomSessionEventDsliteDeletionRe = d.Get("custom_session_event_dslite_deletion_re").(int)
	ret.Inst.CustomSessionEventFw4CreationRecor = d.Get("custom_session_event_fw4_creation_recor").(int)
	ret.Inst.CustomSessionEventFw4DeletionRecor = d.Get("custom_session_event_fw4_deletion_recor").(int)
	ret.Inst.CustomSessionEventFw6CreationRecor = d.Get("custom_session_event_fw6_creation_recor").(int)
	ret.Inst.CustomSessionEventFw6DeletionRecor = d.Get("custom_session_event_fw6_deletion_recor").(int)
	ret.Inst.CustomSessionEventNat44CreationRec = d.Get("custom_session_event_nat44_creation_rec").(int)
	ret.Inst.CustomSessionEventNat44DeletionRec = d.Get("custom_session_event_nat44_deletion_rec").(int)
	ret.Inst.CustomSessionEventNat64CreationRec = d.Get("custom_session_event_nat64_creation_rec").(int)
	ret.Inst.CustomSessionEventNat64DeletionRec = d.Get("custom_session_event_nat64_deletion_rec").(int)
	ret.Inst.DsliteRecordsSentFailure = d.Get("dslite_records_sent_failure").(int)
	ret.Inst.Nat44RecordsSentFailure = d.Get("nat44_records_sent_failure").(int)
	ret.Inst.Nat64RecordsSentFailure = d.Get("nat64_records_sent_failure").(int)
	ret.Inst.NetflowV5ExtRecordsSentFailure = d.Get("netflow_v5_ext_records_sent_failure").(int)
	ret.Inst.NetflowV5RecordsSentFailure = d.Get("netflow_v5_records_sent_failure").(int)
	ret.Inst.PortBatchingDsliteRecordsSentFailu = d.Get("port_batching_dslite_records_sent_failu").(int)
	ret.Inst.PortBatchingNat44RecordsSentFailur = d.Get("port_batching_nat44_records_sent_failur").(int)
	ret.Inst.PortBatchingNat64RecordsSentFailur = d.Get("port_batching_nat64_records_sent_failur").(int)
	ret.Inst.PortBatchingV2DsliteRecordsSentFa = d.Get("port_batching_v2_dslite_records_sent_fa").(int)
	ret.Inst.PortBatchingV2Nat44RecordsSentFai = d.Get("port_batching_v2_nat44_records_sent_fai").(int)
	ret.Inst.PortBatchingV2Nat64RecordsSentFai = d.Get("port_batching_v2_nat64_records_sent_fai").(int)
	ret.Inst.PortMappingDsliteRecordsSentFailur = d.Get("port_mapping_dslite_records_sent_failur").(int)
	ret.Inst.PortMappingNat44RecordsSentFailure = d.Get("port_mapping_nat44_records_sent_failure").(int)
	ret.Inst.PortMappingNat64RecordsSentFailure = d.Get("port_mapping_nat64_records_sent_failure").(int)
	ret.Inst.SessionEventDsliteRecordsSentFailu = d.Get("session_event_dslite_records_sent_failu").(int)
	ret.Inst.SessionEventFw4RecordsSentFailure = d.Get("session_event_fw4_records_sent_failure").(int)
	ret.Inst.SessionEventFw6RecordsSentFailure = d.Get("session_event_fw6_records_sent_failure").(int)
	ret.Inst.SessionEventNat44RecordsSentFailur = d.Get("session_event_nat44_records_sent_failur").(int)
	ret.Inst.SessionEventNat64RecordsSentFailur = d.Get("session_event_nat64_records_sent_failur").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
