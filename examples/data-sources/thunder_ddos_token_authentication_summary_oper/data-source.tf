provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_token_authentication_summary_oper" "thunder_ddos_token_authentication_summary_oper" {
}
output "get_ddos_token_authentication_summary_oper" {
  value = ["${data.thunder_ddos_token_authentication_summary_oper.thunder_ddos_token_authentication_summary_oper}"]
}