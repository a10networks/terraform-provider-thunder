package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_templ_gtp_plcy_tmpl`: Configure template for template.gtp-policy\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplDelete,

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
						"drop_vld_gtp_ie_repeat_count_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP repeated IE count exceeded",
						},
						"drop_vld_reserved_field_set": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Reserved Header Field Set",
						},
						"drop_vld_tunnel_id_flag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Tunnel Header Flag Not Set",
						},
						"drop_vld_invalid_flow_label_v0": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid Flow Label in GTPv0-C Header",
						},
						"drop_vld_invalid_teid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid TEID Value",
						},
						"drop_vld_out_of_state": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Out Of State GTP Message",
						},
						"drop_vld_mandatory_information_element": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Mandatory IE Not Present",
						},
						"drop_vld_mandatory_ie_in_grouped_ie": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Mandatory IE in Grouped IE Not Present",
						},
						"drop_vld_out_of_order_ie": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv1-C Message Out of Order IE",
						},
						"drop_vld_out_of_state_ie": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Unexpected IE Present in Message",
						},
						"drop_vld_reserved_information_element": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Reserved IE Field Present",
						},
						"drop_vld_version_not_supported": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid GTP version",
						},
						"drop_vld_message_length": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Message Length Exceeded",
						},
						"drop_vld_cross_layer_correlation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Cross Layer IP Address Mismatch",
						},
						"drop_vld_country_code_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Country Code Mismatch in IMSI and MSISDN",
						},
						"drop_vld_gtp_u_spoofed_source_address": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-U IP Address Spoofed",
						},
						"drop_vld_gtp_bearer_count_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP Bearer count exceeded max (11)",
						},
						"drop_vld_gtp_v2_wrong_lbi_create_bearer": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV2-C Wrong LBI in Create Bearer Request",
						},
						"drop_vld_v0_reserved_message_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv0-C Reserved Message Drop",
						},
						"drop_vld_v1_reserved_message_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv1-C Reserved Message Drop",
						},
						"drop_vld_v2_reserved_message_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv2-C Reserved Message Drop",
						},
						"drop_vld_invalid_pkt_len_piggyback": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Piggyback message invalid packet length",
						},
						"drop_vld_sanity_failed_piggyback": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: piggyback message anomaly failed",
						},
						"drop_vld_sequence_num_correlation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Sequence number Mismatch",
						},
						"drop_vld_gtpv0_seqnum_buffer_full": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV0-C conn Sequence number Buffer Full",
						},
						"drop_vld_gtpv1_seqnum_buffer_full": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV1-C conn Sequence number Buffer Full",
						},
						"drop_vld_gtpv2_seqnum_buffer_full": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV2-C conn Sequence number Buffer Full",
						},
						"drop_vld_gtp_invalid_imsi_len_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Invalid IMSI Length Drop",
						},
						"drop_vld_gtp_invalid_apn_len_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Invalid APN Length Drop",
						},
						"drop_vld_protocol_flag_unset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Protocol flag in Header Field not Set",
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
						"drop_vld_gtp_ie_repeat_count_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP repeated IE count exceeded",
						},
						"drop_vld_reserved_field_set": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Reserved Header Field Set",
						},
						"drop_vld_tunnel_id_flag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Tunnel Header Flag Not Set",
						},
						"drop_vld_invalid_flow_label_v0": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid Flow Label in GTPv0-C Header",
						},
						"drop_vld_invalid_teid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid TEID Value",
						},
						"drop_vld_out_of_state": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Out Of State GTP Message",
						},
						"drop_vld_mandatory_information_element": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Mandatory IE Not Present",
						},
						"drop_vld_mandatory_ie_in_grouped_ie": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Mandatory IE in Grouped IE Not Present",
						},
						"drop_vld_out_of_order_ie": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv1-C Message Out of Order IE",
						},
						"drop_vld_out_of_state_ie": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Unexpected IE Present in Message",
						},
						"drop_vld_reserved_information_element": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Reserved IE Field Present",
						},
						"drop_vld_version_not_supported": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid GTP version",
						},
						"drop_vld_message_length": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Message Length Exceeded",
						},
						"drop_vld_cross_layer_correlation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Cross Layer IP Address Mismatch",
						},
						"drop_vld_country_code_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Country Code Mismatch in IMSI and MSISDN",
						},
						"drop_vld_gtp_u_spoofed_source_address": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-U IP Address Spoofed",
						},
						"drop_vld_gtp_bearer_count_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP Bearer count exceeded max (11)",
						},
						"drop_vld_gtp_v2_wrong_lbi_create_bearer": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV2-C Wrong LBI in Create Bearer Request",
						},
						"drop_vld_v0_reserved_message_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv0-C Reserved Message Drop",
						},
						"drop_vld_v1_reserved_message_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv1-C Reserved Message Drop",
						},
						"drop_vld_v2_reserved_message_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv2-C Reserved Message Drop",
						},
						"drop_vld_invalid_pkt_len_piggyback": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Piggyback message invalid packet length",
						},
						"drop_vld_sanity_failed_piggyback": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: piggyback message anomaly failed",
						},
						"drop_vld_sequence_num_correlation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Sequence number Mismatch",
						},
						"drop_vld_gtpv0_seqnum_buffer_full": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV0-C conn Sequence number Buffer Full",
						},
						"drop_vld_gtpv1_seqnum_buffer_full": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV1-C conn Sequence number Buffer Full",
						},
						"drop_vld_gtpv2_seqnum_buffer_full": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV2-C conn Sequence number Buffer Full",
						},
						"drop_vld_gtp_invalid_imsi_len_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Invalid IMSI Length Drop",
						},
						"drop_vld_gtp_invalid_apn_len_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Invalid APN Length Drop",
						},
						"drop_vld_protocol_flag_unset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Protocol flag in Header Field not Set",
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
func resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsInc2720(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsInc2720 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsInc2720
	if count1 > 0 {
		in := d[0].(map[string]interface{})
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
		ret.DropVldGtpV2WrongLbiCreateBearer = in["drop_vld_gtp_v2_wrong_lbi_create_bearer"].(int)
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
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate2721(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate2721 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate2721
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
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
		ret.DropVldGtpV2WrongLbiCreateBearer = in["drop_vld_gtp_v2_wrong_lbi_create_bearer"].(int)
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
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsSeverity2722(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsSeverity2722 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsSeverity2722
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

func dataToEndpointVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmpl(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmpl {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmpl
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsInc2720(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate2721(d.Get("trigger_stats_rate").([]interface{}))
	ret.Inst.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsSeverity2722(d.Get("trigger_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
