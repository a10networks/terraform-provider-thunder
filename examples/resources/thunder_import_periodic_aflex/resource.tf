provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_aflex" "thunder_import_periodic_aflex" {
  aflex         = "test"
  period        = 30503109
  use_mgmt_port = 1
  remote_file   = "test"
}