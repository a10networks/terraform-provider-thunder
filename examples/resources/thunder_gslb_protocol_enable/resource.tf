provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_protocol_enable" "thunder_gslb_protocol_enable" {
  type = "controller"
}