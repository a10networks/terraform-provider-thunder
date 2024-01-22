provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_l3_inline_mode_flag" "thunder_vrrp_a_l3_inline_mode_flag" {
  l3_inline_mode = 1
}