provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ipv6_nat_pool_stats" "thunder_ipv6_nat_pool_stats" {

  pool_name = "11"
}
output "get_ipv6_nat_pool_stats" {
  value = ["${data.thunder_ipv6_nat_pool_stats.thunder_ipv6_nat_pool_stats}"]
}