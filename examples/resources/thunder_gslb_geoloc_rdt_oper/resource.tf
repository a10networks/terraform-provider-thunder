provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_geoloc_rdt_oper" "thunder_gslb_geoloc_rdt_oper" {
}
output "get_gslb_geoloc_rdt_oper" {
  value = ["${data.thunder_gslb_geoloc_rdt_oper.thunder_gslb_geoloc_rdt_oper}"]
}