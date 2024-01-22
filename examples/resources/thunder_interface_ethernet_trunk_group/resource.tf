provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_trunk_group" "thunder_interface_ethernet_trunk_group" {
  ifnum        = 2
  trunk_number = 1872
  type         = "static"
  user_tag     = "43"
}