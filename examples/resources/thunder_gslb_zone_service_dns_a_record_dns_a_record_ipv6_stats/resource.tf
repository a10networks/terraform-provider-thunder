provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv6_stats" "thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv6_stats" {

  name              = "a11"
  service_name      = "s25"
  service_port      = 33159
  dns_a_record_ipv6 = "2001:db8:3333:4444:5555:6666:7777:8888"
}
output "get_gslb_zone_service_dns_a_record_dns_a_record_ipv6_stats" {
  value = ["${data.thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv6_stats.thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv6_stats}"]
}