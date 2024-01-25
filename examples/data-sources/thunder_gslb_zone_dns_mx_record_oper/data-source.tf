provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_dns_mx_record_oper" "thunder_gslb_zone_dns_mx_record_oper" {

  name    = "a11"
  mx_name = "69"
}
output "get_gslb_zone_dns_mx_record_oper" {
  value = ["${data.thunder_gslb_zone_dns_mx_record_oper.thunder_gslb_zone_dns_mx_record_oper}"]
}