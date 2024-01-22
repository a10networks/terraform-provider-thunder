provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_ssl_cert" "thunder_import_periodic_ssl_cert" {
  certificate_type = "pem"
  period           = 14543277
  pfx_password     = "25"
  ssl_cert         = "66"
  use_mgmt_port    = 1
  remote_file      = "test"
}