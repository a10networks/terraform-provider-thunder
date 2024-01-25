provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_web_category_reputation_scope_stats" "thunder_web_category_reputation_scope_stats" {

  name = "test"

}
output "get_web_category_reputation_scope_stats" {
  value = ["${data.thunder_web_category_reputation_scope_stats.thunder_web_category_reputation_scope_stats}"]
}