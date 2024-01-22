provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_power_on_self_test_oper" "thunder_system_power_on_self_test_oper" {

}
output "get_system_power_on_self_test_oper" {
  value = ["${data.thunder_system_power_on_self_test_oper.thunder_system_power_on_self_test_oper}"]
}