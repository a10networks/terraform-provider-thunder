provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_clock_show_oper" "thunder_clock_show_oper" {
}
output "get_clock_show_oper" {
  value = ["${data.thunder_clock_show_oper.thunder_clock_show_oper}"]
}