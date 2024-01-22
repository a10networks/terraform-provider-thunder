provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_lldp_management_address_ipv6_addr" "thunder_network_lldp_management_address_ipv6_addr" {
  ipv6 = "f5fd:1fcc:28e2:a236:a282:a957:fc64:32c7"
  interface_ipv6 {
    ipv6_eth = 1
  }
}