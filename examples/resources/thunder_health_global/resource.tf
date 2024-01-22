provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_global" "thunder_health_global" {
  check_rate          = 2
  disable_auto_adjust = 1
  interval            = 6
  retry               = 2
}