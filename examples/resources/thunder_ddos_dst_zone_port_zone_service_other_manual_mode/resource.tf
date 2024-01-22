provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_other_manual_mode" "thunder_ddos_dst_zone_port_zone_service_other_manual_mode" {
  zone_name  = "testZone"
  port_other = "other"
  protocol   = "tcp"
  config     = "configuration"
  user_tag   = "12"
}