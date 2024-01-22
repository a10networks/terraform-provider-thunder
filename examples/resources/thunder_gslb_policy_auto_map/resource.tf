provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_policy_auto_map" "thunder_gslb_policy_auto_map" {
  module_disable = 0
  all            = 0
  ttl            = 300
}