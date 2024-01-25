provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_file_portal_oper" "thunder_aam_authentication_file_portal_oper" {
}
output "get_aam_authentication_file_portal_oper" {
  value = ["${data.thunder_aam_authentication_file_portal_oper.thunder_aam_authentication_file_portal_oper}"]
}