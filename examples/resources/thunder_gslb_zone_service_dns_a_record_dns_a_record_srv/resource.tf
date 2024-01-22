provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_dns_a_record_dns_a_record_srv" "thunder_gslb_zone_service_dns_a_record_dns_a_record_srv" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  admin_ip     = 168
  as_backup    = 1
  as_replace   = 1
  disable      = 1
  no_resp      = 1
  static       = 1
  svrname      = "test-server1"
  ttl          = 32633
  weight       = 44
}