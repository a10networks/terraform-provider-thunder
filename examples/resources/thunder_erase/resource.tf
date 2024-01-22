provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_erase" "thunder_erase" {
  all_partitions      = 0
  grubconfig          = 0
  preserve_accounts   = 0
  preserve_management = 0
  reload              = 0
  service_config      = 0
}