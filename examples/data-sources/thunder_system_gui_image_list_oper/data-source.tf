provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_gui_image_list_oper" "thunder_system_gui_image_list_oper" {
}
output "get_system_gui_image_list_oper" {
  value = ["${data.thunder_system_gui_image_list_oper.thunder_system_gui_image_list_oper}"]
}