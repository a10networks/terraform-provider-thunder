provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_other_level_indicator" "thunder_ddos_dst_zone_port_zone_service_other_level_indicator" {
  zone_name          = "testZone"
  port_other         = "other"
  protocol           = "tcp"
  level_num          = 0
  score              = 921
  src_threshold_num  = 169
  type               = "pkt-rate"
  user_tag           = "60"
  zone_threshold_num = 1302518977
}