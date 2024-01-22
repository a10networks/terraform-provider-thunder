provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_flowspec_operational_mode" "thunder_flowspec_operational_mode" {

  name = "test"
  mode = "enabled"
}