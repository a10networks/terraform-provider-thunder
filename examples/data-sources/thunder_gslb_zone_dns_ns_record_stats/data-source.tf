provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_dns_ns_record_stats" "thunder_gslb_zone_dns_ns_record_stats" {

  name    = "a11"
  ns_name = "98"
}
output "get_gslb_zone_dns_ns_record_stats" {
  value = ["${data.thunder_gslb_zone_dns_ns_record_stats.thunder_gslb_zone_dns_ns_record_stats}"]
}