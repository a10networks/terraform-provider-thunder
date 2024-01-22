provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_dns_ns_record" "thunder_gslb_zone_dns_ns_record" {

  name    = "a11"
  ns_name = "98"
  sampling_enable {
    counters1 = "all"
  }
  ttl = 695136566
}