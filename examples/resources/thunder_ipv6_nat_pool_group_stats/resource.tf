provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ipv6_nat_pool_group_stats" "thunder_ipv6_nat_pool_group_stats" {

  pool_group_name = "35"
}
output "get_ipv6_nat_pool_group_stats" {
  value = ["${data.thunder_ipv6_nat_pool_group_stats.thunder_ipv6_nat_pool_group_stats}"]
}