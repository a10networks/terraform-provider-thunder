provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_tcp_progression_tracking" "thunder_ddos_zone_template_tcp_progression_tracking" {
  name = "test"
  connection_tracking {
    progression_tracking_conn_enabled = "enable-check"
    conn_sent_max                     = 20
    conn_sent_min                     = 10
    conn_rcvd_max                     = 20
    conn_rcvd_min                     = 10
    conn_rcvd_sent_ratio_min          = 15
    conn_rcvd_sent_ratio_max          = 30
    conn_duration_max                 = 44
    conn_duration_min                 = 22
    conn_violation                    = 5
    progression_tracking_conn_action  = "drop"
  }
  progression_tracking_enabled     = "enable-check"
  request_response_model           = "enable"
  violation                        = 33
  ignore_tls_handshake             = 1
  response_length_max              = 20
  request_length_min               = 20
  request_length_max               = 40
  response_request_min_ratio       = 20
  response_request_max_ratio       = 40
  first_request_max_time           = 3
  request_to_response_max_time     = 30
  response_to_request_max_time     = 3
  profiling_request_response_model = 1
  profiling_connection_life_model  = 1
  progression_tracking_action      = "drop"
  time_window_tracking {
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
}