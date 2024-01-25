provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_service_dns_naptr_record_stats" "thunder_gslb_zone_service_dns_naptr_record_stats" {

  name          = "a11"
  service_name  = "s25"
  service_port  = 33159
  service_proto = "40"
  flag          = "1"
  naptr_target  = "29"
}
output "get_gslb_zone_service_dns_naptr_record_stats" {
  value = ["${data.thunder_gslb_zone_service_dns_naptr_record_stats.thunder_gslb_zone_service_dns_naptr_record_stats}"]
}