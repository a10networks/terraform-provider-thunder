provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_site_oper" "thunder_gslb_site_oper" {

  site_name = "3"
}
output "get_gslb_site_oper" {
  value = ["${data.thunder_gslb_site_oper.thunder_gslb_site_oper}"]
}