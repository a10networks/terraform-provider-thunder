provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_class_list_convert" "thunder_file_class_list_convert" {
  host            = "10.64.3.190"
  class_list_type = "ipv4"
  name            = "test_class_list"
  path            = "/class-list-ac1-a10"
  protocol        = "http"
  use_mgmt_port   = 1
}