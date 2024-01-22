provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_chassis_infra" "thunder_chassis_infra" {
  debug_disable      = 0
  debug_enable       = 0
  debug_status       = 0
  detailed           = 0
  sys_sync           = 0
  system_sync_verify = 0
}