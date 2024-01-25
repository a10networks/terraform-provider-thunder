provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_dns_caa_record_oper" "thunder_gslb_zone_dns_caa_record_oper" {

  name          = "a11"
  critical_flag = 229
  property_tag  = "171"
  rdata         = "464"
}
output "get_gslb_zone_dns_caa_record_oper" {
  value = ["${data.thunder_gslb_zone_dns_caa_record_oper.thunder_gslb_zone_dns_caa_record_oper}"]
}