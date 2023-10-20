provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_contact" "Snmp_Server_Contact_Test" {
  contact_name = "test_server"
}