provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_web_category_category_list_stats" "thunder_web_category_category_list_stats" {

  name = "test"
}
output "get_web_category_category_list_stats" {
  value = ["${data.thunder_web_category_category_list_stats.thunder_web_category_category_list_stats}"]
}