provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_site_stats" "thunder_gslb_site_stats" {

  site_name = "3"
}
output "get_gslb_site_stats" {
  value = ["${data.thunder_gslb_site_stats.thunder_gslb_site_stats}"]
}