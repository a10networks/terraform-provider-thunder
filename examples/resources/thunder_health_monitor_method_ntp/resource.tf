provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_ntp" "thunder_health_monitor_method_ntp" {

  name     = "tf_test"
  ntp      = 1
  ntp_port = 12345
}