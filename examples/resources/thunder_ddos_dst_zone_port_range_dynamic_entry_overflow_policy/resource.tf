provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_range_dynamic_entry_overflow_policy" "thunder_ddos_dst_zone_port_range_dynamic_entry_overflow_policy" {
  zone_name        = "testZone"
  port_range_start = 22
  protocol         = "dns-tcp"
  port_range_end   = 24
  action           = "bypass"
  dummy_name       = "configuration"
  log_enable       = 1
  log_periodic     = 1
  user_tag         = "45"

}