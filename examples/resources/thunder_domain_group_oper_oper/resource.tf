provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_domain_group_oper_oper" "thunder_domain_group_oper_oper" {

}
output "get_domain_group_oper_oper" {
  value = ["${data.thunder_domain_group_oper_oper.thunder_domain_group_oper_oper}"]
}