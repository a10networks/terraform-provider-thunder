provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dynamic_class_list" "thunder_ddos_dynamic_class_list" {

  class_list_name = "8"

}