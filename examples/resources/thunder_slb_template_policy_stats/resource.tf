provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_template_policy_stats" "thunder_slb_template_policy_stats" {

  name = "test"

}
output "get_slb_template_policy_stats" {
  value = ["${data.thunder_slb_template_policy_stats.thunder_slb_template_policy_stats}"]
}