provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_geoloc_name_helper_oper" "thunder_system_geoloc_name_helper_oper" {
}
output "get_system_geoloc_name_helper_oper" {
  value = ["${data.thunder_system_geoloc_name_helper_oper.thunder_system_geoloc_name_helper_oper}"]
}