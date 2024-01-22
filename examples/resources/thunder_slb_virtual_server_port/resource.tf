provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_virtual_server_port" "thunder_slb_virtual_server_port" {

  name              = "test"
  protocol          = "tcp"
  port_number       = 343
  stats_data_action = "stats-data-disable"
  extended_stats    = 1
}