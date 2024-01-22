provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_web_service_secure" "thunder_web_service_secure" {
  restart = 1
  wipe    = 1
  generate {
    domain_name = "test.com"
    country     = "IN"
    state       = "MH"
  }
  regenerate {
    domain_name = "test.com"
    country     = "IN"
    state       = "MH"
  }
}