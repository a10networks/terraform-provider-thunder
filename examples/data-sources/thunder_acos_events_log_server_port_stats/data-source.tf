provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_acos_events_log_server_port_stats" "thunder_acos_events_log_server_port_stats" {

  name        = "test"
  port_number = 123
  protocol    = "tcp"

}
output "get_acos_events_log_server_port_stats" {
  value = ["${data.thunder_acos_events_log_server_port_stats.thunder_acos_events_log_server_port_stats}"]
}