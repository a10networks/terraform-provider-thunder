provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fail_safe_config_oper" "thunder_fail_safe_config_oper" {
}
output "get_fail_safe_config_oper" {
  value = ["${data.thunder_fail_safe_config_oper.thunder_fail_safe_config_oper}"]
}