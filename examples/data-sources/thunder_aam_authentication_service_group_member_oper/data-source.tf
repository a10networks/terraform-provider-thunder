provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_service_group_member_oper" "thunder_aam_authentication_service_group_member_oper" {

  name = "test"
  port = 1812

}
output "get_aam_authentication_service_group_member_oper" {
  value = ["${data.thunder_aam_authentication_service_group_member_oper.thunder_aam_authentication_service_group_member_oper}"]
}