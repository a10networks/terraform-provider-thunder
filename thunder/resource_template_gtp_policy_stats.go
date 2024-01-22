package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpPolicyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_template_gtp_policy_stats`: Statistics for the object gtp-policy\n\n__PLACEHOLDER__",
		ReadContext: resourceTemplateGtpPolicyStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Configure the GTP Policy Name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gtp_v0_c_tunnel_created": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Tunnel Created",
						},
						"gtp_v0_c_tunnel_half_open": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Half open tunnel created",
						},
						"gtp_v0_c_tunnel_half_closed": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Tunnel Delete Request",
						},
						"gtp_v0_c_tunnel_closed": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Tunnel Marked Deleted",
						},
						"gtp_v0_c_tunnel_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Tunnel Deleted",
						},
						"gtp_v0_c_half_open_tunnel_closed": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Half open tunnel closed",
						},
						"gtp_v1_c_tunnel_created": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Tunnel Created",
						},
						"gtp_v1_c_tunnel_half_open": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Half open tunnel created",
						},
						"gtp_v1_c_tunnel_half_closed": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Tunnel Delete Request",
						},
						"gtp_v1_c_tunnel_closed": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Tunnel Marked Deleted",
						},
						"gtp_v1_c_tunnel_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Tunnel Deleted",
						},
						"gtp_v1_c_half_open_tunnel_closed": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Half open tunnel closed",
						},
						"gtp_v2_c_tunnel_created": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Tunnel Created",
						},
						"gtp_v2_c_tunnel_half_open": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Half open tunnel created",
						},
						"gtp_v2_c_tunnel_half_closed": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Tunnel Delete Request",
						},
						"gtp_v2_c_tunnel_closed": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Tunnel Marked Deleted",
						},
						"gtp_v2_c_tunnel_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Tunnel Deleted",
						},
						"gtp_v2_c_half_open_tunnel_closed": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Half open tunnel closed",
						},
						"gtp_u_tunnel_created": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Tunnel Created",
						},
						"gtp_u_tunnel_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Tunnel Deleted",
						},
						"gtp_v0_c_update_pdp_resp_unsuccess": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Update PDP Context Response Unsuccessful",
						},
						"gtp_v1_c_update_pdp_resp_unsuccess": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Update PDP Context Response Unsuccessful",
						},
						"gtp_v2_c_mod_bearer_resp_unsuccess": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Modify Bearer Response Unsuccessful",
						},
						"gtp_v0_c_create_pdp_resp_unsuccess": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Create PDP Context Response Unsuccessful",
						},
						"gtp_v1_c_create_pdp_resp_unsuccess": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Create PDP Context Response Unsuccessful",
						},
						"gtp_v2_c_create_sess_resp_unsuccess": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Create Session Response Unsuccessful",
						},
						"gtp_v2_c_piggyback_message": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Piggyback Message",
						},
						"gtp_path_management_message": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Path Management Messages Received",
						},
						"gtp_v0_c_tunnel_deleted_restart": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Tunnel Deleted with Restart/failure",
						},
						"gtp_v1_c_tunnel_deleted_restart": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Tunnel Deleted with Restart/failure",
						},
						"gtp_v2_c_tunnel_deleted_restart": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Tunnel Deleted with Restart/failure",
						},
						"gtp_v0_c_reserved_message_allow": {
							Type: schema.TypeInt, Optional: true, Description: "Permit GTPv0-C Reserved Messages",
						},
						"gtp_v1_c_reserved_message_allow": {
							Type: schema.TypeInt, Optional: true, Description: "Permit GTPv1-C Reserved Messages",
						},
						"gtp_v2_c_reserved_message_allow": {
							Type: schema.TypeInt, Optional: true, Description: "Permit GTPv2-C Reserved Messages",
						},
						"gtp_v2_c_load_contr_info_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Load Control Info IEs in message exceeded 2",
						},
						"gtp_v1_c_pdu_notification_request_forward": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C PDU Notification Request Forward",
						},
						"gtp_v1_c_pdu_notification_reject_request_forward": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C PDU Notification Reject Request Forward",
						},
						"gtp_v0_c_pdu_notification_request_forward": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C PDU Notification Request Forward",
						},
						"gtp_v0_c_pdu_notification_reject_request_forward": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C PDU Notification Reject Request Forward",
						},
						"gtp_v0_c_message_skipped_apn_filtering_no_imsi": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C APN/IMSI Filtering Skipped (No IMSI)",
						},
						"gtp_v1_c_message_skipped_apn_filtering_no_imsi": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C APN/IMSI Filtering Skipped (No IMSI)",
						},
						"gtp_v2_c_message_skipped_apn_filtering_no_imsi": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C APN/IMSI Filtering Skipped (No IMSI)",
						},
						"gtp_v0_c_message_skipped_msisdn_filtering_no_imsi": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C MSISDN Filtering Skipped (No MSISDN)",
						},
						"gtp_v1_c_message_skipped_msisdn_filtering_no_imsi": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C MSISDN Filtering Skipped (No MSISDN)",
						},
						"gtp_v2_c_message_skipped_msisdn_filtering_no_imsi": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C MSISDN Filtering Skipped (No MSISDN)",
						},
						"gtp_v0_c_packet_dummy_msisdn": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Packet With Dummy MSISDN Forwarded",
						},
						"gtp_v1_c_packet_dummy_msisdn": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Packet With Dummy MSISDN Forwarded",
						},
						"gtp_v2_c_packet_dummy_msisdn": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Packet With Dummy MSISDN Forwarded",
						},
						"drop_vld_sanity_gtp_v2_c_message_with_teid_zero_expected": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv2-C Create Session Request with TEID",
						},
						"drop_vld_sanity_gtp_v1_c_message_with_teid_zero_expected": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv1-C PDU Notification Request with TEID",
						},
						"drop_vld_sanity_gtp_v0_c_message_with_teid_zero_expected": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv0-C PDU Notification Request with TEID",
						},
						"drop_vld_gtp_ie_repeat_count_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTP repeated IE count exceeded",
						},
						"drop_vld_reserved_field_set": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Reserved Header Field Set",
						},
						"drop_vld_tunnel_id_flag": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Tunnel Header Flag Not Set",
						},
						"drop_vld_invalid_flow_label_v0": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Invalid Flow Label in GTPv0-C Header",
						},
						"drop_vld_invalid_teid": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Invalid TEID Value",
						},
						"drop_vld_out_of_state": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Out Of State GTP Message",
						},
						"drop_vld_mandatory_information_element": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Mandatory IE Not Present",
						},
						"drop_vld_mandatory_ie_in_grouped_ie": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Mandatory IE in Grouped IE Not Present",
						},
						"drop_vld_out_of_order_ie": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv1-C Message Out of Order IE",
						},
						"drop_vld_out_of_state_ie": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Unexpected IE Present in Message",
						},
						"drop_vld_reserved_information_element": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Reserved IE Field Present",
						},
						"drop_vld_version_not_supported": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Invalid GTP version",
						},
						"drop_vld_message_length": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Message Length Exceeded",
						},
						"drop_vld_cross_layer_correlation": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Cross Layer IP Address Mismatch",
						},
						"drop_vld_country_code_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Country Code Mismatch in IMSI and MSISDN",
						},
						"drop_vld_gtp_u_spoofed_source_address": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTP-U IP Address Spoofed",
						},
						"drop_vld_gtp_bearer_count_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTP Bearer count exceeded max (11)",
						},
						"drop_vld_gtp_v2_wrong_lbi_create_bearer_req": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPV2-C Wrong LBI in Create Bearer Request",
						},
						"gtp_c_handover_in_progress_with_conn": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-C matching a conn with Handover In Progress",
						},
						"drop_vld_v0_reserved_message_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv0-C Reserved Message Drop",
						},
						"drop_vld_v1_reserved_message_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv1-C Reserved Message Drop",
						},
						"drop_vld_v2_reserved_message_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv2-C Reserved Message Drop",
						},
						"drop_vld_invalid_pkt_len_piggyback": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Piggyback message invalid packet length",
						},
						"drop_vld_sanity_failed_piggyback": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: piggyback message anomaly failed",
						},
						"drop_vld_sequence_num_correlation": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTP-C Sequence number Mismatch",
						},
						"drop_vld_gtpv0_seqnum_buffer_full": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPV0-C conn Sequence number Buffer Full",
						},
						"drop_vld_gtpv1_seqnum_buffer_full": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPV1-C conn Sequence number Buffer Full",
						},
						"drop_vld_gtpv2_seqnum_buffer_full": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPV2-C conn Sequence number Buffer Full",
						},
						"drop_vld_gtp_invalid_imsi_len_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTP-C Invalid IMSI Length Drop",
						},
						"drop_vld_gtp_invalid_apn_len_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTP-C Invalid APN Length Drop",
						},
						"drop_vld_protocol_flag_unset": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Protocol flag in Header Field not Set",
						},
						"drop_vld_gtpv0_subscriber_attr_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPV0-c Subscriber Attributes Missing",
						},
						"drop_vld_gtpv1_subscriber_attr_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPV1-c Subscriber Attributes Missing",
						},
						"drop_vld_gtpv2_subscriber_attr_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPV2-c Subscriber Attributes Missing",
						},
						"drop_vld_gtp_v0_c_ie_len_exceed_msg_len": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C IE Length Exceeds Message Length",
						},
						"drop_vld_gtp_v1_c_ie_len_exceed_msg_len": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C IE Length Exceeds Message Length",
						},
						"drop_vld_gtp_v2_c_ie_len_exceed_msg_len": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C IE Length Exceeds Message Length",
						},
						"drop_vld_gtp_v0_c_message_length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Message Length Mismatch Across Layers",
						},
						"drop_vld_gtp_v1_c_message_length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Message Length Mismatch Across Layers",
						},
						"drop_vld_gtp_v2_c_message_length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Message Length Mismatch Across Layers",
						},
						"drop_vld_gtp_v0_c_message_skipped_apn_filtering_no_apn": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv0-C APN/IMSI Filtering Dropped (No APN)",
						},
						"drop_vld_gtp_v1_c_message_skipped_apn_filtering_no_apn": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv1-C APN/IMSI Filtering Dropped (No APN)",
						},
						"drop_vld_gtp_v2_c_message_skipped_apn_filtering_no_apn": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv2-C APN/IMSI Filtering Dropped (No APN)",
						},
						"drop_flt_message_filtering": {
							Type: schema.TypeInt, Optional: true, Description: "Filtering Drop: Message Type Not Permitted on Interface",
						},
						"drop_flt_apn_filtering": {
							Type: schema.TypeInt, Optional: true, Description: "Filtering Drop: APN IMSI Filtering",
						},
						"drop_flt_msisdn_filtering": {
							Type: schema.TypeInt, Optional: true, Description: "Filtering Drop: MSISDN Filtering",
						},
						"drop_flt_rat_type_filtering": {
							Type: schema.TypeInt, Optional: true, Description: "Filtering Drop: RAT Type Filtering",
						},
						"drop_flt_gtp_in_gtp": {
							Type: schema.TypeInt, Optional: true, Description: "Filtering Drop: GTP in GTP Tunnel Present",
						},
						"drop_rl_gtp_v0_c_agg": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: Maximum GTPv0-C Message rate",
						},
						"drop_rl_gtp_v1_c_agg": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: Maximum GTPv1-C Message rate",
						},
						"drop_rl_gtp_v2_c_agg": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: Maximum GTPv2-C Message rate",
						},
						"drop_rl_gtp_v1_c_create_pdp_request": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTPv1-C Create PDP Request rate",
						},
						"drop_rl_gtp_v2_c_create_session_request": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTPv2-C Create Session Request rate",
						},
						"drop_rl_gtp_v1_c_update_pdp_request": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTPv1-C Update PDP Request rate",
						},
						"drop_rl_gtp_v2_c_modify_bearer_request": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTPv2-C Modify Bearer Request rate",
						},
						"drop_rl_gtp_u_tunnel_create": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTP-U Tunnel Creation rate",
						},
						"drop_rl_gtp_u_uplink_byte": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTP-U Uplink byte rate",
						},
						"drop_rl_gtp_u_uplink_packet": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTP-U Uplink packet rate",
						},
						"drop_rl_gtp_u_downlink_byte": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTP-U Downlink byte rate",
						},
						"drop_rl_gtp_u_downlink_packet": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTP-U Downlink packet rate",
						},
						"drop_rl_gtp_u_total_byte": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTP-U Total byte rate",
						},
						"drop_rl_gtp_u_total_packet": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTP-U Total packet rate",
						},
						"drop_rl_gtp_u_max_concurrent_tunnels": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTP-U Concurrent Tunnels",
						},
					},
				},
			},
		},
	}
}

func resourceTemplateGtpPolicyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpPolicyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpPolicyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		TemplateGtpPolicyStatsStats := setObjectTemplateGtpPolicyStatsStats(res)
		d.Set("stats", TemplateGtpPolicyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectTemplateGtpPolicyStatsStats(ret edpt.DataTemplateGtpPolicyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"gtp_v0_c_tunnel_created":                                  ret.DtTemplateGtpPolicyStats.Stats.GtpV0CTunnelCreated,
			"gtp_v0_c_tunnel_half_open":                                ret.DtTemplateGtpPolicyStats.Stats.GtpV0CTunnelHalfOpen,
			"gtp_v0_c_tunnel_half_closed":                              ret.DtTemplateGtpPolicyStats.Stats.GtpV0CTunnelHalfClosed,
			"gtp_v0_c_tunnel_closed":                                   ret.DtTemplateGtpPolicyStats.Stats.GtpV0CTunnelClosed,
			"gtp_v0_c_tunnel_deleted":                                  ret.DtTemplateGtpPolicyStats.Stats.GtpV0CTunnelDeleted,
			"gtp_v0_c_half_open_tunnel_closed":                         ret.DtTemplateGtpPolicyStats.Stats.GtpV0CHalfOpenTunnelClosed,
			"gtp_v1_c_tunnel_created":                                  ret.DtTemplateGtpPolicyStats.Stats.GtpV1CTunnelCreated,
			"gtp_v1_c_tunnel_half_open":                                ret.DtTemplateGtpPolicyStats.Stats.GtpV1CTunnelHalfOpen,
			"gtp_v1_c_tunnel_half_closed":                              ret.DtTemplateGtpPolicyStats.Stats.GtpV1CTunnelHalfClosed,
			"gtp_v1_c_tunnel_closed":                                   ret.DtTemplateGtpPolicyStats.Stats.GtpV1CTunnelClosed,
			"gtp_v1_c_tunnel_deleted":                                  ret.DtTemplateGtpPolicyStats.Stats.GtpV1CTunnelDeleted,
			"gtp_v1_c_half_open_tunnel_closed":                         ret.DtTemplateGtpPolicyStats.Stats.GtpV1CHalfOpenTunnelClosed,
			"gtp_v2_c_tunnel_created":                                  ret.DtTemplateGtpPolicyStats.Stats.GtpV2CTunnelCreated,
			"gtp_v2_c_tunnel_half_open":                                ret.DtTemplateGtpPolicyStats.Stats.GtpV2CTunnelHalfOpen,
			"gtp_v2_c_tunnel_half_closed":                              ret.DtTemplateGtpPolicyStats.Stats.GtpV2CTunnelHalfClosed,
			"gtp_v2_c_tunnel_closed":                                   ret.DtTemplateGtpPolicyStats.Stats.GtpV2CTunnelClosed,
			"gtp_v2_c_tunnel_deleted":                                  ret.DtTemplateGtpPolicyStats.Stats.GtpV2CTunnelDeleted,
			"gtp_v2_c_half_open_tunnel_closed":                         ret.DtTemplateGtpPolicyStats.Stats.GtpV2CHalfOpenTunnelClosed,
			"gtp_u_tunnel_created":                                     ret.DtTemplateGtpPolicyStats.Stats.GtpUTunnelCreated,
			"gtp_u_tunnel_deleted":                                     ret.DtTemplateGtpPolicyStats.Stats.GtpUTunnelDeleted,
			"gtp_v0_c_update_pdp_resp_unsuccess":                       ret.DtTemplateGtpPolicyStats.Stats.GtpV0CUpdatePdpRespUnsuccess,
			"gtp_v1_c_update_pdp_resp_unsuccess":                       ret.DtTemplateGtpPolicyStats.Stats.GtpV1CUpdatePdpRespUnsuccess,
			"gtp_v2_c_mod_bearer_resp_unsuccess":                       ret.DtTemplateGtpPolicyStats.Stats.GtpV2CMod_bearerRespUnsuccess,
			"gtp_v0_c_create_pdp_resp_unsuccess":                       ret.DtTemplateGtpPolicyStats.Stats.GtpV0CCreatePdpRespUnsuccess,
			"gtp_v1_c_create_pdp_resp_unsuccess":                       ret.DtTemplateGtpPolicyStats.Stats.GtpV1CCreatePdpRespUnsuccess,
			"gtp_v2_c_create_sess_resp_unsuccess":                      ret.DtTemplateGtpPolicyStats.Stats.GtpV2CCreateSessRespUnsuccess,
			"gtp_v2_c_piggyback_message":                               ret.DtTemplateGtpPolicyStats.Stats.GtpV2CPiggybackMessage,
			"gtp_path_management_message":                              ret.DtTemplateGtpPolicyStats.Stats.GtpPathManagementMessage,
			"gtp_v0_c_tunnel_deleted_restart":                          ret.DtTemplateGtpPolicyStats.Stats.GtpV0CTunnelDeletedRestart,
			"gtp_v1_c_tunnel_deleted_restart":                          ret.DtTemplateGtpPolicyStats.Stats.GtpV1CTunnelDeletedRestart,
			"gtp_v2_c_tunnel_deleted_restart":                          ret.DtTemplateGtpPolicyStats.Stats.GtpV2CTunnelDeletedRestart,
			"gtp_v0_c_reserved_message_allow":                          ret.DtTemplateGtpPolicyStats.Stats.GtpV0CReservedMessageAllow,
			"gtp_v1_c_reserved_message_allow":                          ret.DtTemplateGtpPolicyStats.Stats.GtpV1CReservedMessageAllow,
			"gtp_v2_c_reserved_message_allow":                          ret.DtTemplateGtpPolicyStats.Stats.GtpV2CReservedMessageAllow,
			"gtp_v2_c_load_contr_info_exceed":                          ret.DtTemplateGtpPolicyStats.Stats.GtpV2CLoadContrInfoExceed,
			"gtp_v1_c_pdu_notification_request_forward":                ret.DtTemplateGtpPolicyStats.Stats.GtpV1CPduNotificationRequestForward,
			"gtp_v1_c_pdu_notification_reject_request_forward":         ret.DtTemplateGtpPolicyStats.Stats.GtpV1CPduNotificationRejectRequestForward,
			"gtp_v0_c_pdu_notification_request_forward":                ret.DtTemplateGtpPolicyStats.Stats.GtpV0CPduNotificationRequestForward,
			"gtp_v0_c_pdu_notification_reject_request_forward":         ret.DtTemplateGtpPolicyStats.Stats.GtpV0CPduNotificationRejectRequestForward,
			"gtp_v0_c_message_skipped_apn_filtering_no_imsi":           ret.DtTemplateGtpPolicyStats.Stats.GtpV0CMessageSkippedApnFilteringNoImsi,
			"gtp_v1_c_message_skipped_apn_filtering_no_imsi":           ret.DtTemplateGtpPolicyStats.Stats.GtpV1CMessageSkippedApnFilteringNoImsi,
			"gtp_v2_c_message_skipped_apn_filtering_no_imsi":           ret.DtTemplateGtpPolicyStats.Stats.GtpV2CMessageSkippedApnFilteringNoImsi,
			"gtp_v0_c_message_skipped_msisdn_filtering_no_imsi":        ret.DtTemplateGtpPolicyStats.Stats.GtpV0CMessageSkippedMsisdnFilteringNoImsi,
			"gtp_v1_c_message_skipped_msisdn_filtering_no_imsi":        ret.DtTemplateGtpPolicyStats.Stats.GtpV1CMessageSkippedMsisdnFilteringNoImsi,
			"gtp_v2_c_message_skipped_msisdn_filtering_no_imsi":        ret.DtTemplateGtpPolicyStats.Stats.GtpV2CMessageSkippedMsisdnFilteringNoImsi,
			"gtp_v0_c_packet_dummy_msisdn":                             ret.DtTemplateGtpPolicyStats.Stats.GtpV0CPacketDummyMsisdn,
			"gtp_v1_c_packet_dummy_msisdn":                             ret.DtTemplateGtpPolicyStats.Stats.GtpV1CPacketDummyMsisdn,
			"gtp_v2_c_packet_dummy_msisdn":                             ret.DtTemplateGtpPolicyStats.Stats.GtpV2CPacketDummyMsisdn,
			"drop_vld_sanity_gtp_v2_c_message_with_teid_zero_expected": ret.DtTemplateGtpPolicyStats.Stats.DropVldSanityGtpV2CMessageWithTeidZeroExpected,
			"drop_vld_sanity_gtp_v1_c_message_with_teid_zero_expected": ret.DtTemplateGtpPolicyStats.Stats.DropVldSanityGtpV1CMessageWithTeidZeroExpected,
			"drop_vld_sanity_gtp_v0_c_message_with_teid_zero_expected": ret.DtTemplateGtpPolicyStats.Stats.DropVldSanityGtpV0CMessageWithTeidZeroExpected,
			"drop_vld_gtp_ie_repeat_count_exceed":                      ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpIeRepeatCountExceed,
			"drop_vld_reserved_field_set":                              ret.DtTemplateGtpPolicyStats.Stats.DropVldReservedFieldSet,
			"drop_vld_tunnel_id_flag":                                  ret.DtTemplateGtpPolicyStats.Stats.DropVldTunnelIdFlag,
			"drop_vld_invalid_flow_label_v0":                           ret.DtTemplateGtpPolicyStats.Stats.DropVldInvalidFlowLabelV0,
			"drop_vld_invalid_teid":                                    ret.DtTemplateGtpPolicyStats.Stats.DropVldInvalidTeid,
			"drop_vld_out_of_state":                                    ret.DtTemplateGtpPolicyStats.Stats.DropVldOutOfState,
			"drop_vld_mandatory_information_element":                   ret.DtTemplateGtpPolicyStats.Stats.DropVldMandatoryInformationElement,
			"drop_vld_mandatory_ie_in_grouped_ie":                      ret.DtTemplateGtpPolicyStats.Stats.DropVldMandatoryIeInGroupedIe,
			"drop_vld_out_of_order_ie":                                 ret.DtTemplateGtpPolicyStats.Stats.DropVldOutOfOrderIe,
			"drop_vld_out_of_state_ie":                                 ret.DtTemplateGtpPolicyStats.Stats.DropVldOutOfStateIe,
			"drop_vld_reserved_information_element":                    ret.DtTemplateGtpPolicyStats.Stats.DropVldReservedInformationElement,
			"drop_vld_version_not_supported":                           ret.DtTemplateGtpPolicyStats.Stats.DropVldVersionNotSupported,
			"drop_vld_message_length":                                  ret.DtTemplateGtpPolicyStats.Stats.DropVldMessageLength,
			"drop_vld_cross_layer_correlation":                         ret.DtTemplateGtpPolicyStats.Stats.DropVldCrossLayerCorrelation,
			"drop_vld_country_code_mismatch":                           ret.DtTemplateGtpPolicyStats.Stats.DropVldCountryCodeMismatch,
			"drop_vld_gtp_u_spoofed_source_address":                    ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpUSpoofedSourceAddress,
			"drop_vld_gtp_bearer_count_exceed":                         ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpBearerCountExceed,
			"drop_vld_gtp_v2_wrong_lbi_create_bearer_req":              ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpV2WrongLbiCreateBearerReq,
			"gtp_c_handover_in_progress_with_conn":                     ret.DtTemplateGtpPolicyStats.Stats.GtpCHandoverInProgressWithConn,
			"drop_vld_v0_reserved_message_drop":                        ret.DtTemplateGtpPolicyStats.Stats.DropVldV0ReservedMessageDrop,
			"drop_vld_v1_reserved_message_drop":                        ret.DtTemplateGtpPolicyStats.Stats.DropVldV1ReservedMessageDrop,
			"drop_vld_v2_reserved_message_drop":                        ret.DtTemplateGtpPolicyStats.Stats.DropVldV2ReservedMessageDrop,
			"drop_vld_invalid_pkt_len_piggyback":                       ret.DtTemplateGtpPolicyStats.Stats.DropVldInvalidPktLenPiggyback,
			"drop_vld_sanity_failed_piggyback":                         ret.DtTemplateGtpPolicyStats.Stats.DropVldSanityFailedPiggyback,
			"drop_vld_sequence_num_correlation":                        ret.DtTemplateGtpPolicyStats.Stats.DropVldSequenceNumCorrelation,
			"drop_vld_gtpv0_seqnum_buffer_full":                        ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpv0SeqnumBufferFull,
			"drop_vld_gtpv1_seqnum_buffer_full":                        ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpv1SeqnumBufferFull,
			"drop_vld_gtpv2_seqnum_buffer_full":                        ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpv2SeqnumBufferFull,
			"drop_vld_gtp_invalid_imsi_len_drop":                       ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpInvalidImsiLenDrop,
			"drop_vld_gtp_invalid_apn_len_drop":                        ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpInvalidApnLenDrop,
			"drop_vld_protocol_flag_unset":                             ret.DtTemplateGtpPolicyStats.Stats.DropVldProtocolFlagUnset,
			"drop_vld_gtpv0_subscriber_attr_miss":                      ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpv0SubscriberAttrMiss,
			"drop_vld_gtpv1_subscriber_attr_miss":                      ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpv1SubscriberAttrMiss,
			"drop_vld_gtpv2_subscriber_attr_miss":                      ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpv2SubscriberAttrMiss,
			"drop_vld_gtp_v0_c_ie_len_exceed_msg_len":                  ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpV0CIeLenExceedMsgLen,
			"drop_vld_gtp_v1_c_ie_len_exceed_msg_len":                  ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpV1CIeLenExceedMsgLen,
			"drop_vld_gtp_v2_c_ie_len_exceed_msg_len":                  ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpV2CIeLenExceedMsgLen,
			"drop_vld_gtp_v0_c_message_length_mismatch":                ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpV0CMessageLengthMismatch,
			"drop_vld_gtp_v1_c_message_length_mismatch":                ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpV1CMessageLengthMismatch,
			"drop_vld_gtp_v2_c_message_length_mismatch":                ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpV2CMessageLengthMismatch,
			"drop_vld_gtp_v0_c_message_skipped_apn_filtering_no_apn":   ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpV0CMessageSkippedApnFilteringNoApn,
			"drop_vld_gtp_v1_c_message_skipped_apn_filtering_no_apn":   ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpV1CMessageSkippedApnFilteringNoApn,
			"drop_vld_gtp_v2_c_message_skipped_apn_filtering_no_apn":   ret.DtTemplateGtpPolicyStats.Stats.DropVldGtpV2CMessageSkippedApnFilteringNoApn,
			"drop_flt_message_filtering":                               ret.DtTemplateGtpPolicyStats.Stats.DropFltMessageFiltering,
			"drop_flt_apn_filtering":                                   ret.DtTemplateGtpPolicyStats.Stats.DropFltApnFiltering,
			"drop_flt_msisdn_filtering":                                ret.DtTemplateGtpPolicyStats.Stats.DropFltMsisdnFiltering,
			"drop_flt_rat_type_filtering":                              ret.DtTemplateGtpPolicyStats.Stats.DropFltRatTypeFiltering,
			"drop_flt_gtp_in_gtp":                                      ret.DtTemplateGtpPolicyStats.Stats.DropFltGtpInGtp,
			"drop_rl_gtp_v0_c_agg":                                     ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpV0CAgg,
			"drop_rl_gtp_v1_c_agg":                                     ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpV1CAgg,
			"drop_rl_gtp_v2_c_agg":                                     ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpV2CAgg,
			"drop_rl_gtp_v1_c_create_pdp_request":                      ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpV1CCreatePdpRequest,
			"drop_rl_gtp_v2_c_create_session_request":                  ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpV2CCreateSessionRequest,
			"drop_rl_gtp_v1_c_update_pdp_request":                      ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpV1CUpdatePdpRequest,
			"drop_rl_gtp_v2_c_modify_bearer_request":                   ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpV2CModifyBearerRequest,
			"drop_rl_gtp_u_tunnel_create":                              ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpUTunnelCreate,
			"drop_rl_gtp_u_uplink_byte":                                ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpUUplinkByte,
			"drop_rl_gtp_u_uplink_packet":                              ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpUUplinkPacket,
			"drop_rl_gtp_u_downlink_byte":                              ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpUDownlinkByte,
			"drop_rl_gtp_u_downlink_packet":                            ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpUDownlinkPacket,
			"drop_rl_gtp_u_total_byte":                                 ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpUTotalByte,
			"drop_rl_gtp_u_total_packet":                               ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpUTotalPacket,
			"drop_rl_gtp_u_max_concurrent_tunnels":                     ret.DtTemplateGtpPolicyStats.Stats.DropRlGtpUMaxConcurrentTunnels,
		},
	}
}

func getObjectTemplateGtpPolicyStatsStats(d []interface{}) edpt.TemplateGtpPolicyStatsStats {

	count1 := len(d)
	var ret edpt.TemplateGtpPolicyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GtpV0CTunnelCreated = in["gtp_v0_c_tunnel_created"].(int)
		ret.GtpV0CTunnelHalfOpen = in["gtp_v0_c_tunnel_half_open"].(int)
		ret.GtpV0CTunnelHalfClosed = in["gtp_v0_c_tunnel_half_closed"].(int)
		ret.GtpV0CTunnelClosed = in["gtp_v0_c_tunnel_closed"].(int)
		ret.GtpV0CTunnelDeleted = in["gtp_v0_c_tunnel_deleted"].(int)
		ret.GtpV0CHalfOpenTunnelClosed = in["gtp_v0_c_half_open_tunnel_closed"].(int)
		ret.GtpV1CTunnelCreated = in["gtp_v1_c_tunnel_created"].(int)
		ret.GtpV1CTunnelHalfOpen = in["gtp_v1_c_tunnel_half_open"].(int)
		ret.GtpV1CTunnelHalfClosed = in["gtp_v1_c_tunnel_half_closed"].(int)
		ret.GtpV1CTunnelClosed = in["gtp_v1_c_tunnel_closed"].(int)
		ret.GtpV1CTunnelDeleted = in["gtp_v1_c_tunnel_deleted"].(int)
		ret.GtpV1CHalfOpenTunnelClosed = in["gtp_v1_c_half_open_tunnel_closed"].(int)
		ret.GtpV2CTunnelCreated = in["gtp_v2_c_tunnel_created"].(int)
		ret.GtpV2CTunnelHalfOpen = in["gtp_v2_c_tunnel_half_open"].(int)
		ret.GtpV2CTunnelHalfClosed = in["gtp_v2_c_tunnel_half_closed"].(int)
		ret.GtpV2CTunnelClosed = in["gtp_v2_c_tunnel_closed"].(int)
		ret.GtpV2CTunnelDeleted = in["gtp_v2_c_tunnel_deleted"].(int)
		ret.GtpV2CHalfOpenTunnelClosed = in["gtp_v2_c_half_open_tunnel_closed"].(int)
		ret.GtpUTunnelCreated = in["gtp_u_tunnel_created"].(int)
		ret.GtpUTunnelDeleted = in["gtp_u_tunnel_deleted"].(int)
		ret.GtpV0CUpdatePdpRespUnsuccess = in["gtp_v0_c_update_pdp_resp_unsuccess"].(int)
		ret.GtpV1CUpdatePdpRespUnsuccess = in["gtp_v1_c_update_pdp_resp_unsuccess"].(int)
		ret.GtpV2CMod_bearerRespUnsuccess = in["gtp_v2_c_mod_bearer_resp_unsuccess"].(int)
		ret.GtpV0CCreatePdpRespUnsuccess = in["gtp_v0_c_create_pdp_resp_unsuccess"].(int)
		ret.GtpV1CCreatePdpRespUnsuccess = in["gtp_v1_c_create_pdp_resp_unsuccess"].(int)
		ret.GtpV2CCreateSessRespUnsuccess = in["gtp_v2_c_create_sess_resp_unsuccess"].(int)
		ret.GtpV2CPiggybackMessage = in["gtp_v2_c_piggyback_message"].(int)
		ret.GtpPathManagementMessage = in["gtp_path_management_message"].(int)
		ret.GtpV0CTunnelDeletedRestart = in["gtp_v0_c_tunnel_deleted_restart"].(int)
		ret.GtpV1CTunnelDeletedRestart = in["gtp_v1_c_tunnel_deleted_restart"].(int)
		ret.GtpV2CTunnelDeletedRestart = in["gtp_v2_c_tunnel_deleted_restart"].(int)
		ret.GtpV0CReservedMessageAllow = in["gtp_v0_c_reserved_message_allow"].(int)
		ret.GtpV1CReservedMessageAllow = in["gtp_v1_c_reserved_message_allow"].(int)
		ret.GtpV2CReservedMessageAllow = in["gtp_v2_c_reserved_message_allow"].(int)
		ret.GtpV2CLoadContrInfoExceed = in["gtp_v2_c_load_contr_info_exceed"].(int)
		ret.GtpV1CPduNotificationRequestForward = in["gtp_v1_c_pdu_notification_request_forward"].(int)
		ret.GtpV1CPduNotificationRejectRequestForward = in["gtp_v1_c_pdu_notification_reject_request_forward"].(int)
		ret.GtpV0CPduNotificationRequestForward = in["gtp_v0_c_pdu_notification_request_forward"].(int)
		ret.GtpV0CPduNotificationRejectRequestForward = in["gtp_v0_c_pdu_notification_reject_request_forward"].(int)
		ret.GtpV0CMessageSkippedApnFilteringNoImsi = in["gtp_v0_c_message_skipped_apn_filtering_no_imsi"].(int)
		ret.GtpV1CMessageSkippedApnFilteringNoImsi = in["gtp_v1_c_message_skipped_apn_filtering_no_imsi"].(int)
		ret.GtpV2CMessageSkippedApnFilteringNoImsi = in["gtp_v2_c_message_skipped_apn_filtering_no_imsi"].(int)
		ret.GtpV0CMessageSkippedMsisdnFilteringNoImsi = in["gtp_v0_c_message_skipped_msisdn_filtering_no_imsi"].(int)
		ret.GtpV1CMessageSkippedMsisdnFilteringNoImsi = in["gtp_v1_c_message_skipped_msisdn_filtering_no_imsi"].(int)
		ret.GtpV2CMessageSkippedMsisdnFilteringNoImsi = in["gtp_v2_c_message_skipped_msisdn_filtering_no_imsi"].(int)
		ret.GtpV0CPacketDummyMsisdn = in["gtp_v0_c_packet_dummy_msisdn"].(int)
		ret.GtpV1CPacketDummyMsisdn = in["gtp_v1_c_packet_dummy_msisdn"].(int)
		ret.GtpV2CPacketDummyMsisdn = in["gtp_v2_c_packet_dummy_msisdn"].(int)
		ret.DropVldSanityGtpV2CMessageWithTeidZeroExpected = in["drop_vld_sanity_gtp_v2_c_message_with_teid_zero_expected"].(int)
		ret.DropVldSanityGtpV1CMessageWithTeidZeroExpected = in["drop_vld_sanity_gtp_v1_c_message_with_teid_zero_expected"].(int)
		ret.DropVldSanityGtpV0CMessageWithTeidZeroExpected = in["drop_vld_sanity_gtp_v0_c_message_with_teid_zero_expected"].(int)
		ret.DropVldGtpIeRepeatCountExceed = in["drop_vld_gtp_ie_repeat_count_exceed"].(int)
		ret.DropVldReservedFieldSet = in["drop_vld_reserved_field_set"].(int)
		ret.DropVldTunnelIdFlag = in["drop_vld_tunnel_id_flag"].(int)
		ret.DropVldInvalidFlowLabelV0 = in["drop_vld_invalid_flow_label_v0"].(int)
		ret.DropVldInvalidTeid = in["drop_vld_invalid_teid"].(int)
		ret.DropVldOutOfState = in["drop_vld_out_of_state"].(int)
		ret.DropVldMandatoryInformationElement = in["drop_vld_mandatory_information_element"].(int)
		ret.DropVldMandatoryIeInGroupedIe = in["drop_vld_mandatory_ie_in_grouped_ie"].(int)
		ret.DropVldOutOfOrderIe = in["drop_vld_out_of_order_ie"].(int)
		ret.DropVldOutOfStateIe = in["drop_vld_out_of_state_ie"].(int)
		ret.DropVldReservedInformationElement = in["drop_vld_reserved_information_element"].(int)
		ret.DropVldVersionNotSupported = in["drop_vld_version_not_supported"].(int)
		ret.DropVldMessageLength = in["drop_vld_message_length"].(int)
		ret.DropVldCrossLayerCorrelation = in["drop_vld_cross_layer_correlation"].(int)
		ret.DropVldCountryCodeMismatch = in["drop_vld_country_code_mismatch"].(int)
		ret.DropVldGtpUSpoofedSourceAddress = in["drop_vld_gtp_u_spoofed_source_address"].(int)
		ret.DropVldGtpBearerCountExceed = in["drop_vld_gtp_bearer_count_exceed"].(int)
		ret.DropVldGtpV2WrongLbiCreateBearerReq = in["drop_vld_gtp_v2_wrong_lbi_create_bearer_req"].(int)
		ret.GtpCHandoverInProgressWithConn = in["gtp_c_handover_in_progress_with_conn"].(int)
		ret.DropVldV0ReservedMessageDrop = in["drop_vld_v0_reserved_message_drop"].(int)
		ret.DropVldV1ReservedMessageDrop = in["drop_vld_v1_reserved_message_drop"].(int)
		ret.DropVldV2ReservedMessageDrop = in["drop_vld_v2_reserved_message_drop"].(int)
		ret.DropVldInvalidPktLenPiggyback = in["drop_vld_invalid_pkt_len_piggyback"].(int)
		ret.DropVldSanityFailedPiggyback = in["drop_vld_sanity_failed_piggyback"].(int)
		ret.DropVldSequenceNumCorrelation = in["drop_vld_sequence_num_correlation"].(int)
		ret.DropVldGtpv0SeqnumBufferFull = in["drop_vld_gtpv0_seqnum_buffer_full"].(int)
		ret.DropVldGtpv1SeqnumBufferFull = in["drop_vld_gtpv1_seqnum_buffer_full"].(int)
		ret.DropVldGtpv2SeqnumBufferFull = in["drop_vld_gtpv2_seqnum_buffer_full"].(int)
		ret.DropVldGtpInvalidImsiLenDrop = in["drop_vld_gtp_invalid_imsi_len_drop"].(int)
		ret.DropVldGtpInvalidApnLenDrop = in["drop_vld_gtp_invalid_apn_len_drop"].(int)
		ret.DropVldProtocolFlagUnset = in["drop_vld_protocol_flag_unset"].(int)
		ret.DropVldGtpv0SubscriberAttrMiss = in["drop_vld_gtpv0_subscriber_attr_miss"].(int)
		ret.DropVldGtpv1SubscriberAttrMiss = in["drop_vld_gtpv1_subscriber_attr_miss"].(int)
		ret.DropVldGtpv2SubscriberAttrMiss = in["drop_vld_gtpv2_subscriber_attr_miss"].(int)
		ret.DropVldGtpV0CIeLenExceedMsgLen = in["drop_vld_gtp_v0_c_ie_len_exceed_msg_len"].(int)
		ret.DropVldGtpV1CIeLenExceedMsgLen = in["drop_vld_gtp_v1_c_ie_len_exceed_msg_len"].(int)
		ret.DropVldGtpV2CIeLenExceedMsgLen = in["drop_vld_gtp_v2_c_ie_len_exceed_msg_len"].(int)
		ret.DropVldGtpV0CMessageLengthMismatch = in["drop_vld_gtp_v0_c_message_length_mismatch"].(int)
		ret.DropVldGtpV1CMessageLengthMismatch = in["drop_vld_gtp_v1_c_message_length_mismatch"].(int)
		ret.DropVldGtpV2CMessageLengthMismatch = in["drop_vld_gtp_v2_c_message_length_mismatch"].(int)
		ret.DropVldGtpV0CMessageSkippedApnFilteringNoApn = in["drop_vld_gtp_v0_c_message_skipped_apn_filtering_no_apn"].(int)
		ret.DropVldGtpV1CMessageSkippedApnFilteringNoApn = in["drop_vld_gtp_v1_c_message_skipped_apn_filtering_no_apn"].(int)
		ret.DropVldGtpV2CMessageSkippedApnFilteringNoApn = in["drop_vld_gtp_v2_c_message_skipped_apn_filtering_no_apn"].(int)
		ret.DropFltMessageFiltering = in["drop_flt_message_filtering"].(int)
		ret.DropFltApnFiltering = in["drop_flt_apn_filtering"].(int)
		ret.DropFltMsisdnFiltering = in["drop_flt_msisdn_filtering"].(int)
		ret.DropFltRatTypeFiltering = in["drop_flt_rat_type_filtering"].(int)
		ret.DropFltGtpInGtp = in["drop_flt_gtp_in_gtp"].(int)
		ret.DropRlGtpV0CAgg = in["drop_rl_gtp_v0_c_agg"].(int)
		ret.DropRlGtpV1CAgg = in["drop_rl_gtp_v1_c_agg"].(int)
		ret.DropRlGtpV2CAgg = in["drop_rl_gtp_v2_c_agg"].(int)
		ret.DropRlGtpV1CCreatePdpRequest = in["drop_rl_gtp_v1_c_create_pdp_request"].(int)
		ret.DropRlGtpV2CCreateSessionRequest = in["drop_rl_gtp_v2_c_create_session_request"].(int)
		ret.DropRlGtpV1CUpdatePdpRequest = in["drop_rl_gtp_v1_c_update_pdp_request"].(int)
		ret.DropRlGtpV2CModifyBearerRequest = in["drop_rl_gtp_v2_c_modify_bearer_request"].(int)
		ret.DropRlGtpUTunnelCreate = in["drop_rl_gtp_u_tunnel_create"].(int)
		ret.DropRlGtpUUplinkByte = in["drop_rl_gtp_u_uplink_byte"].(int)
		ret.DropRlGtpUUplinkPacket = in["drop_rl_gtp_u_uplink_packet"].(int)
		ret.DropRlGtpUDownlinkByte = in["drop_rl_gtp_u_downlink_byte"].(int)
		ret.DropRlGtpUDownlinkPacket = in["drop_rl_gtp_u_downlink_packet"].(int)
		ret.DropRlGtpUTotalByte = in["drop_rl_gtp_u_total_byte"].(int)
		ret.DropRlGtpUTotalPacket = in["drop_rl_gtp_u_total_packet"].(int)
		ret.DropRlGtpUMaxConcurrentTunnels = in["drop_rl_gtp_u_max_concurrent_tunnels"].(int)
	}
	return ret
}

func dataToEndpointTemplateGtpPolicyStats(d *schema.ResourceData) edpt.TemplateGtpPolicyStats {
	var ret edpt.TemplateGtpPolicyStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectTemplateGtpPolicyStatsStats(d.Get("stats").([]interface{}))
	return ret
}
