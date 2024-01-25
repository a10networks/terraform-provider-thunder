provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_notification_template_debug_log_oper" "thunder_ddos_notification_template_debug_log_oper" {
}
output "get_ddos_notification_template_debug_log_oper" {
  value = ["${data.thunder_ddos_notification_template_debug_log_oper.thunder_ddos_notification_template_debug_log_oper}"]
}