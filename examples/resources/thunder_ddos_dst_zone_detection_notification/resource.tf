provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_detection_notification" "thunder_ddos_dst_zone_detection_notification" {
  zone_name     = "testZone"
  configuration = "configuration"

}