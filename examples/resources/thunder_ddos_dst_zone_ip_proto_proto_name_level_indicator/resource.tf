provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_ip_proto_proto_name_level_indicator" "thunder_ddos_dst_zone_ip_proto_proto_name_level_indicator" {
  zone_name          = "testZone"
  protocol           = "icmp-v4"
  level_num          = "0"
  score              = 130432
  src_threshold_num  = 1636636884
  type               = "pkt-rate"
  user_tag           = "21"
  zone_threshold_num = 591299022

}