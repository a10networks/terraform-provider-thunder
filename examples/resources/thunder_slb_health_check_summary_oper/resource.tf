provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_health_check_summary_oper" "thunder_slb_health_check_summary_oper" {
}
output "get_slb_health_check_summary_oper" {
  value = ["${data.thunder_slb_health_check_summary_oper.thunder_slb_health_check_summary_oper}"]
}