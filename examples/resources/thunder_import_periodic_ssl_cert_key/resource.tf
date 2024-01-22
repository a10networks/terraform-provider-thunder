provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_ssl_cert_key" "thunder_import_periodic_ssl_cert_key" {
  period        = 28476762
  secured       = 0
  ssl_cert_key  = "bulk"
  use_mgmt_port = 1
  remote_file   = "test"
}