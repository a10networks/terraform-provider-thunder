provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_netflow_common" "thunder_netflow_common" {
  max_packet_queue_time     = 21
  reset_time_on_flow_record = 1
}