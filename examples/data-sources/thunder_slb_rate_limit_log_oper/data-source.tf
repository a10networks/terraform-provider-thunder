provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_rate_limit_log_oper" "thunder_slb_rate_limit_log_oper" {
}
output "get_slb_rate_limit_log_oper" {
  value = ["${data.thunder_slb_rate_limit_log_oper.thunder_slb_rate_limit_log_oper}"]
}