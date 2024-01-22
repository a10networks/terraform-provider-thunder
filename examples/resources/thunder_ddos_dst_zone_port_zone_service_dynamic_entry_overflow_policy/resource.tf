provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_dynamic_entry_overflow_policy" "thunder_ddos_dst_zone_port_zone_service_dynamic_entry_overflow_policy" {
  zone_name    = "testZone"
  action       = "bypass"
  dummy_name   = "configuration"
  log_enable   = 1
  log_periodic = 1
  user_tag     = "104"
  port_num     = 28712
  protocol     = "dns-tcp"
}