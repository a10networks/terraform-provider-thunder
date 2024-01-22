package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_vpn_error_trigger_stats_rate`: Configure stats to trigger packet capture on increment rate\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateDelete,

		Schema: map[string]*schema.Schema{
			"ah_not_supported_with_gcm_gmac_sha2": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ah_not_supported_with_gcm_gmac_sha2",
			},
			"bad_auth_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_auth_type",
			},
			"bad_checksum": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_checksum",
			},
			"bad_encrypt_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_encrypt_type",
			},
			"bad_encrypt_type_ctr_gcm": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_encrypt_type_ctr_gcm",
			},
			"bad_esp_next_header": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_esp_next_header",
			},
			"bad_frag_size_configuration": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_frag_size_configuration",
			},
			"bad_fragment_size": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_fragment_size",
			},
			"bad_gre_header": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_gre_header",
			},
			"bad_gre_protocol": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_gre_protocol",
			},
			"bad_inline_data": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_inline_data",
			},
			"bad_ip_payload_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ip_payload_type",
			},
			"bad_ip_version": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ip_version",
			},
			"bad_ipcomp_configuration": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipcomp_configuration",
			},
			"bad_ipsec_auth": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_auth",
			},
			"bad_ipsec_context": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_context",
			},
			"bad_ipsec_context_direction": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_context_direction",
			},
			"bad_ipsec_context_flag_mismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_context_flag_mismatch",
			},
			"bad_ipsec_padding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_padding",
			},
			"bad_ipsec_protocol": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_protocol",
			},
			"bad_ipsec_spi": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_spi",
			},
			"bad_ipsec_unknown": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_unknown",
			},
			"bad_len": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_len",
			},
			"bad_min_frag_size_auth_sha384_512": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_min_frag_size_auth_sha384_512",
			},
			"bad_opcode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_opcode",
			},
			"bad_selector_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_selector_match",
			},
			"bad_sg_write_len": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_sg_write_len",
			},
			"bad_srtp_auth_tag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_srtp_auth_tag",
			},
			"dsiv_incorrect_param": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dsiv_incorrect_param",
			},
			"dummy_payload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dummy_payload",
			},
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
			},
			"error_ipv6_decrypt_rh_segs_left_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error_ipv6_decrypt_rh_segs_left_error",
			},
			"error_ipv6_extension_header_bad": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error_IPv6_extension_header_bad",
			},
			"ipcomp_payload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipcomp_payload",
			},
			"ipv6_extension_headers_too_big": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_extension_headers_too_big",
			},
			"ipv6_hop_by_hop_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_hop_by_hop_error",
			},
			"ipv6_outbound_rh_copy_addr_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_outbound_rh_copy_addr_error",
			},
			"ipv6_rh_length_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_rh_length_error",
			},
			"tfc_padding_with_prefrag_not_supported": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tfc_padding_with_prefrag_not_supported",
			},
			"threshold_exceeded_by": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate
	ret.Inst.Ah_not_supported_with_gcm_gmac_sha2 = d.Get("ah_not_supported_with_gcm_gmac_sha2").(int)
	ret.Inst.Bad_auth_type = d.Get("bad_auth_type").(int)
	ret.Inst.Bad_checksum = d.Get("bad_checksum").(int)
	ret.Inst.Bad_encrypt_type = d.Get("bad_encrypt_type").(int)
	ret.Inst.Bad_encrypt_type_ctr_gcm = d.Get("bad_encrypt_type_ctr_gcm").(int)
	ret.Inst.Bad_esp_next_header = d.Get("bad_esp_next_header").(int)
	ret.Inst.Bad_frag_size_configuration = d.Get("bad_frag_size_configuration").(int)
	ret.Inst.Bad_fragment_size = d.Get("bad_fragment_size").(int)
	ret.Inst.Bad_gre_header = d.Get("bad_gre_header").(int)
	ret.Inst.Bad_gre_protocol = d.Get("bad_gre_protocol").(int)
	ret.Inst.Bad_inline_data = d.Get("bad_inline_data").(int)
	ret.Inst.Bad_ip_payload_type = d.Get("bad_ip_payload_type").(int)
	ret.Inst.Bad_ip_version = d.Get("bad_ip_version").(int)
	ret.Inst.Bad_ipcomp_configuration = d.Get("bad_ipcomp_configuration").(int)
	ret.Inst.Bad_ipsec_auth = d.Get("bad_ipsec_auth").(int)
	ret.Inst.Bad_ipsec_context = d.Get("bad_ipsec_context").(int)
	ret.Inst.Bad_ipsec_context_direction = d.Get("bad_ipsec_context_direction").(int)
	ret.Inst.Bad_ipsec_context_flag_mismatch = d.Get("bad_ipsec_context_flag_mismatch").(int)
	ret.Inst.Bad_ipsec_padding = d.Get("bad_ipsec_padding").(int)
	ret.Inst.Bad_ipsec_protocol = d.Get("bad_ipsec_protocol").(int)
	ret.Inst.Bad_ipsec_spi = d.Get("bad_ipsec_spi").(int)
	ret.Inst.Bad_ipsec_unknown = d.Get("bad_ipsec_unknown").(int)
	ret.Inst.Bad_len = d.Get("bad_len").(int)
	ret.Inst.Bad_min_frag_size_auth_sha384_512 = d.Get("bad_min_frag_size_auth_sha384_512").(int)
	ret.Inst.Bad_opcode = d.Get("bad_opcode").(int)
	ret.Inst.Bad_selector_match = d.Get("bad_selector_match").(int)
	ret.Inst.Bad_sg_write_len = d.Get("bad_sg_write_len").(int)
	ret.Inst.Bad_srtp_auth_tag = d.Get("bad_srtp_auth_tag").(int)
	ret.Inst.Dsiv_incorrect_param = d.Get("dsiv_incorrect_param").(int)
	ret.Inst.Dummy_payload = d.Get("dummy_payload").(int)
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.Error_ipv6_decrypt_rh_segs_left_error = d.Get("error_ipv6_decrypt_rh_segs_left_error").(int)
	ret.Inst.Error_ipv6_extension_header_bad = d.Get("error_ipv6_extension_header_bad").(int)
	ret.Inst.Ipcomp_payload = d.Get("ipcomp_payload").(int)
	ret.Inst.Ipv6_extension_headers_too_big = d.Get("ipv6_extension_headers_too_big").(int)
	ret.Inst.Ipv6_hop_by_hop_error = d.Get("ipv6_hop_by_hop_error").(int)
	ret.Inst.Ipv6_outbound_rh_copy_addr_error = d.Get("ipv6_outbound_rh_copy_addr_error").(int)
	ret.Inst.Ipv6_rh_length_error = d.Get("ipv6_rh_length_error").(int)
	ret.Inst.Tfc_padding_with_prefrag_not_supported = d.Get("tfc_padding_with_prefrag_not_supported").(int)
	ret.Inst.ThresholdExceededBy = d.Get("threshold_exceeded_by").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
