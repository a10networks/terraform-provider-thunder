provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_pool_group" "thunder_ip_nat_pool_group" {

  member_list {
    pool_name = "a24"
  }
  pool_group_name = "53"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "116"
}