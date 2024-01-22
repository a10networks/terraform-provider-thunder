provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_slb_change" "thunder_snmp_server_enable_traps_slb_change" {
  all                       = 0
  connection_resource_event = 0
  resource_usage_warning    = 0
  server_port               = 0
  ssl_cert_change           = 0
  ssl_cert_expire           = 0
  system_threshold          = 0
  vip                       = 0
  vip_port                  = 0
}