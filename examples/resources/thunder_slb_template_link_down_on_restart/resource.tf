provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_link_down_on_restart" "thunder_slb_template_link_down_on_restart" {
  enable = 0
}