provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_detection_packet_anomaly_detection" "thunder_ddos_dst_zone_detection_packet_anomaly_detection" {
  zone_name     = "testZone"
  configuration = "configuration"
  indicator_list {
    type          = "port-zero-pkt-rate"
    threshold_num = 100
    user_tag      = "42"
  }
  toggle = "enable"

}