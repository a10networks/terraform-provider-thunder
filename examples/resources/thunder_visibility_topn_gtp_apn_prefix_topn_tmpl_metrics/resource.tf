provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_topn_gtp_apn_prefix_topn_tmpl_metrics" "thunder_visibility_topn_gtp_apn_prefix_topn_tmpl_metrics" {

  name                                    = "test"
  downlink_bytes                          = 1
  downlink_pkts                           = 1
  drop_flt_apn_filtering                  = 1
  drop_flt_gtp_in_gtp                     = 1
  drop_flt_message_filtering              = 1
  drop_flt_msisdn_filtering               = 1
  drop_flt_rat_type_filtering             = 1
  drop_rl_gtp_u_downlink_byte             = 1
  drop_rl_gtp_u_downlink_packet           = 1
  drop_rl_gtp_u_max_concurrent_tunnels    = 1
  drop_rl_gtp_u_total_byte                = 1
  drop_rl_gtp_u_total_packet              = 1
  drop_rl_gtp_u_tunnel_create             = 1
  drop_rl_gtp_u_uplink_byte               = 1
  drop_rl_gtp_u_uplink_packet             = 1
  drop_rl_gtp_v0_c_agg                    = 1
  drop_rl_gtp_v1_c_agg                    = 1
  drop_rl_gtp_v1_c_create_pdp_request     = 1
  drop_rl_gtp_v1_c_update_pdp_request     = 1
  drop_rl_gtp_v2_c_agg                    = 1
  drop_rl_gtp_v2_c_create_session_request = 1
  drop_rl_gtp_v2_c_modify_bearer_request  = 1
  drop_vld_country_code_mismatch          = 1
  drop_vld_cross_layer_correlation        = 1
  drop_vld_gtp_bearer_count_exceed        = 1
  drop_vld_gtp_invalid_apn_len_drop       = 1
  drop_vld_gtp_invalid_imsi_len_drop      = 1
  drop_vld_gtp_u_spoofed_source_address   = 1
  drop_vld_gtp_v2_wrong_lbi_create_bearer = 1
  drop_vld_gtpv0_seqnum_buffer_full       = 1
  drop_vld_gtpv1_seqnum_buffer_full       = 1
  drop_vld_gtpv2_seqnum_buffer_full       = 1
  drop_vld_invalid_flow_label_v0          = 1
  drop_vld_invalid_pkt_len_piggyback      = 1
  drop_vld_invalid_teid                   = 1
}