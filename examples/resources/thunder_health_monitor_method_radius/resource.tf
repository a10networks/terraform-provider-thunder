provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_radius" "thunder_health_monitor_method_radius" {

  name                   = "tf_test"
  radius                 = 1
  radius_expect          = 1
  radius_password_string = "16"
  radius_port            = 18121
  radius_response_code   = "6"
  radius_secret          = "31"
  radius_username        = "5"
}