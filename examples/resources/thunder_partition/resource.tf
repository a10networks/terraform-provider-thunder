provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_partition" "thunder_partition" {
  partition_name = "test"
  id1            = 4
  user_tag       = "test_partition"
}