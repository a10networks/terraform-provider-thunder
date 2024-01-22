provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_tcp_progression_tracking_connection_tracking" "thunder_ddos_template_tcp_progression_tracking_connection_tracking" {
  name                              = "test"
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