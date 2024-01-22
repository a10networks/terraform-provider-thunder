package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityTopnGtpNetworkElementTopnTmplMetrics() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_topn_gtp_network_element_topn_tmpl_metrics`: Configure topn metrics for fw.gtp.network-element\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsCreate,
		UpdateContext: resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsUpdate,
		ReadContext:   resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsRead,
		DeleteContext: resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsDelete,

		Schema: map[string]*schema.Schema{
			"downlink_bytes": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Downlink Bytes",
			},
			"downlink_pkts": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Downlink Packets",
			},
			"drop_flt_apn_filtering": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: APN IMSI Filtering",
			},
			"drop_flt_gtp_in_gtp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: GTP in GTP Tunnel Present",
			},
			"drop_flt_message_filtering": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: Message Type Not Permitted on Interface",
			},
			"drop_flt_msisdn_filtering": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: MSISDN Filtering",
			},
			"drop_flt_rat_type_filtering": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Filtering Drop: RAT Type Filtering",
			},
			"drop_rl_gtp_u_downlink_byte": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Downlink byte rate",
			},
			"drop_rl_gtp_u_downlink_packet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Downlink packet rate",
			},
			"drop_rl_gtp_u_max_concurrent_tunnels": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Concurrent Tunnels",
			},
			"drop_rl_gtp_u_total_byte": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Total byte rate",
			},
			"drop_rl_gtp_u_total_packet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTP-U Total packet rate",
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
			"drop_rl_gtp_v0_c_agg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: Maximum GTPv0-C messages rate",
			},
			"drop_rl_gtp_v1_c_agg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: Maximum GTPv1-C messages rate",
			},
			"drop_rl_gtp_v1_c_create_pdp_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTPv1-C Create PDP Req rate",
			},
			"drop_rl_gtp_v1_c_update_pdp_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTPv1-C Update PDP Req rate",
			},
			"drop_rl_gtp_v2_c_agg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: Maximum GTPv2-C messages rate",
			},
			"drop_rl_gtp_v2_c_create_session_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTPv2-C Create Session Req rate",
			},
			"drop_rl_gtp_v2_c_modify_bearer_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Rate-limit Drop: GTPv2-C Modify Bearer Req rate",
			},
			"drop_vld_country_code_mismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Country Code Mismatch in IMSI and MSISDN",
			},
			"drop_vld_cross_layer_correlation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Cross Layer IP Address Mismatch",
			},
			"drop_vld_gtp_invalid_apn_len_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP-C Invalid APN Length Drop",
			},
			"drop_vld_gtp_invalid_imsi_len_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP-C Invalid IMSI Length Drop",
			},
			"drop_vld_gtp_u_spoofed_source_address": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP-U IP Address Spoofed",
			},
			"drop_vld_gtp_v2_wrong_lbi_create_bearer": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTPV2-C Wrong LBI in Create Bearer Request",
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
			"drop_vld_invalid_flow_label_v0": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Invalid Flow Label in GTPv0-C Header",
			},
			"drop_vld_invalid_pkt_len_piggyback": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Piggyback message invalid packet length",
			},
			"drop_vld_invalid_teid": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Invalid TEID Value",
			},
			"drop_vld_mandatory_information_element": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Mandatory IE Not Present",
			},
			"drop_vld_message_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Message Length Exceeded",
			},
			"drop_vld_out_of_order_ie": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTPv1-C Message Out of Order IE",
			},
			"drop_vld_out_of_state": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Out Of State GTP Message",
			},
			"drop_vld_out_of_state_ie": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Unexpected IE Present in Message",
			},
			"drop_vld_protocol_flag_unset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Protocol flag in Header Field not Set",
			},
			"drop_vld_reserved_field_set": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Reserved Header Field Set",
			},
			"drop_vld_reserved_information_element": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Reserved IE Field Present",
			},
			"drop_vld_sanity_failed_piggyback": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: piggyback message anomaly failed",
			},
			"drop_vld_sequence_num_correlation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP-C Sequence number Mismatch",
			},
			"drop_vld_tunnel_id_flag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Tunnel Header Flag Not Set",
			},
			"drop_vld_unsupported_message_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Message type not supported by GTP Version",
			},
			"drop_vld_version_not_supported": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: Invalid GTP version",
			},
			"drop_vld_gtp_bearer_count_exceed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Validation Drop: GTP Bearer count exceeded max (11)",
			},
			"gtp_c_handover_in_progress_with_conn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP-C matching a conn with Handover In Progress",
			},
			"gtp_path_management_message": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP Path Management Messages Received",
			},
			"gtp_u_tunnel_created": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP-U Tunnel Created",
			},
			"gtp_u_tunnel_deleted": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP-U Tunnel Deleted",
			},
			"gtp_v0_c_create_pdp_resp_unsuccess": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Create PDP Context Response Unsuccessful",
			},
			"gtp_v0_c_half_open_tunnel_closed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Half open tunnel closed",
			},
			"gtp_v0_c_reserved_message_allow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Reserved Message Allow",
			},
			"gtp_v0_c_tunnel_closed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Marked Deleted",
			},
			"gtp_v0_c_tunnel_created": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Created",
			},
			"gtp_v0_c_tunnel_deleted": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Deleted",
			},
			"gtp_v0_c_tunnel_deleted_restart": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Deleted with Restart/failure",
			},
			"gtp_v0_c_tunnel_half_closed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Tunnel Delete Request",
			},
			"gtp_v0_c_tunnel_half_open": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Half open tunnel created",
			},
			"gtp_v0_c_update_pdp_resp_unsuccess": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv0-C Update PDP Context Response Unsuccessful",
			},
			"gtp_v1_c_create_pdp_resp_unsuccess": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Create PDP Context Response Unsuccessful",
			},
			"gtp_v1_c_half_open_tunnel_closed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Half open tunnel closed",
			},
			"gtp_v1_c_reserved_message_allow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Reserved Message Allow",
			},
			"gtp_v1_c_tunnel_closed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Marked Deleted",
			},
			"gtp_v1_c_tunnel_created": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Created",
			},
			"gtp_v1_c_tunnel_deleted": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Deleted",
			},
			"gtp_v1_c_tunnel_deleted_restart": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Deleted with Restart/failure",
			},
			"gtp_v1_c_tunnel_half_closed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Tunnel Delete Request",
			},
			"gtp_v1_c_tunnel_half_open": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Half open tunnel created",
			},
			"gtp_v1_c_update_pdp_resp_unsuccess": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv1-C Update PDP Context Response Unsuccessful",
			},
			"gtp_v2_c_create_sess_resp_unsuccess": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Create Session Response Unsuccessful",
			},
			"gtp_v2_c_half_open_tunnel_closed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Half open tunnel closed",
			},
			"gtp_v2_c_mod_bearer_resp_unsuccess": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Modify Bearer Response Unsuccessful",
			},
			"gtp_v2_c_piggyback_message": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Piggyback Messages",
			},
			"gtp_v2_c_reserved_message_allow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Reserved Message Allow",
			},
			"gtp_v2_c_tunnel_closed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Marked Deleted",
			},
			"gtp_v2_c_tunnel_created": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Created",
			},
			"gtp_v2_c_tunnel_deleted": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Deleted",
			},
			"gtp_v2_c_tunnel_deleted_restart": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Deleted with Restart/failure",
			},
			"gtp_v2_c_tunnel_half_closed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Tunnel Delete Request",
			},
			"gtp_v2_c_tunnel_half_open": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTPv2-C Half open tunnel created",
			},
			"uplink_bytes": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Uplink Bytes",
			},
			"uplink_pkts": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for Uplink Packets",
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
func resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnGtpNetworkElementTopnTmplMetrics(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnGtpNetworkElementTopnTmplMetrics(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnGtpNetworkElementTopnTmplMetrics(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnGtpNetworkElementTopnTmplMetricsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnGtpNetworkElementTopnTmplMetrics(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityTopnGtpNetworkElementTopnTmplMetrics(d *schema.ResourceData) edpt.VisibilityTopnGtpNetworkElementTopnTmplMetrics {
	var ret edpt.VisibilityTopnGtpNetworkElementTopnTmplMetrics
	ret.Inst.DownlinkBytes = d.Get("downlink_bytes").(int)
	ret.Inst.DownlinkPkts = d.Get("downlink_pkts").(int)
	ret.Inst.DropFltApnFiltering = d.Get("drop_flt_apn_filtering").(int)
	ret.Inst.DropFltGtpInGtp = d.Get("drop_flt_gtp_in_gtp").(int)
	ret.Inst.DropFltMessageFiltering = d.Get("drop_flt_message_filtering").(int)
	ret.Inst.DropFltMsisdnFiltering = d.Get("drop_flt_msisdn_filtering").(int)
	ret.Inst.DropFltRatTypeFiltering = d.Get("drop_flt_rat_type_filtering").(int)
	ret.Inst.DropRlGtpUDownlinkByte = d.Get("drop_rl_gtp_u_downlink_byte").(int)
	ret.Inst.DropRlGtpUDownlinkPacket = d.Get("drop_rl_gtp_u_downlink_packet").(int)
	ret.Inst.DropRlGtpUMaxConcurrentTunnels = d.Get("drop_rl_gtp_u_max_concurrent_tunnels").(int)
	ret.Inst.DropRlGtpUTotalByte = d.Get("drop_rl_gtp_u_total_byte").(int)
	ret.Inst.DropRlGtpUTotalPacket = d.Get("drop_rl_gtp_u_total_packet").(int)
	ret.Inst.DropRlGtpUTunnelCreate = d.Get("drop_rl_gtp_u_tunnel_create").(int)
	ret.Inst.DropRlGtpUUplinkByte = d.Get("drop_rl_gtp_u_uplink_byte").(int)
	ret.Inst.DropRlGtpUUplinkPacket = d.Get("drop_rl_gtp_u_uplink_packet").(int)
	ret.Inst.DropRlGtpV0CAgg = d.Get("drop_rl_gtp_v0_c_agg").(int)
	ret.Inst.DropRlGtpV1CAgg = d.Get("drop_rl_gtp_v1_c_agg").(int)
	ret.Inst.DropRlGtpV1CCreatePdpRequest = d.Get("drop_rl_gtp_v1_c_create_pdp_request").(int)
	ret.Inst.DropRlGtpV1CUpdatePdpRequest = d.Get("drop_rl_gtp_v1_c_update_pdp_request").(int)
	ret.Inst.DropRlGtpV2CAgg = d.Get("drop_rl_gtp_v2_c_agg").(int)
	ret.Inst.DropRlGtpV2CCreateSessionRequest = d.Get("drop_rl_gtp_v2_c_create_session_request").(int)
	ret.Inst.DropRlGtpV2CModifyBearerRequest = d.Get("drop_rl_gtp_v2_c_modify_bearer_request").(int)
	ret.Inst.DropVldCountryCodeMismatch = d.Get("drop_vld_country_code_mismatch").(int)
	ret.Inst.DropVldCrossLayerCorrelation = d.Get("drop_vld_cross_layer_correlation").(int)
	ret.Inst.DropVldGtpInvalidApnLenDrop = d.Get("drop_vld_gtp_invalid_apn_len_drop").(int)
	ret.Inst.DropVldGtpInvalidImsiLenDrop = d.Get("drop_vld_gtp_invalid_imsi_len_drop").(int)
	ret.Inst.DropVldGtpUSpoofedSourceAddress = d.Get("drop_vld_gtp_u_spoofed_source_address").(int)
	ret.Inst.DropVldGtpV2WrongLbiCreateBearer = d.Get("drop_vld_gtp_v2_wrong_lbi_create_bearer").(int)
	ret.Inst.DropVldGtpv0SeqnumBufferFull = d.Get("drop_vld_gtpv0_seqnum_buffer_full").(int)
	ret.Inst.DropVldGtpv1SeqnumBufferFull = d.Get("drop_vld_gtpv1_seqnum_buffer_full").(int)
	ret.Inst.DropVldGtpv2SeqnumBufferFull = d.Get("drop_vld_gtpv2_seqnum_buffer_full").(int)
	ret.Inst.DropVldInvalidFlowLabelV0 = d.Get("drop_vld_invalid_flow_label_v0").(int)
	ret.Inst.DropVldInvalidPktLenPiggyback = d.Get("drop_vld_invalid_pkt_len_piggyback").(int)
	ret.Inst.DropVldInvalidTeid = d.Get("drop_vld_invalid_teid").(int)
	ret.Inst.DropVldMandatoryInformationElement = d.Get("drop_vld_mandatory_information_element").(int)
	ret.Inst.DropVldMessageLength = d.Get("drop_vld_message_length").(int)
	ret.Inst.DropVldOutOfOrderIe = d.Get("drop_vld_out_of_order_ie").(int)
	ret.Inst.DropVldOutOfState = d.Get("drop_vld_out_of_state").(int)
	ret.Inst.DropVldOutOfStateIe = d.Get("drop_vld_out_of_state_ie").(int)
	ret.Inst.DropVldProtocolFlagUnset = d.Get("drop_vld_protocol_flag_unset").(int)
	ret.Inst.DropVldReservedFieldSet = d.Get("drop_vld_reserved_field_set").(int)
	ret.Inst.DropVldReservedInformationElement = d.Get("drop_vld_reserved_information_element").(int)
	ret.Inst.DropVldSanityFailedPiggyback = d.Get("drop_vld_sanity_failed_piggyback").(int)
	ret.Inst.DropVldSequenceNumCorrelation = d.Get("drop_vld_sequence_num_correlation").(int)
	ret.Inst.DropVldTunnelIdFlag = d.Get("drop_vld_tunnel_id_flag").(int)
	ret.Inst.DropVldUnsupportedMessageType = d.Get("drop_vld_unsupported_message_type").(int)
	ret.Inst.DropVldVersionNotSupported = d.Get("drop_vld_version_not_supported").(int)
	ret.Inst.Drop_vldGtpBearerCountExceed = d.Get("drop_vld_gtp_bearer_count_exceed").(int)
	ret.Inst.GtpCHandoverInProgressWithConn = d.Get("gtp_c_handover_in_progress_with_conn").(int)
	ret.Inst.GtpPathManagementMessage = d.Get("gtp_path_management_message").(int)
	ret.Inst.GtpUTunnelCreated = d.Get("gtp_u_tunnel_created").(int)
	ret.Inst.GtpUTunnelDeleted = d.Get("gtp_u_tunnel_deleted").(int)
	ret.Inst.GtpV0CCreatePdpRespUnsuccess = d.Get("gtp_v0_c_create_pdp_resp_unsuccess").(int)
	ret.Inst.GtpV0CHalfOpenTunnelClosed = d.Get("gtp_v0_c_half_open_tunnel_closed").(int)
	ret.Inst.GtpV0CReservedMessageAllow = d.Get("gtp_v0_c_reserved_message_allow").(int)
	ret.Inst.GtpV0CTunnelClosed = d.Get("gtp_v0_c_tunnel_closed").(int)
	ret.Inst.GtpV0CTunnelCreated = d.Get("gtp_v0_c_tunnel_created").(int)
	ret.Inst.GtpV0CTunnelDeleted = d.Get("gtp_v0_c_tunnel_deleted").(int)
	ret.Inst.GtpV0CTunnelDeletedRestart = d.Get("gtp_v0_c_tunnel_deleted_restart").(int)
	ret.Inst.GtpV0CTunnelHalfClosed = d.Get("gtp_v0_c_tunnel_half_closed").(int)
	ret.Inst.GtpV0CTunnelHalfOpen = d.Get("gtp_v0_c_tunnel_half_open").(int)
	ret.Inst.GtpV0CUpdatePdpRespUnsuccess = d.Get("gtp_v0_c_update_pdp_resp_unsuccess").(int)
	ret.Inst.GtpV1CCreatePdpRespUnsuccess = d.Get("gtp_v1_c_create_pdp_resp_unsuccess").(int)
	ret.Inst.GtpV1CHalfOpenTunnelClosed = d.Get("gtp_v1_c_half_open_tunnel_closed").(int)
	ret.Inst.GtpV1CReservedMessageAllow = d.Get("gtp_v1_c_reserved_message_allow").(int)
	ret.Inst.GtpV1CTunnelClosed = d.Get("gtp_v1_c_tunnel_closed").(int)
	ret.Inst.GtpV1CTunnelCreated = d.Get("gtp_v1_c_tunnel_created").(int)
	ret.Inst.GtpV1CTunnelDeleted = d.Get("gtp_v1_c_tunnel_deleted").(int)
	ret.Inst.GtpV1CTunnelDeletedRestart = d.Get("gtp_v1_c_tunnel_deleted_restart").(int)
	ret.Inst.GtpV1CTunnelHalfClosed = d.Get("gtp_v1_c_tunnel_half_closed").(int)
	ret.Inst.GtpV1CTunnelHalfOpen = d.Get("gtp_v1_c_tunnel_half_open").(int)
	ret.Inst.GtpV1CUpdatePdpRespUnsuccess = d.Get("gtp_v1_c_update_pdp_resp_unsuccess").(int)
	ret.Inst.GtpV2CCreateSessRespUnsuccess = d.Get("gtp_v2_c_create_sess_resp_unsuccess").(int)
	ret.Inst.GtpV2CHalfOpenTunnelClosed = d.Get("gtp_v2_c_half_open_tunnel_closed").(int)
	ret.Inst.GtpV2CMod_bearerRespUnsuccess = d.Get("gtp_v2_c_mod_bearer_resp_unsuccess").(int)
	ret.Inst.GtpV2CPiggybackMessage = d.Get("gtp_v2_c_piggyback_message").(int)
	ret.Inst.GtpV2CReservedMessageAllow = d.Get("gtp_v2_c_reserved_message_allow").(int)
	ret.Inst.GtpV2CTunnelClosed = d.Get("gtp_v2_c_tunnel_closed").(int)
	ret.Inst.GtpV2CTunnelCreated = d.Get("gtp_v2_c_tunnel_created").(int)
	ret.Inst.GtpV2CTunnelDeleted = d.Get("gtp_v2_c_tunnel_deleted").(int)
	ret.Inst.GtpV2CTunnelDeletedRestart = d.Get("gtp_v2_c_tunnel_deleted_restart").(int)
	ret.Inst.GtpV2CTunnelHalfClosed = d.Get("gtp_v2_c_tunnel_half_closed").(int)
	ret.Inst.GtpV2CTunnelHalfOpen = d.Get("gtp_v2_c_tunnel_half_open").(int)
	ret.Inst.UplinkBytes = d.Get("uplink_bytes").(int)
	ret.Inst.UplinkPkts = d.Get("uplink_pkts").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
