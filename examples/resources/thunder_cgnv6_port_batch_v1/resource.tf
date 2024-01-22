provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_port_batch_v1" "thunder_cgnv6_port_batch_v1" {
  enable_port_batch_v1 = 1
}