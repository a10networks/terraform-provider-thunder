provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_fixed_nat_disabled_config_oper" "thunder_cgnv6_fixed_nat_disabled_config_oper" {

}
output "get_cgnv6_fixed_nat_disabled_config_oper" {
  value = ["${data.thunder_cgnv6_fixed_nat_disabled_config_oper.thunder_cgnv6_fixed_nat_disabled_config_oper}"]
}