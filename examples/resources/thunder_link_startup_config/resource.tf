provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_link_startup_config" "thunder_link_startup_config" {
  all_partitions = 1
  default        = 1
  file_name      = "11"
  primary        = 1
  secondary      = 0
}