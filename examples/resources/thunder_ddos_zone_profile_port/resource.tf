provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_profile_port" "thunder_ddos_zone_profile_port" {
  profile_name = "test"
  indicator_list {
    indicator_name = "pkt-rate"
    src_threshold_cfg {
      src_threshold_num = 2106520166
    }
    zone_threshold_cfg {
      zone_threshold_num = 1957526267
    }
    user_tag = "67"
  }
  port_num      = 37352
  port_protocol = "dns-tcp"
  user_tag      = "26"

}