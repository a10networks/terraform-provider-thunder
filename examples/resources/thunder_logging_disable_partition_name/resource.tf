provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_disable_partition_name" "thunder_logging_disable_partition_name" {
  disable_partition_name = 0
}