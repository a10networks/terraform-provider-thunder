provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_service_dns_a_record_dns_a_record_srv_stats" "thunder_gslb_zone_service_dns_a_record_dns_a_record_srv_stats" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  svrname      = "test-server1"
}
output "get_gslb_zone_service_dns_a_record_dns_a_record_srv_stats" {
  value = ["${data.thunder_gslb_zone_service_dns_a_record_dns_a_record_srv_stats.thunder_gslb_zone_service_dns_a_record_dns_a_record_srv_stats}"]
}