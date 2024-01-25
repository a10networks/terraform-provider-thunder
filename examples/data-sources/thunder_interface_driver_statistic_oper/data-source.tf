provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_driver_statistic_oper" "thunder_interface_driver_statistic_oper" {

}
output "get_interface_driver_statistic_oper" {
  value = ["${data.thunder_interface_driver_statistic_oper.thunder_interface_driver_statistic_oper}"]
}