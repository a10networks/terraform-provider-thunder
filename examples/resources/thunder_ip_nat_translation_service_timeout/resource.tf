provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_translation_service_timeout" "thunder_ip_nat_translation_service_timeout" {
  port         = 37642
  service_type = "tcp"
  timeout_type = "age"
  timeout_val  = 13303
}