provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_class_list_oper" "thunder_class_list_oper" {

  name = "test"
}
output "get_class_list_oper" {
  value = ["${data.thunder_class_list_oper.thunder_class_list_oper}"]
}