provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_vpn_error_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_vpn_error_trigger_stats_inc" {

  name                                   = "test"
  ah_not_supported_with_gcm_gmac_sha2    = 1
  bad_auth_type                          = 1
  bad_checksum                           = 1
  bad_encrypt_type                       = 1
  bad_encrypt_type_ctr_gcm               = 1
  bad_esp_next_header                    = 1
  bad_frag_size_configuration            = 1
  bad_fragment_size                      = 1
  bad_gre_header                         = 1
  bad_gre_protocol                       = 1
  bad_inline_data                        = 1
  bad_ip_payload_type                    = 1
  bad_ip_version                         = 1
  bad_ipcomp_configuration               = 1
  bad_ipsec_auth                         = 1
  bad_ipsec_context                      = 1
  bad_ipsec_context_direction            = 1
  bad_ipsec_context_flag_mismatch        = 1
  bad_ipsec_padding                      = 1
  bad_ipsec_protocol                     = 1
  bad_ipsec_spi                          = 1
  bad_ipsec_unknown                      = 1
  bad_len                                = 1
  bad_min_frag_size_auth_sha384_512      = 1
  bad_opcode                             = 1
  bad_selector_match                     = 1
  bad_sg_write_len                       = 1
  bad_srtp_auth_tag                      = 1
  dsiv_incorrect_param                   = 1
  dummy_payload                          = 1
  error_ipv6_decrypt_rh_segs_left_error  = 1
  error_ipv6_extension_header_bad        = 1
  ipcomp_payload                         = 1
  ipv6_extension_headers_too_big         = 1
  ipv6_hop_by_hop_error                  = 1
  ipv6_outbound_rh_copy_addr_error       = 1
  ipv6_rh_length_error                   = 1
  tfc_padding_with_prefrag_not_supported = 1
}