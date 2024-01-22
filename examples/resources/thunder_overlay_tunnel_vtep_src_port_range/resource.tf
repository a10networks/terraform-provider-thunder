provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_overlay_tunnel_vtep_src_port_range" "thunder_overlay_tunnel_vtep_src_port_range" {

  id1      = 56
  max_port = 25
  min_port = 23
}