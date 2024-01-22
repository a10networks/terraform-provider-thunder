provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_http" "thunder_health_monitor_method_http" {

  name               = "tf_test"
  http               = 1
  http_expect        = 1
  http_host          = "www.example.com"
  http_kerberos_auth = 1
  http_kerberos_kdc {
    http_kerberos_hostip = "1.1.1.1"
    http_kerberos_port   = 50248
  }
  http_kerberos_realm   = "example.com"
  http_maintenance_code = "800"
  http_password         = 1
  http_password_string  = "pd17"
  http_port             = 80
  http_response_code    = "121"
  http_url              = 1
  http_username         = "16"
  maintenance           = 1
  maintenance_text      = "health_main"
  response_code_regex   = "189"
  text_regex            = "33"
  url_path              = "/www.example.com"
  url_type              = "GET"
  version2              = 1
}