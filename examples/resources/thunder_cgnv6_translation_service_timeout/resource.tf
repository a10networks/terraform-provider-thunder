provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_translation_service_timeout" "thunder_cgnv6_translation_service_timeout" {
  port         = 12970
  port_end     = 4026
  service_type = "tcp"
  timeout_val  = 4068
}