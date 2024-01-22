provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_ssl" "thunder_snmp_server_enable_traps_ssl" {
  server_certificate_error = 0
}