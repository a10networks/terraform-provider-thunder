package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_templ_gtp_plcy_tmpl_trigger_stats_inc`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"drop_vld_country_code_mismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Country Code Mismatch in IMSI and MSISDN",
			},
			"drop_vld_cross_layer_correlation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Cross Layer IP Address Mismatch",
			},
			"drop_vld_gtp_bearer_count_exceed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP Bearer count exceeded max (11)",
			},
			"drop_vld_gtp_ie_repeat_count_exceed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP repeated IE count exceeded",
			},
			"drop_vld_gtp_invalid_apn_len_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Invalid APN Length Drop",
			},
			"drop_vld_gtp_invalid_imsi_len_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Invalid IMSI Length Drop",
			},
			"drop_vld_gtp_u_spoofed_source_address": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-U IP Address Spoofed",
			},
			"drop_vld_gtp_v2_wrong_lbi_create_bearer": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV2-C Wrong LBI in Create Bearer Request",
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
			"drop_vld_invalid_flow_label_v0": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid Flow Label in GTPv0-C Header",
			},
			"drop_vld_invalid_pkt_len_piggyback": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Piggyback message invalid packet length",
			},
			"drop_vld_invalid_teid": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid TEID Value",
			},
			"drop_vld_mandatory_ie_in_grouped_ie": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Mandatory IE in Grouped IE Not Present",
			},
			"drop_vld_mandatory_information_element": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Mandatory IE Not Present",
			},
			"drop_vld_message_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Message Length Exceeded",
			},
			"drop_vld_out_of_order_ie": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv1-C Message Out of Order IE",
			},
			"drop_vld_out_of_state": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Out Of State GTP Message",
			},
			"drop_vld_out_of_state_ie": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Unexpected IE Present in Message",
			},
			"drop_vld_protocol_flag_unset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Protocol flag in Header Field not Set",
			},
			"drop_vld_reserved_field_set": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Reserved Header Field Set",
			},
			"drop_vld_reserved_information_element": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Reserved IE Field Present",
			},
			"drop_vld_sanity_failed_piggyback": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: piggyback message anomaly failed",
			},
			"drop_vld_sequence_num_correlation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Sequence number Mismatch",
			},
			"drop_vld_tunnel_id_flag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Tunnel Header Flag Not Set",
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
			"drop_vld_version_not_supported": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid GTP version",
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
func resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsInc
	ret.Inst.DropVldCountryCodeMismatch = d.Get("drop_vld_country_code_mismatch").(int)
	ret.Inst.DropVldCrossLayerCorrelation = d.Get("drop_vld_cross_layer_correlation").(int)
	ret.Inst.DropVldGtpBearerCountExceed = d.Get("drop_vld_gtp_bearer_count_exceed").(int)
	ret.Inst.DropVldGtpIeRepeatCountExceed = d.Get("drop_vld_gtp_ie_repeat_count_exceed").(int)
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
	ret.Inst.DropVldMandatoryIeInGroupedIe = d.Get("drop_vld_mandatory_ie_in_grouped_ie").(int)
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
	ret.Inst.DropVldV0ReservedMessageDrop = d.Get("drop_vld_v0_reserved_message_drop").(int)
	ret.Inst.DropVldV1ReservedMessageDrop = d.Get("drop_vld_v1_reserved_message_drop").(int)
	ret.Inst.DropVldV2ReservedMessageDrop = d.Get("drop_vld_v2_reserved_message_drop").(int)
	ret.Inst.DropVldVersionNotSupported = d.Get("drop_vld_version_not_supported").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
