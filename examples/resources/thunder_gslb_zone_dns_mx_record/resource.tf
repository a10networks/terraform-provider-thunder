provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_dns_mx_record" "thunder_gslb_zone_dns_mx_record" {

  name     = "a11"
  mx_name  = "69"
  priority = 62684
  sampling_enable {
    counters1 = "all"
  }
  ttl = 903257279
}