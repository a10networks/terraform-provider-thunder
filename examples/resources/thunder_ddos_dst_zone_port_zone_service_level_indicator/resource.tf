provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_level_indicator" "thunder_ddos_dst_zone_port_zone_service_level_indicator" {
  zone_name          = "testZone"
  port_num           = 28712
  protocol           = "dns-tcp"
  level_num          = 0
  score              = 412
  src_threshold_num  = 36146637
  type               = "pkt-rate"
  user_tag           = "40"
  zone_threshold_num = 1034
}