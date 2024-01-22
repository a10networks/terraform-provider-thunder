provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_notification_template_api" "thunder_ddos_notification_template_api" {
  name = "test"
  authentication {
    api_key        = 1
    api_key_string = "15"
  }
  host_ipv4_address = "10.10.10.10"
  http_protocol     = "https"
  https_port        = 443
  relative_uri      = "/"
  timeout           = 10
  use_mgmt_port     = 1

}