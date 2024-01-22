provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_rate_limit_summary_oper" "thunder_fw_rate_limit_summary_oper" {

}
output "get_fw_rate_limit_summary_oper" {
  value = ["${data.thunder_fw_rate_limit_summary_oper.thunder_fw_rate_limit_summary_oper}"]
}