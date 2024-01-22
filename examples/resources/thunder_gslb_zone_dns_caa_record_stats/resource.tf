provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_dns_caa_record_stats" "thunder_gslb_zone_dns_caa_record_stats" {

  name          = "a11"
  critical_flag = 229
  property_tag  = "171"
  rdata         = "464"
}
output "get_gslb_zone_dns_caa_record_stats" {
  value = ["${data.thunder_gslb_zone_dns_caa_record_stats.thunder_gslb_zone_dns_caa_record_stats}"]
}