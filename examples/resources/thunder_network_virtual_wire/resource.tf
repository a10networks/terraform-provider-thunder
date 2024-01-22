provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_virtual_wire" "thunder_network_virtual_wire" {
  user_tag        = "33"
  virtual_wire_id = 30
}