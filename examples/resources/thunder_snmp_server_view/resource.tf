provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_view" "resourceSnmpServerViewTest" {
  type     = "excluded"
  oid      = "123"
  viewname = "a10"
}
