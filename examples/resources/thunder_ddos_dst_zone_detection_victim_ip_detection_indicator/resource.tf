provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_detection_victim_ip_detection_indicator" "thunder_ddos_dst_zone_detection_victim_ip_detection_indicator" {
  zone_name        = "testZone"
  ip_threshold_num = 1290129697
  type             = "pkt-rate"
  user_tag         = "51"

}