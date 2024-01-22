provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_one_to_one_shared_pool_group" "thunder_cgnv6_one_to_one_shared_pool_group" {
  members {
  }
}