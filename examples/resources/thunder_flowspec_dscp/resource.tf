provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_flowspec_dscp" "thunder_flowspec_dscp" {

  name           = "test"
  dscp_attribute = "eq"
  dscp_val       = 32
}