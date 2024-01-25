provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_nptv6_domain_stats" "thunder_cgnv6_nptv6_domain_stats" {

  name = "test"
}
output "get_cgnv6_nptv6_domain_stats" {
  value = ["${data.thunder_cgnv6_nptv6_domain_stats.thunder_cgnv6_nptv6_domain_stats}"]
}