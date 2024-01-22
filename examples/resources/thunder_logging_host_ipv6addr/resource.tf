provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_host_ipv6addr" "v6_1" {
  host_ipv6     = "fd01::11:2"
  use_mgmt_port = 1
  port          = 6551
  tcp           = 1
}