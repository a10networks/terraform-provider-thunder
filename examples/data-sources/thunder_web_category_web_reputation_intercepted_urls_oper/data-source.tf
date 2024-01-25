provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_web_category_web_reputation_intercepted_urls_oper" "thunder_web_category_web_reputation_intercepted_urls_oper" {

}
output "get_web_category_web_reputation_intercepted_urls_oper" {
  value = ["${data.thunder_web_category_web_reputation_intercepted_urls_oper.thunder_web_category_web_reputation_intercepted_urls_oper}"]
}