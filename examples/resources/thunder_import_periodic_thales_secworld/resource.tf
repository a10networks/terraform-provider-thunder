provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_thales_secworld" "thunder_import_periodic_thales_secworld" {
  overwrite       = 1
  period          = 30204088
  thales_secworld = "21"
  use_mgmt_port   = 1
  remote_file     = "test"
}