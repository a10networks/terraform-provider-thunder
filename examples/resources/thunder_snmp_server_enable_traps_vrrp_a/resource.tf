provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_vrrp_a" "thunder_snmp_server_enable_traps_vrrp_a" {
  active  = 0
  all     = 0
  standby = 0
}