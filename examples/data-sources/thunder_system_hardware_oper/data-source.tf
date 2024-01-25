provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_hardware_oper" "thunder_system_hardware_oper" {
}
output "get_system_hardware_oper" {
  value = ["${data.thunder_system_hardware_oper.thunder_system_hardware_oper}"]
}