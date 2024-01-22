provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_dns_response_rate_limiting_stats" "thunder_slb_dns_response_rate_limiting_stats" {
}
output "get_slb_dns_response_rate_limiting_stats" {
  value = ["${data.thunder_slb_dns_response_rate_limiting_stats.thunder_slb_dns_response_rate_limiting_stats}"]
}