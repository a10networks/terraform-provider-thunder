provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dynamic_class_list_oper" "thunder_ddos_dynamic_class_list_oper" {

}
output "get_ddos_dynamic_class_list_oper" {
  value = ["${data.thunder_ddos_dynamic_class_list_oper.thunder_ddos_dynamic_class_list_oper}"]
}