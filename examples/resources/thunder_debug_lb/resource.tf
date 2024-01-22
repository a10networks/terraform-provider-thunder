provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_lb" "thunder_debug_lb" {
  all  = 0
  cfg  = 0
  clb  = 0
  flow = 0
  fwlb = 0
  llb  = 0
  slb  = 0
}