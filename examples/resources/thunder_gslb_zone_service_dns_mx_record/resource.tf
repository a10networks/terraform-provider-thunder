provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_dns_mx_record" "thunder_gslb_zone_service_dns_mx_record" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  mx_name      = "2"
  priority     = 21510
  sampling_enable {
    counters1 = "all"
  }
  ttl = 0
}