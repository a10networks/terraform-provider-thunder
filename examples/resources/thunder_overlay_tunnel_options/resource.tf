provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_overlay_tunnel_options" "thunder_overlay_tunnel_options" {

  fragmentation_mode_inner = 1
  ip_dscp_preserve         = 1
  nvgre_disable_flow_id    = 1
  nvgre_key_mode_lower24   = 1
  src_port_range {
    min_port = 10
    max_port = 200
  }
  tcp_mss_adjust_disable = 0
  vxlan_dest_port        = 15390
}