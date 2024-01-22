provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_dns_rrl" "thunder_debug_dns_rrl" {
  dumy = 0
}