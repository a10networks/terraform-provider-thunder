provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_hm" "thunder_debug_hm" {
  level       = 2
  method_type = "icmp"
  pin_uid     = 667918429
}