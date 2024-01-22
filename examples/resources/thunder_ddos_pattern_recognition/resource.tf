provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_pattern_recognition" "thunder_ddos_pattern_recognition" {

  sflow_event_periodic_interval = 6
}