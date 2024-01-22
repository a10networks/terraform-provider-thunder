provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_stats" "thunder_gslb_zone_stats" {

  name = "a11"
}
output "get_gslb_zone_stats" {
  value = ["${data.thunder_gslb_zone_stats.thunder_gslb_zone_stats}"]
}