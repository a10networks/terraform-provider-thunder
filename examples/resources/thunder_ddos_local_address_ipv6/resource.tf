provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_local_address_ipv6" "thunder_ddos_local_address_ipv6" {
  ipv6_addr = "2001:db8:3333:4444:5555:6666:7777:8888"
}