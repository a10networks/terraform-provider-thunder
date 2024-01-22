provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_ip_proto_proto_number_level_indicator" "thunder_ddos_dst_zone_ip_proto_proto_number_level_indicator" {
  zone_name          = "testZone"
  protocol_num       = 223
  level_num          = "0"
  score              = 961380
  src_threshold_num  = 1092226694
  type               = "pkt-rate"
  user_tag           = "105"
  zone_threshold_num = 1409247422

}