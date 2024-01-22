provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_view" "thunder_snmp_server_view" {
  mask     = "29"
  oid      = "84"
  viewname = "test-view"
}