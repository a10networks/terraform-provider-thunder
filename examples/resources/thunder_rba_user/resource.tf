provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_rba_user" "thunder_rba_user" {
  name = "test"
  partition_list {
    partition_name = "shared"
    user_tag       = "test"
  }
  user_tag = "test"
}