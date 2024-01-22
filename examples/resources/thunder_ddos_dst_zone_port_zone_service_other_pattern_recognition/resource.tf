provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_other_pattern_recognition" "thunder_ddos_dst_zone_port_zone_service_other_pattern_recognition" {
  zone_name       = "test"
  port_other      = "other"
  protocol        = "tcp"
  algorithm       = "heuristic"
  triggered_by    = "zone-escalation"
  capture_traffic = "all"

}