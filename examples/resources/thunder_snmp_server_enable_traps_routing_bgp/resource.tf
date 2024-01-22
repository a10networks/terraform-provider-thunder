provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_routing_bgp" "thunder_snmp_server_enable_traps_routing_bgp" {
  ax {
    bgpestablishednotification             = 0
    bgpbackwardtransnotification           = 0
    bgpprefixthresholdexceedednotification = 0
    bgpprefixthresholdclearnotification    = 0
  }
  bgpbackwardtransnotification = 0
  bgpestablishednotification   = 0
}