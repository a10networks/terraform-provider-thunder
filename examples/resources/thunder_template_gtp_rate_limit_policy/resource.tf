provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_rate_limit_policy" "thunder_template_gtp_rate_limit_policy" {
  name                           = "test_temp"
  gtp_u_downlink_byte_rate       = 124470908
  gtp_u_downlink_packet_rate     = 693212351
  gtp_u_max_concurrent_tunnels   = 741612703
  gtp_u_total_byte_rate          = 1866963973
  gtp_u_total_packet_rate        = 130289302
  gtp_u_tunnel_create_rate       = 672907037
  gtp_u_uplink_byte_rate         = 1246050005
  gtp_u_uplink_packet_rate       = 1015968531
  lockout                        = 151
  rate_limit_action              = "drop"
  user_tag                       = "59"
  v0_agg_message_type_rate       = 1253314264
  v1_agg_message_type_rate       = 1341412669
  v1_create_pdp_request_rate     = 33670369
  v1_update_pdp_request_rate     = 1690475847
  v2_agg_message_type_rate       = 963565887
  v2_create_session_request_rate = 398134586
  v2_modify_bearer_request_rate  = 1763872288
}