package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwGtpNetworkElementStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_gtp_network_element_stats`: Statistics for the object network-element\n\n__PLACEHOLDER__",
		ReadContext: resourceFwGtpNetworkElementStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"key_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"uplink_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Uplink Bytes",
						},
						"downlink_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Downlink Bytes",
						},
						"uplink_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Uplink Packets",
						},
						"downlink_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Downlink Packets",
						},
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
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Piggyback Messages",
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
						"gtp_node_restart_gtp_c": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Node Restoration due to Recovery IE in GTP-C Message",
						},
						"gtp_v0_c_reserved_message_allow": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Reserved Message Allow",
						},
						"gtp_v1_c_reserved_message_allow": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Reserved Message Allow",
						},
						"gtp_v2_c_reserved_message_allow": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Reserved Message Allow",
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
						"gtp_v0_c_message_skipped_msisdn_filtering_no_msisdn": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C MSISDN Filtering Skipped (No MSISDN)",
						},
						"gtp_v1_c_message_skipped_msisdn_filtering_no_msisdn": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C MSISDN Filtering Skipped (No MSISDN)",
						},
						"gtp_v2_c_message_skipped_msisdn_filtering_no_msisdn": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C MSISDN Filtering Skipped (No MSISDN)",
						},
						"gtp_v0_c_packet_dummy_msisdn": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Packet With Dummy MSISDN Forwarded",
						},
						"gtp_v1_c_packet_dummy_msisdn": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Packet With Dummy MSISDN Forwarded",
						},
						"gtp_v2_c_packet_dummy_msisdn": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Packet With Dummy MSISDN Forwarded",
						},
						"drop_vld_gtp_v2_c_message_with_teid_zero_expected": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv2-C Create Session Request with TEID",
						},
						"drop_vld_gtp_v1_c_message_with_teid_zero_expected": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv1-C PDU Notification Request with TEID",
						},
						"drop_vld_gtp_v0_c_message_with_teid_zero_expected": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv0-C PDU Notification Request with TEID",
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
						"drop_vld_unsupported_message_type": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Message type not supported by GTP Version",
						},
						"drop_vld_out_of_state": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Out Of State GTP Message",
						},
						"drop_vld_mandatory_information_element": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: Mandatory IE Not Present",
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
						"drop_vld_gtp_v0_c_message_dropped_apn_filtering_no_apn": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv0-C APN/IMSI Filtering dropped (No APN)",
						},
						"drop_vld_gtp_v1_c_message_dropped_apn_filtering_no_apn": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv1-C APN/IMSI Filtering dropped (No APN)",
						},
						"drop_vld_gtp_v2_c_message_dropped_apn_filtering_no_apn": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv2-C APN/IMSI Filtering dropped (No APN)",
						},
						"drop_vld_gtp_v0_c_ie_len_exceed_msg_len": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv0-C IE Length Exceeds Message Length",
						},
						"drop_vld_gtp_v1_c_ie_len_exceed_msg_len": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv1-C IE Length Exceeds Message Length",
						},
						"drop_vld_gtp_v2_c_ie_len_exceed_msg_len": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv2-C IE Length Exceeds Message Length",
						},
						"drop_vld_gtp_v0_c_message_length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv0-C Message Length Mismatch Across Layers",
						},
						"drop_vld_gtp_v1_c_message_length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv1-C Message Length Mismatch Across Layers",
						},
						"drop_vld_gtp_v2_c_message_length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Validation Drop: GTPv2-C Message Length Mismatch Across Layers",
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
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: Maximum GTPv0-C messages rate",
						},
						"drop_rl_gtp_v1_c_agg": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: Maximum GTPv1-C messages rate",
						},
						"drop_rl_gtp_v2_c_agg": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: Maximum GTPv2-C messages rate",
						},
						"drop_rl_gtp_v1_c_create_pdp_request": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTPv1-C Create PDP Req rate",
						},
						"drop_rl_gtp_v2_c_create_session_request": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTPv2-C Create Session Req rate",
						},
						"drop_rl_gtp_v1_c_update_pdp_request": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTPv1-C Update PDP Req rate",
						},
						"drop_rl_gtp_v2_c_modify_bearer_request": {
							Type: schema.TypeInt, Optional: true, Description: "Rate-limit Drop: GTPv2-C Modify Bearer Req rate",
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

func resourceFwGtpNetworkElementStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwGtpNetworkElementStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwGtpNetworkElementStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwGtpNetworkElementStatsStats := setObjectFwGtpNetworkElementStatsStats(res)
		d.Set("stats", FwGtpNetworkElementStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwGtpNetworkElementStatsStats(ret edpt.DataFwGtpNetworkElementStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"key_name":                                               ret.DtFwGtpNetworkElementStats.Stats.KeyName,
			"key_type":                                               ret.DtFwGtpNetworkElementStats.Stats.KeyType,
			"uplink_bytes":                                           ret.DtFwGtpNetworkElementStats.Stats.UplinkBytes,
			"downlink_bytes":                                         ret.DtFwGtpNetworkElementStats.Stats.DownlinkBytes,
			"uplink_pkts":                                            ret.DtFwGtpNetworkElementStats.Stats.UplinkPkts,
			"downlink_pkts":                                          ret.DtFwGtpNetworkElementStats.Stats.DownlinkPkts,
			"gtp_v0_c_tunnel_created":                                ret.DtFwGtpNetworkElementStats.Stats.GtpV0CTunnelCreated,
			"gtp_v0_c_tunnel_half_open":                              ret.DtFwGtpNetworkElementStats.Stats.GtpV0CTunnelHalfOpen,
			"gtp_v0_c_tunnel_half_closed":                            ret.DtFwGtpNetworkElementStats.Stats.GtpV0CTunnelHalfClosed,
			"gtp_v0_c_tunnel_closed":                                 ret.DtFwGtpNetworkElementStats.Stats.GtpV0CTunnelClosed,
			"gtp_v0_c_tunnel_deleted":                                ret.DtFwGtpNetworkElementStats.Stats.GtpV0CTunnelDeleted,
			"gtp_v0_c_half_open_tunnel_closed":                       ret.DtFwGtpNetworkElementStats.Stats.GtpV0CHalfOpenTunnelClosed,
			"gtp_v1_c_tunnel_created":                                ret.DtFwGtpNetworkElementStats.Stats.GtpV1CTunnelCreated,
			"gtp_v1_c_tunnel_half_open":                              ret.DtFwGtpNetworkElementStats.Stats.GtpV1CTunnelHalfOpen,
			"gtp_v1_c_tunnel_half_closed":                            ret.DtFwGtpNetworkElementStats.Stats.GtpV1CTunnelHalfClosed,
			"gtp_v1_c_tunnel_closed":                                 ret.DtFwGtpNetworkElementStats.Stats.GtpV1CTunnelClosed,
			"gtp_v1_c_tunnel_deleted":                                ret.DtFwGtpNetworkElementStats.Stats.GtpV1CTunnelDeleted,
			"gtp_v1_c_half_open_tunnel_closed":                       ret.DtFwGtpNetworkElementStats.Stats.GtpV1CHalfOpenTunnelClosed,
			"gtp_v2_c_tunnel_created":                                ret.DtFwGtpNetworkElementStats.Stats.GtpV2CTunnelCreated,
			"gtp_v2_c_tunnel_half_open":                              ret.DtFwGtpNetworkElementStats.Stats.GtpV2CTunnelHalfOpen,
			"gtp_v2_c_tunnel_half_closed":                            ret.DtFwGtpNetworkElementStats.Stats.GtpV2CTunnelHalfClosed,
			"gtp_v2_c_tunnel_closed":                                 ret.DtFwGtpNetworkElementStats.Stats.GtpV2CTunnelClosed,
			"gtp_v2_c_tunnel_deleted":                                ret.DtFwGtpNetworkElementStats.Stats.GtpV2CTunnelDeleted,
			"gtp_v2_c_half_open_tunnel_closed":                       ret.DtFwGtpNetworkElementStats.Stats.GtpV2CHalfOpenTunnelClosed,
			"gtp_u_tunnel_created":                                   ret.DtFwGtpNetworkElementStats.Stats.GtpUTunnelCreated,
			"gtp_u_tunnel_deleted":                                   ret.DtFwGtpNetworkElementStats.Stats.GtpUTunnelDeleted,
			"gtp_v0_c_update_pdp_resp_unsuccess":                     ret.DtFwGtpNetworkElementStats.Stats.GtpV0CUpdatePdpRespUnsuccess,
			"gtp_v1_c_update_pdp_resp_unsuccess":                     ret.DtFwGtpNetworkElementStats.Stats.GtpV1CUpdatePdpRespUnsuccess,
			"gtp_v2_c_mod_bearer_resp_unsuccess":                     ret.DtFwGtpNetworkElementStats.Stats.GtpV2CMod_bearerRespUnsuccess,
			"gtp_v0_c_create_pdp_resp_unsuccess":                     ret.DtFwGtpNetworkElementStats.Stats.GtpV0CCreatePdpRespUnsuccess,
			"gtp_v1_c_create_pdp_resp_unsuccess":                     ret.DtFwGtpNetworkElementStats.Stats.GtpV1CCreatePdpRespUnsuccess,
			"gtp_v2_c_create_sess_resp_unsuccess":                    ret.DtFwGtpNetworkElementStats.Stats.GtpV2CCreateSessRespUnsuccess,
			"gtp_v2_c_piggyback_message":                             ret.DtFwGtpNetworkElementStats.Stats.GtpV2CPiggybackMessage,
			"gtp_path_management_message":                            ret.DtFwGtpNetworkElementStats.Stats.GtpPathManagementMessage,
			"gtp_v0_c_tunnel_deleted_restart":                        ret.DtFwGtpNetworkElementStats.Stats.GtpV0CTunnelDeletedRestart,
			"gtp_v1_c_tunnel_deleted_restart":                        ret.DtFwGtpNetworkElementStats.Stats.GtpV1CTunnelDeletedRestart,
			"gtp_v2_c_tunnel_deleted_restart":                        ret.DtFwGtpNetworkElementStats.Stats.GtpV2CTunnelDeletedRestart,
			"gtp_node_restart_gtp_c":                                 ret.DtFwGtpNetworkElementStats.Stats.GtpNodeRestartGtpC,
			"gtp_v0_c_reserved_message_allow":                        ret.DtFwGtpNetworkElementStats.Stats.GtpV0CReservedMessageAllow,
			"gtp_v1_c_reserved_message_allow":                        ret.DtFwGtpNetworkElementStats.Stats.GtpV1CReservedMessageAllow,
			"gtp_v2_c_reserved_message_allow":                        ret.DtFwGtpNetworkElementStats.Stats.GtpV2CReservedMessageAllow,
			"gtp_v1_c_pdu_notification_request_forward":              ret.DtFwGtpNetworkElementStats.Stats.GtpV1CPduNotificationRequestForward,
			"gtp_v1_c_pdu_notification_reject_request_forward":       ret.DtFwGtpNetworkElementStats.Stats.GtpV1CPduNotificationRejectRequestForward,
			"gtp_v0_c_pdu_notification_request_forward":              ret.DtFwGtpNetworkElementStats.Stats.GtpV0CPduNotificationRequestForward,
			"gtp_v0_c_pdu_notification_reject_request_forward":       ret.DtFwGtpNetworkElementStats.Stats.GtpV0CPduNotificationRejectRequestForward,
			"gtp_v0_c_message_skipped_apn_filtering_no_imsi":         ret.DtFwGtpNetworkElementStats.Stats.GtpV0CMessageSkippedApnFilteringNoImsi,
			"gtp_v1_c_message_skipped_apn_filtering_no_imsi":         ret.DtFwGtpNetworkElementStats.Stats.GtpV1CMessageSkippedApnFilteringNoImsi,
			"gtp_v2_c_message_skipped_apn_filtering_no_imsi":         ret.DtFwGtpNetworkElementStats.Stats.GtpV2CMessageSkippedApnFilteringNoImsi,
			"gtp_v0_c_message_skipped_msisdn_filtering_no_msisdn":    ret.DtFwGtpNetworkElementStats.Stats.GtpV0CMessageSkippedMsisdnFilteringNoMsisdn,
			"gtp_v1_c_message_skipped_msisdn_filtering_no_msisdn":    ret.DtFwGtpNetworkElementStats.Stats.GtpV1CMessageSkippedMsisdnFilteringNoMsisdn,
			"gtp_v2_c_message_skipped_msisdn_filtering_no_msisdn":    ret.DtFwGtpNetworkElementStats.Stats.GtpV2CMessageSkippedMsisdnFilteringNoMsisdn,
			"gtp_v0_c_packet_dummy_msisdn":                           ret.DtFwGtpNetworkElementStats.Stats.GtpV0CPacketDummyMsisdn,
			"gtp_v1_c_packet_dummy_msisdn":                           ret.DtFwGtpNetworkElementStats.Stats.GtpV1CPacketDummyMsisdn,
			"gtp_v2_c_packet_dummy_msisdn":                           ret.DtFwGtpNetworkElementStats.Stats.GtpV2CPacketDummyMsisdn,
			"drop_vld_gtp_v2_c_message_with_teid_zero_expected":      ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV2CMessageWithTeidZeroExpected,
			"drop_vld_gtp_v1_c_message_with_teid_zero_expected":      ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV1CMessageWithTeidZeroExpected,
			"drop_vld_gtp_v0_c_message_with_teid_zero_expected":      ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV0CMessageWithTeidZeroExpected,
			"drop_vld_reserved_field_set":                            ret.DtFwGtpNetworkElementStats.Stats.DropVldReservedFieldSet,
			"drop_vld_tunnel_id_flag":                                ret.DtFwGtpNetworkElementStats.Stats.DropVldTunnelIdFlag,
			"drop_vld_invalid_flow_label_v0":                         ret.DtFwGtpNetworkElementStats.Stats.DropVldInvalidFlowLabelV0,
			"drop_vld_invalid_teid":                                  ret.DtFwGtpNetworkElementStats.Stats.DropVldInvalidTeid,
			"drop_vld_unsupported_message_type":                      ret.DtFwGtpNetworkElementStats.Stats.DropVldUnsupportedMessageType,
			"drop_vld_out_of_state":                                  ret.DtFwGtpNetworkElementStats.Stats.DropVldOutOfState,
			"drop_vld_mandatory_information_element":                 ret.DtFwGtpNetworkElementStats.Stats.DropVldMandatoryInformationElement,
			"drop_vld_out_of_order_ie":                               ret.DtFwGtpNetworkElementStats.Stats.DropVldOutOfOrderIe,
			"drop_vld_out_of_state_ie":                               ret.DtFwGtpNetworkElementStats.Stats.DropVldOutOfStateIe,
			"drop_vld_reserved_information_element":                  ret.DtFwGtpNetworkElementStats.Stats.DropVldReservedInformationElement,
			"drop_vld_version_not_supported":                         ret.DtFwGtpNetworkElementStats.Stats.DropVldVersionNotSupported,
			"drop_vld_message_length":                                ret.DtFwGtpNetworkElementStats.Stats.DropVldMessageLength,
			"drop_vld_cross_layer_correlation":                       ret.DtFwGtpNetworkElementStats.Stats.DropVldCrossLayerCorrelation,
			"drop_vld_country_code_mismatch":                         ret.DtFwGtpNetworkElementStats.Stats.DropVldCountryCodeMismatch,
			"drop_vld_gtp_u_spoofed_source_address":                  ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpUSpoofedSourceAddress,
			"drop_vld_gtp_bearer_count_exceed":                       ret.DtFwGtpNetworkElementStats.Stats.Drop_vldGtpBearerCountExceed,
			"drop_vld_gtp_v2_wrong_lbi_create_bearer_req":            ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV2WrongLbiCreateBearerReq,
			"gtp_c_handover_in_progress_with_conn":                   ret.DtFwGtpNetworkElementStats.Stats.GtpCHandoverInProgressWithConn,
			"drop_vld_v0_reserved_message_drop":                      ret.DtFwGtpNetworkElementStats.Stats.DropVldV0ReservedMessageDrop,
			"drop_vld_v1_reserved_message_drop":                      ret.DtFwGtpNetworkElementStats.Stats.DropVldV1ReservedMessageDrop,
			"drop_vld_v2_reserved_message_drop":                      ret.DtFwGtpNetworkElementStats.Stats.DropVldV2ReservedMessageDrop,
			"drop_vld_invalid_pkt_len_piggyback":                     ret.DtFwGtpNetworkElementStats.Stats.DropVldInvalidPktLenPiggyback,
			"drop_vld_sanity_failed_piggyback":                       ret.DtFwGtpNetworkElementStats.Stats.DropVldSanityFailedPiggyback,
			"drop_vld_sequence_num_correlation":                      ret.DtFwGtpNetworkElementStats.Stats.DropVldSequenceNumCorrelation,
			"drop_vld_gtpv0_seqnum_buffer_full":                      ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpv0SeqnumBufferFull,
			"drop_vld_gtpv1_seqnum_buffer_full":                      ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpv1SeqnumBufferFull,
			"drop_vld_gtpv2_seqnum_buffer_full":                      ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpv2SeqnumBufferFull,
			"drop_vld_gtp_invalid_imsi_len_drop":                     ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpInvalidImsiLenDrop,
			"drop_vld_gtp_invalid_apn_len_drop":                      ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpInvalidApnLenDrop,
			"drop_vld_protocol_flag_unset":                           ret.DtFwGtpNetworkElementStats.Stats.DropVldProtocolFlagUnset,
			"drop_vld_gtpv0_subscriber_attr_miss":                    ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpv0SubscriberAttrMiss,
			"drop_vld_gtpv1_subscriber_attr_miss":                    ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpv1SubscriberAttrMiss,
			"drop_vld_gtpv2_subscriber_attr_miss":                    ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpv2SubscriberAttrMiss,
			"drop_vld_gtp_v0_c_message_dropped_apn_filtering_no_apn": ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV0CMessageDroppedApnFilteringNoApn,
			"drop_vld_gtp_v1_c_message_dropped_apn_filtering_no_apn": ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV1CMessageDroppedApnFilteringNoApn,
			"drop_vld_gtp_v2_c_message_dropped_apn_filtering_no_apn": ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV2CMessageDroppedApnFilteringNoApn,
			"drop_vld_gtp_v0_c_ie_len_exceed_msg_len":                ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV0CIeLenExceedMsgLen,
			"drop_vld_gtp_v1_c_ie_len_exceed_msg_len":                ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV1CIeLenExceedMsgLen,
			"drop_vld_gtp_v2_c_ie_len_exceed_msg_len":                ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV2CIeLenExceedMsgLen,
			"drop_vld_gtp_v0_c_message_length_mismatch":              ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV0CMessageLengthMismatch,
			"drop_vld_gtp_v1_c_message_length_mismatch":              ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV1CMessageLengthMismatch,
			"drop_vld_gtp_v2_c_message_length_mismatch":              ret.DtFwGtpNetworkElementStats.Stats.DropVldGtpV2CMessageLengthMismatch,
			"drop_flt_message_filtering":                             ret.DtFwGtpNetworkElementStats.Stats.DropFltMessageFiltering,
			"drop_flt_apn_filtering":                                 ret.DtFwGtpNetworkElementStats.Stats.DropFltApnFiltering,
			"drop_flt_msisdn_filtering":                              ret.DtFwGtpNetworkElementStats.Stats.DropFltMsisdnFiltering,
			"drop_flt_rat_type_filtering":                            ret.DtFwGtpNetworkElementStats.Stats.DropFltRatTypeFiltering,
			"drop_flt_gtp_in_gtp":                                    ret.DtFwGtpNetworkElementStats.Stats.DropFltGtpInGtp,
			"drop_rl_gtp_v0_c_agg":                                   ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpV0CAgg,
			"drop_rl_gtp_v1_c_agg":                                   ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpV1CAgg,
			"drop_rl_gtp_v2_c_agg":                                   ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpV2CAgg,
			"drop_rl_gtp_v1_c_create_pdp_request":                    ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpV1CCreatePdpRequest,
			"drop_rl_gtp_v2_c_create_session_request":                ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpV2CCreateSessionRequest,
			"drop_rl_gtp_v1_c_update_pdp_request":                    ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpV1CUpdatePdpRequest,
			"drop_rl_gtp_v2_c_modify_bearer_request":                 ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpV2CModifyBearerRequest,
			"drop_rl_gtp_u_tunnel_create":                            ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpUTunnelCreate,
			"drop_rl_gtp_u_uplink_byte":                              ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpUUplinkByte,
			"drop_rl_gtp_u_uplink_packet":                            ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpUUplinkPacket,
			"drop_rl_gtp_u_downlink_byte":                            ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpUDownlinkByte,
			"drop_rl_gtp_u_downlink_packet":                          ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpUDownlinkPacket,
			"drop_rl_gtp_u_total_byte":                               ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpUTotalByte,
			"drop_rl_gtp_u_total_packet":                             ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpUTotalPacket,
			"drop_rl_gtp_u_max_concurrent_tunnels":                   ret.DtFwGtpNetworkElementStats.Stats.DropRlGtpUMaxConcurrentTunnels,
		},
	}
}

func getObjectFwGtpNetworkElementStatsStats(d []interface{}) edpt.FwGtpNetworkElementStatsStats {

	count1 := len(d)
	var ret edpt.FwGtpNetworkElementStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyName = in["key_name"].(string)
		ret.KeyType = in["key_type"].(string)
		ret.UplinkBytes = in["uplink_bytes"].(int)
		ret.DownlinkBytes = in["downlink_bytes"].(int)
		ret.UplinkPkts = in["uplink_pkts"].(int)
		ret.DownlinkPkts = in["downlink_pkts"].(int)
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
		ret.GtpNodeRestartGtpC = in["gtp_node_restart_gtp_c"].(int)
		ret.GtpV0CReservedMessageAllow = in["gtp_v0_c_reserved_message_allow"].(int)
		ret.GtpV1CReservedMessageAllow = in["gtp_v1_c_reserved_message_allow"].(int)
		ret.GtpV2CReservedMessageAllow = in["gtp_v2_c_reserved_message_allow"].(int)
		ret.GtpV1CPduNotificationRequestForward = in["gtp_v1_c_pdu_notification_request_forward"].(int)
		ret.GtpV1CPduNotificationRejectRequestForward = in["gtp_v1_c_pdu_notification_reject_request_forward"].(int)
		ret.GtpV0CPduNotificationRequestForward = in["gtp_v0_c_pdu_notification_request_forward"].(int)
		ret.GtpV0CPduNotificationRejectRequestForward = in["gtp_v0_c_pdu_notification_reject_request_forward"].(int)
		ret.GtpV0CMessageSkippedApnFilteringNoImsi = in["gtp_v0_c_message_skipped_apn_filtering_no_imsi"].(int)
		ret.GtpV1CMessageSkippedApnFilteringNoImsi = in["gtp_v1_c_message_skipped_apn_filtering_no_imsi"].(int)
		ret.GtpV2CMessageSkippedApnFilteringNoImsi = in["gtp_v2_c_message_skipped_apn_filtering_no_imsi"].(int)
		ret.GtpV0CMessageSkippedMsisdnFilteringNoMsisdn = in["gtp_v0_c_message_skipped_msisdn_filtering_no_msisdn"].(int)
		ret.GtpV1CMessageSkippedMsisdnFilteringNoMsisdn = in["gtp_v1_c_message_skipped_msisdn_filtering_no_msisdn"].(int)
		ret.GtpV2CMessageSkippedMsisdnFilteringNoMsisdn = in["gtp_v2_c_message_skipped_msisdn_filtering_no_msisdn"].(int)
		ret.GtpV0CPacketDummyMsisdn = in["gtp_v0_c_packet_dummy_msisdn"].(int)
		ret.GtpV1CPacketDummyMsisdn = in["gtp_v1_c_packet_dummy_msisdn"].(int)
		ret.GtpV2CPacketDummyMsisdn = in["gtp_v2_c_packet_dummy_msisdn"].(int)
		ret.DropVldGtpV2CMessageWithTeidZeroExpected = in["drop_vld_gtp_v2_c_message_with_teid_zero_expected"].(int)
		ret.DropVldGtpV1CMessageWithTeidZeroExpected = in["drop_vld_gtp_v1_c_message_with_teid_zero_expected"].(int)
		ret.DropVldGtpV0CMessageWithTeidZeroExpected = in["drop_vld_gtp_v0_c_message_with_teid_zero_expected"].(int)
		ret.DropVldReservedFieldSet = in["drop_vld_reserved_field_set"].(int)
		ret.DropVldTunnelIdFlag = in["drop_vld_tunnel_id_flag"].(int)
		ret.DropVldInvalidFlowLabelV0 = in["drop_vld_invalid_flow_label_v0"].(int)
		ret.DropVldInvalidTeid = in["drop_vld_invalid_teid"].(int)
		ret.DropVldUnsupportedMessageType = in["drop_vld_unsupported_message_type"].(int)
		ret.DropVldOutOfState = in["drop_vld_out_of_state"].(int)
		ret.DropVldMandatoryInformationElement = in["drop_vld_mandatory_information_element"].(int)
		ret.DropVldOutOfOrderIe = in["drop_vld_out_of_order_ie"].(int)
		ret.DropVldOutOfStateIe = in["drop_vld_out_of_state_ie"].(int)
		ret.DropVldReservedInformationElement = in["drop_vld_reserved_information_element"].(int)
		ret.DropVldVersionNotSupported = in["drop_vld_version_not_supported"].(int)
		ret.DropVldMessageLength = in["drop_vld_message_length"].(int)
		ret.DropVldCrossLayerCorrelation = in["drop_vld_cross_layer_correlation"].(int)
		ret.DropVldCountryCodeMismatch = in["drop_vld_country_code_mismatch"].(int)
		ret.DropVldGtpUSpoofedSourceAddress = in["drop_vld_gtp_u_spoofed_source_address"].(int)
		ret.Drop_vldGtpBearerCountExceed = in["drop_vld_gtp_bearer_count_exceed"].(int)
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
		ret.DropVldGtpV0CMessageDroppedApnFilteringNoApn = in["drop_vld_gtp_v0_c_message_dropped_apn_filtering_no_apn"].(int)
		ret.DropVldGtpV1CMessageDroppedApnFilteringNoApn = in["drop_vld_gtp_v1_c_message_dropped_apn_filtering_no_apn"].(int)
		ret.DropVldGtpV2CMessageDroppedApnFilteringNoApn = in["drop_vld_gtp_v2_c_message_dropped_apn_filtering_no_apn"].(int)
		ret.DropVldGtpV0CIeLenExceedMsgLen = in["drop_vld_gtp_v0_c_ie_len_exceed_msg_len"].(int)
		ret.DropVldGtpV1CIeLenExceedMsgLen = in["drop_vld_gtp_v1_c_ie_len_exceed_msg_len"].(int)
		ret.DropVldGtpV2CIeLenExceedMsgLen = in["drop_vld_gtp_v2_c_ie_len_exceed_msg_len"].(int)
		ret.DropVldGtpV0CMessageLengthMismatch = in["drop_vld_gtp_v0_c_message_length_mismatch"].(int)
		ret.DropVldGtpV1CMessageLengthMismatch = in["drop_vld_gtp_v1_c_message_length_mismatch"].(int)
		ret.DropVldGtpV2CMessageLengthMismatch = in["drop_vld_gtp_v2_c_message_length_mismatch"].(int)
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

func dataToEndpointFwGtpNetworkElementStats(d *schema.ResourceData) edpt.FwGtpNetworkElementStats {
	var ret edpt.FwGtpNetworkElementStats

	ret.Stats = getObjectFwGtpNetworkElementStatsStats(d.Get("stats").([]interface{}))
	return ret
}
