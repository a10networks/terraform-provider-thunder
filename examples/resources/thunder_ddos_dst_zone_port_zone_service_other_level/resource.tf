provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_other_level" "thunder_ddos_dst_zone_port_zone_service_other_level" {
  zone_name  = "testZone"
  port_other = "other"
  protocol   = "tcp"
  indicator_list {
    type               = "pkt-rate"
    score              = 206
    src_threshold_num  = 180577924
    zone_threshold_num = 2016321448
    user_tag           = "59"
  }
  level_num                 = "0"
  src_escalation_score      = 642
  start_pattern_recognition = 1
  user_tag                  = "57"
  zone_escalation_score     = 137
}