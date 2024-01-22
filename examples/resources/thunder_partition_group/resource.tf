provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_partition_group" "thunder_partition_group" {

  member_list {
    member = "test"
  }
  partition_group_name = "test"
  user_tag             = "test_part_group"
}