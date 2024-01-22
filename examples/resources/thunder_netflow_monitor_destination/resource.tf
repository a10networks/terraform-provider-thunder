provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_netflow_monitor_destination" "thunder_netflow_monitor_destination" {

  name = "a11"
  ip_cfg {
    ip    = "10.0.2.7"
    port4 = 9996
  }
}