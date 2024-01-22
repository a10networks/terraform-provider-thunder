provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_topn_gtp_apn_prefix_topn_tmpl" "thunder_visibility_topn_gtp_apn_prefix_topn_tmpl" {
  interval = "5"
  metrics {
    uplink_bytes                     = 1
    downlink_bytes                   = 1
    uplink_pkts                      = 1
    downlink_pkts                    = 1
    gtp_v0_c_tunnel_created          = 1
    gtp_v0_c_tunnel_half_open        = 1
    gtp_v0_c_tunnel_half_closed      = 1
    gtp_v0_c_tunnel_closed           = 1
    gtp_v0_c_tunnel_deleted          = 1
    gtp_v0_c_half_open_tunnel_closed = 1
    gtp_v1_c_tunnel_created          = 1
    gtp_v1_c_tunnel_half_open        = 1
    gtp_v1_c_tunnel_half_closed      = 1
  }
  name     = "test"
  user_tag = "68"
}