provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_lsn_alg_mgcp_stats" "thunder_cgnv6_lsn_alg_mgcp_stats" {
}
output "get_cgnv6_lsn_alg_mgcp_stats" {
  value = ["${data.thunder_cgnv6_lsn_alg_mgcp_stats.thunder_cgnv6_lsn_alg_mgcp_stats}"]
}