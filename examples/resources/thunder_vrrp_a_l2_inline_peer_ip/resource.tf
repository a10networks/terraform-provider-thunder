provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_l2_inline_peer_ip" "thunder_vrrp_a_l2_inline_peer_ip" {
  ip_address = "10.10.10.10"
}