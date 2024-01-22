provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_detection_packet_anomaly_detection_indicator" "thunder_ddos_dst_zone_detection_packet_anomaly_detection_indicator" {
  zone_name     = "testZone"
  threshold_num = 100
  type          = "port-zero-pkt-rate"
  user_tag      = "114"

}