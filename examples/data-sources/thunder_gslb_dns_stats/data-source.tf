provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_dns_stats" "thunder_gslb_dns_stats" {
}
output "get_gslb_dns_stats" {
  value = ["${data.thunder_gslb_dns_stats.thunder_gslb_dns_stats}"]
}