provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_overlay_tunnel_vtep_local_ip_address" "thunder_overlay_tunnel_vtep_local_ip_address" {

  id1        = 56
  ip_address = "10.10.101.10"
  vni_list {
    segment = 3
    lif     = "test"
  }
}