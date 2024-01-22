provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_dns_ns_record_oper" "thunder_gslb_zone_dns_ns_record_oper" {

  name    = "a11"
  ns_name = "98"
}
output "get_gslb_zone_dns_ns_record_oper" {
  value = ["${data.thunder_gslb_zone_dns_ns_record_oper.thunder_gslb_zone_dns_ns_record_oper}"]
}