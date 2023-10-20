provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_management_index" "resourceSnmpServerManagementIndexTest" {
  mgmt_index = 99
}