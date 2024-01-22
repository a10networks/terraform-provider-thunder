provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_thales_kmdata" "thunder_import_periodic_thales_kmdata" {
  overwrite     = 1
  period        = 27792761
  thales_kmdata = "62"
  use_mgmt_port = 1
  remote_file   = "test"
}