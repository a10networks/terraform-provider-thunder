provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_range_level_indicator" "thunder_ddos_dst_zone_port_range_level_indicator" {
  zone_name          = "testZone"
  port_range_start   = 22
  port_range_end     = 24
  protocol           = "dns-tcp"
  level_num          = "2"
  score              = 126845
  src_threshold_num  = 970247381
  type               = "pkt-rate"
  user_tag           = "75"
  zone_threshold_num = 405414918

}