provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_reporting_template_notification_template_name_stats" "thunder_visibility_reporting_template_notification_template_name_stats" {

  name = "test"
}
output "get_visibility_reporting_template_notification_template_name_stats" {
  value = ["${data.thunder_visibility_reporting_template_notification_template_name_stats.thunder_visibility_reporting_template_notification_template_name_stats}"]
}