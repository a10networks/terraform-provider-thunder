provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_gslb" "thunder_snmp_server_enable_traps_gslb" {
  all        = 0
  group      = 0
  service_ip = 0
  site       = 0
  zone       = 0
}