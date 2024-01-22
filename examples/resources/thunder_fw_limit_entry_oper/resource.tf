provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_limit_entry_oper" "thunder_fw_limit_entry_oper" {

}
output "get_fw_limit_entry_oper" {
  value = ["${data.thunder_fw_limit_entry_oper.thunder_fw_limit_entry_oper}"]
}