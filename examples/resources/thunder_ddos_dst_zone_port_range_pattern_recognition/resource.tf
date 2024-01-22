provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_range_pattern_recognition" "thunder_ddos_dst_zone_port_range_pattern_recognition" {
  zone_name          = "testZone"
  port_range_start   = 22
  port_range_end     = 24
  protocol           = "dns-tcp"
  algorithm          = "heuristic"
  app_payload_offset = 1107
  capture_traffic    = "all"
  filter_threshold   = 63
  mode               = "capture-never-expire"
  sensitivity        = "high"
  triggered_by       = "zone-escalation"

}