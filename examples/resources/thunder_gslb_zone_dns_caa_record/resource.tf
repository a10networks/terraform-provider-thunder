provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_dns_caa_record" "thunder_gslb_zone_dns_caa_record" {


  name          = "a11"
  critical_flag = 229
  property_tag  = "171"
  rdata         = "464"
  sampling_enable {
    counters1 = "all"
  }
  ttl = 0
}