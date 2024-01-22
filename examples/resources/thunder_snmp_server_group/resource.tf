provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_group" "thunder_snmp_server_group" {

  groupname = "test-group"
  read      = "test-view"
  v3        = "noauth"
}