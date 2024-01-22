provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_dns_caa_record" "thunder_gslb_zone_service_dns_caa_record" {

  name          = "a11"
  service_name  = "s25"
  service_port  = 33159
  critical_flag = 35
  property_tag  = "165"
  rdata         = "200"
  sampling_enable {
    counters1 = "all"
  }
  ttl = 3473
}