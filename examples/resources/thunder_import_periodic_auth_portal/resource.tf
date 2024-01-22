provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_auth_portal" "thunder_import_periodic_auth_portal" {
  auth_portal   = "test"
  period        = 27578005
  use_mgmt_port = 1
  remote_file   = "test"
}