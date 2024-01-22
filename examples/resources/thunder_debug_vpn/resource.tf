provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_vpn" "thunder_debug_vpn" {
  ike_gateway = "25"
  level       = 4
  strict      = 0
}