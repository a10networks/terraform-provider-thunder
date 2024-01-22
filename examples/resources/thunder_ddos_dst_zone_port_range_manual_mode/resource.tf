provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_range_manual_mode" "thunder_ddos_dst_zone_port_range_manual_mode" {
  zone_name                         = "testZone"
  port_range_start                  = 22
  port_range_end                    = 24
  protocol                          = "dns-tcp"
  close_sessions_for_unauth_sources = 1
  config                            = "configuration"
  user_tag                          = "25"
}