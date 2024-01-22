provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_range_level" "thunder_ddos_dst_zone_port_range_level" {
  zone_name                         = "testZone"
  port_range_start                  = 22
  port_range_end                    = 24
  protocol                          = "dns-tcp"
  apply_extracted_filters           = 1
  close_sessions_for_unauth_sources = 1
  indicator_list {
    type               = "pkt-rate"
    score              = 107878
    src_threshold_num  = 202787013
    zone_threshold_num = 464740406
    user_tag           = "125"
  }
  level_num                 = "2"
  src_escalation_score      = 259190
  start_pattern_recognition = 1
  user_tag                  = "11"
  zone_escalation_score     = 923683
}