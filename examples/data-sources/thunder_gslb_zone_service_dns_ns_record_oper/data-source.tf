provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_service_dns_ns_record_oper" "thunder_gslb_zone_service_dns_ns_record_oper" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  ns_name      = "4"
}
output "get_gslb_zone_service_dns_ns_record_oper" {
  value = ["${data.thunder_gslb_zone_service_dns_ns_record_oper.thunder_gslb_zone_service_dns_ns_record_oper}"]
}