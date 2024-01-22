provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_service_dns_caa_record_stats" "thunder_gslb_zone_service_dns_caa_record_stats" {

  name          = "a11"
  service_name  = "s25"
  service_port  = 33159
  critical_flag = 35
  property_tag  = "165"
  rdata         = "200"

}
output "get_gslb_zone_service_dns_caa_record_stats" {
  value = ["${data.thunder_gslb_zone_service_dns_caa_record_stats.thunder_gslb_zone_service_dns_caa_record_stats}"]
}