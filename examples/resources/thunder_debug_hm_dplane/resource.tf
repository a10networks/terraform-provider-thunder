provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_hm_dplane" "thunder_debug_hm_dplane" {
  dumy = 0
}