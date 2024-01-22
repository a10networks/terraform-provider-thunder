provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat_exclude_port_tcp" "thunder_cgnv6_nat_exclude_port_tcp" {
  port_list {
    port     = 1234
    port_end = 1237
  }
}