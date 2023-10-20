provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_interface_ethernet_trunk_group" "trunkgroup" {
  ifnum         = 5
  mode          = "active"
  user_tag      = "trunk"
  port_priority = 100
  admin_key     = 20000
  type          = "lacp"
  timeout       = "long"
  trunk_number  = 4
}