provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ip_nat_pool_oper" "thunder_ip_nat_pool_oper" {

  pool_name = "a24"
}
output "get_ip_nat_pool_oper" {
  value = ["${data.thunder_ip_nat_pool_oper.thunder_ip_nat_pool_oper}"]
}