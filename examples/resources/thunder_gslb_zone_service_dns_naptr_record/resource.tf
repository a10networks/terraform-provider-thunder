provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_dns_naptr_record" "thunder_gslb_zone_service_dns_naptr_record" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  flag         = "1"
  naptr_target = "29"
  order        = 48506
  preference   = 55779
  regexp       = 0
  sampling_enable {
    counters1 = "all"
  }
  service_proto = "40"
  ttl           = 0
}