provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_one_to_one_pool_group" "thunder_cgnv6_one_to_one_pool_group" {

  member_list {
    pool_name = "test"
  }
  pool_group_name = "testing"
  user_tag        = "13"
  vrid            = 0
}