provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_active_rdt" "thunder_gslb_active_rdt" {
  icmp     = 0
  interval = 2
  port     = 123
  retry    = 4
  sleep    = 4
  timeout  = 3010
  track    = 70
}