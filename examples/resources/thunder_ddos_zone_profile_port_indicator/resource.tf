provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_profile_port_indicator" "thunder_ddos_zone_profile_port_indicator" {
  profile_name   = "test"
  indicator_name = "pkt-rate"
  src_threshold_cfg {
    src_threshold_num = 337249241
  }
  user_tag = "54"
  zone_threshold_cfg {
    zone_threshold_num = 1848492589
  }
  port_num      = 37352
  port_protocol = "dns-tcp"
}