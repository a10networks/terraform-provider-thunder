provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_nat_pool_group" "thunder_ipv6_nat_pool_group" {

  member_list {
    pool_name = "11"
  }
  pool_group_name = "35"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "8"
}