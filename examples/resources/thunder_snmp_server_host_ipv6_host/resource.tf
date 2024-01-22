provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_host_ipv6_host" "thunder_snmp_server_host_ipv6_host" {
  ipv6_addr   = "48e5:808a:8fe9:17eb:16d6:203b:4fa1:67e2"
  udp_port    = 162
  v1_v2c_comm = "18"
  version     = "v1"
}