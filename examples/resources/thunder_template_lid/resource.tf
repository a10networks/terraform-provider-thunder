provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_lid" "thunder_template_lid" {
  ddos_protection_factor = 3
  downlink_pps           = 1330848596
  downlink_throughput    = 2954
  lid_number             = 917
  limit_cps              = 855734971
  limit_rate             = "limit-pps"
  respond_to_user_mac    = 0
  src_ip {
    concurrent_sessions           = 1171659635
    prefix_length                 = 105
    burstsize_downlink_pps        = 1431948678
    burstsize_uplink_pps          = 2048055075
    burstsize_total_pps           = 1674508940
    burstsize_downlink_throughput = 35553
    burstsize_uplink_throughput   = 39457
    burstsize_total_throughput    = 71307
    burstsize_cps                 = 1857364811
  }
  total_pps         = 611328894
  total_throughput  = 12025
  uplink_pps        = 50647451
  uplink_throughput = 36917
  user_tag          = "24"
}