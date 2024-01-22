provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_snmp" "thunder_snmp_server_enable_traps_snmp" {
  all      = 0
  linkdown = 0
  linkup   = 0
}