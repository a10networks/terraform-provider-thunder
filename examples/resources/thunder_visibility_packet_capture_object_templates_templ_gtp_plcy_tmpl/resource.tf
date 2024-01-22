provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_templ_gtp_plcy_tmpl" "thunder_visibility_packet_capture_object_templates_templ_gtp_plcy_tmpl" {

  name           = "test"
  capture_config = "test"
  trigger_stats_rate {
    threshold_exceeded_by                   = 5
    duration                                = 60
    drop_vld_gtp_ie_repeat_count_exceed     = 1
    drop_vld_reserved_field_set             = 1
    drop_vld_tunnel_id_flag                 = 1
    drop_vld_invalid_flow_label_v0          = 1
    drop_vld_invalid_teid                   = 1
    drop_vld_out_of_state                   = 1
    drop_vld_mandatory_information_element  = 1
    drop_vld_mandatory_ie_in_grouped_ie     = 1
    drop_vld_out_of_order_ie                = 1
    drop_vld_out_of_state_ie                = 1
    drop_vld_reserved_information_element   = 1
    drop_vld_version_not_supported          = 1
    drop_vld_message_length                 = 1
    drop_vld_cross_layer_correlation        = 1
    drop_vld_country_code_mismatch          = 1
    drop_vld_gtp_u_spoofed_source_address   = 1
    drop_vld_gtp_bearer_count_exceed        = 1
    drop_vld_gtp_v2_wrong_lbi_create_bearer = 1
    drop_vld_v0_reserved_message_drop       = 1
    drop_vld_v1_reserved_message_drop       = 1
    drop_vld_v2_reserved_message_drop       = 1
    drop_vld_invalid_pkt_len_piggyback      = 1
    drop_vld_sanity_failed_piggyback        = 1
    drop_vld_sequence_num_correlation       = 1
    drop_vld_gtpv0_seqnum_buffer_full       = 1
    drop_vld_gtpv1_seqnum_buffer_full       = 1
    drop_vld_gtpv2_seqnum_buffer_full       = 1
    drop_vld_gtp_invalid_imsi_len_drop      = 1
    drop_vld_gtp_invalid_apn_len_drop       = 1
    drop_vld_protocol_flag_unset            = 1
  }
  trigger_stats_severity {
    error          = 1
    error_alert    = 1
    error_warning  = 1
    error_critical = 1
    drop           = 1
    drop_alert     = 1
    drop_warning   = 1
    drop_critical  = 1
  }
  user_tag = "2"
}