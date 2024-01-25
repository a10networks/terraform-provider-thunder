provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_nat_pool_oper" "thunder_cgnv6_nat_pool_oper" {

  pool_name = "testPool"
}
output "get_cgnv6_nat_pool_oper" {
  value = ["${data.thunder_cgnv6_nat_pool_oper.thunder_cgnv6_nat_pool_oper}"]
}