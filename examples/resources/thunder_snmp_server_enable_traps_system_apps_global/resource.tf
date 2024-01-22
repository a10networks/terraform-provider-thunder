provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_system_apps_global" "thunder_snmp_server_enable_traps_system_apps_global" {
  cps_threshold      = 0
  sessions_threshold = 0
}