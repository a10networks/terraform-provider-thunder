provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_SNMPv3_user" "thunder_snmp_server_SNMPv3_user" {

  group    = "test-group"
  username = "test-user"
  v3       = "noauth"

}