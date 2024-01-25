provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_access_list_summary_oper" "thunder_access_list_summary_oper" {
}
output "get_access_list_summary_oper" {
  value = ["${data.thunder_access_list_summary_oper.thunder_access_list_summary_oper}"]
}