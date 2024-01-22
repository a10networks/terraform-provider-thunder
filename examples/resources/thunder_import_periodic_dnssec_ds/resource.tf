provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_dnssec_ds" "thunder_import_periodic_dnssec_ds" {
  dnssec_ds     = "test"
  period        = 3123315
  use_mgmt_port = 1
  remote_file   = "test"
}