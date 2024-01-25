provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ip_nat_pool_group_stats" "thunder_ip_nat_pool_group_stats" {

  pool_group_name = "53"
}
output "get_ip_nat_pool_group_stats" {
  value = ["${data.thunder_ip_nat_pool_group_stats.thunder_ip_nat_pool_group_stats}"]
}