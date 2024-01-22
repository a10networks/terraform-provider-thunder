provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_service_group_oper" "thunder_fw_service_group_oper" {

  name = "sg1"
}
output "get_fw_service_group_oper" {
  value = ["${data.thunder_fw_service_group_oper.thunder_fw_service_group_oper}"]
}