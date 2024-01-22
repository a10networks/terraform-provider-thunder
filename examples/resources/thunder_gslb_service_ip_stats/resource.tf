provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_service_ip_stats" "thunder_gslb_service_ip_stats" {

  node_name = "s27"
}
output "get_gslb_service_ip_stats" {
  value = ["${data.thunder_gslb_service_ip_stats.thunder_gslb_service_ip_stats}"]
}