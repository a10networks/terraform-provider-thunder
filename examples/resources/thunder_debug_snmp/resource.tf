provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_snmp" "thunder_debug_snmp" {
  all   = 0
  error = 0
  event = 0
}