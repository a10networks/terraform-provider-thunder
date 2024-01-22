provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_local_uri_file" "thunder_import_periodic_local_uri_file" {
  local_uri_file = "26"
  period         = 24293212
  use_mgmt_port  = 1
  remote_file    = "test"
}