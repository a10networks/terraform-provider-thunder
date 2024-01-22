provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_manual_mode" "thunder_ddos_dst_zone_port_zone_service_manual_mode" {
  zone_name = "testZone"
  port_num  = 28712
  protocol  = "dns-tcp"
  config    = "configuration"
  user_tag  = "96"
}