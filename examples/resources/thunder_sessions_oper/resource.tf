provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_sessions_oper" "thunder_sessions_oper" {
}
output "get_sessions_oper" {
  value = ["${data.thunder_sessions_oper.thunder_sessions_oper}"]
}