provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_nat_inside_source_statistics_oper" "thunder_cgnv6_nat_inside_source_statistics_oper" {
}
output "get_cgnv6_nat_inside_source_statistics_oper" {
  value = ["${data.thunder_cgnv6_nat_inside_source_statistics_oper.thunder_cgnv6_nat_inside_source_statistics_oper}"]
}