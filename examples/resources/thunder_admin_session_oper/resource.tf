provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_admin_session_oper" "thunder_admin_session_oper" {
}
output "get_admin_session_oper" {
  value = ["${data.thunder_admin_session_oper.thunder_admin_session_oper}"]
}