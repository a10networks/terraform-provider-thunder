provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_lldp_management_address_ipv4_addr" "thunder_network_lldp_management_address_ipv4_addr" {
  ipv4 = "10.10.10.10"
  interface_ipv4 {
    ipv4_eth = 1
  }
}