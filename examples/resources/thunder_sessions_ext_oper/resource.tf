provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_sessions_ext_oper" "thunder_sessions_ext_oper" {
}
output "get_sessions_ext_oper" {
  value = ["${data.thunder_sessions_ext_oper.thunder_sessions_ext_oper}"]
}