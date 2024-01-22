package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_vpn_error`: Configure triggers for vpn.error object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bad_opcode": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_opcode",
						},
						"bad_sg_write_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_sg_write_len",
						},
						"bad_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_len",
						},
						"bad_ipsec_protocol": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_protocol",
						},
						"bad_ipsec_auth": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_auth",
						},
						"bad_ipsec_padding": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_padding",
						},
						"bad_ip_version": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ip_version",
						},
						"bad_auth_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_auth_type",
						},
						"bad_encrypt_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_encrypt_type",
						},
						"bad_ipsec_spi": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_spi",
						},
						"bad_checksum": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_checksum",
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
						"ipcomp_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipcomp_payload",
						},
						"bad_selector_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_selector_match",
						},
						"bad_fragment_size": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_fragment_size",
						},
						"bad_inline_data": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_inline_data",
						},
						"bad_frag_size_configuration": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_frag_size_configuration",
						},
						"dummy_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dummy_payload",
						},
						"bad_ip_payload_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ip_payload_type",
						},
						"bad_min_frag_size_auth_sha384_512": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_min_frag_size_auth_sha384_512",
						},
						"bad_esp_next_header": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_esp_next_header",
						},
						"bad_gre_header": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_gre_header",
						},
						"bad_gre_protocol": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_gre_protocol",
						},
						"ipv6_extension_headers_too_big": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_extension_headers_too_big",
						},
						"ipv6_hop_by_hop_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_hop_by_hop_error",
						},
						"error_ipv6_decrypt_rh_segs_left_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error_ipv6_decrypt_rh_segs_left_error",
						},
						"ipv6_rh_length_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_rh_length_error",
						},
						"ipv6_outbound_rh_copy_addr_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_outbound_rh_copy_addr_error",
						},
						"error_ipv6_extension_header_bad": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error_IPv6_extension_header_bad",
						},
						"bad_encrypt_type_ctr_gcm": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_encrypt_type_ctr_gcm",
						},
						"ah_not_supported_with_gcm_gmac_sha2": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ah_not_supported_with_gcm_gmac_sha2",
						},
						"tfc_padding_with_prefrag_not_supported": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tfc_padding_with_prefrag_not_supported",
						},
						"bad_srtp_auth_tag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_srtp_auth_tag",
						},
						"bad_ipcomp_configuration": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipcomp_configuration",
						},
						"dsiv_incorrect_param": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dsiv_incorrect_param",
						},
						"bad_ipsec_unknown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_unknown",
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
						"bad_opcode": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_opcode",
						},
						"bad_sg_write_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_sg_write_len",
						},
						"bad_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_len",
						},
						"bad_ipsec_protocol": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_protocol",
						},
						"bad_ipsec_auth": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_auth",
						},
						"bad_ipsec_padding": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_padding",
						},
						"bad_ip_version": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ip_version",
						},
						"bad_auth_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_auth_type",
						},
						"bad_encrypt_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_encrypt_type",
						},
						"bad_ipsec_spi": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_spi",
						},
						"bad_checksum": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_checksum",
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
						"ipcomp_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipcomp_payload",
						},
						"bad_selector_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_selector_match",
						},
						"bad_fragment_size": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_fragment_size",
						},
						"bad_inline_data": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_inline_data",
						},
						"bad_frag_size_configuration": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_frag_size_configuration",
						},
						"dummy_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dummy_payload",
						},
						"bad_ip_payload_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ip_payload_type",
						},
						"bad_min_frag_size_auth_sha384_512": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_min_frag_size_auth_sha384_512",
						},
						"bad_esp_next_header": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_esp_next_header",
						},
						"bad_gre_header": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_gre_header",
						},
						"bad_gre_protocol": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_gre_protocol",
						},
						"ipv6_extension_headers_too_big": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_extension_headers_too_big",
						},
						"ipv6_hop_by_hop_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_hop_by_hop_error",
						},
						"error_ipv6_decrypt_rh_segs_left_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error_ipv6_decrypt_rh_segs_left_error",
						},
						"ipv6_rh_length_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_rh_length_error",
						},
						"ipv6_outbound_rh_copy_addr_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_outbound_rh_copy_addr_error",
						},
						"error_ipv6_extension_header_bad": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error_IPv6_extension_header_bad",
						},
						"bad_encrypt_type_ctr_gcm": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_encrypt_type_ctr_gcm",
						},
						"ah_not_supported_with_gcm_gmac_sha2": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ah_not_supported_with_gcm_gmac_sha2",
						},
						"tfc_padding_with_prefrag_not_supported": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tfc_padding_with_prefrag_not_supported",
						},
						"bad_srtp_auth_tag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_srtp_auth_tag",
						},
						"bad_ipcomp_configuration": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipcomp_configuration",
						},
						"dsiv_incorrect_param": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dsiv_incorrect_param",
						},
						"bad_ipsec_unknown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_unknown",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2103(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2103 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2103
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bad_opcode = in["bad_opcode"].(int)
		ret.Bad_sg_write_len = in["bad_sg_write_len"].(int)
		ret.Bad_len = in["bad_len"].(int)
		ret.Bad_ipsec_protocol = in["bad_ipsec_protocol"].(int)
		ret.Bad_ipsec_auth = in["bad_ipsec_auth"].(int)
		ret.Bad_ipsec_padding = in["bad_ipsec_padding"].(int)
		ret.Bad_ip_version = in["bad_ip_version"].(int)
		ret.Bad_auth_type = in["bad_auth_type"].(int)
		ret.Bad_encrypt_type = in["bad_encrypt_type"].(int)
		ret.Bad_ipsec_spi = in["bad_ipsec_spi"].(int)
		ret.Bad_checksum = in["bad_checksum"].(int)
		ret.Bad_ipsec_context = in["bad_ipsec_context"].(int)
		ret.Bad_ipsec_context_direction = in["bad_ipsec_context_direction"].(int)
		ret.Bad_ipsec_context_flag_mismatch = in["bad_ipsec_context_flag_mismatch"].(int)
		ret.Ipcomp_payload = in["ipcomp_payload"].(int)
		ret.Bad_selector_match = in["bad_selector_match"].(int)
		ret.Bad_fragment_size = in["bad_fragment_size"].(int)
		ret.Bad_inline_data = in["bad_inline_data"].(int)
		ret.Bad_frag_size_configuration = in["bad_frag_size_configuration"].(int)
		ret.Dummy_payload = in["dummy_payload"].(int)
		ret.Bad_ip_payload_type = in["bad_ip_payload_type"].(int)
		ret.Bad_min_frag_size_auth_sha384_512 = in["bad_min_frag_size_auth_sha384_512"].(int)
		ret.Bad_esp_next_header = in["bad_esp_next_header"].(int)
		ret.Bad_gre_header = in["bad_gre_header"].(int)
		ret.Bad_gre_protocol = in["bad_gre_protocol"].(int)
		ret.Ipv6_extension_headers_too_big = in["ipv6_extension_headers_too_big"].(int)
		ret.Ipv6_hop_by_hop_error = in["ipv6_hop_by_hop_error"].(int)
		ret.Error_ipv6_decrypt_rh_segs_left_error = in["error_ipv6_decrypt_rh_segs_left_error"].(int)
		ret.Ipv6_rh_length_error = in["ipv6_rh_length_error"].(int)
		ret.Ipv6_outbound_rh_copy_addr_error = in["ipv6_outbound_rh_copy_addr_error"].(int)
		ret.Error_ipv6_extension_header_bad = in["error_ipv6_extension_header_bad"].(int)
		ret.Bad_encrypt_type_ctr_gcm = in["bad_encrypt_type_ctr_gcm"].(int)
		ret.Ah_not_supported_with_gcm_gmac_sha2 = in["ah_not_supported_with_gcm_gmac_sha2"].(int)
		ret.Tfc_padding_with_prefrag_not_supported = in["tfc_padding_with_prefrag_not_supported"].(int)
		ret.Bad_srtp_auth_tag = in["bad_srtp_auth_tag"].(int)
		ret.Bad_ipcomp_configuration = in["bad_ipcomp_configuration"].(int)
		ret.Dsiv_incorrect_param = in["dsiv_incorrect_param"].(int)
		ret.Bad_ipsec_unknown = in["bad_ipsec_unknown"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2104(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2104 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2104
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Bad_opcode = in["bad_opcode"].(int)
		ret.Bad_sg_write_len = in["bad_sg_write_len"].(int)
		ret.Bad_len = in["bad_len"].(int)
		ret.Bad_ipsec_protocol = in["bad_ipsec_protocol"].(int)
		ret.Bad_ipsec_auth = in["bad_ipsec_auth"].(int)
		ret.Bad_ipsec_padding = in["bad_ipsec_padding"].(int)
		ret.Bad_ip_version = in["bad_ip_version"].(int)
		ret.Bad_auth_type = in["bad_auth_type"].(int)
		ret.Bad_encrypt_type = in["bad_encrypt_type"].(int)
		ret.Bad_ipsec_spi = in["bad_ipsec_spi"].(int)
		ret.Bad_checksum = in["bad_checksum"].(int)
		ret.Bad_ipsec_context = in["bad_ipsec_context"].(int)
		ret.Bad_ipsec_context_direction = in["bad_ipsec_context_direction"].(int)
		ret.Bad_ipsec_context_flag_mismatch = in["bad_ipsec_context_flag_mismatch"].(int)
		ret.Ipcomp_payload = in["ipcomp_payload"].(int)
		ret.Bad_selector_match = in["bad_selector_match"].(int)
		ret.Bad_fragment_size = in["bad_fragment_size"].(int)
		ret.Bad_inline_data = in["bad_inline_data"].(int)
		ret.Bad_frag_size_configuration = in["bad_frag_size_configuration"].(int)
		ret.Dummy_payload = in["dummy_payload"].(int)
		ret.Bad_ip_payload_type = in["bad_ip_payload_type"].(int)
		ret.Bad_min_frag_size_auth_sha384_512 = in["bad_min_frag_size_auth_sha384_512"].(int)
		ret.Bad_esp_next_header = in["bad_esp_next_header"].(int)
		ret.Bad_gre_header = in["bad_gre_header"].(int)
		ret.Bad_gre_protocol = in["bad_gre_protocol"].(int)
		ret.Ipv6_extension_headers_too_big = in["ipv6_extension_headers_too_big"].(int)
		ret.Ipv6_hop_by_hop_error = in["ipv6_hop_by_hop_error"].(int)
		ret.Error_ipv6_decrypt_rh_segs_left_error = in["error_ipv6_decrypt_rh_segs_left_error"].(int)
		ret.Ipv6_rh_length_error = in["ipv6_rh_length_error"].(int)
		ret.Ipv6_outbound_rh_copy_addr_error = in["ipv6_outbound_rh_copy_addr_error"].(int)
		ret.Error_ipv6_extension_header_bad = in["error_ipv6_extension_header_bad"].(int)
		ret.Bad_encrypt_type_ctr_gcm = in["bad_encrypt_type_ctr_gcm"].(int)
		ret.Ah_not_supported_with_gcm_gmac_sha2 = in["ah_not_supported_with_gcm_gmac_sha2"].(int)
		ret.Tfc_padding_with_prefrag_not_supported = in["tfc_padding_with_prefrag_not_supported"].(int)
		ret.Bad_srtp_auth_tag = in["bad_srtp_auth_tag"].(int)
		ret.Bad_ipcomp_configuration = in["bad_ipcomp_configuration"].(int)
		ret.Dsiv_incorrect_param = in["dsiv_incorrect_param"].(int)
		ret.Bad_ipsec_unknown = in["bad_ipsec_unknown"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2103(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2104(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
