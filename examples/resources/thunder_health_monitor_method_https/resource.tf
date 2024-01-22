provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_https" "thunder_health_monitor_method_https" {

  name                = "tf_test"
  disable_sslv2hello  = 1
  http_version        = "http-version2"
  https               = 1
  https_expect        = 1
  https_host          = "www.example.com"
  https_kerberos_auth = 1
  https_kerberos_kdc {
    https_kerberos_hostip = "14.10.10.1"
    https_kerberos_port   = 10053
  }
  https_kerberos_realm   = "example.com"
  https_maintenance_code = "800"
  https_password         = 1
  https_password_string  = "a1031"
  https_response_code    = "200"
  https_url              = 1
  https_username         = "1"
  maintenance            = 1
  maintenance_text       = "75"
  response_code_regex    = "121"
  sni                    = 1
  text_regex             = "53"
  url_path               = "/www.example.com"
  url_type               = "GET"
  web_port               = 443
}