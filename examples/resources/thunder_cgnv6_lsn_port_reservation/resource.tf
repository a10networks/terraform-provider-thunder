provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_port_reservation" "thunder_cgnv6_lsn_port_reservation" {
  inside            = "192.168.1.3"
  inside_port_start = 22
  inside_port_end   = 33
  nat               = "192.168.1.10"
  nat_port_start    = 22
  nat_port_end      = 33
}