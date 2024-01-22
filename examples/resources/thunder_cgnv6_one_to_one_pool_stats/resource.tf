provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_one_to_one_pool_stats" "thunder_cgnv6_one_to_one_pool_stats" {

  pool_name = "test"

}
output "get_cgnv6_one_to_one_pool_stats" {
  value = ["${data.thunder_cgnv6_one_to_one_pool_stats.thunder_cgnv6_one_to_one_pool_stats}"]
}