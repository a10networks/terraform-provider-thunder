provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_list_all_cli" "thunder_debug_list_all_cli" {
  dumy = 0
}