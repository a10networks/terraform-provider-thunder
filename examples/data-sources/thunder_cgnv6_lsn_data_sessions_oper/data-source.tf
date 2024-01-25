provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_lsn_data_sessions_oper" "thunder_cgnv6_lsn_data_sessions_oper" {

}
output "get_cgnv6_lsn_data_sessions_oper" {
  value = ["${data.thunder_cgnv6_lsn_data_sessions_oper.thunder_cgnv6_lsn_data_sessions_oper}"]
}