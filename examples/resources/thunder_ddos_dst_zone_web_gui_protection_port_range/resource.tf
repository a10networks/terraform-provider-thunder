provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_web_gui_protection_port_range" "thunder_ddos_dst_zone_web_gui_protection_port_range" {
  zone_name        = "testZone"
  pbe              = "85"
  port_range_end   = 52143
  port_range_start = 4820
  protocol         = "dns-tcp"
  user_tag         = "49"

}