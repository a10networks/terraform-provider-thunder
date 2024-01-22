provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_management_index" "thunder_snmp_server_management_index" {
  mgmt_index = 1178168705
}