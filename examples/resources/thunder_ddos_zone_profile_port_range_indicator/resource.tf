provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_profile_port_range_indicator" "thunder_ddos_zone_profile_port_range_indicator" {
  profile_name   = "test"
  indicator_name = "pkt-rate"
  src_threshold_cfg {
    src_threshold_num = 215509240
  }
  user_tag = "33"
  zone_threshold_cfg {
    zone_threshold_num = 1713351456
  }
  port_range_end   = 25
  port_range_start = 20
  protocol         = "dns-tcp"
}