provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_backup_system" "thunder_backup_system" {
  encrypt       = 0
  store_name    = "14"
  use_mgmt_port = 1
}