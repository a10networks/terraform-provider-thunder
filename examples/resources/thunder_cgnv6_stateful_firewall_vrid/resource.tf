provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_vrid" "thunder_cgnv6_stateful_firewall_vrid" {

  vrid_value = 3

}