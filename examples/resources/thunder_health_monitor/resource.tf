provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_health_monitor" "common" {
  method {
    radius {
      radius                 = 1
      radius_expect          = 1
      radius_password_string = "radius"
      radius_port            = 345
      radius_response_code   = 13
      radius_secret          = "secret"
      radius_username        = "radius"
    }
    http {
      http                  = 1
      http_port             = 100
      http_expect           = 1
      http_response_code    = "200"
      http_host             = "www.example.com"
      http_maintenance_code = "800"
      http_url              = 1
      url_type              = "GET"
      maintenance           = 1
      maintenance_text      = "health_main"
      url_path              = "/www.example.com"
      http_username         = "admin"
      http_password         = 1
      http_password_string  = "password"
      http_kerberos_auth    = 1
      http_kerberos_realm   = "example.com"
      http_kerberos_kdc {
        http_kerberos_hostip = "1.1.1.1"
        http_kerberos_port   = 65535
      }
    }
  }
  ssl_version = 33
  name        = "tf_test"
}
