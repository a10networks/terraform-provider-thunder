provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_class_list_convert" "thunder_import_periodic_class_list_convert" {
  class_list_convert = "test"
  class_list_type    = "ac"
  period             = 13074728
  use_mgmt_port      = 1
  remote_file        = "test"
}