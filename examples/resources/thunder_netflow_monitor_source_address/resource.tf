provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_netflow_monitor_source_address" "thunder_netflow_monitor_source_address" {

  name = "a11"
  ip   = "10.0.2.5"
}