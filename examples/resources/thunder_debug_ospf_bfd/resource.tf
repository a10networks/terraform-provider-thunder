provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ospf_bfd" "thunder_debug_ospf_bfd" {
  dumy = 0
}