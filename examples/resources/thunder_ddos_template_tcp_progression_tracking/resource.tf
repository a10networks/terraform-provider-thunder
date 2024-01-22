provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_tcp_progression_tracking" "thunder_ddos_template_tcp_progression_tracking" {
  name = "test"
  connection_tracking {
    progression_tracking_conn_enabled = "enable-check"
    conn_sent_max                     = 2635
    conn_sent_min                     = 334
    conn_rcvd_max                     = 27961
    conn_rcvd_min                     = 642
    conn_rcvd_sent_ratio_min          = 427
    conn_rcvd_sent_ratio_max          = 30371
    conn_duration_max                 = 594449
    conn_duration_min                 = 1063
    conn_violation                    = 6
    progression_tracking_conn_action  = "drop"
  }
  first_request_max_time           = 301
  profiling_connection_life_model  = 1
  profiling_request_response_model = 1
  profiling_time_window_model      = 1
  progression_tracking_action      = "drop"
  progression_tracking_enabled     = "enable-check"
  request_length_max               = 17682
  request_length_min               = 3607
  request_response_model           = "enable"
  request_to_response_max_time     = 4920
  time_window_tracking {
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
  violation = 197

}