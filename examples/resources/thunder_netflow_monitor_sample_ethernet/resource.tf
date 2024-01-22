provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_netflow_monitor_sample_ethernet" "thunder_netflow_monitor_sample_ethernet" {

  name    = "a11"
  ifindex = 2
}