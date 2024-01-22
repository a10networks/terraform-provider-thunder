provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_rba_group" "thunder_rba_group" {
  name = "test"
  partition_list {
    partition_name = "shared"
    role_list {
      role = "NetworkOperator"
    }
  }
  user_tag = "108"
}