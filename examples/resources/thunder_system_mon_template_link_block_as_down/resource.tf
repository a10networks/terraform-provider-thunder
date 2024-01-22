provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_mon_template_link_block_as_down" "thunder_system_mon_template_link_block_as_down" {
}