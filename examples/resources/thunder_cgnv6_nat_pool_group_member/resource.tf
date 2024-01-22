provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat_pool_group_member" "thunder_cgnv6_nat_pool_group_member" {

  pool_group_name = "test"
  pool_name       = "testPool"
}