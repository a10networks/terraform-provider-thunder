provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_trunk_map" "thunder_interface_trunk_map" {

  ifnum         = 11
  inside        = 0
  map_t_inside  = 1
  map_t_outside = 0
  outside       = 0
}