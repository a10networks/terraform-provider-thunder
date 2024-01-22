provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_allow_slb_cfg" "thunder_allow_slb_cfg" {
  enable = 0
}