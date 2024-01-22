provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_network" "thunder_snmp_server_enable_traps_network" {
  trunk_port_threshold = 0
}