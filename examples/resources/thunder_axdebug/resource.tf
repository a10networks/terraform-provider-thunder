provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_axdebug" "thunder_axdebug" {
  capture {
    brief        = 0
    detail       = 0
    save         = "test3"
    current_slot = 0
    no_stop      = 0
  }
  count1 = 3000
  exit {
    stop_capture = 0
  }
  save_config {
    config_file = "test3"
    default     = 0
  }
  sess_filter_dis = 0
  timeout         = 5
}