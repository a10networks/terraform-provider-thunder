provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_profile_ip_proto_proto_number_indicator" "thunder_ddos_zone_profile_ip_proto_proto_number_indicator" {
  protocol_num   = 15
  profile_name   = "test"
  indicator_name = "pkt-rate"
  src_threshold_cfg {
    src_threshold_num = 412681828
  }
  user_tag = "52"
  zone_threshold_cfg {
    zone_threshold_num = 546314870
  }
}