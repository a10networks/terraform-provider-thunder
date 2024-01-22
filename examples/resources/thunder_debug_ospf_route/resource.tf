provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ospf_route" "thunder_debug_ospf_route" {
  ase     = 0
  ia      = 0
  install = 0
  spf     = 0
}