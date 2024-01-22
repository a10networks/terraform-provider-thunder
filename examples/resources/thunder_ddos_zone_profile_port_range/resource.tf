provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_profile_port_range" "thunder_ddos_zone_profile_port_range" {
  profile_name = "test"
  indicator_list {
    indicator_name = "pkt-rate"
    src_threshold_cfg {
      src_threshold_num = 151421017
    }
    zone_threshold_cfg {
      zone_threshold_num = 2037026289
    }
    user_tag = "31"
  }
  port_range_end   = 25
  port_range_start = 20
  protocol         = "dns-tcp"
  user_tag         = "60"

}