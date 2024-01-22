provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_lldp_management_address_dns" "thunder_network_lldp_management_address_dns" {
  dns = "20"
  interface {
    ethernet = 1
  }
}