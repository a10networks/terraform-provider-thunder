provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_detection_service_discovery" "thunder_ddos_dst_zone_detection_service_discovery" {
  zone_name          = "testZone"
  configuration      = "configuration"
  pkt_rate_threshold = 10
  toggle             = "disable"
}