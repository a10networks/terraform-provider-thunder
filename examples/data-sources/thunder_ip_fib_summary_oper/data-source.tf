provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ip_fib_summary_oper" "thunder_ip_fib_summary_oper" {
}
output "get_ip_fib_summary_oper" {
  value = ["${data.thunder_ip_fib_summary_oper.thunder_ip_fib_summary_oper}"]
}