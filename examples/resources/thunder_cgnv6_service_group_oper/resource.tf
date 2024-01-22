provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_service_group_oper" "thunder_cgnv6_service_group_oper" {

  name = "testing"
}
output "get_cgnv6_service_group_oper" {
  value = ["${data.thunder_cgnv6_service_group_oper.thunder_cgnv6_service_group_oper}"]
}