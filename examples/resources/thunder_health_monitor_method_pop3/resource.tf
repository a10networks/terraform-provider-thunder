provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_pop3" "thunder_health_monitor_method_pop3" {

  name                 = "tf_test"
  pop3                 = 1
  pop3_password        = 1
  pop3_password_string = "a10net"
  pop3_port            = 11011
  pop3_username        = "admin1"
}