provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_overlay_tunnel_vtep_remote_ip_address_vni" "thunder_overlay_tunnel_vtep_remote_ip_address_vni" {

  id1        = 56
  ip_address = "30.20.20.20"
  segment    = 3
}