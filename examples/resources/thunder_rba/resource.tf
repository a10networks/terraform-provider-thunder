provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_rba" "thunder_rba" {

  action = "enable"
  group_list {
    name = "test"
    user_list {
      user = "test"
    }
    user_tag = "10"
    partition_list {
      partition_name = "shared"
      user_tag       = "test"
    }
  }
}