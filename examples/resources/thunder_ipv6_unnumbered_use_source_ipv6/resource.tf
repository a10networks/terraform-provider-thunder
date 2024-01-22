provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_unnumbered_use_source_ipv6" "thunder_ipv6_unnumbered_use_source_ipv6" {
  update_source_ipv6 = "2001:db8:3333:4444:5555:6666:7777:8888"
}