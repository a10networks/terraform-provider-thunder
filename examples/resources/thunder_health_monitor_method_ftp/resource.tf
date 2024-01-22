provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_ftp" "thunder_health_monitor_method_ftp" {

  name                = "tf_test"
  ftp                 = 1
  ftp_password        = 1
  ftp_password_string = "root24"
  ftp_port            = 210
  ftp_username        = "kp24"
}