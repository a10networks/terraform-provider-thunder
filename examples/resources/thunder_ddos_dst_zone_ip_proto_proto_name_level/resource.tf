provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_ip_proto_proto_name_level" "thunder_ddos_dst_zone_ip_proto_proto_name_level" {
  zone_name = "testZone"
  protocol  = "icmp-v4"
  indicator_list {
    type              = "pkt-rate"
    score             = 831362
    src_threshold_num = 242826774
    user_tag          = "118"
  }
  level_num             = "0"
  src_escalation_score  = 445560
  user_tag              = "86"
  zone_escalation_score = 125932
}