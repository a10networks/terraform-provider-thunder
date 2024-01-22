provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ddos" "thunder_debug_ddos" {
  control_var          = 1
  dns_cache            = 1
  event                = 1
  event_filter         = "44"
  flow_based_detection = 1
  level                = 1
  zbar                 = 1
}