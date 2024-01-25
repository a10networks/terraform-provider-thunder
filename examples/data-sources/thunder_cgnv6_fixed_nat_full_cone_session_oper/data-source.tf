provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_fixed_nat_full_cone_session_oper" "thunder_cgnv6_fixed_nat_full_cone_session_oper" {

}
output "get_cgnv6_fixed_nat_full_cone_session_oper" {
  value = ["${data.thunder_cgnv6_fixed_nat_full_cone_session_oper.thunder_cgnv6_fixed_nat_full_cone_session_oper}"]
}