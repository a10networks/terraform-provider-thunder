provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_geoloc_stats" "thunder_system_geoloc_stats" {
}
output "get_system_geoloc_stats" {
  value = ["${data.thunder_system_geoloc_stats.thunder_system_geoloc_stats}"]
}