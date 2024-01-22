provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_admin_lockout" "thunder_admin_lockout" {
  duration   = 10
  enable     = 1
  reset_time = 10
  threshold  = 5
}