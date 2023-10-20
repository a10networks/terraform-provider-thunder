provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_host_ipv6_host" "resourceSnmpServerHostIpv6HostTest" {
  ipv6_addr   = "2003::1"
  udp_port    = 179
  v1_v2c_comm = "testing"
  version1    = "v1"
}