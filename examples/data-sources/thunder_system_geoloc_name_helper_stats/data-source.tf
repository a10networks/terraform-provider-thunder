provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_geoloc_name_helper_stats" "thunder_system_geoloc_name_helper_stats" {
}
output "get_system_geoloc_name_helper_stats" {
  value = ["${data.thunder_system_geoloc_name_helper_stats.thunder_system_geoloc_name_helper_stats}"]
}