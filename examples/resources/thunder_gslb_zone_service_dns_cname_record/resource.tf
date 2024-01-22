provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_dns_cname_record" "thunder_gslb_zone_service_dns_cname_record" {

  name             = "a11"
  service_name     = "s25"
  service_port     = 33159
  admin_preference = 100
  alias_name       = "3"
  as_backup        = 0
  sampling_enable {
    counters1 = "all"
  }
  weight = 99
}