provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_geo_location" "thunder_gslb_zone_service_geo_location" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  action       = 1
  action_type  = "forward"
  forward_type = "both"
  geo_name     = "116"
  user_tag     = "79"
}