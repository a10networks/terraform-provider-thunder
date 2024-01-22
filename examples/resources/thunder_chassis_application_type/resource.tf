provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_chassis_application_type" "thunder_chassis_application_type" {
  type = "cgnv6"
}