provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_ospf_inline" "thunder_vrrp_a_ospf_inline" {
  vlan = 1
}