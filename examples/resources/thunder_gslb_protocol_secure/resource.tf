provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_protocol_secure" "thunder_gslb_protocol_secure" {
  action = "enable"
}