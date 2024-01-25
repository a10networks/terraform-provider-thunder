provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_server_oper" "thunder_cgnv6_server_oper" {

  name = "test12345"
}
output "get_cgnv6_server_oper" {
  value = ["${data.thunder_cgnv6_server_oper.thunder_cgnv6_server_oper}"]
}