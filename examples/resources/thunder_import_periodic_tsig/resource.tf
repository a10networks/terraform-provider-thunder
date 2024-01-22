provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_tsig" "thunder_import_periodic_tsig" {
  period        = 26830136
  tsig          = "50"
  use_mgmt_port = 1
  remote_file   = "test"
}