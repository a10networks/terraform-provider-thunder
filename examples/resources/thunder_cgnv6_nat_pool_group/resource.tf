provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat_pool_group" "thunder_cgnv6_nat_pool_group" {

  member_list {
    pool_name = "testPool"
  }
  pool_group_name = "test"
  user_tag        = "test"
  vrid            = 0
}