provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_ssl_error" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_ssl_error" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by                 = 5
    duration                              = 60
    app_data_in_handshake                 = 1
    attempt_to_reuse_sess_in_diff_context = 1
    bad_alert_record                      = 1
    bad_authentication_type               = 1
    bad_change_cipher_spec                = 1
    bad_checksum                          = 1
    bad_data_returned_by_callback         = 1
    bad_decompression                     = 1
    bad_dh_g_length                       = 1
    bad_dh_pub_key_length                 = 1
    bad_dh_p_length                       = 1
    bad_digest_length                     = 1
    bad_dsa_signature                     = 1
    bad_hello_request                     = 1
    bad_length                            = 1
    bad_mac_decode                        = 1
    bad_message_type                      = 1
    bad_packet_length                     = 1
    bad_protocol_version_counter          = 1
    bad_response_argument                 = 1
    bad_rsa_decrypt                       = 1
    bad_rsa_encrypt                       = 1
    bad_rsa_e_length                      = 1
    bad_rsa_modulus_length                = 1
    bad_rsa_signature                     = 1
    bad_signature                         = 1
    bad_ssl_filetype                      = 1
    bad_ssl_session_id_length             = 1
    bad_state                             = 1
    bad_write_retry                       = 1
    bio_not_set                           = 1
    block_cipher_pad_is_wrong             = 1
    bn_lib                                = 1
    ca_dn_length_mismatch                 = 1
    ca_dn_too_long                        = 1
    ccs_received_early                    = 1
    certificate_verify_failed             = 1
    cert_length_mismatch                  = 1
    challenge_is_different                = 1
    cipher_code_wrong_length              = 1
    cipher_or_hash_unavailable            = 1
    cipher_table_src_error                = 1
  }
}