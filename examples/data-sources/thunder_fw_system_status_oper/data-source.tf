provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_system_status_oper" "thunder_fw_system_status_oper" {

}
output "get_fw_system_status_oper" {
  value = ["${data.thunder_fw_system_status_oper.thunder_fw_system_status_oper}"]
}