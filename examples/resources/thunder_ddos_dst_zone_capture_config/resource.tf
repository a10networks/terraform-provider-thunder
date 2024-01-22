provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_capture_config" "thunder_ddos_dst_zone_capture_config" {
  zone_name = "testZone"
  name      = "10"
  mode      = "all"
}