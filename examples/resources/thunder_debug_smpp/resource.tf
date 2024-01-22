provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_smpp" "thunder_debug_smpp" {
  dumy = 0
}