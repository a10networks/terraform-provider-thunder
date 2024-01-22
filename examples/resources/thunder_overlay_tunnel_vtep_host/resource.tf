provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_overlay_tunnel_vtep_host" "thunder_overlay_tunnel_vtep_host" {

  id1              = 56
  vni              = 3
  remote_vtep      = "20.20.20.20"
  ip_addr          = "10.10.10.10"
  overlay_mac_addr = "aabb.ccdd.eeff"
}