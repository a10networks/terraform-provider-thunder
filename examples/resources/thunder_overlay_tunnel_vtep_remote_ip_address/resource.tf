provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_overlay_tunnel_vtep_remote_ip_address" "thunder_overlay_tunnel_vtep_remote_ip_address" {

  id1        = 56
  ip_address = "10.10.10.10"
  user_tag   = "test"
}