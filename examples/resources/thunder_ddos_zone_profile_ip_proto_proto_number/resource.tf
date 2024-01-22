provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_profile_ip_proto_proto_number" "thunder_ddos_zone_profile_ip_proto_proto_number" {

  indicator_list {
    indicator_name = "pkt-rate"
    src_threshold_cfg {
      src_threshold_num = 2039028515
    }
    zone_threshold_cfg {
      zone_threshold_num = 293088462
    }
    user_tag = "122"
  }
  protocol_num = 15
  profile_name = "test"
}