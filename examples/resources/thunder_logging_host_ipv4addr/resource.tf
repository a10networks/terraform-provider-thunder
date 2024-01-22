provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_host_ipv4addr" "v4_1" {
  host_ipv4     = "1.1.1.1"
  use_mgmt_port = 1
  tcp           = 1
  port          = 3456
}