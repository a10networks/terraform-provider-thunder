package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityTopn() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_topn`: Configure topn\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityTopnCreate,
		UpdateContext: resourceVisibilityTopnUpdate,
		ReadContext:   resourceVisibilityTopnRead,
		DeleteContext: resourceVisibilityTopnDelete,

		Schema: map[string]*schema.Schema{
			"cgnv6_nat_pool_topn_node": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activate": {
							Type: schema.TypeString, Optional: true, Description: "Name of the templated to be activated",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"cgnv6_nat_pool_topn_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Template Name",
						},
						"topn_size": {
							Type: schema.TypeInt, Optional: true, Description: "Congure value of N for topn",
						},
						"interval": {
							Type: schema.TypeString, Optional: true, Description: "'5': 5 minutes; '15': 15 minutes; '30': 30 minutes; '60': 60 minutes; 'all-time': Since template is activated;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"metrics": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"udp_total": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for UDP Total",
									},
									"tcp_total": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for TCP total",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"gtp_apn_prefix_topn_node": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activate": {
							Type: schema.TypeString, Optional: true, Description: "Name of the templated to be activated",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"gtp_apn_prefix_topn_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Template Name",
						},
						"topn_size": {
							Type: schema.TypeInt, Optional: true, Description: "Congure value of N for topn",
						},
						"interval": {
							Type: schema.TypeString, Optional: true, Description: "'5': 5 minutes; '15': 15 minutes; '30': 30 minutes; '60': 60 minutes; 'all-time': Since template is activated;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"metrics": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uplink_bytes": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Uplink Bytes",
									},
									"downlink_bytes": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Downlink Bytes",
									},
									"uplink_pkts": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Uplink Packets",
									},
									"downlink_pkts": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Downlink Packets",
									},
									"gtp_v0_c_tunnel_created": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Created",
									},
									"gtp_v0_c_tunnel_half_open": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Half open tunnel created",
									},
									"gtp_v0_c_tunnel_half_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Delete Request",
									},
									"gtp_v0_c_tunnel_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Marked Deleted",
									},
									"gtp_v0_c_tunnel_deleted": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Deleted",
									},
									"gtp_v0_c_half_open_tunnel_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Half open tunnel closed",
									},
									"gtp_v1_c_tunnel_created": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Created",
									},
									"gtp_v1_c_tunnel_half_open": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Half open tunnel created",
									},
									"gtp_v1_c_tunnel_half_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Delete Request",
									},
									"gtp_v1_c_tunnel_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Marked Deleted",
									},
									"gtp_v1_c_tunnel_deleted": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Deleted",
									},
									"gtp_v1_c_half_open_tunnel_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Half open tunnel closed",
									},
									"gtp_v2_c_tunnel_created": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Created",
									},
									"gtp_v2_c_tunnel_half_open": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Half open tunnel created",
									},
									"gtp_v2_c_tunnel_half_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Delete Request",
									},
									"gtp_v2_c_tunnel_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Marked Deleted",
									},
									"gtp_v2_c_tunnel_deleted": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Deleted",
									},
									"gtp_v2_c_half_open_tunnel_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Half open tunnel closed",
									},
									"gtp_u_tunnel_created": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP-U Tunnel Created",
									},
									"gtp_u_tunnel_deleted": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP-U Tunnel Deleted",
									},
									"gtp_v0_c_update_pdp_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPV0-C Update PDP Context Response Unsuccessful",
									},
									"gtp_v1_c_update_pdp_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPV1-C Update PDP Context response Unsuccessful",
									},
									"gtp_v2_c_mod_bearer_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPV2-C Modify Bearer Response Unsuccessful",
									},
									"gtp_v0_c_create_pdp_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPV0-C Create PDP Context Response Unsuccessful",
									},
									"gtp_v1_c_create_pdp_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPV1-C Create PDP Context Response Unsuccessful",
									},
									"gtp_v2_c_create_sess_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPV2-C Create Session Response Unsuccessful",
									},
									"gtp_v2_c_piggyback_message": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Piggyback Messages seen",
									},
									"gtp_path_management_message": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP Path Management Messages Received",
									},
									"gtp_v0_c_tunnel_deleted_restart": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Deleted with Restart/failure",
									},
									"gtp_v1_c_tunnel_deleted_restart": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Deleted with Restart/failure",
									},
									"gtp_v2_c_tunnel_deleted_restart": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Deleted with Restart/failure",
									},
									"drop_vld_reserved_field_set": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Reserved Header Field Set",
									},
									"drop_vld_tunnel_id_flag": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Tunnel Header Flag Not Set",
									},
									"drop_vld_invalid_flow_label_v0": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Invalid Flow Label in GTPv0-C Header",
									},
									"drop_vld_invalid_teid": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Invalid TEID Value",
									},
									"drop_vld_unsupported_message_type": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Message type not supported by GTP Version",
									},
									"drop_vld_out_of_state": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Out Of State GTP Message",
									},
									"drop_vld_mandatory_information_element": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Mandatory IE Not Present",
									},
									"drop_vld_out_of_order_ie": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTPv1-C Message Out of Order IE",
									},
									"drop_vld_out_of_state_ie": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Unexpected IE Present in Message",
									},
									"drop_vld_reserved_information_element": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Reserved IE Field Present",
									},
									"drop_vld_version_not_supported": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Invalid GTP version",
									},
									"drop_vld_message_length": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Message Length Exceeded",
									},
									"drop_vld_cross_layer_correlation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Cross Layer IP Address Mismatch",
									},
									"drop_vld_country_code_mismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Country Code Mismatch in IMSI and MSISDN",
									},
									"drop_vld_gtp_u_spoofed_source_address": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP-U IP Address Spoofed",
									},
									"drop_vld_gtp_bearer_count_exceed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP Bearer count exceeded max (11)",
									},
									"drop_vld_gtp_v2_wrong_lbi_create_bearer": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTPV2-C Wrong LBI in Create Bearer Request",
									},
									"gtp_c_handover_in_progress_with_conn": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP-C matching a conn with Handover In Progress",
									},
									"drop_vld_invalid_pkt_len_piggyback": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Piggyback message invalid packet length",
									},
									"drop_vld_sanity_failed_piggyback": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: piggyback message anomaly failed",
									},
									"drop_vld_sequence_num_correlation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP-C Sequence number Mismatch",
									},
									"drop_vld_gtpv0_seqnum_buffer_full": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTPV0-C conn Sequence number Buffer Full",
									},
									"drop_vld_gtpv1_seqnum_buffer_full": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTPV1-C conn Sequence number Buffer Full",
									},
									"drop_vld_gtpv2_seqnum_buffer_full": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTPV2-C conn Sequence number Buffer Full",
									},
									"drop_vld_gtp_invalid_imsi_len_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP-C Invalid IMSI Length Drop",
									},
									"drop_vld_gtp_invalid_apn_len_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP-C Invalid APN Length Drop",
									},
									"drop_vld_protocol_flag_unset": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Protocol flag in Header Field not Set",
									},
									"drop_flt_message_filtering": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: Message Type Not Permitted on Interface",
									},
									"drop_flt_apn_filtering": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: APN IMSI Filtering",
									},
									"drop_flt_msisdn_filtering": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: MSISDN Filtering",
									},
									"drop_flt_rat_type_filtering": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: RAT Type Filtering",
									},
									"drop_flt_gtp_in_gtp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: GTP in GTP Tunnel Present",
									},
									"drop_rl_gtp_v0_c_agg": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: Maximum GTPv0-C messages rate",
									},
									"drop_rl_gtp_v1_c_agg": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: Maximum GTPv1-C messages rate",
									},
									"drop_rl_gtp_v2_c_agg": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: Maximum GTPv2-C messages rate",
									},
									"drop_rl_gtp_v1_c_create_pdp_request": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTPv1-C Create PDP Req rate",
									},
									"drop_rl_gtp_v2_c_create_session_request": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTPv2-C Create Session Req rate",
									},
									"drop_rl_gtp_v1_c_update_pdp_request": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTPv1-C Update PDP Req rate",
									},
									"drop_rl_gtp_v2_c_modify_bearer_request": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTPv2-C Modify Bearer Req rate",
									},
									"drop_rl_gtp_u_tunnel_create": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Tunnel Creation rate",
									},
									"drop_rl_gtp_u_uplink_byte": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Uplink byte rate",
									},
									"drop_rl_gtp_u_uplink_packet": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Uplink packet rate",
									},
									"drop_rl_gtp_u_downlink_byte": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Downlink byte rate",
									},
									"drop_rl_gtp_u_downlink_packet": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Downlink packet rate",
									},
									"drop_rl_gtp_u_total_byte": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Total byte rate",
									},
									"drop_rl_gtp_u_total_packet": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Total packet rate",
									},
									"drop_rl_gtp_u_max_concurrent_tunnels": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Concurrent Tunnels",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"gtp_network_element_topn_node": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activate": {
							Type: schema.TypeString, Optional: true, Description: "Name of the templated to be activated",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"gtp_network_element_topn_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Template Name",
						},
						"topn_size": {
							Type: schema.TypeInt, Optional: true, Description: "Congure value of N for topn",
						},
						"interval": {
							Type: schema.TypeString, Optional: true, Description: "'5': 5 minutes; '15': 15 minutes; '30': 30 minutes; '60': 60 minutes; 'all-time': Since template is activated;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"metrics": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uplink_bytes": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Uplink Bytes",
									},
									"downlink_bytes": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Downlink Bytes",
									},
									"uplink_pkts": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Uplink Packets",
									},
									"downlink_pkts": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Downlink Packets",
									},
									"gtp_v0_c_tunnel_created": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Created",
									},
									"gtp_v0_c_tunnel_half_open": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Half open tunnel created",
									},
									"gtp_v0_c_tunnel_half_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Delete Request",
									},
									"gtp_v0_c_tunnel_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Marked Deleted",
									},
									"gtp_v0_c_tunnel_deleted": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Deleted",
									},
									"gtp_v0_c_half_open_tunnel_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Half open tunnel closed",
									},
									"gtp_v1_c_tunnel_created": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Created",
									},
									"gtp_v1_c_tunnel_half_open": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Half open tunnel created",
									},
									"gtp_v1_c_tunnel_half_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Delete Request",
									},
									"gtp_v1_c_tunnel_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Marked Deleted",
									},
									"gtp_v1_c_tunnel_deleted": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Deleted",
									},
									"gtp_v1_c_half_open_tunnel_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Half open tunnel closed",
									},
									"gtp_v2_c_tunnel_created": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Created",
									},
									"gtp_v2_c_tunnel_half_open": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Half open tunnel created",
									},
									"gtp_v2_c_tunnel_half_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Delete Request",
									},
									"gtp_v2_c_tunnel_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Marked Deleted",
									},
									"gtp_v2_c_tunnel_deleted": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Deleted",
									},
									"gtp_v2_c_half_open_tunnel_closed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Half open tunnel closed",
									},
									"gtp_u_tunnel_created": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP-U Tunnel Created",
									},
									"gtp_u_tunnel_deleted": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP-U Tunnel Deleted",
									},
									"gtp_v0_c_update_pdp_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Update PDP Context Response Unsuccessful",
									},
									"gtp_v1_c_update_pdp_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Update PDP Context Response Unsuccessful",
									},
									"gtp_v2_c_mod_bearer_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Modify Bearer Response Unsuccessful",
									},
									"gtp_v0_c_create_pdp_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Create PDP Context Response Unsuccessful",
									},
									"gtp_v1_c_create_pdp_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Create PDP Context Response Unsuccessful",
									},
									"gtp_v2_c_create_sess_resp_unsuccess": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Create Session Response Unsuccessful",
									},
									"gtp_v2_c_piggyback_message": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Piggyback Messages",
									},
									"gtp_path_management_message": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP Path Management Messages Received",
									},
									"gtp_v0_c_tunnel_deleted_restart": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Deleted with Restart/failure",
									},
									"gtp_v1_c_tunnel_deleted_restart": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Deleted with Restart/failure",
									},
									"gtp_v2_c_tunnel_deleted_restart": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Deleted with Restart/failure",
									},
									"gtp_v0_c_reserved_message_allow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Reserved Message Allow",
									},
									"gtp_v1_c_reserved_message_allow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Reserved Message Allow",
									},
									"gtp_v2_c_reserved_message_allow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Reserved Message Allow",
									},
									"drop_vld_reserved_field_set": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Reserved Header Field Set",
									},
									"drop_vld_tunnel_id_flag": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Tunnel Header Flag Not Set",
									},
									"drop_vld_invalid_flow_label_v0": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Invalid Flow Label in GTPv0-C Header",
									},
									"drop_vld_invalid_teid": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Invalid TEID Value",
									},
									"drop_vld_unsupported_message_type": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Message type not supported by GTP Version",
									},
									"drop_vld_out_of_state": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Out Of State GTP Message",
									},
									"drop_vld_mandatory_information_element": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Mandatory IE Not Present",
									},
									"drop_vld_out_of_order_ie": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTPv1-C Message Out of Order IE",
									},
									"drop_vld_out_of_state_ie": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Unexpected IE Present in Message",
									},
									"drop_vld_reserved_information_element": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Reserved IE Field Present",
									},
									"drop_vld_version_not_supported": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Invalid GTP version",
									},
									"drop_vld_message_length": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Message Length Exceeded",
									},
									"drop_vld_cross_layer_correlation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Cross Layer IP Address Mismatch",
									},
									"drop_vld_country_code_mismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Country Code Mismatch in IMSI and MSISDN",
									},
									"drop_vld_gtp_u_spoofed_source_address": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP-U IP Address Spoofed",
									},
									"drop_vld_gtp_bearer_count_exceed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP Bearer count exceeded max (11)",
									},
									"drop_vld_gtp_v2_wrong_lbi_create_bearer": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTPV2-C Wrong LBI in Create Bearer Request",
									},
									"gtp_c_handover_in_progress_with_conn": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP-C matching a conn with Handover In Progress",
									},
									"drop_vld_invalid_pkt_len_piggyback": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Piggyback message invalid packet length",
									},
									"drop_vld_sanity_failed_piggyback": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: piggyback message anomaly failed",
									},
									"drop_vld_sequence_num_correlation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP-C Sequence number Mismatch",
									},
									"drop_vld_gtpv0_seqnum_buffer_full": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTPV0-C conn Sequence number Buffer Full",
									},
									"drop_vld_gtpv1_seqnum_buffer_full": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTPV1-C conn Sequence number Buffer Full",
									},
									"drop_vld_gtpv2_seqnum_buffer_full": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTPV2-C conn Sequence number Buffer Full",
									},
									"drop_vld_gtp_invalid_imsi_len_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP-C Invalid IMSI Length Drop",
									},
									"drop_vld_gtp_invalid_apn_len_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP-C Invalid APN Length Drop",
									},
									"drop_vld_protocol_flag_unset": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Protocol flag in Header Field not Set",
									},
									"drop_flt_message_filtering": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: Message Type Not Permitted on Interface",
									},
									"drop_flt_apn_filtering": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: APN IMSI Filtering",
									},
									"drop_flt_msisdn_filtering": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: MSISDN Filtering",
									},
									"drop_flt_rat_type_filtering": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: RAT Type Filtering",
									},
									"drop_flt_gtp_in_gtp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: GTP in GTP Tunnel Present",
									},
									"drop_rl_gtp_v0_c_agg": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: Maximum GTPv0-C messages rate",
									},
									"drop_rl_gtp_v1_c_agg": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: Maximum GTPv1-C messages rate",
									},
									"drop_rl_gtp_v2_c_agg": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: Maximum GTPv2-C messages rate",
									},
									"drop_rl_gtp_v1_c_create_pdp_request": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTPv1-C Create PDP Req rate",
									},
									"drop_rl_gtp_v2_c_create_session_request": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTPv2-C Create Session Req rate",
									},
									"drop_rl_gtp_v1_c_update_pdp_request": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTPv1-C Update PDP Req rate",
									},
									"drop_rl_gtp_v2_c_modify_bearer_request": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTPv2-C Modify Bearer Req rate",
									},
									"drop_rl_gtp_u_tunnel_create": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Tunnel Creation rate",
									},
									"drop_rl_gtp_u_uplink_byte": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Uplink byte rate",
									},
									"drop_rl_gtp_u_uplink_packet": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Uplink packet rate",
									},
									"drop_rl_gtp_u_downlink_byte": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Downlink byte rate",
									},
									"drop_rl_gtp_u_downlink_packet": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Downlink packet rate",
									},
									"drop_rl_gtp_u_total_byte": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Total byte rate",
									},
									"drop_rl_gtp_u_total_packet": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Total packet rate",
									},
									"drop_rl_gtp_u_max_concurrent_tunnels": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Concurrent Tunnels",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'heap-alloc-success': Total heap node allocated; 'heap-alloc-failed': Total heap node alloc failed; 'heap-alloc-oom': Total heap node alloc failed Out of Memory; 'obj-reg-success': Total object node allocated; 'obj-reg-failed': Total object node alloc failed; 'obj-reg-oom': Total object node alloc failed Out of Memory; 'heap-deleted': Total Heap node deleted; 'obj-deleted': Total Object node deleted; 'heap-metric-alloc-success': Total heap metric node allocated; 'heap-metric-alloc-oom': Total heap metric node alloc failed Out of Memory; 'heap-move-to-delq': Total heap node moved to delq; 'heap-metric-deleted': Total Heap metric node deleted; 'obj-metric-reg-success': Total object Metric node allocated; 'obj-metric-reg-oom': Total object Metric node alloc failed Out of Memory; 'obj-move-to-delq': Total object node moved to delq; 'obj-metric-deleted': Total Object metric node deleted; 'hc-obj-alloc-failed': Send failed to HC, Out of Memory;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityTopnCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopn(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityTopnUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopn(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityTopnDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopn(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityTopnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopn(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityTopnCgnv6NatPoolTopnNode3134(d []interface{}) edpt.VisibilityTopnCgnv6NatPoolTopnNode3134 {

	count1 := len(d)
	var ret edpt.VisibilityTopnCgnv6NatPoolTopnNode3134
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Activate = in["activate"].(string)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityTopnCgnv6NatPoolTopnTmplList(d []interface{}) []edpt.VisibilityTopnCgnv6NatPoolTopnTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityTopnCgnv6NatPoolTopnTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityTopnCgnv6NatPoolTopnTmplList
		oi.Name = in["name"].(string)
		oi.TopnSize = in["topn_size"].(int)
		oi.Interval = in["interval"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.Metrics = getObjectVisibilityTopnCgnv6NatPoolTopnTmplListMetrics(in["metrics"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityTopnCgnv6NatPoolTopnTmplListMetrics(d []interface{}) edpt.VisibilityTopnCgnv6NatPoolTopnTmplListMetrics {

	count1 := len(d)
	var ret edpt.VisibilityTopnCgnv6NatPoolTopnTmplListMetrics
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UdpTotal = in["udp_total"].(int)
		ret.TcpTotal = in["tcp_total"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityTopnGtpApnPrefixTopnNode3135(d []interface{}) edpt.VisibilityTopnGtpApnPrefixTopnNode3135 {

	count1 := len(d)
	var ret edpt.VisibilityTopnGtpApnPrefixTopnNode3135
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Activate = in["activate"].(string)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityTopnGtpApnPrefixTopnTmplList(d []interface{}) []edpt.VisibilityTopnGtpApnPrefixTopnTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityTopnGtpApnPrefixTopnTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityTopnGtpApnPrefixTopnTmplList
		oi.Name = in["name"].(string)
		oi.TopnSize = in["topn_size"].(int)
		oi.Interval = in["interval"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.Metrics = getObjectVisibilityTopnGtpApnPrefixTopnTmplListMetrics(in["metrics"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityTopnGtpApnPrefixTopnTmplListMetrics(d []interface{}) edpt.VisibilityTopnGtpApnPrefixTopnTmplListMetrics {

	count1 := len(d)
	var ret edpt.VisibilityTopnGtpApnPrefixTopnTmplListMetrics
	if count1 > 0 {
		in := d[0].(map[string]interface{})
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
		ret.DropVldGtpV2WrongLbiCreateBearer = in["drop_vld_gtp_v2_wrong_lbi_create_bearer"].(int)
		ret.GtpCHandoverInProgressWithConn = in["gtp_c_handover_in_progress_with_conn"].(int)
		ret.DropVldInvalidPktLenPiggyback = in["drop_vld_invalid_pkt_len_piggyback"].(int)
		ret.DropVldSanityFailedPiggyback = in["drop_vld_sanity_failed_piggyback"].(int)
		ret.DropVldSequenceNumCorrelation = in["drop_vld_sequence_num_correlation"].(int)
		ret.DropVldGtpv0SeqnumBufferFull = in["drop_vld_gtpv0_seqnum_buffer_full"].(int)
		ret.DropVldGtpv1SeqnumBufferFull = in["drop_vld_gtpv1_seqnum_buffer_full"].(int)
		ret.DropVldGtpv2SeqnumBufferFull = in["drop_vld_gtpv2_seqnum_buffer_full"].(int)
		ret.DropVldGtpInvalidImsiLenDrop = in["drop_vld_gtp_invalid_imsi_len_drop"].(int)
		ret.DropVldGtpInvalidApnLenDrop = in["drop_vld_gtp_invalid_apn_len_drop"].(int)
		ret.DropVldProtocolFlagUnset = in["drop_vld_protocol_flag_unset"].(int)
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
		//omit uuid
	}
	return ret
}

func getObjectVisibilityTopnGtpNetworkElementTopnNode3136(d []interface{}) edpt.VisibilityTopnGtpNetworkElementTopnNode3136 {

	count1 := len(d)
	var ret edpt.VisibilityTopnGtpNetworkElementTopnNode3136
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Activate = in["activate"].(string)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityTopnGtpNetworkElementTopnTmplList(d []interface{}) []edpt.VisibilityTopnGtpNetworkElementTopnTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityTopnGtpNetworkElementTopnTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityTopnGtpNetworkElementTopnTmplList
		oi.Name = in["name"].(string)
		oi.TopnSize = in["topn_size"].(int)
		oi.Interval = in["interval"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.Metrics = getObjectVisibilityTopnGtpNetworkElementTopnTmplListMetrics(in["metrics"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityTopnGtpNetworkElementTopnTmplListMetrics(d []interface{}) edpt.VisibilityTopnGtpNetworkElementTopnTmplListMetrics {

	count1 := len(d)
	var ret edpt.VisibilityTopnGtpNetworkElementTopnTmplListMetrics
	if count1 > 0 {
		in := d[0].(map[string]interface{})
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
		ret.DropVldGtpV2WrongLbiCreateBearer = in["drop_vld_gtp_v2_wrong_lbi_create_bearer"].(int)
		ret.GtpCHandoverInProgressWithConn = in["gtp_c_handover_in_progress_with_conn"].(int)
		ret.DropVldInvalidPktLenPiggyback = in["drop_vld_invalid_pkt_len_piggyback"].(int)
		ret.DropVldSanityFailedPiggyback = in["drop_vld_sanity_failed_piggyback"].(int)
		ret.DropVldSequenceNumCorrelation = in["drop_vld_sequence_num_correlation"].(int)
		ret.DropVldGtpv0SeqnumBufferFull = in["drop_vld_gtpv0_seqnum_buffer_full"].(int)
		ret.DropVldGtpv1SeqnumBufferFull = in["drop_vld_gtpv1_seqnum_buffer_full"].(int)
		ret.DropVldGtpv2SeqnumBufferFull = in["drop_vld_gtpv2_seqnum_buffer_full"].(int)
		ret.DropVldGtpInvalidImsiLenDrop = in["drop_vld_gtp_invalid_imsi_len_drop"].(int)
		ret.DropVldGtpInvalidApnLenDrop = in["drop_vld_gtp_invalid_apn_len_drop"].(int)
		ret.DropVldProtocolFlagUnset = in["drop_vld_protocol_flag_unset"].(int)
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
		//omit uuid
	}
	return ret
}

func getSliceVisibilityTopnSamplingEnable(d []interface{}) []edpt.VisibilityTopnSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VisibilityTopnSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityTopnSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityTopn(d *schema.ResourceData) edpt.VisibilityTopn {
	var ret edpt.VisibilityTopn
	ret.Inst.Cgnv6NatPoolTopnNode = getObjectVisibilityTopnCgnv6NatPoolTopnNode3134(d.Get("cgnv6_nat_pool_topn_node").([]interface{}))
	ret.Inst.Cgnv6NatPoolTopnTmplList = getSliceVisibilityTopnCgnv6NatPoolTopnTmplList(d.Get("cgnv6_nat_pool_topn_tmpl_list").([]interface{}))
	ret.Inst.GtpApnPrefixTopnNode = getObjectVisibilityTopnGtpApnPrefixTopnNode3135(d.Get("gtp_apn_prefix_topn_node").([]interface{}))
	ret.Inst.GtpApnPrefixTopnTmplList = getSliceVisibilityTopnGtpApnPrefixTopnTmplList(d.Get("gtp_apn_prefix_topn_tmpl_list").([]interface{}))
	ret.Inst.GtpNetworkElementTopnNode = getObjectVisibilityTopnGtpNetworkElementTopnNode3136(d.Get("gtp_network_element_topn_node").([]interface{}))
	ret.Inst.GtpNetworkElementTopnTmplList = getSliceVisibilityTopnGtpNetworkElementTopnTmplList(d.Get("gtp_network_element_topn_tmpl_list").([]interface{}))
	ret.Inst.SamplingEnable = getSliceVisibilityTopnSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
