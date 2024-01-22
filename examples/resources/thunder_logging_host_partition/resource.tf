provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_host_partition" "thunder_logging_host_partition" {
}