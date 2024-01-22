provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_oper" "thunder_gslb_zone_oper" {

  name = "a11"
}
output "get_gslb_zone_oper" {
  value = ["${data.thunder_gslb_zone_oper.thunder_gslb_zone_oper}"]
}