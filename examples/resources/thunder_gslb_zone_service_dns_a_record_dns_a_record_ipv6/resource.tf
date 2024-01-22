provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv6" "thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv6" {

  name              = "a11"
  service_name      = "s25"
  service_port      = 33159
  admin_ip          = 234
  as_backup         = 1
  as_replace        = 0
  disable           = 0
  dns_a_record_ipv6 = "2001:db8:3333:4444:5555:6666:7777:8888"
  no_resp           = 0
  static            = 1
  ttl               = 35173
  weight            = 95
}