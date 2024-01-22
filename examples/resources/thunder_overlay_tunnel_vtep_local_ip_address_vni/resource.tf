provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_overlay_tunnel_vtep_local_ip_address_vni" "thunder_overlay_tunnel_vtep_local_ip_address_vni" {

  id1     = 56
  segment = 3
  lif     = "test"
}