provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_detection_victim_ip_detection" "thunder_ddos_dst_zone_detection_victim_ip_detection" {
  zone_name        = "testZone"
  configuration    = "configuration"
  histogram_toggle = "histogram-disable"
  indicator_list {
    type             = "pkt-rate"
    ip_threshold_num = 183807983
    user_tag         = "83"
  }
  toggle = "disable"
}