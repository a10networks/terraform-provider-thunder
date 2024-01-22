provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_ssl_crl" "thunder_import_periodic_ssl_crl" {
  period        = 23792473
  ssl_crl       = "214"
  use_mgmt_port = 1
  remote_file   = "test"
}