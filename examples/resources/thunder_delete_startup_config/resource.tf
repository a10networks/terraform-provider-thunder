provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_delete_startup_config" "thunder_delete_startup_config" {
  file_name = "test"
}