provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_system" "thunder_gslb_system" {

  age_interval      = 10
  geo_location_iana = 1
  ip_ttl            = 0
  module            = 0
  ttl               = 350
  wait              = 10
}