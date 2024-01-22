provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_reporting" "thunder_visibility_reporting" {
  session_logging = "disable"
  template {
    notification {
      template_name_list {
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
    }
  }
}