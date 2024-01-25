provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ipv6_nat_pool_oper" "thunder_ipv6_nat_pool_oper" {

  pool_name = "11"
}
output "get_ipv6_nat_pool_oper" {
  value = ["${data.thunder_ipv6_nat_pool_oper.thunder_ipv6_nat_pool_oper}"]
}