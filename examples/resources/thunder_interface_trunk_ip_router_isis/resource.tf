provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_trunk_ip_router_isis" "thunder_interface_trunk_ip_router_isis" {

  ifnum = 11
  tag   = "126"
}