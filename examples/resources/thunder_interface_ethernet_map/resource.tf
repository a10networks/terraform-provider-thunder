provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_map" "thunder_interface_ethernet_map" {

  ifnum         = 2
  inside        = 1
  map_t_inside  = 1
  map_t_outside = 1
  outside       = 1
}