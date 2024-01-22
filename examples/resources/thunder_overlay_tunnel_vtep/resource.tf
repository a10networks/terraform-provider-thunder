provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_overlay_tunnel_vtep" "thunder_overlay_tunnel_vtep" {
  id1   = 56
  encap = "nvgre"
  local_ip_address {
    ip_address = "10.2.43.8"
  }
}