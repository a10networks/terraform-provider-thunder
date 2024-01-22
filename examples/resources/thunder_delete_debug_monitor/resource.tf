provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_delete_debug_monitor" "thunder_delete_debug_monitor" {
  file_name = "27"
}