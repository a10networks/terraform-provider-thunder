provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_flowspec_destination_port" "thunder_flowspec_destination_port" {

  name           = "test"
  port_attribute = "eq"
  port_num       = 20
}