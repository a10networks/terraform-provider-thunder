provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_detection_outbound_detection_indicator" "thunder_ddos_dst_zone_detection_outbound_detection_indicator" {
  threshold_num = 470644108
  type          = "pkt-rate"
  user_tag      = "86"
  zone_name     = "test"

}