provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_lw_4o6" "thunder_interface_ethernet_lw_4o6" {
  ifnum   = 2
  inside  = 1
  outside = 1
}