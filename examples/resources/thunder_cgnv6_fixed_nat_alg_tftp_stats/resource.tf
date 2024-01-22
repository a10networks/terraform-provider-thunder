provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_fixed_nat_alg_tftp_stats" "thunder_cgnv6_fixed_nat_alg_tftp_stats" {
}
output "get_cgnv6_fixed_nat_alg_tftp_stats" {
  value = ["${data.thunder_cgnv6_fixed_nat_alg_tftp_stats.thunder_cgnv6_fixed_nat_alg_tftp_stats}"]
}