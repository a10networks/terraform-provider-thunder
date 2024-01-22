provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_domain_list" "thunder_import_periodic_domain_list" {
  domain_list   = "test"
  period        = 2349332
  use_mgmt_port = 1
  remote_file   = "test"
}