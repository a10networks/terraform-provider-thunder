provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_table_sync_event" "thunder_debug_table_sync_event" {
  all = 0
}