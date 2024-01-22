provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_axdebug_capture" "thunder_axdebug_capture" {
  brief        = 0
  current_slot = 0
  detail       = 0
  no_stop      = 0
  save         = "54"
}