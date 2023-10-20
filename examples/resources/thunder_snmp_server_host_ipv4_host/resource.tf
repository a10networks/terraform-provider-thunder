provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_host_ipv4_host" "resourceSnmpServerHostIpv4HostTest" {
  ipv4_addr   = "192.168.20.3"
  udp_port    = 65535
  v1_v2c_comm = "testing"
  version1    = "v1"
}