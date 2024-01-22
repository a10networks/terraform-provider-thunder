provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_location" "thunder_snmp_server_location" {
  loc = "230"
}