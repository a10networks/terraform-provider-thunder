provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_web_category_statistics_stats" "thunder_web_category_statistics_stats" {

}
output "get_web_category_statistics_stats" {
  value = ["${data.thunder_web_category_statistics_stats.thunder_web_category_statistics_stats}"]
}