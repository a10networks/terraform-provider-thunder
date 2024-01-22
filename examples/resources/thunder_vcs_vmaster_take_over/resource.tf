provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vcs_vmaster_take_over" "thunder_vcs_vmaster_take_over" {
  vmaster_take_over = 213
}