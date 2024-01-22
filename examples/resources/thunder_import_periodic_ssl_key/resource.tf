provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_ssl_key" "thunder_import_periodic_ssl_key" {
  period        = 2796745
  secured       = 0
  ssl_key       = "70"
  use_mgmt_port = 1
  remote_file   = "test"
}