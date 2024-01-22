provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_geo_location_db_oper" "thunder_ddos_geo_location_db_oper" {
}
output "get_ddos_geo_location_db_oper" {
  value = ["${data.thunder_ddos_geo_location_db_oper.thunder_ddos_geo_location_db_oper}"]
}