provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vcs_reload" "thunder_vcs_reload" {
  disable_merge = 1
}