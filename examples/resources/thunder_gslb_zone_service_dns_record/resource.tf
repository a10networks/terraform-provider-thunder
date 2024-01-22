provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_dns_record" "thunder_gslb_zone_service_dns_record" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  data         = "442"
  type         = 26622
}