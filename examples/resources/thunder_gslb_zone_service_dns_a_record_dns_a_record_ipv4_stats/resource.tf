provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv4_stats" "thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv4_stats" {

  name            = "a11"
  service_name    = "s25"
  service_port    = 33159
  dns_a_record_ip = "10.10.10.10"

}
output "get_gslb_zone_service_dns_a_record_dns_a_record_ipv4_stats" {
  value = ["${data.thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv4_stats.thunder_gslb_zone_service_dns_a_record_dns_a_record_ipv4_stats}"]
}