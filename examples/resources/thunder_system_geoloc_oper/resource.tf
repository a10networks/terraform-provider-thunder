provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_geoloc_oper" "thunder_system_geoloc_oper" {
}
output "get_system_geoloc_oper" {
  value = ["${data.thunder_system_geoloc_oper.thunder_system_geoloc_oper}"]
}