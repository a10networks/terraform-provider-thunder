provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ipv6_ospf_lsa" "thunder_debug_ipv6_ospf_lsa" {
  flooding = 0
  gererate = 0
  install  = 0
  maxage   = 0
  refresh  = 0
}