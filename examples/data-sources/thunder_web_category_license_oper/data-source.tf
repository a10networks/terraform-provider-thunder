provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_web_category_license_oper" "thunder_web_category_license_oper" {

}
output "get_web_category_license_oper" {
  value = ["${data.thunder_web_category_license_oper.thunder_web_category_license_oper}"]
}