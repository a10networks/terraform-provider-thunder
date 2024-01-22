package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwGtpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_gtp_stats`: Statistics for the object gtp\n\n__PLACEHOLDER__",
		ReadContext: resourceFwGtpStatsRead,

		Schema: map[string]*schema.Schema{
			"apn_prefix": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_name": {
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
										Type: schema.TypeInt, Optional: true, Description: "GTPV0-C Update PDP Context Response Unsuccessful",
									},
									"gtp_v1_c_update_pdp_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Description: "GTPV1-C Update PDP Context response Unsuccessful",
									},
									"gtp_v2_c_mod_bearer_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Description: "GTPV2-C Modify Bearer Response Unsuccessful",
									},
									"gtp_v0_c_create_pdp_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Description: "GTPV0-C Create PDP Context Response Unsuccessful",
									},
									"gtp_v1_c_create_pdp_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Description: "GTPV1-C Create PDP Context Response Unsuccessful",
									},
									"gtp_v2_c_create_sess_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Description: "GTPV2-C Create Session Response Unsuccessful",
									},
									"gtp_v2_c_piggyback_message": {
										Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Piggyback Messages seen",
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
										Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Packet With Dummy MSISDN Forwarded",
									},
									"gtp_v2_c_packet_dummy_msisdn": {
										Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Packet With Dummy MSISDN Forwarded",
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
									"drop_vld_gtp_v0_c_message_dropped_apn_filtering_no_apn": {
										Type: schema.TypeInt, Optional: true, Description: "GTPv0-C APN/IMSI Filtering dropped (No APN)",
									},
									"drop_vld_gtp_v1_c_message_dropped_apn_filtering_no_apn": {
										Type: schema.TypeInt, Optional: true, Description: "GTPv1-C APN/IMSI Filtering dropped (No APN)",
									},
									"drop_vld_gtp_v2_c_message_dropped_apn_filtering_no_apn": {
										Type: schema.TypeInt, Optional: true, Description: "GTPv2-C APN/IMSI Filtering dropped (No APN)",
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
				},
			},
			"network_element": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"out_of_session_memory": {
							Type: schema.TypeInt, Optional: true, Description: "Out of Tunnel Memory",
						},
						"no_fwd_route": {
							Type: schema.TypeInt, Optional: true, Description: "No Forward Route",
						},
						"no_rev_route": {
							Type: schema.TypeInt, Optional: true, Description: "No Reverse Route",
						},
						"gtp_smp_path_check_failed": {
							Type: schema.TypeInt, Optional: true, Description: "GTP SMP PATH check Failed",
						},
						"gtp_smp_check_failed": {
							Type: schema.TypeInt, Optional: true, Description: "GTP SMP check Failed",
						},
						"gtp_smp_session_count_check_failed": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U session count is not in range of 0-11 in GTP-C SMP",
						},
						"gtp_c_ref_count_smp_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-C session count on C-smp exceeded 2",
						},
						"gtp_u_smp_in_rml_with_sess": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U smp is marked RML with U-session",
						},
						"gtp_tunnel_rate_limit_entry_create_failure": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Tunnel Level Rate Limit Entry Create Failure",
						},
						"gtp_rate_limit_smp_create_failure": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Rate Limit SMP Create Failure",
						},
						"gtp_rate_limit_t3_ctr_create_failure": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Rate Limit Dynamic Counters Create Failure",
						},
						"gtp_rate_limit_entry_create_failure": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Rate Limit Entry Create Failure",
						},
						"gtp_node_restart_echo": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Node Restoration due to Recovery IE in Echo",
						},
						"gtp_c_echo_path_failure": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-C Path Failure due to Echo",
						},
						"drop_vld_gtp_echo_out_of_state_": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Echo Out of State Drop",
						},
						"drop_vld_gtp_echo_ie_len_exceed_msg_len": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Echo IE Length Exceeds Message Length",
						},
						"gtp_del_bearer_request_retransmit": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-C Retransmitted Delete Bearer Request",
						},
						"gtp_add_bearer_response_retransmit": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-C Retransmitted Add Bearer Response",
						},
						"gtp_u_out_of_state_drop": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Out of state Drop",
						},
						"gtp_c_handover_request_out_of_state_drop": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-C Handover Request Out of state Drop",
						},
						"gtp_v1_c_nsapi_not_found_in_delete_req": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C NSAPI Not Found in GTP Request",
						},
						"gtp_v2_c_bearer_not_found_in_delete_req": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Bearer Not Found in GTP Request",
						},
						"gtp_v2_c_bearer_not_found_in_delete_resp": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Bearer Not Found in GTP Response",
						},
						"gtp_rr_message_drop": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Message Dropped in RR Mode",
						},
						"drop_gtp_frag_or_jumbo_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Fragmented or JUMBO packet Drop",
						},
						"gtp_c_handover_in_progress_with_conn": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-C mesg matching conn with HO In Progress",
						},
						"gtp_smp_dec_sess_count_check_failed": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U session count is 0 in GTP-C SMP",
						},
						"gtp_v0_c_uplink_ingress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Uplink Ingress Packets",
						},
						"gtp_v0_c_uplink_egress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Uplink Egress Packets",
						},
						"gtp_v0_c_downlink_ingress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Downlink Ingress Packets",
						},
						"gtp_v0_c_downlink_egress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Downlink Egress Packets",
						},
						"gtp_v0_c_uplink_ingress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Uplink Ingress Bytes",
						},
						"gtp_v0_c_uplink_egress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Uplink Egress Bytes",
						},
						"gtp_v0_c_downlink_ingress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Downlink Ingress Bytes",
						},
						"gtp_v0_c_downlink_egress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv0-C Downlink Egress Bytes",
						},
						"gtp_v1_c_uplink_ingress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Uplink Ingress Packets",
						},
						"gtp_v1_c_uplink_egress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Uplink Egress Packets",
						},
						"gtp_v1_c_downlink_ingress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Downlink Ingress Packets",
						},
						"gtp_v1_c_downlink_egress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Downlink Egress Packets",
						},
						"gtp_v1_c_uplink_ingress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Uplink Ingress Bytes",
						},
						"gtp_v1_c_uplink_egress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Uplink Egress Bytes",
						},
						"gtp_v1_c_downlink_ingress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Downlink Ingress Bytes",
						},
						"gtp_v1_c_downlink_egress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv1-C Downlink Egress Bytes",
						},
						"gtp_v2_c_uplink_ingress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Uplink Ingress Packets",
						},
						"gtp_v2_c_uplink_egress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Uplink Egress Packets",
						},
						"gtp_v2_c_downlink_ingress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Downlink Ingress Packets",
						},
						"gtp_v2_c_downlink_egress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Downlink Egress Packets",
						},
						"gtp_v2_c_uplink_ingress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Uplink Ingress Bytes",
						},
						"gtp_v2_c_uplink_egress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Uplink Egress Bytes",
						},
						"gtp_v2_c_downlink_ingress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Downlink Ingress Bytes",
						},
						"gtp_v2_c_downlink_egress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTPv2-C Downlink Egress Bytes",
						},
						"gtp_u_uplink_ingress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Uplink Ingress Packets",
						},
						"gtp_u_uplink_egress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Uplink Egress Packets",
						},
						"gtp_u_downlink_ingress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Downlink Ingress Packets",
						},
						"gtp_u_downlink_egress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Downlink Egress Packets",
						},
						"gtp_u_uplink_ingress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Uplink Ingress Bytes",
						},
						"gtp_u_uplink_egress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Uplink Egress Bytes",
						},
						"gtp_u_downlink_ingress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Downlink Ingress Bytes",
						},
						"gtp_u_downlink_egress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Downlink Egress Bytes",
						},
						"gtp_u_message_length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Message Length Mismatch Across Layers",
						},
						"gtp_path_message_length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-Path Message Length Mismatch Across Layers",
						},
						"drop_gtp_missing_cond_ie_bearer_ctx": {
							Type: schema.TypeInt, Optional: true, Description: "Missing conditional IE in bearer context Drop",
						},
						"drop_gtp_bearer_not_found_in_resp": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Bearer not found in response",
						},
						"gtp_stateless_forward": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Stateless Forward",
						},
						"gtp_monitor_forward": {
							Type: schema.TypeInt, Optional: true, Description: "GTP messages forwarded via monitor mode",
						},
					},
				},
			},
		},
	}
}

func resourceFwGtpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwGtpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwGtpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwGtpStatsApnPrefix := setObjectFwGtpStatsApnPrefix(res)
		d.Set("apn_prefix", FwGtpStatsApnPrefix)
		FwGtpStatsNetworkElement := setObjectFwGtpStatsNetworkElement(res)
		d.Set("network_element", FwGtpStatsNetworkElement)
		FwGtpStatsStats := setObjectFwGtpStatsStats(res)
		d.Set("stats", FwGtpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwGtpStatsApnPrefix(ret edpt.DataFwGtpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectFwGtpStatsApnPrefixStats(ret.DtFwGtpStats.ApnPrefix.Stats),
		},
	}
}

func setObjectFwGtpStatsApnPrefixStats(d edpt.FwGtpStatsApnPrefixStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["key_name"] = d.KeyName

	in["uplink_bytes"] = d.UplinkBytes

	in["downlink_bytes"] = d.DownlinkBytes

	in["uplink_pkts"] = d.UplinkPkts

	in["downlink_pkts"] = d.DownlinkPkts

	in["gtp_v0_c_tunnel_created"] = d.GtpV0CTunnelCreated

	in["gtp_v0_c_tunnel_half_open"] = d.GtpV0CTunnelHalfOpen

	in["gtp_v0_c_tunnel_half_closed"] = d.GtpV0CTunnelHalfClosed

	in["gtp_v0_c_tunnel_closed"] = d.GtpV0CTunnelClosed

	in["gtp_v0_c_tunnel_deleted"] = d.GtpV0CTunnelDeleted

	in["gtp_v0_c_half_open_tunnel_closed"] = d.GtpV0CHalfOpenTunnelClosed

	in["gtp_v1_c_tunnel_created"] = d.GtpV1CTunnelCreated

	in["gtp_v1_c_tunnel_half_open"] = d.GtpV1CTunnelHalfOpen

	in["gtp_v1_c_tunnel_half_closed"] = d.GtpV1CTunnelHalfClosed

	in["gtp_v1_c_tunnel_closed"] = d.GtpV1CTunnelClosed

	in["gtp_v1_c_tunnel_deleted"] = d.GtpV1CTunnelDeleted

	in["gtp_v1_c_half_open_tunnel_closed"] = d.GtpV1CHalfOpenTunnelClosed

	in["gtp_v2_c_tunnel_created"] = d.GtpV2CTunnelCreated

	in["gtp_v2_c_tunnel_half_open"] = d.GtpV2CTunnelHalfOpen

	in["gtp_v2_c_tunnel_half_closed"] = d.GtpV2CTunnelHalfClosed

	in["gtp_v2_c_tunnel_closed"] = d.GtpV2CTunnelClosed

	in["gtp_v2_c_tunnel_deleted"] = d.GtpV2CTunnelDeleted

	in["gtp_v2_c_half_open_tunnel_closed"] = d.GtpV2CHalfOpenTunnelClosed

	in["gtp_u_tunnel_created"] = d.GtpUTunnelCreated

	in["gtp_u_tunnel_deleted"] = d.GtpUTunnelDeleted

	in["gtp_v0_c_update_pdp_resp_unsuccess"] = d.GtpV0CUpdatePdpRespUnsuccess

	in["gtp_v1_c_update_pdp_resp_unsuccess"] = d.GtpV1CUpdatePdpRespUnsuccess

	in["gtp_v2_c_mod_bearer_resp_unsuccess"] = d.GtpV2CMod_bearerRespUnsuccess

	in["gtp_v0_c_create_pdp_resp_unsuccess"] = d.GtpV0CCreatePdpRespUnsuccess

	in["gtp_v1_c_create_pdp_resp_unsuccess"] = d.GtpV1CCreatePdpRespUnsuccess

	in["gtp_v2_c_create_sess_resp_unsuccess"] = d.GtpV2CCreateSessRespUnsuccess

	in["gtp_v2_c_piggyback_message"] = d.GtpV2CPiggybackMessage

	in["gtp_path_management_message"] = d.GtpPathManagementMessage

	in["gtp_v0_c_tunnel_deleted_restart"] = d.GtpV0CTunnelDeletedRestart

	in["gtp_v1_c_tunnel_deleted_restart"] = d.GtpV1CTunnelDeletedRestart

	in["gtp_v2_c_tunnel_deleted_restart"] = d.GtpV2CTunnelDeletedRestart

	in["gtp_v0_c_reserved_message_allow"] = d.GtpV0CReservedMessageAllow

	in["gtp_v1_c_reserved_message_allow"] = d.GtpV1CReservedMessageAllow

	in["gtp_v2_c_reserved_message_allow"] = d.GtpV2CReservedMessageAllow

	in["gtp_v1_c_pdu_notification_request_forward"] = d.GtpV1CPduNotificationRequestForward

	in["gtp_v1_c_pdu_notification_reject_request_forward"] = d.GtpV1CPduNotificationRejectRequestForward

	in["gtp_v0_c_pdu_notification_request_forward"] = d.GtpV0CPduNotificationRequestForward

	in["gtp_v0_c_pdu_notification_reject_request_forward"] = d.GtpV0CPduNotificationRejectRequestForward

	in["gtp_v0_c_message_skipped_apn_filtering_no_imsi"] = d.GtpV0CMessageSkippedApnFilteringNoImsi

	in["gtp_v1_c_message_skipped_apn_filtering_no_imsi"] = d.GtpV1CMessageSkippedApnFilteringNoImsi

	in["gtp_v2_c_message_skipped_apn_filtering_no_imsi"] = d.GtpV2CMessageSkippedApnFilteringNoImsi

	in["gtp_v0_c_message_skipped_msisdn_filtering_no_msisdn"] = d.GtpV0CMessageSkippedMsisdnFilteringNoMsisdn

	in["gtp_v1_c_message_skipped_msisdn_filtering_no_msisdn"] = d.GtpV1CMessageSkippedMsisdnFilteringNoMsisdn

	in["gtp_v2_c_message_skipped_msisdn_filtering_no_msisdn"] = d.GtpV2CMessageSkippedMsisdnFilteringNoMsisdn

	in["gtp_v0_c_packet_dummy_msisdn"] = d.GtpV0CPacketDummyMsisdn

	in["gtp_v1_c_packet_dummy_msisdn"] = d.GtpV1CPacketDummyMsisdn

	in["gtp_v2_c_packet_dummy_msisdn"] = d.GtpV2CPacketDummyMsisdn

	in["drop_vld_gtp_v2_c_message_with_teid_zero_expected"] = d.DropVldGtpV2CMessageWithTeidZeroExpected

	in["drop_vld_gtp_v1_c_message_with_teid_zero_expected"] = d.DropVldGtpV1CMessageWithTeidZeroExpected

	in["drop_vld_gtp_v0_c_message_with_teid_zero_expected"] = d.DropVldGtpV0CMessageWithTeidZeroExpected

	in["drop_vld_reserved_field_set"] = d.DropVldReservedFieldSet

	in["drop_vld_tunnel_id_flag"] = d.DropVldTunnelIdFlag

	in["drop_vld_invalid_flow_label_v0"] = d.DropVldInvalidFlowLabelV0

	in["drop_vld_invalid_teid"] = d.DropVldInvalidTeid

	in["drop_vld_unsupported_message_type"] = d.DropVldUnsupportedMessageType

	in["drop_vld_out_of_state"] = d.DropVldOutOfState

	in["drop_vld_mandatory_information_element"] = d.DropVldMandatoryInformationElement

	in["drop_vld_out_of_order_ie"] = d.DropVldOutOfOrderIe

	in["drop_vld_out_of_state_ie"] = d.DropVldOutOfStateIe

	in["drop_vld_reserved_information_element"] = d.DropVldReservedInformationElement

	in["drop_vld_version_not_supported"] = d.DropVldVersionNotSupported

	in["drop_vld_message_length"] = d.DropVldMessageLength

	in["drop_vld_cross_layer_correlation"] = d.DropVldCrossLayerCorrelation

	in["drop_vld_country_code_mismatch"] = d.DropVldCountryCodeMismatch

	in["drop_vld_gtp_u_spoofed_source_address"] = d.DropVldGtpUSpoofedSourceAddress

	in["drop_vld_gtp_bearer_count_exceed"] = d.DropVldGtpBearerCountExceed

	in["drop_vld_gtp_v2_wrong_lbi_create_bearer_req"] = d.DropVldGtpV2WrongLbiCreateBearerReq

	in["gtp_c_handover_in_progress_with_conn"] = d.GtpCHandoverInProgressWithConn

	in["drop_vld_v0_reserved_message_drop"] = d.DropVldV0ReservedMessageDrop

	in["drop_vld_v1_reserved_message_drop"] = d.DropVldV1ReservedMessageDrop

	in["drop_vld_v2_reserved_message_drop"] = d.DropVldV2ReservedMessageDrop

	in["drop_vld_invalid_pkt_len_piggyback"] = d.DropVldInvalidPktLenPiggyback

	in["drop_vld_sanity_failed_piggyback"] = d.DropVldSanityFailedPiggyback

	in["drop_vld_sequence_num_correlation"] = d.DropVldSequenceNumCorrelation

	in["drop_vld_gtpv0_seqnum_buffer_full"] = d.DropVldGtpv0SeqnumBufferFull

	in["drop_vld_gtpv1_seqnum_buffer_full"] = d.DropVldGtpv1SeqnumBufferFull

	in["drop_vld_gtpv2_seqnum_buffer_full"] = d.DropVldGtpv2SeqnumBufferFull

	in["drop_vld_gtp_invalid_imsi_len_drop"] = d.DropVldGtpInvalidImsiLenDrop

	in["drop_vld_gtp_invalid_apn_len_drop"] = d.DropVldGtpInvalidApnLenDrop

	in["drop_vld_protocol_flag_unset"] = d.DropVldProtocolFlagUnset

	in["drop_vld_gtpv0_subscriber_attr_miss"] = d.DropVldGtpv0SubscriberAttrMiss

	in["drop_vld_gtpv1_subscriber_attr_miss"] = d.DropVldGtpv1SubscriberAttrMiss

	in["drop_vld_gtpv2_subscriber_attr_miss"] = d.DropVldGtpv2SubscriberAttrMiss

	in["drop_vld_gtp_v0_c_ie_len_exceed_msg_len"] = d.DropVldGtpV0CIeLenExceedMsgLen

	in["drop_vld_gtp_v1_c_ie_len_exceed_msg_len"] = d.DropVldGtpV1CIeLenExceedMsgLen

	in["drop_vld_gtp_v2_c_ie_len_exceed_msg_len"] = d.DropVldGtpV2CIeLenExceedMsgLen

	in["drop_vld_gtp_v0_c_message_length_mismatch"] = d.DropVldGtpV0CMessageLengthMismatch

	in["drop_vld_gtp_v1_c_message_length_mismatch"] = d.DropVldGtpV1CMessageLengthMismatch

	in["drop_vld_gtp_v2_c_message_length_mismatch"] = d.DropVldGtpV2CMessageLengthMismatch

	in["drop_vld_gtp_v0_c_message_dropped_apn_filtering_no_apn"] = d.DropVldGtpV0CMessageDroppedApnFilteringNoApn

	in["drop_vld_gtp_v1_c_message_dropped_apn_filtering_no_apn"] = d.DropVldGtpV1CMessageDroppedApnFilteringNoApn

	in["drop_vld_gtp_v2_c_message_dropped_apn_filtering_no_apn"] = d.DropVldGtpV2CMessageDroppedApnFilteringNoApn

	in["drop_flt_message_filtering"] = d.DropFltMessageFiltering

	in["drop_flt_apn_filtering"] = d.DropFltApnFiltering

	in["drop_flt_msisdn_filtering"] = d.DropFltMsisdnFiltering

	in["drop_flt_rat_type_filtering"] = d.DropFltRatTypeFiltering

	in["drop_flt_gtp_in_gtp"] = d.DropFltGtpInGtp

	in["drop_rl_gtp_v0_c_agg"] = d.DropRlGtpV0CAgg

	in["drop_rl_gtp_v1_c_agg"] = d.DropRlGtpV1CAgg

	in["drop_rl_gtp_v2_c_agg"] = d.DropRlGtpV2CAgg

	in["drop_rl_gtp_v1_c_create_pdp_request"] = d.DropRlGtpV1CCreatePdpRequest

	in["drop_rl_gtp_v2_c_create_session_request"] = d.DropRlGtpV2CCreateSessionRequest

	in["drop_rl_gtp_v1_c_update_pdp_request"] = d.DropRlGtpV1CUpdatePdpRequest

	in["drop_rl_gtp_v2_c_modify_bearer_request"] = d.DropRlGtpV2CModifyBearerRequest

	in["drop_rl_gtp_u_tunnel_create"] = d.DropRlGtpUTunnelCreate

	in["drop_rl_gtp_u_uplink_byte"] = d.DropRlGtpUUplinkByte

	in["drop_rl_gtp_u_uplink_packet"] = d.DropRlGtpUUplinkPacket

	in["drop_rl_gtp_u_downlink_byte"] = d.DropRlGtpUDownlinkByte

	in["drop_rl_gtp_u_downlink_packet"] = d.DropRlGtpUDownlinkPacket

	in["drop_rl_gtp_u_total_byte"] = d.DropRlGtpUTotalByte

	in["drop_rl_gtp_u_total_packet"] = d.DropRlGtpUTotalPacket

	in["drop_rl_gtp_u_max_concurrent_tunnels"] = d.DropRlGtpUMaxConcurrentTunnels
	result = append(result, in)
	return result
}

func setObjectFwGtpStatsNetworkElement(ret edpt.DataFwGtpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectFwGtpStatsNetworkElementStats(ret.DtFwGtpStats.NetworkElement.Stats),
		},
	}
}

func setObjectFwGtpStatsNetworkElementStats(d edpt.FwGtpStatsNetworkElementStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["key_name"] = d.KeyName

	in["key_type"] = d.KeyType

	in["uplink_bytes"] = d.UplinkBytes

	in["downlink_bytes"] = d.DownlinkBytes

	in["uplink_pkts"] = d.UplinkPkts

	in["downlink_pkts"] = d.DownlinkPkts

	in["gtp_v0_c_tunnel_created"] = d.GtpV0CTunnelCreated

	in["gtp_v0_c_tunnel_half_open"] = d.GtpV0CTunnelHalfOpen

	in["gtp_v0_c_tunnel_half_closed"] = d.GtpV0CTunnelHalfClosed

	in["gtp_v0_c_tunnel_closed"] = d.GtpV0CTunnelClosed

	in["gtp_v0_c_tunnel_deleted"] = d.GtpV0CTunnelDeleted

	in["gtp_v0_c_half_open_tunnel_closed"] = d.GtpV0CHalfOpenTunnelClosed

	in["gtp_v1_c_tunnel_created"] = d.GtpV1CTunnelCreated

	in["gtp_v1_c_tunnel_half_open"] = d.GtpV1CTunnelHalfOpen

	in["gtp_v1_c_tunnel_half_closed"] = d.GtpV1CTunnelHalfClosed

	in["gtp_v1_c_tunnel_closed"] = d.GtpV1CTunnelClosed

	in["gtp_v1_c_tunnel_deleted"] = d.GtpV1CTunnelDeleted

	in["gtp_v1_c_half_open_tunnel_closed"] = d.GtpV1CHalfOpenTunnelClosed

	in["gtp_v2_c_tunnel_created"] = d.GtpV2CTunnelCreated

	in["gtp_v2_c_tunnel_half_open"] = d.GtpV2CTunnelHalfOpen

	in["gtp_v2_c_tunnel_half_closed"] = d.GtpV2CTunnelHalfClosed

	in["gtp_v2_c_tunnel_closed"] = d.GtpV2CTunnelClosed

	in["gtp_v2_c_tunnel_deleted"] = d.GtpV2CTunnelDeleted

	in["gtp_v2_c_half_open_tunnel_closed"] = d.GtpV2CHalfOpenTunnelClosed

	in["gtp_u_tunnel_created"] = d.GtpUTunnelCreated

	in["gtp_u_tunnel_deleted"] = d.GtpUTunnelDeleted

	in["gtp_v0_c_update_pdp_resp_unsuccess"] = d.GtpV0CUpdatePdpRespUnsuccess

	in["gtp_v1_c_update_pdp_resp_unsuccess"] = d.GtpV1CUpdatePdpRespUnsuccess

	in["gtp_v2_c_mod_bearer_resp_unsuccess"] = d.GtpV2CMod_bearerRespUnsuccess

	in["gtp_v0_c_create_pdp_resp_unsuccess"] = d.GtpV0CCreatePdpRespUnsuccess

	in["gtp_v1_c_create_pdp_resp_unsuccess"] = d.GtpV1CCreatePdpRespUnsuccess

	in["gtp_v2_c_create_sess_resp_unsuccess"] = d.GtpV2CCreateSessRespUnsuccess

	in["gtp_v2_c_piggyback_message"] = d.GtpV2CPiggybackMessage

	in["gtp_path_management_message"] = d.GtpPathManagementMessage

	in["gtp_v0_c_tunnel_deleted_restart"] = d.GtpV0CTunnelDeletedRestart

	in["gtp_v1_c_tunnel_deleted_restart"] = d.GtpV1CTunnelDeletedRestart

	in["gtp_v2_c_tunnel_deleted_restart"] = d.GtpV2CTunnelDeletedRestart

	in["gtp_node_restart_gtp_c"] = d.GtpNodeRestartGtpC

	in["gtp_v0_c_reserved_message_allow"] = d.GtpV0CReservedMessageAllow

	in["gtp_v1_c_reserved_message_allow"] = d.GtpV1CReservedMessageAllow

	in["gtp_v2_c_reserved_message_allow"] = d.GtpV2CReservedMessageAllow

	in["gtp_v1_c_pdu_notification_request_forward"] = d.GtpV1CPduNotificationRequestForward

	in["gtp_v1_c_pdu_notification_reject_request_forward"] = d.GtpV1CPduNotificationRejectRequestForward

	in["gtp_v0_c_pdu_notification_request_forward"] = d.GtpV0CPduNotificationRequestForward

	in["gtp_v0_c_pdu_notification_reject_request_forward"] = d.GtpV0CPduNotificationRejectRequestForward

	in["gtp_v0_c_message_skipped_apn_filtering_no_imsi"] = d.GtpV0CMessageSkippedApnFilteringNoImsi

	in["gtp_v1_c_message_skipped_apn_filtering_no_imsi"] = d.GtpV1CMessageSkippedApnFilteringNoImsi

	in["gtp_v2_c_message_skipped_apn_filtering_no_imsi"] = d.GtpV2CMessageSkippedApnFilteringNoImsi

	in["gtp_v0_c_message_skipped_msisdn_filtering_no_msisdn"] = d.GtpV0CMessageSkippedMsisdnFilteringNoMsisdn

	in["gtp_v1_c_message_skipped_msisdn_filtering_no_msisdn"] = d.GtpV1CMessageSkippedMsisdnFilteringNoMsisdn

	in["gtp_v2_c_message_skipped_msisdn_filtering_no_msisdn"] = d.GtpV2CMessageSkippedMsisdnFilteringNoMsisdn

	in["gtp_v0_c_packet_dummy_msisdn"] = d.GtpV0CPacketDummyMsisdn

	in["gtp_v1_c_packet_dummy_msisdn"] = d.GtpV1CPacketDummyMsisdn

	in["gtp_v2_c_packet_dummy_msisdn"] = d.GtpV2CPacketDummyMsisdn

	in["drop_vld_gtp_v2_c_message_with_teid_zero_expected"] = d.DropVldGtpV2CMessageWithTeidZeroExpected

	in["drop_vld_gtp_v1_c_message_with_teid_zero_expected"] = d.DropVldGtpV1CMessageWithTeidZeroExpected

	in["drop_vld_gtp_v0_c_message_with_teid_zero_expected"] = d.DropVldGtpV0CMessageWithTeidZeroExpected

	in["drop_vld_reserved_field_set"] = d.DropVldReservedFieldSet

	in["drop_vld_tunnel_id_flag"] = d.DropVldTunnelIdFlag

	in["drop_vld_invalid_flow_label_v0"] = d.DropVldInvalidFlowLabelV0

	in["drop_vld_invalid_teid"] = d.DropVldInvalidTeid

	in["drop_vld_unsupported_message_type"] = d.DropVldUnsupportedMessageType

	in["drop_vld_out_of_state"] = d.DropVldOutOfState

	in["drop_vld_mandatory_information_element"] = d.DropVldMandatoryInformationElement

	in["drop_vld_out_of_order_ie"] = d.DropVldOutOfOrderIe

	in["drop_vld_out_of_state_ie"] = d.DropVldOutOfStateIe

	in["drop_vld_reserved_information_element"] = d.DropVldReservedInformationElement

	in["drop_vld_version_not_supported"] = d.DropVldVersionNotSupported

	in["drop_vld_message_length"] = d.DropVldMessageLength

	in["drop_vld_cross_layer_correlation"] = d.DropVldCrossLayerCorrelation

	in["drop_vld_country_code_mismatch"] = d.DropVldCountryCodeMismatch

	in["drop_vld_gtp_u_spoofed_source_address"] = d.DropVldGtpUSpoofedSourceAddress

	in["drop_vld_gtp_bearer_count_exceed"] = d.Drop_vldGtpBearerCountExceed

	in["drop_vld_gtp_v2_wrong_lbi_create_bearer_req"] = d.DropVldGtpV2WrongLbiCreateBearerReq

	in["gtp_c_handover_in_progress_with_conn"] = d.GtpCHandoverInProgressWithConn

	in["drop_vld_v0_reserved_message_drop"] = d.DropVldV0ReservedMessageDrop

	in["drop_vld_v1_reserved_message_drop"] = d.DropVldV1ReservedMessageDrop

	in["drop_vld_v2_reserved_message_drop"] = d.DropVldV2ReservedMessageDrop

	in["drop_vld_invalid_pkt_len_piggyback"] = d.DropVldInvalidPktLenPiggyback

	in["drop_vld_sanity_failed_piggyback"] = d.DropVldSanityFailedPiggyback

	in["drop_vld_sequence_num_correlation"] = d.DropVldSequenceNumCorrelation

	in["drop_vld_gtpv0_seqnum_buffer_full"] = d.DropVldGtpv0SeqnumBufferFull

	in["drop_vld_gtpv1_seqnum_buffer_full"] = d.DropVldGtpv1SeqnumBufferFull

	in["drop_vld_gtpv2_seqnum_buffer_full"] = d.DropVldGtpv2SeqnumBufferFull

	in["drop_vld_gtp_invalid_imsi_len_drop"] = d.DropVldGtpInvalidImsiLenDrop

	in["drop_vld_gtp_invalid_apn_len_drop"] = d.DropVldGtpInvalidApnLenDrop

	in["drop_vld_protocol_flag_unset"] = d.DropVldProtocolFlagUnset

	in["drop_vld_gtpv0_subscriber_attr_miss"] = d.DropVldGtpv0SubscriberAttrMiss

	in["drop_vld_gtpv1_subscriber_attr_miss"] = d.DropVldGtpv1SubscriberAttrMiss

	in["drop_vld_gtpv2_subscriber_attr_miss"] = d.DropVldGtpv2SubscriberAttrMiss

	in["drop_vld_gtp_v0_c_message_dropped_apn_filtering_no_apn"] = d.DropVldGtpV0CMessageDroppedApnFilteringNoApn

	in["drop_vld_gtp_v1_c_message_dropped_apn_filtering_no_apn"] = d.DropVldGtpV1CMessageDroppedApnFilteringNoApn

	in["drop_vld_gtp_v2_c_message_dropped_apn_filtering_no_apn"] = d.DropVldGtpV2CMessageDroppedApnFilteringNoApn

	in["drop_vld_gtp_v0_c_ie_len_exceed_msg_len"] = d.DropVldGtpV0CIeLenExceedMsgLen

	in["drop_vld_gtp_v1_c_ie_len_exceed_msg_len"] = d.DropVldGtpV1CIeLenExceedMsgLen

	in["drop_vld_gtp_v2_c_ie_len_exceed_msg_len"] = d.DropVldGtpV2CIeLenExceedMsgLen

	in["drop_vld_gtp_v0_c_message_length_mismatch"] = d.DropVldGtpV0CMessageLengthMismatch

	in["drop_vld_gtp_v1_c_message_length_mismatch"] = d.DropVldGtpV1CMessageLengthMismatch

	in["drop_vld_gtp_v2_c_message_length_mismatch"] = d.DropVldGtpV2CMessageLengthMismatch

	in["drop_flt_message_filtering"] = d.DropFltMessageFiltering

	in["drop_flt_apn_filtering"] = d.DropFltApnFiltering

	in["drop_flt_msisdn_filtering"] = d.DropFltMsisdnFiltering

	in["drop_flt_rat_type_filtering"] = d.DropFltRatTypeFiltering

	in["drop_flt_gtp_in_gtp"] = d.DropFltGtpInGtp

	in["drop_rl_gtp_v0_c_agg"] = d.DropRlGtpV0CAgg

	in["drop_rl_gtp_v1_c_agg"] = d.DropRlGtpV1CAgg

	in["drop_rl_gtp_v2_c_agg"] = d.DropRlGtpV2CAgg

	in["drop_rl_gtp_v1_c_create_pdp_request"] = d.DropRlGtpV1CCreatePdpRequest

	in["drop_rl_gtp_v2_c_create_session_request"] = d.DropRlGtpV2CCreateSessionRequest

	in["drop_rl_gtp_v1_c_update_pdp_request"] = d.DropRlGtpV1CUpdatePdpRequest

	in["drop_rl_gtp_v2_c_modify_bearer_request"] = d.DropRlGtpV2CModifyBearerRequest

	in["drop_rl_gtp_u_tunnel_create"] = d.DropRlGtpUTunnelCreate

	in["drop_rl_gtp_u_uplink_byte"] = d.DropRlGtpUUplinkByte

	in["drop_rl_gtp_u_uplink_packet"] = d.DropRlGtpUUplinkPacket

	in["drop_rl_gtp_u_downlink_byte"] = d.DropRlGtpUDownlinkByte

	in["drop_rl_gtp_u_downlink_packet"] = d.DropRlGtpUDownlinkPacket

	in["drop_rl_gtp_u_total_byte"] = d.DropRlGtpUTotalByte

	in["drop_rl_gtp_u_total_packet"] = d.DropRlGtpUTotalPacket

	in["drop_rl_gtp_u_max_concurrent_tunnels"] = d.DropRlGtpUMaxConcurrentTunnels
	result = append(result, in)
	return result
}

func setObjectFwGtpStatsStats(ret edpt.DataFwGtpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"out_of_session_memory":                      ret.DtFwGtpStats.Stats.OutOfSessionMemory,
			"no_fwd_route":                               ret.DtFwGtpStats.Stats.NoFwdRoute,
			"no_rev_route":                               ret.DtFwGtpStats.Stats.NoRevRoute,
			"gtp_smp_path_check_failed":                  ret.DtFwGtpStats.Stats.GtpSmpPathCheckFailed,
			"gtp_smp_check_failed":                       ret.DtFwGtpStats.Stats.GtpSmpCheckFailed,
			"gtp_smp_session_count_check_failed":         ret.DtFwGtpStats.Stats.GtpSmpSessionCountCheckFailed,
			"gtp_c_ref_count_smp_exceeded":               ret.DtFwGtpStats.Stats.GtpCRefCountSmpExceeded,
			"gtp_u_smp_in_rml_with_sess":                 ret.DtFwGtpStats.Stats.GtpUSmpInRmlWithSess,
			"gtp_tunnel_rate_limit_entry_create_failure": ret.DtFwGtpStats.Stats.GtpTunnelRateLimitEntryCreateFailure,
			"gtp_rate_limit_smp_create_failure":          ret.DtFwGtpStats.Stats.GtpRateLimitSmpCreateFailure,
			"gtp_rate_limit_t3_ctr_create_failure":       ret.DtFwGtpStats.Stats.GtpRateLimitT3CtrCreateFailure,
			"gtp_rate_limit_entry_create_failure":        ret.DtFwGtpStats.Stats.GtpRateLimitEntryCreateFailure,
			"gtp_node_restart_echo":                      ret.DtFwGtpStats.Stats.GtpNodeRestartEcho,
			"gtp_c_echo_path_failure":                    ret.DtFwGtpStats.Stats.GtpCEchoPathFailure,
			"drop_vld_gtp_echo_out_of_state_":            ret.DtFwGtpStats.Stats.DropVldGtpEchoOutOfState,
			"drop_vld_gtp_echo_ie_len_exceed_msg_len":    ret.DtFwGtpStats.Stats.DropVldGtpEchoIeLenExceedMsgLen,
			"gtp_del_bearer_request_retransmit":          ret.DtFwGtpStats.Stats.GtpDelBearerRequestRetransmit,
			"gtp_add_bearer_response_retransmit":         ret.DtFwGtpStats.Stats.GtpAddBearerResponseRetransmit,
			"gtp_u_out_of_state_drop":                    ret.DtFwGtpStats.Stats.GtpUOutOfStateDrop,
			"gtp_c_handover_request_out_of_state_drop":   ret.DtFwGtpStats.Stats.GtpCHandoverRequestOutOfStateDrop,
			"gtp_v1_c_nsapi_not_found_in_delete_req":     ret.DtFwGtpStats.Stats.GtpV1CNsapiNotFoundInDeleteReq,
			"gtp_v2_c_bearer_not_found_in_delete_req":    ret.DtFwGtpStats.Stats.GtpV2CBearerNotFoundInDeleteReq,
			"gtp_v2_c_bearer_not_found_in_delete_resp":   ret.DtFwGtpStats.Stats.GtpV2CBearerNotFoundInDeleteResp,
			"gtp_rr_message_drop":                        ret.DtFwGtpStats.Stats.GtpRrMessageDrop,
			"drop_gtp_frag_or_jumbo_pkt":                 ret.DtFwGtpStats.Stats.DropGtpFragOrJumboPkt,
			"gtp_c_handover_in_progress_with_conn":       ret.DtFwGtpStats.Stats.GtpCHandoverInProgressWithConn,
			"gtp_smp_dec_sess_count_check_failed":        ret.DtFwGtpStats.Stats.GtpSmpDecSessCountCheckFailed,
			"gtp_v0_c_uplink_ingress_packets":            ret.DtFwGtpStats.Stats.GtpV0CUplinkIngressPackets,
			"gtp_v0_c_uplink_egress_packets":             ret.DtFwGtpStats.Stats.GtpV0CUplinkEgressPackets,
			"gtp_v0_c_downlink_ingress_packets":          ret.DtFwGtpStats.Stats.GtpV0CDownlinkIngressPackets,
			"gtp_v0_c_downlink_egress_packets":           ret.DtFwGtpStats.Stats.GtpV0CDownlinkEgressPackets,
			"gtp_v0_c_uplink_ingress_bytes":              ret.DtFwGtpStats.Stats.GtpV0CUplinkIngressBytes,
			"gtp_v0_c_uplink_egress_bytes":               ret.DtFwGtpStats.Stats.GtpV0CUplinkEgressBytes,
			"gtp_v0_c_downlink_ingress_bytes":            ret.DtFwGtpStats.Stats.GtpV0CDownlinkIngressBytes,
			"gtp_v0_c_downlink_egress_bytes":             ret.DtFwGtpStats.Stats.GtpV0CDownlinkEgressBytes,
			"gtp_v1_c_uplink_ingress_packets":            ret.DtFwGtpStats.Stats.GtpV1CUplinkIngressPackets,
			"gtp_v1_c_uplink_egress_packets":             ret.DtFwGtpStats.Stats.GtpV1CUplinkEgressPackets,
			"gtp_v1_c_downlink_ingress_packets":          ret.DtFwGtpStats.Stats.GtpV1CDownlinkIngressPackets,
			"gtp_v1_c_downlink_egress_packets":           ret.DtFwGtpStats.Stats.GtpV1CDownlinkEgressPackets,
			"gtp_v1_c_uplink_ingress_bytes":              ret.DtFwGtpStats.Stats.GtpV1CUplinkIngressBytes,
			"gtp_v1_c_uplink_egress_bytes":               ret.DtFwGtpStats.Stats.GtpV1CUplinkEgressBytes,
			"gtp_v1_c_downlink_ingress_bytes":            ret.DtFwGtpStats.Stats.GtpV1CDownlinkIngressBytes,
			"gtp_v1_c_downlink_egress_bytes":             ret.DtFwGtpStats.Stats.GtpV1CDownlinkEgressBytes,
			"gtp_v2_c_uplink_ingress_packets":            ret.DtFwGtpStats.Stats.GtpV2CUplinkIngressPackets,
			"gtp_v2_c_uplink_egress_packets":             ret.DtFwGtpStats.Stats.GtpV2CUplinkEgressPackets,
			"gtp_v2_c_downlink_ingress_packets":          ret.DtFwGtpStats.Stats.GtpV2CDownlinkIngressPackets,
			"gtp_v2_c_downlink_egress_packets":           ret.DtFwGtpStats.Stats.GtpV2CDownlinkEgressPackets,
			"gtp_v2_c_uplink_ingress_bytes":              ret.DtFwGtpStats.Stats.GtpV2CUplinkIngressBytes,
			"gtp_v2_c_uplink_egress_bytes":               ret.DtFwGtpStats.Stats.GtpV2CUplinkEgressBytes,
			"gtp_v2_c_downlink_ingress_bytes":            ret.DtFwGtpStats.Stats.GtpV2CDownlinkIngressBytes,
			"gtp_v2_c_downlink_egress_bytes":             ret.DtFwGtpStats.Stats.GtpV2CDownlinkEgressBytes,
			"gtp_u_uplink_ingress_packets":               ret.DtFwGtpStats.Stats.GtpUUplinkIngressPackets,
			"gtp_u_uplink_egress_packets":                ret.DtFwGtpStats.Stats.GtpUUplinkEgressPackets,
			"gtp_u_downlink_ingress_packets":             ret.DtFwGtpStats.Stats.GtpUDownlinkIngressPackets,
			"gtp_u_downlink_egress_packets":              ret.DtFwGtpStats.Stats.GtpUDownlinkEgressPackets,
			"gtp_u_uplink_ingress_bytes":                 ret.DtFwGtpStats.Stats.GtpUUplinkIngressBytes,
			"gtp_u_uplink_egress_bytes":                  ret.DtFwGtpStats.Stats.GtpUUplinkEgressBytes,
			"gtp_u_downlink_ingress_bytes":               ret.DtFwGtpStats.Stats.GtpUDownlinkIngressBytes,
			"gtp_u_downlink_egress_bytes":                ret.DtFwGtpStats.Stats.GtpUDownlinkEgressBytes,
			"gtp_u_message_length_mismatch":              ret.DtFwGtpStats.Stats.GtpUMessageLengthMismatch,
			"gtp_path_message_length_mismatch":           ret.DtFwGtpStats.Stats.GtpPathMessageLengthMismatch,
			"drop_gtp_missing_cond_ie_bearer_ctx":        ret.DtFwGtpStats.Stats.DropGtpMissingCondIeBearerCtx,
			"drop_gtp_bearer_not_found_in_resp":          ret.DtFwGtpStats.Stats.DropGtpBearerNotFoundInResp,
			"gtp_stateless_forward":                      ret.DtFwGtpStats.Stats.GtpStatelessForward,
			"gtp_monitor_forward":                        ret.DtFwGtpStats.Stats.GtpMonitorForward,
		},
	}
}

func getObjectFwGtpStatsApnPrefix(d []interface{}) edpt.FwGtpStatsApnPrefix {

	count1 := len(d)
	var ret edpt.FwGtpStatsApnPrefix
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectFwGtpStatsApnPrefixStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectFwGtpStatsApnPrefixStats(d []interface{}) edpt.FwGtpStatsApnPrefixStats {

	count1 := len(d)
	var ret edpt.FwGtpStatsApnPrefixStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyName = in["key_name"].(string)
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
		ret.DropVldGtpV0CMessageDroppedApnFilteringNoApn = in["drop_vld_gtp_v0_c_message_dropped_apn_filtering_no_apn"].(int)
		ret.DropVldGtpV1CMessageDroppedApnFilteringNoApn = in["drop_vld_gtp_v1_c_message_dropped_apn_filtering_no_apn"].(int)
		ret.DropVldGtpV2CMessageDroppedApnFilteringNoApn = in["drop_vld_gtp_v2_c_message_dropped_apn_filtering_no_apn"].(int)
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

func getObjectFwGtpStatsNetworkElement(d []interface{}) edpt.FwGtpStatsNetworkElement {

	count1 := len(d)
	var ret edpt.FwGtpStatsNetworkElement
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectFwGtpStatsNetworkElementStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectFwGtpStatsNetworkElementStats(d []interface{}) edpt.FwGtpStatsNetworkElementStats {

	count1 := len(d)
	var ret edpt.FwGtpStatsNetworkElementStats
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

func getObjectFwGtpStatsStats(d []interface{}) edpt.FwGtpStatsStats {

	count1 := len(d)
	var ret edpt.FwGtpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OutOfSessionMemory = in["out_of_session_memory"].(int)
		ret.NoFwdRoute = in["no_fwd_route"].(int)
		ret.NoRevRoute = in["no_rev_route"].(int)
		ret.GtpSmpPathCheckFailed = in["gtp_smp_path_check_failed"].(int)
		ret.GtpSmpCheckFailed = in["gtp_smp_check_failed"].(int)
		ret.GtpSmpSessionCountCheckFailed = in["gtp_smp_session_count_check_failed"].(int)
		ret.GtpCRefCountSmpExceeded = in["gtp_c_ref_count_smp_exceeded"].(int)
		ret.GtpUSmpInRmlWithSess = in["gtp_u_smp_in_rml_with_sess"].(int)
		ret.GtpTunnelRateLimitEntryCreateFailure = in["gtp_tunnel_rate_limit_entry_create_failure"].(int)
		ret.GtpRateLimitSmpCreateFailure = in["gtp_rate_limit_smp_create_failure"].(int)
		ret.GtpRateLimitT3CtrCreateFailure = in["gtp_rate_limit_t3_ctr_create_failure"].(int)
		ret.GtpRateLimitEntryCreateFailure = in["gtp_rate_limit_entry_create_failure"].(int)
		ret.GtpNodeRestartEcho = in["gtp_node_restart_echo"].(int)
		ret.GtpCEchoPathFailure = in["gtp_c_echo_path_failure"].(int)
		ret.DropVldGtpEchoOutOfState = in["drop_vld_gtp_echo_out_of_state_"].(int)
		ret.DropVldGtpEchoIeLenExceedMsgLen = in["drop_vld_gtp_echo_ie_len_exceed_msg_len"].(int)
		ret.GtpDelBearerRequestRetransmit = in["gtp_del_bearer_request_retransmit"].(int)
		ret.GtpAddBearerResponseRetransmit = in["gtp_add_bearer_response_retransmit"].(int)
		ret.GtpUOutOfStateDrop = in["gtp_u_out_of_state_drop"].(int)
		ret.GtpCHandoverRequestOutOfStateDrop = in["gtp_c_handover_request_out_of_state_drop"].(int)
		ret.GtpV1CNsapiNotFoundInDeleteReq = in["gtp_v1_c_nsapi_not_found_in_delete_req"].(int)
		ret.GtpV2CBearerNotFoundInDeleteReq = in["gtp_v2_c_bearer_not_found_in_delete_req"].(int)
		ret.GtpV2CBearerNotFoundInDeleteResp = in["gtp_v2_c_bearer_not_found_in_delete_resp"].(int)
		ret.GtpRrMessageDrop = in["gtp_rr_message_drop"].(int)
		ret.DropGtpFragOrJumboPkt = in["drop_gtp_frag_or_jumbo_pkt"].(int)
		ret.GtpCHandoverInProgressWithConn = in["gtp_c_handover_in_progress_with_conn"].(int)
		ret.GtpSmpDecSessCountCheckFailed = in["gtp_smp_dec_sess_count_check_failed"].(int)
		ret.GtpV0CUplinkIngressPackets = in["gtp_v0_c_uplink_ingress_packets"].(int)
		ret.GtpV0CUplinkEgressPackets = in["gtp_v0_c_uplink_egress_packets"].(int)
		ret.GtpV0CDownlinkIngressPackets = in["gtp_v0_c_downlink_ingress_packets"].(int)
		ret.GtpV0CDownlinkEgressPackets = in["gtp_v0_c_downlink_egress_packets"].(int)
		ret.GtpV0CUplinkIngressBytes = in["gtp_v0_c_uplink_ingress_bytes"].(int)
		ret.GtpV0CUplinkEgressBytes = in["gtp_v0_c_uplink_egress_bytes"].(int)
		ret.GtpV0CDownlinkIngressBytes = in["gtp_v0_c_downlink_ingress_bytes"].(int)
		ret.GtpV0CDownlinkEgressBytes = in["gtp_v0_c_downlink_egress_bytes"].(int)
		ret.GtpV1CUplinkIngressPackets = in["gtp_v1_c_uplink_ingress_packets"].(int)
		ret.GtpV1CUplinkEgressPackets = in["gtp_v1_c_uplink_egress_packets"].(int)
		ret.GtpV1CDownlinkIngressPackets = in["gtp_v1_c_downlink_ingress_packets"].(int)
		ret.GtpV1CDownlinkEgressPackets = in["gtp_v1_c_downlink_egress_packets"].(int)
		ret.GtpV1CUplinkIngressBytes = in["gtp_v1_c_uplink_ingress_bytes"].(int)
		ret.GtpV1CUplinkEgressBytes = in["gtp_v1_c_uplink_egress_bytes"].(int)
		ret.GtpV1CDownlinkIngressBytes = in["gtp_v1_c_downlink_ingress_bytes"].(int)
		ret.GtpV1CDownlinkEgressBytes = in["gtp_v1_c_downlink_egress_bytes"].(int)
		ret.GtpV2CUplinkIngressPackets = in["gtp_v2_c_uplink_ingress_packets"].(int)
		ret.GtpV2CUplinkEgressPackets = in["gtp_v2_c_uplink_egress_packets"].(int)
		ret.GtpV2CDownlinkIngressPackets = in["gtp_v2_c_downlink_ingress_packets"].(int)
		ret.GtpV2CDownlinkEgressPackets = in["gtp_v2_c_downlink_egress_packets"].(int)
		ret.GtpV2CUplinkIngressBytes = in["gtp_v2_c_uplink_ingress_bytes"].(int)
		ret.GtpV2CUplinkEgressBytes = in["gtp_v2_c_uplink_egress_bytes"].(int)
		ret.GtpV2CDownlinkIngressBytes = in["gtp_v2_c_downlink_ingress_bytes"].(int)
		ret.GtpV2CDownlinkEgressBytes = in["gtp_v2_c_downlink_egress_bytes"].(int)
		ret.GtpUUplinkIngressPackets = in["gtp_u_uplink_ingress_packets"].(int)
		ret.GtpUUplinkEgressPackets = in["gtp_u_uplink_egress_packets"].(int)
		ret.GtpUDownlinkIngressPackets = in["gtp_u_downlink_ingress_packets"].(int)
		ret.GtpUDownlinkEgressPackets = in["gtp_u_downlink_egress_packets"].(int)
		ret.GtpUUplinkIngressBytes = in["gtp_u_uplink_ingress_bytes"].(int)
		ret.GtpUUplinkEgressBytes = in["gtp_u_uplink_egress_bytes"].(int)
		ret.GtpUDownlinkIngressBytes = in["gtp_u_downlink_ingress_bytes"].(int)
		ret.GtpUDownlinkEgressBytes = in["gtp_u_downlink_egress_bytes"].(int)
		ret.GtpUMessageLengthMismatch = in["gtp_u_message_length_mismatch"].(int)
		ret.GtpPathMessageLengthMismatch = in["gtp_path_message_length_mismatch"].(int)
		ret.DropGtpMissingCondIeBearerCtx = in["drop_gtp_missing_cond_ie_bearer_ctx"].(int)
		ret.DropGtpBearerNotFoundInResp = in["drop_gtp_bearer_not_found_in_resp"].(int)
		ret.GtpStatelessForward = in["gtp_stateless_forward"].(int)
		ret.GtpMonitorForward = in["gtp_monitor_forward"].(int)
	}
	return ret
}

func dataToEndpointFwGtpStats(d *schema.ResourceData) edpt.FwGtpStats {
	var ret edpt.FwGtpStats

	ret.ApnPrefix = getObjectFwGtpStatsApnPrefix(d.Get("apn_prefix").([]interface{}))

	ret.NetworkElement = getObjectFwGtpStatsNetworkElement(d.Get("network_element").([]interface{}))

	ret.Stats = getObjectFwGtpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
