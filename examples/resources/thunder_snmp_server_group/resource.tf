provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_group" "resourceSnmpServerGroupTest" {
  read      = "a10"
  groupname = "a"
  v3        = "auth"
}