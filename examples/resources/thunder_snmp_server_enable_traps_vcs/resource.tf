provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_vcs" "thunder_snmp_server_enable_traps_vcs" {
  state_change = 0
}