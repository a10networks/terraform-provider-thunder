provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_web_gui_protection_port_zone_service_other" "thunder_ddos_dst_zone_web_gui_protection_port_zone_service_other" {
  zone_name  = "testZone"
  pbe        = "91"
  port_other = "other"
  protocol   = "tcp"

}