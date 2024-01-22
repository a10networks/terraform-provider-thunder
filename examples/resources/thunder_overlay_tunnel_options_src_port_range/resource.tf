provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_overlay_tunnel_options_src_port_range" "thunder_overlay_tunnel_options_src_port_range" {
  min_port = 10
  max_port = 200
}