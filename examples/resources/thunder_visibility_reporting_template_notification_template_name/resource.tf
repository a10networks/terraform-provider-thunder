provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_reporting_template_notification_template_name" "thunder_visibility_reporting_template_notification_template_name" {
  name         = "test"
  protocol     = "http"
  relative_uri = "test.com"
  action       = "enable"
  debug_mode   = 1
  authentication {
    relative_login_uri   = "test.com"
    relative_logoff_uri  = "test.com"
    auth_username        = "test"
    auth_password        = 1
    auth_password_string = "test123"
  }
}