provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_tacplus" "thunder_health_monitor_method_tacplus" {

  name                    = "tf_test"
  tacplus                 = 1
  tacplus_password        = 1
  tacplus_password_string = "a12"
  tacplus_port            = 491
  tacplus_secret          = 1
  tacplus_secret_string   = "6"
  tacplus_type            = "inbound-ascii-login"
  tacplus_username        = "13"
}