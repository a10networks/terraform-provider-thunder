provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_routing_bgp_ax" "thunder_snmp_server_enable_traps_routing_bgp_ax" {
  bgpbackwardtransnotification           = 0
  bgpestablishednotification             = 0
  bgpprefixthresholdclearnotification    = 0
  bgpprefixthresholdexceedednotification = 0
}