provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_detection" "thunder_ddos_detection" {
  disable = 1
  settings {
    detector_mode             = "standalone"
    detection_window_size     = 2
    initial_learning_interval = 23
    notification_debug_log    = "enable"
    de_escalation_quiet_time  = 22
  }
}