provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_multi_queue_support" "thunder_system_multi_queue_support" {
}