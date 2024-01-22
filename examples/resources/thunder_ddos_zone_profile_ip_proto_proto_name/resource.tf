provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_profile_ip_proto_proto_name" "thunder_ddos_zone_profile_ip_proto_proto_name" {
  profile_name = "test"
  indicator_list {
    indicator_name = "pkt-rate"
    src_threshold_cfg {
      src_threshold_num = 129
    }
    zone_threshold_cfg {
      zone_threshold_num = 45
    }
    user_tag = "79"
  }
  protocol = "icmp-v4"

}