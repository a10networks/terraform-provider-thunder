provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_web_service_secure_regenerate" "thunder_web_service_secure_regenerate" {
  country     = "IN"
  domain_name = "test.dom"
  state       = "MH"
}