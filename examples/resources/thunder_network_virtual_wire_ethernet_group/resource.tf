provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_virtual_wire_ethernet_group" "thunder_network_virtual_wire_ethernet_group" {
  group_id = 1
  user_tag = "118"
}