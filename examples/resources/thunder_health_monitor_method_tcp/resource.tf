provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_tcp" "thunder_health_monitor_method_tcp" {

  name             = "tf_test"
  maintenance      = 1
  maintenance_text = "36"
  method_tcp       = 1
  port_send        = 521
  port_resp {
    port_contains = "124"
  }
  tcp_port = 9851
}