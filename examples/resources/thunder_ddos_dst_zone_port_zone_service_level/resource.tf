provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_level" "thunder_ddos_dst_zone_port_zone_service_level" {
  zone_name               = "testZone"
  port_num                = 28712
  protocol                = "dns-tcp"
  apply_extracted_filters = 1
  indicator_list {
    type               = "pkt-rate"
    score              = 128
    src_threshold_num  = 717839624
    zone_threshold_num = 222
    user_tag           = "91"
  }
  level_num                 = "0"
  src_escalation_score      = 741
  start_pattern_recognition = 1
  user_tag                  = "8"
}