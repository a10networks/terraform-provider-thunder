provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_udp" "thunder_health_monitor_method_udp" {

  name                             = "tf_test"
  udp                              = 1
  force_up_with_single_healthcheck = 1
  udp_port                         = 61932
}