provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_host_ipv4_host" "thunder_snmp_server_host_ipv4_host" {
  ipv4_addr   = "10.10.10.10"
  udp_port    = 162
  v1_v2c_comm = "20"
  version     = "v1"
}