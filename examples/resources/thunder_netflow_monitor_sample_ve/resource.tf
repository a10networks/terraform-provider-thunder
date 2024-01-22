provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_netflow_monitor_sample_ve" "thunder_netflow_monitor_sample_ve" {
  ve_num = 2196
}