provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_cpu" "thunder_debug_cpu" {
  dumy = 0
}