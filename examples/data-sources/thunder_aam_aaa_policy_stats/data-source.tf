provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_aaa_policy_stats" "thunder_aam_aaa_policy_stats" {

  name = "test"
}
output "get_aam_aaa_policy_stats" {
  value = ["${data.thunder_aam_aaa_policy_stats.thunder_aam_aaa_policy_stats}"]
}