provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_template_gtp_policy_stats" "thunder_template_gtp_policy_stats" {

  name = "test"
}
output "get_template_gtp_policy_stats" {
  value = ["${data.thunder_template_gtp_policy_stats.thunder_template_gtp_policy_stats}"]
}