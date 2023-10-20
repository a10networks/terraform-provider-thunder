provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_host_host_name" "resourceSnmpServerHostHostNameTest" {
  hostname    = "test.in"
  udp_port    = 180
  v1_v2c_comm = "incedo1"
  version1    = "v1"
}