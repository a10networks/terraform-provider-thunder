provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_rba_group_partition" "thunder_rba_group_partition" {

  name           = "test"
  partition_name = "shared"
  role_list {
    role = "NetworkOperator"
  }
  user_tag = "57"
}