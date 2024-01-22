provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_tcp_progression_tracking_time_window_tracking" "thunder_ddos_template_tcp_progression_tracking_time_window_tracking" {
  name                                = "test"
  progression_tracking_win_enabled    = "enable-check"
  window_sent_max                     = 21795
  window_sent_min                     = 485
  window_rcvd_max                     = 53441
  window_rcvd_min                     = 5219
  window_rcvd_sent_ratio_min          = 624
  window_rcvd_sent_ratio_max          = 42396
  window_violation                    = 233
  progression_tracking_windows_action = "drop"

}