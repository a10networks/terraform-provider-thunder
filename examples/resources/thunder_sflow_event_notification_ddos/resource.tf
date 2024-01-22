provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_event_notification_ddos" "thunder_sflow_event_notification_ddos" {
  toggle = "disable"
}