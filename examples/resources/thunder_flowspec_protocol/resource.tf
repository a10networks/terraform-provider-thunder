provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_flowspec_protocol" "thunder_flowspec_protocol" {

  name            = "test"
  proto_attribute = "eq"
  proto_num       = 21
}