provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_ip_proto_proto_number_level" "thunder_ddos_dst_zone_ip_proto_proto_number_level" {
  zone_name    = "testZone"
  protocol_num = 223
  indicator_list {
    type               = "pkt-rate"
    score              = 49845
    src_threshold_num  = 1220097657
    zone_threshold_num = 1420327810
    user_tag           = "101"
  }
  level_num             = "0"
  src_escalation_score  = 605190
  user_tag              = "42"
  zone_escalation_score = 897139

}