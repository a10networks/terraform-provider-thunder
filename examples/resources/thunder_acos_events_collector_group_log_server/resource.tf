provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_collector_group_log_server" "thunder_acos_events_collector_group_log_server" {

  name = "test"
  port = 346

}