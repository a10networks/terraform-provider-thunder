provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_tcp_progression_tracking_time_window_tracking" "thunder_ddos_zone_template_tcp_progression_tracking_time_window_tracking" {
  name                                = "test"
  progression_tracking_win_enabled    = "enable-check"
  window_sent_max                     = 20
  window_sent_min                     = 10
  window_rcvd_max                     = 20
  window_rcvd_min                     = 10
  window_rcvd_sent_ratio_min          = 10
  window_rcvd_sent_ratio_max          = 20
  window_violation                    = 2
  progression_tracking_windows_action = "drop"
}