provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_sip" "thunder_health_monitor_method_sip" {

  name         = "tf_test"
  register     = 1
  sip          = 1
  sip_hostname = "156"
  sip_port     = 50601
  sip_tcp      = 1
}