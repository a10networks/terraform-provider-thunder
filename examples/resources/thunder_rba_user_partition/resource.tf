provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_rba_user_partition" "thunder_rba_user_partition" {

  name           = "test"
  partition_name = "shared"
  user_tag       = "40"
}