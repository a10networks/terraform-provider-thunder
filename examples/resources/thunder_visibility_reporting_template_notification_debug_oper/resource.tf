provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_reporting_template_notification_debug_oper" "thunder_visibility_reporting_template_notification_debug_oper" {
}
output "get_visibility_reporting_template_notification_debug_oper" {
  value = ["${data.thunder_visibility_reporting_template_notification_debug_oper.thunder_visibility_reporting_template_notification_debug_oper}"]
}