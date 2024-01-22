provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv4" "thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv4" {

  name            = "a11"
  service_name    = "s25"
  service_port    = 33159
  as_backup       = 1
  as_replace      = 0
  disable         = 0
  dns_a_record_ip = "10.10.10.10"
  no_resp         = 0
  static          = 1
  admin_ip        = 1
  ttl             = 35284
  weight          = 65
}