provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_dns_cache" "thunder_debug_dns_cache" {
  level = 1
}