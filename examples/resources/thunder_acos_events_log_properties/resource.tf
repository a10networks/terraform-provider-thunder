provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_log_properties" "thunder_acos_events_log_properties" {
  add_msgid_in_header  = 0
  enable_8k_tcp_syslog = 0
}