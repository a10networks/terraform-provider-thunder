provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_global" "IpNatGlobal" {
  reset_idle_tcp_conn = 1
}