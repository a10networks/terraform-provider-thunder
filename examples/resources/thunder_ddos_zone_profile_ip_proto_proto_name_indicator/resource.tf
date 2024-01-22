provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_profile_ip_proto_proto_name_indicator" "thunder_ddos_zone_profile_ip_proto_proto_name_indicator" {
  profile_name   = "test"
  indicator_name = "pkt-rate"
  src_threshold_cfg {
    src_threshold_num = 1539838742
  }
  protocol = "icmp-v4"
  user_tag = "89"
  zone_threshold_cfg {
    zone_threshold_num = 1563939234
  }

}