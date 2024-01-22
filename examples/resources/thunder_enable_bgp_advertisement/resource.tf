provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_enable_bgp_advertisement" "thunder_enable_bgp_advertisement" {
  enable_bgp_advertisement = 0
}