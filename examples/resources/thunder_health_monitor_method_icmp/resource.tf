provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_icmp" "thunder_health_monitor_method_icmp" {

  name        = "tf_test"
  icmp        = 0
  transparent = 0
}