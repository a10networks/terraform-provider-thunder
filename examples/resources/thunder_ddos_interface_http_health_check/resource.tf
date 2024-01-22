provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_interface_http_health_check" "thunder_ddos_interface_http_health_check" {
  challenge_method        = "http-redirect"
  challenge_redirect_code = "302"
  enable                  = "enable"
}