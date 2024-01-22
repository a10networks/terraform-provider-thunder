provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_web_gui_protection_port_zone_service" "thunder_ddos_dst_zone_web_gui_protection_port_zone_service" {
  pbe       = "web_gui"
  port_num  = 23551
  protocol  = "dns-tcp"
  zone_name = "test"

}