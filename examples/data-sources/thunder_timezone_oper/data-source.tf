provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_timezone_oper" "thunder_timezone_oper" {
}
output "get_timezone_oper" {
  value = ["${data.thunder_timezone_oper.thunder_timezone_oper}"]
}