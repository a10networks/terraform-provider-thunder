provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_notification_template" "thunder_ddos_notification_template" {
  name = "test"
  api {
    host_ipv4_address = "10.10.10.10"
    http_protocol     = "https"
    http_port         = 80
    timeout           = 10
    relative_uri      = "/"
    use_mgmt_port     = 1
    authentication {
      api_key        = 1
      api_key_string = "53"
    }
  }
  user_tag = "28"
}