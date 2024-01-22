provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_tcp_progression_tracking_connection_tracking" "thunder_ddos_zone_template_tcp_progression_tracking_connection_tracking" {
  name                              = "test"
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