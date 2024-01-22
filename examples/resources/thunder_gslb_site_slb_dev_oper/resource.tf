provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_site_slb_dev_oper" "thunder_gslb_site_slb_dev_oper" {

  device_name = "25"
  site_name   = "3"
}
output "get_gslb_site_slb_dev_oper" {
  value = ["${data.thunder_gslb_site_slb_dev_oper.thunder_gslb_site_slb_dev_oper}"]
}