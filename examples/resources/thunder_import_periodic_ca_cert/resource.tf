provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_ca_cert" "thunder_import_periodic_ca_cert" {
  ca_cert          = "test"
  certificate_type = "pem"
  period           = 12438129
  pfx_password     = "19"
  use_mgmt_port    = 1
  remote_file      = "test"
}